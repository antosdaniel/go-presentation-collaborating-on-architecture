package no

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no/api"
)

func New() (NoSetup, error) {
	return NoSetup{
		API: api.APIImplementation{},
	}, nil
}

type NoSetup struct {
	API api.API
}

func (n NoSetup) PeriodicRetry() {
	// If there are any failed notifications, retry them
}
