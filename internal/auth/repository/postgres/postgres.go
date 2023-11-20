package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"primarykey"`
	Username  string    `gorm:"size:255"`
	Password  string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserSession struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	User
}
