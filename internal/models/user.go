package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	ChatId   int64     `json:"chatId"` // telegram chatId (userId)
	Password string    `json:"password"`
}
