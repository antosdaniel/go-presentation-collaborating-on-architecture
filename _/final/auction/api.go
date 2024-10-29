package auction

import (
	"bridge/auction/internal"
	"bridge/auction/internal/infra"
)

type API struct{}

func (a API) FindContract() Contract {
	return infra.New().Find().Contract()
}

type Contract = internal.Contract
