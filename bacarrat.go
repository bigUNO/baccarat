package baccarat

import (
	"fmt"
	"time"

	"github.com/bigUNO/baccarat/croupier"
)

// Baccarat holds information about the hand.
//
// Represents all the acticities that would take place at a table in a casino.
type Baccarat struct {
	croupier croupier.Croupier
	ID       int64
	// State I can't remember what I wanted to do with this field.
	State      string
	digitalPad DigitalPad
}

// New retuns a new game of bacarrat
func New() (*Baccarat, error) {
	baccarat := &Baccarat{
		croupier:   croupier.NewCroupier(),
		ID:         time.Now().UnixNano(),
		digitalPad: DigitalPad{},
	}
	return baccarat, nil
}

// Start deals a hand of barrarat.
func (b *Baccarat) Start() {
	// TODO: Collect bets
	b.croupier.Deal()
}

// CloseHand closes the current hand
func (b *Baccarat) CloseHand() {
	// scoreMessage := fmt.Sprintf("PLAYER shows %v, BANKER has %v",
	// PhScore, BhScore)
	// record results
	results := b.croupier.Call()

	fmt.Printf("Game goes to %s; SCORE: %v - %v\n", results.HandState(),
		results.PlayerHand, results.BankerHand)

	//b.digitalPad.EnterResults(results.HandState())
	b.croupier.ClearCards()
}

// ShowDigitalPad prints history of hands dealt.
func (b *Baccarat) ShowDigitalPad() {
	fmt.Print(b.digitalPad)
}
