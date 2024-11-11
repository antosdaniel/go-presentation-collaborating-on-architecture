package no_setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/api"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/infra"
)

func NewSetup() (NoSetup, error) {
	return NoSetup{
		API: api.APIImplementation{
			TemplateRepo: infra.TemplateRepo{},
			UserRepo:     infra.UserRepo{},
			SmsSender:    infra.SmsSender{},
		},
	}, nil
}

type NoSetup struct {
	API api.API
}

func (n NoSetup) PeriodicRetry() {
	// If there are any failed notifications, retry them
}
