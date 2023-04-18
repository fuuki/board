package daifugo

import (
	"github.com/fuuki/board/board"
	"github.com/fuuki/board/host"
)

const (
	DealPhase board.PhaseName = "deal"
	PlayPhase board.PhaseName = "play"
)

type jGame = board.Game[*daifugoBoardProfile, *daifugoPlayerAction]
type jPhase = board.Phase[*daifugoBoardProfile, *daifugoPlayerAction]
type jAction = board.ActionProfile[*daifugoPlayerAction]
type jActionReq = board.ActionRequest[*daifugoPlayerAction]

func Play() {
	g := daifugoGame(2)
	h := host.NewTerminalHost(g)
	h.Play()
}

// daifugoGame returns a game of rock-paper-scissors.
func daifugoGame(totalPlayer uint) *jGame {
	rp := resourceProfile(totalPlayer)

	p1 := dealPhase()
	p2 := playPhase()
	g := board.NewGame(totalPlayer, DealPhase, []*jPhase{p1, p2}, rp, resultFn)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile(totalPlayer uint) *daifugoBoardProfile {
	rp := NewDaifugoBoardProfile(totalPlayer)
	return rp
}

func resultFn(g *jGame) *board.Result {
	r := board.NewResult()
	rank := func(score int) uint {
		if score == 0 {
			return 1
		}
		return 2
	}
	for _, p := range g.Players() {
		score := -len(g.BoardProfile().Player(p).Cards())
		r.AddPlayerResult(
			board.PlayerResult{
				Player: p,
				Score:  score,
				Rank:   rank(score),
			})
	}
	return r
}
