package main

import (
	"fmt"
	"log/slog"
	"medodsTest/config"
	"medodsTest/storage"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Config loaded")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Database)

	strg, err := storage.NewStorage(dbUrl)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer strg.Close()
	log.Info("DB created", slog.String("DB URL:", dbUrl))
}
