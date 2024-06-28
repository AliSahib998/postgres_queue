package controller

import (
	"encoding/json"
	"go.dataddo.com/pgq"
	"io"
	"strconv"

	"github.com/go-gomail/gomail"
)

const (
	queueName         = "mail_queue"
	emailType         = "X-Email-Type"
	emailTypeStandard = "standard"
)

// SendEmail ...
func (c *Controller) SendEmail(req *EmailRequest) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", req.Sender)
	msg.SetHeader("Bcc", req.Receivers...)
	msg.SetHeader("Subject", req.Subject)

	if req.Body != "" {
		msg.SetBody("text/html", req.Body)
	}

	var err = attachAttachments(msg, req.Attachments)
	if err != nil {
		return err
	}

	smtpHost := c.configs.SMTP.Host
	smtpUsername := c.configs.SMTP.Username
	smtpPassword := c.configs.SMTP.Password

	smtpPort, err := strconv.Atoi(c.configs.SMTP.Port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	return d.DialAndSend(msg)
}

func attachAttachments(msg *gomail.Message, attachments []*RequestAttachment) error {
	for _, attachment := range attachments {
		msg.Attach(attachment.FileName+"."+attachment.FileType, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(attachment.Data)
			return err
		}))
	}
	return nil
}

func (c *Controller) PublishMails(requests []*EmailRequest) error {
	var messages []*pgq.MessageOutgoing

	for _, request := range requests {
		payload, err := json.Marshal(request)
		if err != nil {
			return err
		}

		message := &pgq.MessageOutgoing{
			Metadata: pgq.Metadata{
				emailType: emailTypeStandard,
			},
			Payload: json.RawMessage(payload),
		}
		messages = append(messages, message)
	}

	err := c.publisher.PublishMessage(queueName, messages)
	if err != nil {
		return err
	}

	return nil
}
