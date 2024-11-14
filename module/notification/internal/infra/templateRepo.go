package infra

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
)

type TemplateRepo struct{}

func (r TemplateRepo) Find(id string) notification.Template {
	return notification.Template{
		ID: id,
	}
}
