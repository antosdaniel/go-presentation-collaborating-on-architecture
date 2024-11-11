package api

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/domain"
)

type API interface {
	SendNotification(templateID string, userID string) error
}

type APIImplementation struct {
	TemplateRepo TemplateRepo
	UserRepo     UserRepo
	SmsSender    SmsSender
}

func (a APIImplementation) SendNotification(templateID string, userID string) error {
	template := a.TemplateRepo.Find(templateID)
	user := a.UserRepo.Find(userID)

	message := template.FillWithName(user.FirstName)

	a.SmsSender.Send(user.Phone, message)

	return nil
}

type UserRepo interface {
	Find(id string) domain.User
}

type TemplateRepo interface {
	Find(id string) domain.Template
}

type SmsSender interface {
	Send(phone string, message string)
}
