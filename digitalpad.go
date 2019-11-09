package baccarat

// DigitalPad represents the digital scoreboard that displays game statistics.
type DigitalPad struct {
	// Number of times that BANKER hand was closest to nine in currrent game.
	banker int

	// Number of times that PLAYER hand was closest to nine in currrent game.
	player int

	// Number of NATURALs in a current game.
	natural int

	// Number of TIEs in current game.
	tie int

	//	previousHands []string
}

// EnterResults saves the reults of the hand after bets have been paid out.
func (d *DigitalPad) EnterResults(result string) {
	switch result {
	case "banker":
		d.banker++
	case "player":
		d.player++
	case "natural":
		d.natural++
	case "tie":
		d.tie++
	}
}
