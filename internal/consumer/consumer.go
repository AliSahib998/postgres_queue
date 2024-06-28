package consumer

import (
	"database/sql"
	"mail-service/internal/controller"
)

type PGQConsumer struct {
	db         *sql.DB
	controller *controller.Controller
}

func NewPGQConsumer(db *sql.DB, controller *controller.Controller) (*PGQConsumer, error) {
	return &PGQConsumer{db: db, controller: controller}, nil
}
