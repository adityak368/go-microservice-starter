package clients

import (
	"commons/proto/email"

	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

type Clients struct {
	Email email.EmailService
}

var clients *Clients = nil

func InitClients(service micro.Service) {
	clients = new(Clients)
	clients.Email = email.NewEmailService("micro.EmailService", service.Client())
	logger.Info.Println("Initialized Microservice Clients")
}

func GetClients() *Clients {
	return clients
}
