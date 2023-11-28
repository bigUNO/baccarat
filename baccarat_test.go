package baccarat

import (
	"testing"
)

func TestStart(t *testing.T) {
	testBaccarat, err := New()
	if err != nil {
		t.Error("Could no create new baccarat")
	}
	testBaccarat.Start()
	got := testBaccarat.croupier.Call()

	if got.PlayerHand == 0 && got.BankerHand == 0 {
		t.Errorf("not a good test. could tie with 0:0.")
	}
}
