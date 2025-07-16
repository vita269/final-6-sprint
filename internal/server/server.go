package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server - структура, представляющая HTTP-сервер
type Server struct {
	logger *log.Logger
	Server *http.Server
}

// NewServer - конструктор сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем маршрутизатор
	router := http.NewServeMux()

	// Добавляем корневой маршрут
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет"))
	})

	// Настраиваем параметры сервера
	return &Server{
		logger: logger,
		Server: &http.Server{
			Addr:         ":8080", // Исправлен формат адреса
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

// Start - запуск сервера
func (s *Server) Start() error {
	s.logger.Println("Запуск сервера на порту 8080")
	return s.Server.ListenAndServe()
}

// Stop - корректная остановка сервера
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.logger.Println("Остановка сервера")
	if err := s.Server.Shutdown(ctx); err != nil {
		s.logger.Printf("Ошибка при остановке сервера: %v", err)
		return err
	}
	s.logger.Println("Сервер остановлен")
	return nil
}
