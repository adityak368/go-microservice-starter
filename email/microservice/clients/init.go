package clients

import (
	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

type Clients struct {
}

var clients *Clients = nil

func InitClients(service micro.Service) {
	clients = new(Clients)
	logger.Info.Println("Initialized Microservice Clients")
}

func GetClients() *Clients {
	return clients
}
