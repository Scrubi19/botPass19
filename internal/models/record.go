package models

import "github.com/google/uuid"

type Record struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	FolderId    uuid.UUID `json:"folderId"`
	Login       string    `json:"login"`
	Pass        string    `json:"pass"`
	Description string    `json:"description"`
}
