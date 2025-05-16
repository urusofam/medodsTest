package main

import (
	"log/slog"
	"medodsTest/config"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err.Error())
	}

	log.Info("Config loaded", slog.String("Address:", cfg.ServerConfig.Address))
}
