package daifugo

import (
	"github.com/fuuki/board/board"
)

// dealPhase returns a phase of deal cards.
func dealPhase() *jPhase {
	prepare := func(_ *jGame) jActionReq {
		apr := &daifugoActionRequest{}
		return apr
	}

	execute := func(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
		// Deal cards
		bp.PrepareNewRound(g.Players(), 0)
		return PlayPhase, bp
	}

	p := board.NewPhase(DealPhase, prepare, execute)
	return p
}
