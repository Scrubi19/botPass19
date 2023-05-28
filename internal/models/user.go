package models

import "github.com/google/uuid"

type User struct {
	id     uuid.UUID `json:"id"`
	chatId int64     `json:"chatId"`
	pass   string    `json:"pass"`
}
