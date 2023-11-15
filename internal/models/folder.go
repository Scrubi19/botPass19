package models

import "github.com/google/uuid"

type Folder struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
	Icone  string    `json:"icone"`
}
