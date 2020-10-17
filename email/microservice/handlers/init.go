package handlers

import (
	"commons/proto/email"

	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

func InitHandlers(service micro.Service) {
	email.RegisterEmailHandler(service.Server(), new(Email))
	logger.Info.Println("Initialized Microservice Handlers")
}
