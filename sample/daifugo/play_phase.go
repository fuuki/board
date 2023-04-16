package daifugo

import "github.com/fuuki/board/board"

func playPhase() *jPhase {
	p := board.NewPhase(PlayPhase, playPhasePrepare, playPhaseExecute)
	return p
}

func playPhasePrepare(g *jGame) jActionReq {
	// Define action profile
	apr := &daifugoActionRequest{}
	return apr
}

func playPhaseExecute(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
	p := bp.turn.Current()
	a := ap.Player(p)
	if a.Pass {
		bp.passMarker[p] = true
		// Find next player
		var next board.Player
		for {
			next = bp.turn.Next()
			if !bp.passMarker[next] {
				break
			}
		}
		// All players pass
		if next == bp.turn.Current() {
			return DealPhase, bp
		}
		return PlayPhase, bp
	}
	cards := bp.Player(p).PickMulti((*a).Select)
	bp.PlayArea.AddMulti(cards)
	if isFinished(g, bp) {
		return "", bp
	}
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
