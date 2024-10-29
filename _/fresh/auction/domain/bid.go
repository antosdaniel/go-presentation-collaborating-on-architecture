package domain

type Bids []Bid

type Bid struct {
	PlayerID string
	Value    string
}

func (bs Bids) Contract() Bid {
	return bs[len(bs)-1]
}
