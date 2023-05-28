package models

import "github.com/google/uuid"

type Folder struct {
	id     uuid.UUID `json:"id"`
	userId uuid.UUID `json:"userId"`
	name   string    `json:"name"`
	icone  string    `json:"icone"`
}
