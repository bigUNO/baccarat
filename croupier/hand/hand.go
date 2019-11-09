package hand

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Hander holds the total amount of points and the number of cards dealt.
type Hander interface {
	GetCards() []string
}

// BankerHand holds the Banker cards
type BankerHand struct {
	First  string
	Second string
	Third  string
}

// GetCards returns the Banker cards
func (bh BankerHand) GetCards() []string {
	cards := []string{}

	if bh.First != "" {
		cards = append(cards, bh.First)
	}
	if bh.Second != "" {
		cards = append(cards, bh.Second)
	}
	if bh.Third != "" {
		cards = append(cards, bh.Third)
	}
	return cards
}

// Action calculates the action from the Player's hand.
func (bh BankerHand) Action(ph PlayerHand) (string, error) {
	if bh.First == "" || bh.Second == "" {
		return "", errors.New("Looks ma no hands")
	}

	total := 0

	firstTwoCard := cardValue(bh.First) + cardValue(bh.Second)
	if firstTwoCard < 10 {
		total = firstTwoCard
	}

	if firstTwoCard >= 10 {
		total = firstTwoCard % 10
	}
	switch total {
	case 0, 1, 2, 3, 4, 5:
		return "DRAW", nil
	case 6, 7:
		return "STANDS", nil
	case 8, 9:
		return "NATURAL", nil
	}
	return "", nil
}

// PlayerHand holds the Player cards
type PlayerHand struct {
	First  string
	Second string
	Third  string
}

// GetCards returns the Banker cards
func (ph PlayerHand) GetCards() []string {
	cards := []string{}

	if ph.First != "" {
		cards = append(cards, ph.First)
	}
	if ph.Second != "" {
		cards = append(cards, ph.Second)
	}
	if ph.Third != "" {
		cards = append(cards, ph.Third)
	}
	return cards
}

// Action calculates the action from the Player's hand.
func (ph PlayerHand) Action() (string, error) {
	if ph.First == "" || ph.Second == "" {
		return "", errors.New("Looks ma no hands")
	}

	total := 0

	firstTwoCard := cardValue(ph.First) + cardValue(ph.Second)
	if firstTwoCard < 10 {
		total = firstTwoCard
	}

	if firstTwoCard >= 10 {
		total = firstTwoCard % 10
	}
	switch total {
	case 0, 1, 2, 3, 4, 5:
		return "DRAW", nil
	case 6, 7:
		return "STANDS", nil
	case 8, 9:
		return "NATURAL", nil
	}
	return "", nil
}

// CalcValue sets the value of the hand. Hands are valued according to the
// rightmost digit of the sum of their constituent cards. For example, a hand
// consisting of 2 and 3 is worth 5, but a hand consisting of 6 and 7 is worth
// 3 (i.e., the 3 being the rightmost digit in the combined points total of 13).
// The highest possible hand value in baccarat is, therefore, 9.
func CalcValue(cards []string) int {
	total := 0
	for _, v := range cards {
		r := []rune(v)

		// If card is a number
		if unicode.IsDigit(r[0]) {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Could not convert string to int.")
			}
			//fmt.Printf("Card value: %v\n", v)
			total += n
			//fmt.Printf("Total: %v\n", total)
		}
		// Ace is worth one point
		if strings.Contains("a", strings.ToLower(v)) {
			total++
		}
	}
	total %= 10
	return total
}

// cardValue is like calcValue, but with only one card
func cardValue(card string) int {
	value := 0
	r := []rune(card)

	// If card is a number
	if unicode.IsDigit(r[0]) {
		n, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Could not convert string to int.")
		}
		//fmt.Printf("Card value: %v\n", v)
		value = n
		//fmt.Printf("Total: %v\n", total)
	}
	// Ace is worth one point
	if strings.Contains("a", strings.ToLower(card)) {
		value = 1
	}
	return value
}
