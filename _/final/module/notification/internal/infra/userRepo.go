package infra

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no"
)

type UserRepo struct{}

func (r UserRepo) Find(id string) no.User {
	return no.User{
		ID:        id,
		FirstName: "Jon",
		LastName:  "Doe",
	}
}
