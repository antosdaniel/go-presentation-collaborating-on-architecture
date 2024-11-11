package domain

import "fmt"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
}

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
