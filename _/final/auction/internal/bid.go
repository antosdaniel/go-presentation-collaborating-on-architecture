package internal

type Bids []Bid

type Bid struct {
	PlayerID string
	Value    string
}

func (bs Bids) Contract() Contract {
	return Contract(bs[len(bs)-1])
}

type Contract Bid
