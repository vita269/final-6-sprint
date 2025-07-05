package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет"))
	})
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		logger: logger,
		server: server,
	}
}

func (s *Server) Start() error {
	s.logger.Println("Запуск сервпщера на порту 8080")
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.logger.Println("Остановка сервера")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Printf("Ошибка при остановке сервера: %v", err)
		return err
	}
	s.logger.Println("Сервер остановлен")
	return nil
}
