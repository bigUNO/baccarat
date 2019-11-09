package wager

const (
	// Main bets payout 1:1
	mainBetPayout = 1

	// Ties payout 9:1
	mainPayoutTie = 9

	// pairs payout 12:1
	sideBetPair = 12
)

// Bets holds the current bets. The first wager in the Pair slice goes to the
// player.
type Bets struct {
	Main MainBets
	Side SideBets
}

// MainBets defines the three main bets in Baccarat
type MainBets struct {
	Player int
	Banker int
	Tie    int
}

// PlaceMainBets records wager for primary bets
func (m *MainBets) PlaceMainBets(playerBetAmount, bankBetAmount, tieBetAmount int) {
	m.Player = playerBetAmount
	m.Banker = bankBetAmount
	m.Tie = tieBetAmount
}

func (m MainBets) payoutPlayer() int {
	return (m.Player*mainBetPayout)/1 + m.Player
}

func (m MainBets) payoutBanker() int {
	return (m.Banker*mainBetPayout)/1 + m.Banker
}

func (m MainBets) payoutTie() int {
	return (m.Tie*mainPayoutTie)/1 + m.Tie
}

// SideBets defines the three main bets in Baccarat
type SideBets struct {
	PairPlayer int
	PairBanker int
	Dragon     dragon
}

// PlaceSideBets records wagers for side bets
func (s *SideBets) PlaceSideBets(pairP int, pairB int) {
	s.PairPlayer = pairP
	s.PairBanker = pairB
}

// PlaceDragonBet records wager for the dragon
func (s *SideBets) PlaceDragonBet(side string, amount int) {
	switch side {
	case "player":
		s.Dragon.player = true
		s.Dragon.amount = amount
	case "banker":
		s.Dragon.banker = true
		s.Dragon.amount = amount
	}
}

func (s SideBets) payoutPair() int {
	bankerBet := (s.PairBanker*sideBetPair)/1 + s.PairBanker
	playerBet := (s.PairPlayer*sideBetPair)/1 + s.PairPlayer
	return bankerBet + playerBet
}

type dragon struct {
	amount         int
	player, banker bool
}

func (dragon dragon) Payout() {

}
