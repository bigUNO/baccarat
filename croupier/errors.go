package croupier

import "errors"

var (
	// ErrAlreadyThirdCard is returned when the PLAYER or BANKER already has
	// a third card.
	ErrAlreadyThirdCard = errors.New("Already has a third card")
)
