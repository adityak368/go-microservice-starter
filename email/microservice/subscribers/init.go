package subscribers

import (
	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

func InitSubscribers(service micro.Service) {
	micro.RegisterSubscriber("auth.userCreated", service.Server(), new(UserCreated))
	logger.Info.Println("Initialized Microservice Subscribers")
}
