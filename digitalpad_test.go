package baccarat

import "testing"

func TestEnterResults(t *testing.T) {
	type test struct {
		digitalpad DigitalPad
		results    string
		want       int
	}

	tc := []test{
		test{DigitalPad{}, "", 0},
		test{DigitalPad{}, "banker", 1},
		test{DigitalPad{player: 45}, "player", 46},
		test{DigitalPad{natural: 11}, "natural", 12},
		test{DigitalPad{tie: 3}, "tie", 4},
	}

	for _, test := range tc {
		test.digitalpad.EnterResults(test.results)

		switch test.results {
		case "":
			if test.digitalpad.banker != 0 &&
				test.digitalpad.player != 0 &&
				test.digitalpad.natural != 0 &&
				test.digitalpad.tie != 0 {
				t.Error("saved a results when nothing happen")
			}
		case "banker":
			if test.digitalpad.banker != test.want {
				t.Errorf("got the wrong number of banker wins: %v",
					test.digitalpad.banker)
			}
		case "player":
			if test.digitalpad.player != test.want {
				t.Errorf("got the wrong number of player wins: %v",
					test.digitalpad.player)
			}
		case "natural":
			if test.digitalpad.natural != test.want {
				t.Errorf("got the wrong number of naturals: %v",
					test.digitalpad.natural)
			}
		case "tie":
			if test.digitalpad.tie != test.want {
				t.Errorf("got the wrong number of ties: %v",
					test.digitalpad.tie)
			}
		}
	}
}
