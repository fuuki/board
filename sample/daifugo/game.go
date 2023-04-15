package daifugo

import (
	"github.com/fuuki/board/board"
)

const (
	DealPhase board.PhaseName = "deal"
	PlayPhase board.PhaseName = "play"
)

type jGame = board.Game[*daifugoBoardProfile, *daifugoActionProfile]
type jPhase = board.Phase[*daifugoBoardProfile, *daifugoActionProfile]
type jAction = board.ActionProfile[*daifugoActionProfile]
type jActionReq = board.ActionRequest[*daifugoActionProfile]

func Play() {
	g := daifugoGame()
	inputer := &board.InteractiveActionInputer[*daifugoActionProfile]{}
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

// dealPhase returns a phase of deal cards.
func dealPhase() *jPhase {
	prepare := func(_ *jGame) jActionReq {
		apr := &daifugoActionRequest{}
		return apr
	}

	execute := func(g *jGame, bp *daifugoBoardProfile, ap *jAction) (board.PhaseName, *daifugoBoardProfile) {
		return PlayPhase, bp
	}

	p := board.NewPhase(DealPhase, prepare, execute)
	return p
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
