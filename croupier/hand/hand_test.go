package hand

import (
	"testing"
)

func TestCalcValue(t *testing.T) {
	type testHand struct {
		cards []string
		want  int
	}
	tc := []testHand{
		testHand{[]string{"2", "4"}, 6},
		testHand{[]string{"2", "k"}, 2},
		testHand{[]string{"3", "8"}, 1},
		testHand{[]string{"3", "8"}, 1},
	}

	for _, test := range tc {
		got := CalcValue(test.cards)
		if got != test.want {
			t.Errorf("Failed to calculate hand: got %v want %v", got, test.want)
		}
	}
}

func TestCardValue(t *testing.T) {
	type testCards struct {
		Card    string
		Want    int
		Message string
	}

	tc := []testCards{
		testCards{"a", 1, "Wrong value for Ace"},
		testCards{"k", 0, "Wrong value for King"},
		testCards{"2", 2, "Wrong value for Duece"},
	}

	for _, test := range tc {
		got := cardValue(test.Card)
		if got != test.Want {
			t.Errorf("%v: got %v want %v", test.Message, got, test.Want)
		}
	}
}

func TestBankerGetCards(t *testing.T) {
	type testCards struct {
		bh   BankerHand
		want []string
	}

	tc := []testCards{
		testCards{BankerHand{First: "k", Second: "10"}, []string{"k", "10"}},
		testCards{BankerHand{First: "a", Second: "j", Third: "2"}, []string{"a", "j", "2"}},
		testCards{BankerHand{First: "5", Second: "q"}, []string{"5", "q"}},
	}

	for _, test := range tc {
		got := test.bh.GetCards()
		for k := range test.want {
			if got[k] != test.want[k] {
				t.Errorf("got %v wanted \"\"", got[k])
			}
			if got[k] != test.want[k] {
				t.Errorf("got %v want %v", got[k], test.want[k])
			}
		}

	}
}

func TestPlayerGetCards(t *testing.T) {
	type testCards struct {
		ph   PlayerHand
		want []string
	}

	tc := []testCards{
		testCards{PlayerHand{First: "j", Second: "2"}, []string{"j", "2"}},
		testCards{PlayerHand{First: "k", Second: "9", Third: "2"}, []string{"k", "9", "2"}},
		testCards{PlayerHand{First: "2", Second: "7"}, []string{"2", "7"}},
	}

	for _, test := range tc {
		got := test.ph.GetCards()
		for k := range test.want {
			if got[k] != test.want[k] {
				t.Errorf("got %v wanted \"\"", got[k])
			}
			if got[k] != test.want[k] {
				t.Errorf("got %v want %v", got[k], test.want[k])
			}
		}

	}
}

func TestPlayerHandAction(t *testing.T) {
	type testPHs struct {
		pHand   PlayerHand
		Want    string
		Message string
	}

	tc := []testPHs{
		testPHs{PlayerHand{}, "", "Wrong action for empty hands"},
		testPHs{PlayerHand{First: "a", Second: "4"}, "DRAW", "Wrong action for hand valued 5"},
		testPHs{PlayerHand{First: "2", Second: "4"}, "STANDS", "Wrong action for hand valued 6"},
		testPHs{PlayerHand{First: "4", Second: "5"}, "NATURAL", "Wrong action for hand valued 9"},
		testPHs{PlayerHand{First: "5", Second: "9"}, "DRAW", "Wrong action for hand valued 4"},
	}

	for _, test := range tc {
		got, _ := test.pHand.Action()
		if got != test.Want {
			t.Errorf("%v: got %v want %v", test.Message, got, test.Want)
		}
	}
}
