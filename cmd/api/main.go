// Package main
package main

import (
	"log"
	"net/http"

	"github.com/johernandezvaz/devboard/internal/server"
)

func main() {
	// Aqui arranca el servidor
	srv := server.New(":8080")

	srv.RegisterRoutes("GET /health", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"status": "ok"}`))
	}))

	log.Println("Servicio funcionando en :8080")

	if err := srv.Start(); err != nil {
		log.Fatalf("Error al iniciar el servidor %v", err)
	}

}
