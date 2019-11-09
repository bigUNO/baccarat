package wager

import "testing"

func TestPayoutBanker(t *testing.T) {
	type testBets struct {
		bet  MainBets
		want int
	}

	tc := []testBets{
		testBets{MainBets{Banker: 2}, 4},
		testBets{MainBets{Banker: 10}, 20},
		testBets{MainBets{Banker: 2000}, 4000},
	}

	for _, test := range tc {
		got := test.bet.payoutBanker()
		if got != test.want {
			t.Errorf("got the wrong payout for banker\ngot: %v want %v",
				got, test.want)
		}
	}
}

func TestPayoutPlayer(t *testing.T) {
	type testBets struct {
		bet  MainBets
		want int
	}

	tc := []testBets{
		testBets{MainBets{Player: 10}, 20},
		testBets{MainBets{Player: 25}, 50},
		testBets{MainBets{Player: 2000}, 4000},
	}

	for _, test := range tc {
		got := test.bet.payoutPlayer()
		if got != test.want {
			t.Errorf("got the wrong payout for player\ngot: %v want %v",
				got, test.want)
		}
	}
}

func TestPayoutTie(t *testing.T) {
	type testBets struct {
		bet  MainBets
		want int
	}

	tc := []testBets{
		testBets{MainBets{Tie: 1}, 10},
		testBets{MainBets{Tie: 10}, 100},
		testBets{MainBets{Tie: 25}, 250},
		testBets{MainBets{Tie: 50}, 500},
		testBets{MainBets{Tie: 100}, 1000},
	}

	for _, test := range tc {
		got := test.bet.payoutTie()
		if got != test.want {
			t.Errorf("wrong payout for %v bet on tie\ngot: %v want %v",
				test.bet.Tie, got, test.want)
		}
	}
}

func TestPayoutPair(t *testing.T) {
	type testBets struct {
		bet  SideBets
		want int
	}

	tc := []testBets{
		testBets{SideBets{PairPlayer: 1}, 13},
		testBets{SideBets{PairBanker: 10}, 130},
		testBets{SideBets{PairPlayer: 25}, 325},
		testBets{SideBets{PairBanker: 50}, 650},
		testBets{SideBets{PairPlayer: 100}, 1300},
	}

	for _, test := range tc {
		got := test.bet.payoutPair()
		if got != test.want {
			t.Errorf("wrong payout for side bet on player/banker pair\ngot: %v want %v",
				got, test.want)
		}
	}
}

func TestDragonPayout(t *testing.T) {
	type testBets struct {
		bet  dragon
		want int
	}

	//tc := []testBets{
	//	testBets{dragon{amount: 1, banker: true}, 2},
	//}
}

func TestPlaceMainBets(t *testing.T) {
	type testBet struct {
		mb     MainBets
		player int
		banker int
		tie    int
		want   []int
	}

	tc := []testBet{
		testBet{MainBets{}, 0, 50, 100, []int{0, 50, 100}},
		testBet{MainBets{}, 0, 50, 100, []int{0, 50, 100}},
	}

	for _, test := range tc {
		test.mb.PlaceMainBets(test.player, test.banker, test.tie)
		if test.mb.Player != test.want[0] {
			t.Errorf("got the wrong player main bet: got %v want %v",
				test.mb.Player, test.want[0])
		}

		if test.mb.Banker != test.want[1] {
			t.Errorf("got the wrong banker main bet: got %v want %v",
				test.mb.Banker, test.want[1])
		}

		if test.mb.Tie != test.want[2] {
			t.Errorf("got the wrong tie main bet: got %v want %v",
				test.mb.Tie, test.want[2])
		}
	}
}

func TestPlaceSideBets(t *testing.T) {
	type testBet struct {
		sb     SideBets
		player int
		banker int
		want   []int
	}

	tc := []testBet{
		testBet{SideBets{}, 5, 10, []int{5, 10}},
	}

	for _, test := range tc {
		test.sb.PlaceSideBets(test.player, test.banker)
		if test.sb.PairPlayer != test.want[0] {
			t.Errorf("got the wrong player pair bet: got %v want %v",
				test.sb.PairPlayer, test.want[0])
		}

		if test.sb.PairBanker != test.want[1] {
			t.Errorf("got the wrong banker pair bet: got %v want %v",
				test.sb.PairBanker, test.want[1])
		}
	}
}

func TestPlaceDragonBet(t *testing.T) {
	type testBet struct {
		sb     SideBets
		side   string
		amount int
		want   int
	}
	tc := []testBet{
		testBet{SideBets{}, "player", 10, 10},
		testBet{SideBets{}, "banker", 15, 15},
	}

	for _, test := range tc {
		test.sb.PlaceDragonBet(test.side, test.amount)
		switch test.side {
		case "player":
			if !test.sb.Dragon.player {
				t.Errorf("no dragon bet placed on player")
			}
			if test.sb.Dragon.amount != test.want {
				t.Errorf("wrong dragon bet placed on player\n got: %v want %v",
					test.sb.Dragon.amount, test.want)
			}
		case "banker":
			if !test.sb.Dragon.banker {
				t.Errorf("no dragon bet placed on banker")
			}
			if test.sb.Dragon.amount != test.want {
				t.Errorf("wrong dragon bet placed on banker\n got: %v want %v",
					test.sb.Dragon.amount, test.want)
			}
		}
	}
}
