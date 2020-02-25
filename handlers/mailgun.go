package handlers

import (
	"context"
	"fmt"
	"github.com/matcornic/hermes/v2"
	"time"
)

func (handlers Handlers) sendHermesMail(email *hermes.Email, subject string, recipient string) error {
	emailBodyHTML, err := handlers.hermes.GenerateHTML(*email)
	if err != nil {
		return fmt.Errorf("error generating the mail: %w", err)
	}

	emailBodyPlainText, err := handlers.hermes.GeneratePlainText(*email)
	if err != nil {
		return fmt.Errorf("error generating the mail: %w", err)
	}

	message := handlers.mailgun.NewMessage(handlers.config.MailgunSender, subject, emailBodyPlainText, recipient)
	message.SetHtml(emailBodyHTML)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if _, _, err = handlers.mailgun.Send(ctx, message); err != nil {
		return fmt.Errorf("error sending the mail: %w", err)
	}

	return nil
}
