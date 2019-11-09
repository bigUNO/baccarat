package croupier

import (
	"fmt"

	"github.com/bigUNO/baccarat/croupier/deck"
	"github.com/bigUNO/baccarat/croupier/hand"
	"github.com/bigUNO/baccarat/croupier/wager"
)

// Croupier is the interface that represents all the activities performed by a
// baccarat dealer.
type Croupier interface {
	// Deal cards to start a new game.
	Deal()

	// Call returns the status of the game.
	Call() HandResult

	// ClearCards resets the current hand in preparation for dealing the next
	// round.
	ClearCards()
}

// Dealer holds the information for managing the deck of cards as well as
// running the table.
type Dealer struct {
	// Array of strings that represents a shuffled deck of cards.
	shoe deck.Deck

	// currentGame represents the current pair of cards on the table
	currentGame Game

	// bets
	bets wager.Bets
}

// NewCroupier returns a dealer to satisfy the Croupier interface.
//
// Handles management of all card functionality. Takes payout values from the
// table (casino).
func NewCroupier() *Dealer {
	d := &Dealer{
		shoe: deck.NewDeck(),
	}
	return d
}

// Call returns a HandResult struct with the score of the hand at the current stage
// in the game.
func (d Dealer) Call() HandResult {
	result := HandResult{
		PlayerHand: hand.CalcValue(d.currentGame.player.GetCards()),
		BankerHand: hand.CalcValue(d.currentGame.banker.GetCards()),
	}

	return result
}

// Deal draws two cards for both the Player and the Banker to start the game.
//
// Cards are woven between the Player and the Banker, starting with the player.
func (d *Dealer) Deal() {
	if d.currentGame.player.First == "" {
		playerCard, err := d.shoe.Draw()
		if err != nil {
			fmt.Println(err)
		}
		d.currentGame.player.First = playerCard
	}

	if d.currentGame.banker.First == "" {
		playerCard, err := d.shoe.Draw()
		if err != nil {
			fmt.Println(err)
		}
		d.currentGame.banker.First = playerCard
	}

	if d.currentGame.player.Second == "" {
		playerCard, err := d.shoe.Draw()
		if err != nil {
			fmt.Println(err)
		}
		d.currentGame.player.Second = playerCard
	}

	if d.currentGame.banker.Second == "" {
		playerCard, err := d.shoe.Draw()
		if err != nil {
			fmt.Println(err)
		}
		d.currentGame.banker.Second = playerCard
	}
}

// DealCardToPlayer deals a card to the PLAYER. This is a separate function
// because I'm too lazy to come up with a more elagant way of combining this
// with the Draw() function.
func (d *Dealer) DealCardToPlayer() error {
	var err error
	if d.currentGame.player.Third != "" {
		return ErrAlreadyThirdCard
	}
	d.currentGame.player.Third, err = d.shoe.Draw()
	if err != nil {
		return err
	}
	return nil
}

// DealCardToBanker deals a card to the BANKER. This is a separate function
// because I'm too lazy to come up with a more elagant way of combining this
// with the Draw() function.
func (d *Dealer) DealCardToBanker() error {
	var err error
	if d.currentGame.banker.Third != "" {
		return ErrAlreadyThirdCard
	}
	d.currentGame.banker.Third, err = d.shoe.Draw()
	if err != nil {
		return err
	}
	return nil
}

// ClearCards removes the current cards from both hands.
func (d *Dealer) ClearCards() {
	d.currentGame.reset()
}

// Game represents the cards in a single hand in a game of Baccarat.
type Game struct {
	player hand.PlayerHand
	banker hand.BankerHand
}

func (g *Game) reset() {
	g.banker = hand.BankerHand{}
	g.player = hand.PlayerHand{}
}

// HandResult contains the totals of the hands for both the PLAYER and the
// BANKER.
type HandResult struct {
	PlayerHand int
	BankerHand int
}

// HandState returns a string that represents the game state.
func (h HandResult) HandState() string {
	var state string

	switch {
	case h.PlayerHand > h.BankerHand:
		state = "player"
	case h.BankerHand > h.PlayerHand:
		state = "banker"

	}
	return state
}
