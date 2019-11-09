package croupier

import (
	"testing"

	"github.com/bigUNO/baccarat/croupier/deck"
	"github.com/bigUNO/baccarat/croupier/hand"
)

func TestNewCroupier(t *testing.T) {
	want := NewCroupier()

	if want == nil {
		t.Errorf("could not get new croupier: %v", want)
	}
}

func TestDeal(t *testing.T) {
	c := NewCroupier()
	c.Deal()

	if c.currentGame.player.First == "" {
		t.Errorf("Player didn't get a first card")
	}

	if c.currentGame.player.Second == "" {
		t.Errorf("Player didn't get a second card")
	}

	if c.currentGame.banker.First == "" {
		t.Errorf("Banker didn't get a first card")
	}

	if c.currentGame.banker.Second == "" {
		t.Errorf("Banker didn't get a second card")
	}
}

func TestCroupierCall(t *testing.T) {
	type testDealer struct {
		dealer Dealer
		want   HandResult
	}

	tc := []testDealer{
		testDealer{Dealer{
			shoe: deck.Deck{},
			currentGame: Game{
				player: hand.PlayerHand{First: "10", Second: "5", Third: "6"},
				banker: hand.BankerHand{First: "a", Second: "q"},
			},
		}, HandResult{PlayerHand: 1, BankerHand: 1}},
		testDealer{Dealer{
			shoe: deck.Deck{},
			currentGame: Game{
				player: hand.PlayerHand{First: "2", Second: "8", Third: "6"},
				banker: hand.BankerHand{First: "q", Second: "q", Third: "5"},
			},
		}, HandResult{PlayerHand: 6, BankerHand: 5}},
	}

	for _, test := range tc {
		got := test.dealer.Call()
		if got.BankerHand != test.want.BankerHand {
			t.Errorf("got banker hand: %v, want %v",
				got.BankerHand, test.want.BankerHand)
		}
		if got.PlayerHand != test.want.PlayerHand {
			t.Errorf("got player hand: %v, want %v",
				got.PlayerHand, test.want.PlayerHand)
		}
	}
}

func TestDealCardToPlayer(t *testing.T) {
	type testDealer struct {
		dealer Dealer
	}

	tc := []testDealer{
		testDealer{Dealer{
			shoe: deck.Deck{"2"},
			currentGame: Game{
				player: hand.PlayerHand{First: "a", Second: "q"},
				banker: hand.BankerHand{},
			},
		}},
		testDealer{Dealer{
			shoe: deck.Deck{},
			currentGame: Game{
				player: hand.PlayerHand{First: "q", Second: "q", Third: "5"},
				banker: hand.BankerHand{},
			},
		}},
	}

	for _, test := range tc {
		err := test.dealer.DealCardToPlayer()
		if err != nil && err != ErrAlreadyThirdCard {
			t.Errorf("failed to draw player card: %v", err)
		}

		if test.dealer.currentGame.player.Third == "" {
			t.Error("Player drew card but it was not assigned\n")
		}
	}
}

func TestDealCardToer(t *testing.T) {
	type testDealer struct {
		dealer Dealer
	}

	tc := []testDealer{
		testDealer{Dealer{
			shoe: deck.Deck{"2"},
			currentGame: Game{
				player: hand.PlayerHand{},
				banker: hand.BankerHand{First: "a", Second: "q"},
			},
		}},
		testDealer{Dealer{
			shoe: deck.Deck{},
			currentGame: Game{
				player: hand.PlayerHand{},
				banker: hand.BankerHand{First: "q", Second: "q", Third: "5"},
			},
		}},
	}

	for _, test := range tc {
		err := test.dealer.DealCardToBanker()
		if err != nil && err != ErrAlreadyThirdCard {
			t.Errorf("failed to draw banker card: %v", err)
		}

		if test.dealer.currentGame.banker.Third == "" {
			t.Error("Player drew card but it was not assigned\n")
		}
	}
}
