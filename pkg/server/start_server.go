package server

import (
	"log"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/connection"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberConnURL, _ := connection.ConnectionURLBuilder("fiber")
	fiberListen := fmt.Sprintf(":%s",fiberConnURL)

	// Run server.
	if err := a.Listen(fiberListen); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := connection.ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
