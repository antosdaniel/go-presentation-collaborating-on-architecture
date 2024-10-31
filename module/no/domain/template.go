package domain

import "fmt"

type Template struct {
	ID   string
	Name string
	Path string
}

func (t Template) FillWithName(fullName string) string {
	return fmt.Sprintf("<h1>Hi %s!</h1>", fullName)
}
