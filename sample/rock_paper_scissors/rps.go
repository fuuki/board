package rock_paper_scissors

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/player"
	"github.com/fuuki/board/result"
)

//go:generate stringer -type=Hand
type Hand int

const (
	ROCK Hand = iota
	PAPER
	SCISSORS
)

const (
	PLAY_PHASE game.PhaseName = "play"
)

func Play() {
	g := rockPaperScissorsGame()
	g.Play(&action.InteractiveActionInputer{})
}

// rockPaperScissorsGame returns a game of rock-paper-scissors.
func rockPaperScissorsGame() *game.Game[*JankenBoardProfile] {
	rp := resourceProfile()

	p1 := playPhase()
	g := game.NewGame(PLAY_PHASE, []*game.Phase[*JankenBoardProfile]{p1}, rp, resultFn)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile() *JankenBoardProfile {
	rp := NewJankenBoardProfile(2)
	return rp
}

// playPhase returns a phase of rock-paper-scissors.
func playPhase() *game.Phase[*JankenBoardProfile] {
	prepare := func(_ *game.Game[*JankenBoardProfile]) *action.ActionProfileDefinition {
		// Define action profile
		apd := profileDef()
		return apd
	}

	execute := func(g *game.Game[*JankenBoardProfile], ap *action.ActionProfile) game.PhaseName {
		getReward(ap, g.BoardProfile())
		if isFinished(g.BoardProfile()) {
			return ""
		}
		return PLAY_PHASE
	}

	p := game.NewPhase(PLAY_PHASE, prepare, execute)
	return p
}

func profileDef() *action.ActionProfileDefinition {
	def := action.NewActionDefinition()
	def.AddChoice([]action.ActionChoiceOption{ROCK, PAPER, SCISSORS})
	apd := action.NewActionProfileDefinition()
	apd.AddActionDefinition(0, def)
	apd.AddActionDefinition(1, def)
	return apd
}

// getReward returns the result of the game.
func getReward(ap *action.ActionProfile, rp *JankenBoardProfile) {
	aa1, _ := ap.GetAction(0)
	aa2, _ := ap.GetAction(1)
	a1 := aa1.(Hand)
	a2 := aa2.(Hand)

	switch (a1 - a2 + 3) % 3 {
	case 0:
		return
	case 1:
		// Player 0 win
		rp.Player(0).AddPoint(1)
		return
	case 2:
		// Player 1 win
		rp.Player(1).AddPoint(1)
		return
	}
}

func isFinished(jp *JankenBoardProfile) bool {
	for i := 0; i < jp.PlayerNum(); i++ {
		if jp.Player(player.Player(i)).Point() >= 3 {
			return true
		}
	}
	return false
}

func resultFn(g *game.Game[*JankenBoardProfile]) *result.Result {
	r := result.NewResult()
	rank := func(point int) uint {
		if point == 3 {
			return 1
		}
		return 2
	}
	for i := 0; i < g.BoardProfile().PlayerNum(); i++ {
		score := g.BoardProfile().Player(player.Player(i)).Point()
		r.AddPlayerResult(
			result.PlayerResult{
				Player: player.Player(i),
				Score:  score,
				Rank:   rank(score),
			})
	}
	return r
}
