package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "morse-converter: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("The server is running on http://localhost:8080 and in the file")
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
