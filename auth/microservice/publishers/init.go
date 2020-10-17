package publishers

import (
	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

type Publishers struct {
	UserCreated micro.Event
	UserUpdated micro.Event
}

var publishers *Publishers

func InitPublishers(service micro.Service) {
	publishers = new(Publishers)
	publishers.UserCreated = micro.NewEvent("auth.userCreated", service.Client())
	publishers.UserUpdated = micro.NewEvent("auth.userUpdated", service.Client())
	logger.Info.Println("Initialized Microservice Publishers")
}

func GetPublishers() *Publishers {
	return publishers
}
