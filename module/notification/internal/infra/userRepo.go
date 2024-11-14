package infra

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
)

type UserRepo struct{}

func (r UserRepo) Find(id string) notification.User {
	return notification.User{
		ID:        id,
		FirstName: "Jon",
		LastName:  "Doe",
		Phone:     "+48 123 456 789",
	}
}
