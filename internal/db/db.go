package db

import (
	"botPass19/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUsername,
		cfg.DbName,
		cfg.DbPassword,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: cfg.DbPreferSimpleProtocol,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
