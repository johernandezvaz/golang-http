// Package server creates a new HTTP server
package server

import (
	"net/http"
	"time"
)

// Server struct
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// New create a new server
func New(addr string) *Server {
	mux := http.NewServeMux() // Entrutador HTTP

	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,

		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		mux:        mux,
	}

}

// Start server
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// RegisterRoutes register a new route
func (s *Server) RegisterRoutes(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// ServeHTTP server
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}
