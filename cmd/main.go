package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	// Создаем и запускаем сервер
	srv := server.NewServer(logger)
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
