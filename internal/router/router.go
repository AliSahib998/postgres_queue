package router

import (
	"github.com/gin-gonic/gin"
	"mail-service/internal/configs"
	"mail-service/internal/controller"
	"mail-service/internal/publisher"
)

type Router struct {
	config     *configs.Router
	engine     *gin.Engine
	controller *controller.Controller
}

func NewRouter(configs *configs.Configs, pgqPublisher *publisher.PGQPublisher) *Router {
	appController := controller.NewController(configs, pgqPublisher)
	return &Router{
		config:     configs.Router,
		engine:     gin.Default(),
		controller: appController,
	}
}

func (r *Router) Run() error {
	v1 := r.engine.Group("/api/v1/email")
	{
		v1.POST("/send",
			r.SendMail)
	}
	return r.engine.Run(":" + r.config.Port)
}
