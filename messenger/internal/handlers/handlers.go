package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"messenger/internal/models"
	"messenger/internal/repository"
	ws "messenger/internal/websocket"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Handler struct {
	repo     *repository.Repository
	hub      *ws.Hub
	upgrader websocket.Upgrader
}

func NewHandler(repo *repository.Repository, hub *ws.Hub) *Handler {
	return &Handler{
		repo: repo,
		hub:  hub,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // In production, implement proper origin checking
			},
		},
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/api/chats", h.CreateChat).Methods("POST")
	r.HandleFunc("/api/chats/{chatID}/messages", h.GetMessages).Methods("GET")
	r.HandleFunc("/api/users/{userID}/chats", h.GetUserChats).Methods("GET")
	r.HandleFunc("/ws/{userID}/{chatID}", h.HandleWebSocket)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(chat.Userto, chat.UserFrom)
	if err := h.repo.CreateChat(&chat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.repo.AddUserToChat(chat.Userto, chat.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.repo.AddUserToChat(chat.UserFrom, chat.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chat)
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID, err := strconv.Atoi(vars["chatID"])
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	messages, err := h.repo.GetChatMessages(chatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) GetUserChats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// if err != nil {
	// 	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }

	chats, err := h.repo.GetUserChats(vars["userID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(chats)
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])
	chatID, _ := strconv.Atoi(vars["chatID"])

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &ws.Client{
		Hub:    h.hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		UserID: userID,
		ChatID: chatID,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
