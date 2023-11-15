package auth

import (
	"botPass19/internal/models"
	"context"
)

type Repository interface {
	AuthRepository()
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string) (*models.User, error)
}
