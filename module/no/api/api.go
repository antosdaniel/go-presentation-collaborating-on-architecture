package api

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no/infra"
)

type API interface {
	SendNotification(templateID string, userID string) error
}

type APIImplementation struct {
	templateRepo infra.TemplateRepo
	userRepo     infra.UserRepo
	smsSender    infra.SmsSender
}

func (a APIImplementation) SendNotification(templateID string, userID string) error {
	template := a.templateRepo.Find(templateID)
	user := a.userRepo.Find(userID)

	message := template.FillWithName(user.FirstName)

	a.smsSender.Send(user.Phone, message)

	return nil
}
