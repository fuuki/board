package daifugo

import (
	"github.com/fuuki/board/board"
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
	g := daifugoGame()
	inputer := &board.InteractiveActionInputer[*daifugoPlayerAction]{}
	g.Play(inputer)
}

// daifugoGame returns a game of rock-paper-scissors.
func daifugoGame() *jGame {
	rp := resourceProfile()

	p1 := dealPhase()
	p2 := playPhase()
	g := board.NewGame(2, DealPhase, []*jPhase{p1, p2}, rp, resultFn)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile() *daifugoBoardProfile {
	rp := NewDaifugoBoardProfile(2)
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
