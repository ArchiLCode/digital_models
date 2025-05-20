package repository

import (
	"database/sql"
	"messenger/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *models.User) error {
	query := `
        INSERT INTO users (username, password_hash)
        VALUES ($1, $2)
        RETURNING id, created_at`
	return r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *Repository) UpdateUser(user *models.User) error {
	query := `
        UPDATE users 
        SET username = $1
        WHERE id = $2
        RETURNING created_at`
	return r.db.QueryRow(query, user.Username, user.ID).Scan(&user.CreatedAt)
}

func (r *Repository) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, password_hash, created_at FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, password_hash, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) CreateChat(chat *models.Chat) error {
	query := `
        INSERT INTO chats (name, is_private)
        VALUES ($1, $2)
        RETURNING id, created_at`
	return r.db.QueryRow(query, chat.Name, chat.IsPrivate).Scan(&chat.ID, &chat.CreatedAt)
}

func (r *Repository) AddUserToChat(userID string, chatID int) error {
	query := `INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, chatID, userID)
	return err
}

func (r *Repository) SaveMessage(msg *models.Message) error {
	query := `
        INSERT INTO messages (chat_id, user_id, content)
        VALUES ($1, $2, $3)
        RETURNING id, created_at`
	return r.db.QueryRow(query, msg.ChatID, msg.UserID, msg.Content).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *Repository) GetChatMessages(chatID int) ([]models.Message, error) {	query := `
        SELECT id, chat_id, user_id, content, created_at
        FROM messages
        WHERE chat_id = $1
        ORDER BY id ASC
        LIMIT 50`

	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserID, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *Repository) GetUserChats(userID string) ([]models.Chat, error) {
	// Получаем все чаты пользователя вместе с их участниками
	query := `
		WITH user_chats AS (
			SELECT c.id, c.name, c.is_private, c.created_at
			FROM chats c
			JOIN chat_participants cp ON c.id = cp.chat_id
			WHERE cp.user_id = $1
		)
		SELECT 
			uc.id, 
			uc.name, 
			uc.is_private, 
			uc.created_at,
			(
				SELECT cp.user_id 
				FROM chat_participants cp 
				WHERE cp.chat_id = uc.id AND cp.user_id != $1 
				LIMIT 1
			) as other_user_id
		FROM user_chats uc`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		var otherUserID sql.NullString
		if err := rows.Scan(
			&chat.ID,
			&chat.Name,
			&chat.IsPrivate,
			&chat.CreatedAt,
			&otherUserID,
		); err != nil {
			return nil, err
		}
		if otherUserID.Valid {
			chat.Userto = otherUserID.String
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chats, nil
}
