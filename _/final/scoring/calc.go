package scoring

import (
	"bridge/auction"
)

func init() {
	contract := auction.API{}.FindContract()

	tricks := 10

	CalculatePoints(contract, tricks)
}

func CalculatePoints(contract auction.Contract, tricks int) int {
	return 0
}
