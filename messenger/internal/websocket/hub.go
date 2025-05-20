package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"messenger/internal/models"
	"messenger/internal/repository"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID int
	ChatID int
}

type Message struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mutex      sync.Mutex
	Repo       repository.Repository
}

func NewHub(repo *repository.Repository) *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Repo:       *repo,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()

		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.Mutex.Unlock()

		case message := <-h.Broadcast:
			h.Mutex.Lock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			h.Mutex.Unlock()
		}
	}
}

func (h *Hub) BroadcastToChat(chatID int, message []byte) {
	h.Mutex.Lock()
	for client := range h.Clients {
		if client.ChatID == chatID {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
	h.Mutex.Unlock()
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("error unmarshaling message: %v", err)
			continue
		}

		// Handle different message types
		switch msg.Type {		case "chat_message":
			var chatMsg models.Message
			if err := json.Unmarshal(msg.Content, &chatMsg); err != nil {
				log.Printf("error unmarshaling chat message: %v", err)
				continue
			}
			
			// Save message to database
			if err := c.Hub.Repo.SaveMessage(&chatMsg); err != nil {
				log.Printf("error saving message to database: %v", err)
				continue
			}

			// Broadcast message to all clients in the chat
			messageBytes, err := json.Marshal(Message{
				Type:    "chat_message",
				Content: msg.Content,
			})
			if err != nil {
				log.Printf("error marshaling message for broadcast: %v", err)
				continue
			}
			c.Hub.BroadcastToChat(chatMsg.ChatID, messageBytes)
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
