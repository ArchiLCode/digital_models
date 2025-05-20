package models

import "github.com/google/uuid"

// Определяем тип для цели пользователя
type Goal string

// Goal Возможные значения для цели
const (
	GoalModel Goal = "model"
	GoalScout Goal = "scout"
)

// Структура User
type User struct {
	ID       uuid.UUID
	Name     string
	LastName string
	Email    string
	Goal     Goal
	PassHash []byte
}

func (g Goal) String() string {
	return string(g)
}
