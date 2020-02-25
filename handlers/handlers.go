package handlers

import (
	"github.com/VegimagDevs/vegimag-api/storage"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/matcornic/hermes/v2"
)

type Config struct {
	Storage              *storage.Storage
	MailgunDomain        string
	MailgunPrivateAPIKey string
	MailgunSender        string
}

type Handlers struct {
	config *Config

	mailgun *mailgun.MailgunImpl
	hermes  *hermes.Hermes
}

func New(config *Config) *Handlers {
	handlers := &Handlers{
		config: config,
	}

	handlers.mailgun = mailgun.NewMailgun(config.MailgunDomain, config.MailgunPrivateAPIKey)
	handlers.mailgun.SetAPIBase(mailgun.APIBaseEU)

	handlers.hermes = &hermes.Hermes{
		Product: hermes.Product{
			Name:      "Vegimag",
			Link:      "https://vegimag.org/",
			Copyright: "The content of this mail is licensed under the GNU General Public License v3.0.",
		},
	}

	return handlers
}
