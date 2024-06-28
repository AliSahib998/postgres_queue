package publisher

import (
	"context"
	"database/sql"
	"fmt"
	"go.dataddo.com/pgq"
)

type Publisher interface {
	PublishMessage(queueName string, messages []*pgq.MessageOutgoing) error
}

type PGQPublisher struct {
	db *sql.DB
	pgq.Publisher
}

func NewPGQPublisher(db *sql.DB) (*PGQPublisher, error) {
	publisher := pgq.NewPublisher(db)
	return &PGQPublisher{db: db, Publisher: publisher}, nil
}

func (p *PGQPublisher) PublishMessage(queueName string, messages []*pgq.MessageOutgoing) error {
	_, err := p.Publisher.Publish(context.Background(), queueName, messages...)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
