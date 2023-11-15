package main

import (
	"botPass19/config"
	"botPass19/internal/db"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(cfg)

	db, err := db.InitDB(cfg)
}
