package infra

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no"
)

type TemplateRepo struct{}

func (r TemplateRepo) Find(id string) no.Template {
	return no.Template{
		ID: id,
	}
}
