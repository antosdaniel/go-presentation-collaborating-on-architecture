package infra

import (
	"bridge/auction/internal"
)

func New() Repo {
	return logs{Repo: repo{}}
}

type Repo interface {
	Find() internal.Bids
}

type repo struct{}

func (r repo) Find() internal.Bids {
	return nil
}

type logs struct {
	Repo Repo
}

func (l logs) Find() internal.Bids {
	println("log!")
	return l.Repo.Find()
}
