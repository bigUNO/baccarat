package baccarat

// Bankroll defines purse
type Bankroll struct {
	amount int
}

// NewBankroll sets the initial bankroll to
func NewBankroll() *Bankroll {
	purse := &Bankroll{
		amount: 2500,
	}

	return purse
}
