package deck

import (
	"testing"
)

func TestDeckLenth(t *testing.T) {
	want := 416
	d := NewDeck()
	if len(d) != want {
		t.Errorf("wrong number of cards: got %v want %v", len(d), want)
	}
}
