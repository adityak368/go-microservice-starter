package microservice

import (
	"email/microservice/clients"
	"email/microservice/handlers"
	"email/microservice/publishers"
	"email/microservice/subscribers"

	"github.com/micro/go-micro/v2"
)

// Initialize the service
func InitService(service micro.Service) {

	// Register Handlers
	handlers.InitHandlers(service)

	// Initialize microservice client
	clients.InitClients(service)

	// Initialize the microservice publishers
	publishers.InitPublishers(service)

	// Register Struct as Subscribers
	subscribers.InitSubscribers(service)
}
