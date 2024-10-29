package scoring

import (
	"bridge/auction/internal/domain"
)

func init() {
	bids := domain.Bids{{}, {}, {}}
	contract := bids.Contract()

	tricks := 10

	CalculatePoints(contract, tricks)
}

func CalculatePoints(contract domain.Bid, tricks int) int {
	return 0
}
