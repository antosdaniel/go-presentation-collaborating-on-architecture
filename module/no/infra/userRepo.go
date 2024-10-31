package infra

import "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no/domain"

type UserRepo struct{}

func (r UserRepo) Find(id string) domain.User {
	return domain.User{
		ID:        id,
		FirstName: "Jon",
		LastName:  "Doe",
	}
}