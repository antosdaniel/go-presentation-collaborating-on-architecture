package internal

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/internal/infra"
)

func New(
	userRepo infra.UserRepo,
	templateRepo infra.TemplateRepo,
	smsSender infra.SmsSender,
) notification.API {
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
