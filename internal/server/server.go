package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server - структура, представляющая HTTP-сервер
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer - конструктор сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем маршрутизатор
	router := http.NewServeMux()

	// Добавляем корневой маршрут и маршрут для загрузки
	router.HandleFunc("/", handlers.HandleHTML)
	router.HandleFunc("/upload", handlers.HandleUpload)

	// Настраиваем параметры сервера
	return &Server{
		Logger: logger,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
