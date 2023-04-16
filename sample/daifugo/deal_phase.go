package daifugo

import (
	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

// dealPhase returns a phase of deal cards.
func dealPhase() *jPhase {
	prepare := func(_ *jGame) jActionReq {
		apr := &daifugoActionRequest{}
		return apr
	}

	execute := func(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
		// Deal cards
		dealCards(g, bp)
		return PlayPhase, bp
	}

	p := board.NewPhase(DealPhase, prepare, execute)
	return p
}

// dealCards deals cards to players.
func dealCards(g *jGame, bp *daifugoBoardProfile) {
	// Shuffle cards
	cards := shortDeck

	// Deal cards
	c := len(cards) / len(g.Players())
	for i, p := range g.Players() {
		bp.playerHands[p] = resource.NewCardLine(cards[i*c : (i+1)*c])
	}
	bp.PlayArea = resource.NewCardLine([]*Card{})
}
