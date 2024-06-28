package consumer

import (
	"database/sql"
	"mail-service/internal/configs"
	"mail-service/internal/controller"
	"mail-service/internal/database"
)

type PGQConsumer struct {
	db         *sql.DB
	controller *controller.Controller
}

func NewPGQConsumer(config *configs.DB, controller *controller.Controller) (*PGQConsumer, error) {
	db, err := database.ConnectDb(config)
	if err != nil {
		return nil, err
	}
	return &PGQConsumer{db: db, controller: controller}, nil
}
