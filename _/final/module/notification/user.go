package notification

import "fmt"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
}

type FullName string

func (u User) FullName() FullName {
	return FullName(fmt.Sprintf("%s %s", u.FirstName, u.LastName))
}
