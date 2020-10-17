package microservice

import (
	"auth/microservice/clients"
	"auth/microservice/handlers"
	"auth/microservice/publishers"
	"auth/microservice/subscribers"

	"github.com/micro/go-micro/v2"
)

// Initialize the service
func InitService(service micro.Service, runOnlyHTTPServer bool) {

	if !runOnlyHTTPServer {
		// Register Handlers
		handlers.InitHandlers(service)
	}

	// Initialize microservice client
	clients.InitClients(service)

	// Initialize the microservice publishers
	publishers.InitPublishers(service)

	if !runOnlyHTTPServer {
		// Register Struct as Subscribers
		subscribers.InitSubscribers(service)
	}
}
