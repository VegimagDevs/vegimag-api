package handlers

import (
	"github.com/VegimagDevs/vegimag-api/storage"
)

type Config struct {
	Storage *storage.Storage
}

type Handlers struct {
	config *Config
}

func New(config *Config) *Handlers {
	return &Handlers{
		config: config,
	}
}
