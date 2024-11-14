package no_setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
	infra2 "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/internal/infra"
)

func NewSetup() (NoSetup, error) {
	return NoSetup{
		API: notification.New(infra2.TemplateRepo{}, infra2.UserRepo{}, infra2.SmsSender{}),
	}, nil
}

type NoSetup struct {
	API notification.API
}

func (n NoSetup) PeriodicRetry() {
	// If there are any failed notifications, retry them
}
