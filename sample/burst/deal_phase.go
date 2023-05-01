package burst

import (
	"github.com/fuuki/board"
	"github.com/fuuki/board/resource"
)

// dealPhase returns a phase of deal cards.
func dealPhase() *jPhase {
	prepare := func(st *board.Status, bp *burstBoardProfile) (*jActionReq, *burstBoardProfile) {
		apr := board.NewActionRequest[*burstPlayerAction](st.TotalPlayer())
		return apr, bp
	}

	execute := func(st *board.Status, bp *burstBoardProfile, ap *jAction) (board.PhaseName, *burstBoardProfile) {
		bp.dealCards(st.TotalPlayer())
		return PlayPhase, bp
	}

	p := board.NewPhase(DealPhase, prepare, execute)
	return p
}

// dealCards prepares a new round.
func (bp *burstBoardProfile) dealCards(
	totalPlayer uint,
) {
	// dealCards deals cards to players.
	// Shuffle cards
	cards := newDeck()

	// Deal cards
	for p := 0; p < int(totalPlayer); p++ {
		bp.PlayerHands[p] = resource.NewCardLine(cards[p*3 : p*3+3])
	}
	bp.Deck = resource.NewCardLine(cards[totalPlayer*3:])
}
