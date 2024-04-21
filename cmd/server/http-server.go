// Package server provides HTTP server functionalities to handle incoming
// web requests. It encapsulates all necessary settings and utilities
// required to start, run, and gracefully shut down an HTTP server.
package server

import (
	"context"
	"net/http"
	"time"
)

// Server encapsulates an HTTP server with configuration settings.
// It allows starting the server and handling requests with specified routes and middleware.
type Server struct {
	httpServer *http.Server
}

// Run starts the HTTP server and listens on the specified port for incoming requests.
//
// Parameters:
// - port: the port number on which the server will listen.
// - handler: the HTTP handler that will handle incoming requests.
//
// Returns:
// - error: an error if the server fails to start or an error occurs during the server's operation.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Close gracefully shuts down the HTTP server.
//
// Parameters:
// - ctx: the context.Context used to control the cancellation of the server shutdown.
//
// Returns:
// - error: an error if the server fails to shut down gracefully.
func (s *Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
