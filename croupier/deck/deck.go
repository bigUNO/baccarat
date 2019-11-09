package deck

/*
"Lifted" from https://github.com/x-color/gocard
*/

import (
	"errors"
	"math/rand"
	"time"
)

const (
	// Number of decks in a shoe
	numOfDecks = 8
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Deck is a set of cards
type Deck []string

// Draw pulls a card from the top of the deck.
func (deck *Deck) Draw() (card string, err error) {
	if len(*deck) == 0 {
		err = errors.New("could not draw, deck is empty")
		return "", err
	}
	card, *deck = (*deck)[0], (*deck)[1:]
	return card, err
}

// Shuffle suffles the deck
func (deck Deck) shuffle() {
	for i := len(deck); i > 0; i-- {
		randIndex := rand.Intn(i)
		deck[i-1], deck[randIndex] = deck[randIndex], deck[i-1]
	}
}

// NewDeck returns a shuffled deck of cards
func NewDeck() (deck Deck) {
	// TODO: Add a cut card to the deck
	// https://www.youtube.com/watch?v=oj6uvmQAtEE&t=1121s
	var rank = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "j", "q", "k"}
	for deckNumber := 1; deckNumber <= numOfDecks; deckNumber++ {
		for suitNumber := 1; suitNumber <= 4; suitNumber++ {
			for _, v := range rank {
				deck = append(deck, v)
			}
		}
	}
	deck.shuffle()
	return deck
}
