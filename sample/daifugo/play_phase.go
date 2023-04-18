package daifugo

import (
	"errors"

	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

var ErrNoCardSelected = errors.New("no card selected")

func playPhase() *jPhase {
	p := board.NewPhase(PlayPhase, playPhasePrepare, playPhaseExecute)
	return p
}

func playPhasePrepare(g *jGame) *jActionReq {
	// Define action profile
	apr := board.NewActionRequest[*daifugoPlayerAction](g.TotalPlayer())
	apr.RegisterValidator(g.BoardProfile().turn.Current(), func(dpa *daifugoPlayerAction) error {
		if dpa.Pass {
			return nil
		}
		if len(dpa.Select) == 0 {
			return ErrNoCardSelected
		}
		return nil
	})
	return apr
}

func playPhaseExecute(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
	p := bp.turn.Current()
	a := ap.Player(p)
	if a.Pass {
		bp.turn.Pass()
		if len(bp.turn.Order()) == 1 {
			bp.PrepareNewSequence(g.Players(), 0)
			return PlayPhase, bp
		}
	} else {
		cards := bp.Player(p).PickMulti((*a).Select)
		bp.PlayArea = resource.NewCardLine(cards)
		if isFinished(g, bp) {
			return "", bp
		}
		bp.turn.Next()
	}
	return PlayPhase, bp
}

func isFinished(g *jGame, bp *daifugoBoardProfile) bool {
	for _, p := range g.Players() {
		if len(bp.Player(p).Cards()) == 0 {
			return true
		}
	}
	return false
}
