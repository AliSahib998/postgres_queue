package controller

import (
	"mail-service/internal/configs"
	"mail-service/internal/publisher"
)

type Controller struct {
	configs   *configs.Configs
	publisher publisher.Publisher
}

func NewController(configs *configs.Configs, pgqPublisher *publisher.PGQPublisher) *Controller {
	return &Controller{
		configs:   configs,
		publisher: pgqPublisher,
	}
}
