package daifugo

import "github.com/fuuki/board/board"

func playPhase() *jPhase {
	p := board.NewPhase(PlayPhase, playPhasePrepare, playPhaseExecute)
	return p
}

func playPhasePrepare(g *jGame) jActionReq {
	// Define action profile
	apr := &daifugoActionRequest{
		currentPlayer: g.BoardProfile().turn.Current(),
	}
	return apr
}

func playPhaseExecute(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
	p := bp.turn.Current()
	a := ap.Player(p)
	if a.Pass {
		bp.turn.Pass()
		if len(bp.turn.Order()) == 1 {
			// TODO: Go to next round.
			return PlayPhase, bp
		}
	} else {
		cards := bp.Player(p).PickMulti((*a).Select)
		bp.PlayArea.AddMulti(cards)
		if isFinished(g, bp) {
			return "", bp
		}
	}
	bp.turn.Next()
	return PlayPhase, bp
}

func isFinished(g *jGame, jp *daifugoBoardProfile) bool {
	for _, p := range g.Players() {
		if len(jp.Player(p).Cards()) == 0 {
			return true
		}
	}
	return false
}
