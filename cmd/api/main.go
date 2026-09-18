// Package main
package main

import (
	"log"

	"github.com/johernandezvaz/devboard/internal/server"
)

func main() {
	// Aqui arranca el servidor
	srv := server.New(":8080")

	log.Println("Servicio funcionando en :8080")

	if err := srv.Start(); err != nil {
		log.Fatalf("Error al iniciar el servidor %v", err)
	}
}
