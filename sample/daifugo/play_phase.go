package daifugo

import "github.com/fuuki/board/board"

func playPhase() *jPhase {
	p := board.NewPhase(PlayPhase, playPhasePrepare, playPhaseExecute)
	return p
}

func playPhasePrepare(g *jGame) *jActionReq {
	_ = g.BoardProfile().turn.Next()
	// Define action profile
	apr := &jActionReq{}
	return apr
}

func playPhaseExecute(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
	p := bp.turn.Current()
	a := ap.Player(p)
	cards := bp.Player(p).PickMulti((**a).Select)
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