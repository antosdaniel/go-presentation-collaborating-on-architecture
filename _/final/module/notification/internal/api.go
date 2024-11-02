package internal

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no/internal/infra"
)

func New(
	userRepo infra.UserRepo,
	templateRepo infra.TemplateRepo,
	smsSender infra.SmsSender,
) no.API {
	return ApiWithLogs{
		API: api{
			templateRepo: templateRepo,
			userRepo:     userRepo,
			smsSender:    smsSender,
		},
	}
}

type api struct {
	templateRepo infra.TemplateRepo
	userRepo     infra.UserRepo
	smsSender    infra.SmsSender
}

func (a api) SendNotification(templateID string, userID string) error {
	template := a.templateRepo.Find(templateID)
	user := a.userRepo.Find(userID)

	message := template.FillWithName(user.FullName())

	a.smsSender.Send(user.Phone, message)

	return nil
}
