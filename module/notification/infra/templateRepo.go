package infra

import "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/domain"

type TemplateRepo struct{}

func (r TemplateRepo) Find(id string) domain.Template {
	return domain.Template{
		ID: id,
	}
}
