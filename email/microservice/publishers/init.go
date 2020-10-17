package publishers

import (
	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

type Publishers struct {
}

var publishers *Publishers

func InitPublishers(service micro.Service) {
	publishers = new(Publishers)
	logger.Info.Println("Initialized Microservice Publishers")
}

func GetPublishers() *Publishers {
	return publishers
}
