package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dataddo.com/pgq"
	"mail-service/internal/controller"
)

const (
	queueName         = "mail_queue"
	emailType         = "X-Email-Type"
	emailTypeStandard = "standard"
)

type handler struct {
	controller *controller.Controller
}

func (p *PGQConsumer) ConsumeMailMessage() error {
	h := &handler{controller: p.controller}
	consumer, err := pgq.NewConsumer(p.db, queueName, h)
	if err != nil {
		return err
	}
	go func() {
		err = consumer.Run(context.Background())
		if err != nil {
			panic(err.Error())
		}
	}()

	return nil
}

func (h *handler) HandleMessage(_ context.Context, msg *pgq.MessageIncoming) (processed bool, err error) {
	metadata := msg.Metadata
	if metadata == nil || metadata[emailType] == emailTypeStandard {
		return h.handleStandardEmail(msg)
	}
	return false, fmt.Errorf("unknown email type: %v", metadata[emailType])
}

func (h *handler) handleStandardEmail(msg *pgq.MessageIncoming) (processed bool, err error) {
	var emailReq *controller.EmailRequest
	if err := h.unmarshalPayload(msg.Payload, &emailReq); err != nil {
		return false, err
	}
	if err := h.controller.SendEmail(emailReq); err != nil {
		return false, err
	}
	return true, nil
}

func (h *handler) unmarshalPayload(payload json.RawMessage, v interface{}) error {
	if err := json.Unmarshal(payload, v); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}
	return nil
}
