package infra

import "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification/domain"

type UserRepo struct{}

func (r UserRepo) Find(id string) domain.User {
	return domain.User{
		ID:        id,
		FirstName: "Jon",
		LastName:  "Doe",
		Phone:     "+48 123 456 789",
	}
}
