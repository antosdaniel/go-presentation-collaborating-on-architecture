package setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/internal"
	infra2 "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/internal/infra"
)

func New() (NoSetup, error) {
	userRepo := infra2.UserRepo{}
	templateRepo := infra2.TemplateRepo{}
	smsSender := infra2.SmsSender{}
	return NoSetup{
		API: internal.New(userRepo, templateRepo, smsSender),
	}, nil
}

type NoSetup struct {
	API notification.API
}

func (n NoSetup) PeriodicRetry() {
	// If there are any failed notifications, retry them
}
