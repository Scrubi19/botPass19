package models

import "github.com/google/uuid"

type Record struct {
	id       uuid.UUID `json:"id"`
	userId   uuid.UUID `json:"userId"`
	folderId uuid.UUID `json:"folderId"`
	login    string    `json:"login"`
	pass     string    `json:"pass"`
	notes    string    `json:"notes"`
}
