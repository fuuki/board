package rock_paper_scissors

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/resource"
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
func rockPaperScissorsGame() *game.Game {
	rp := resourceProfile()

	p1 := playPhase()
	g := game.NewGame(PLAY_PHASE, []*game.Phase{p1}, rp)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile() *resource.ResourceProfile {
	rp := resource.NewResourceProfile()
	rp.AddResource(0, resource.NewResource())
	rp.AddResource(1, resource.NewResource())
	return rp
}

// playPhase returns a phase of rock-paper-scissors.
func playPhase() *game.Phase {
	prepare := func(_ *game.Game) *action.ActionProfileDefinition {
		// Define action profile
		apd := profileDef()
		return apd
	}

	execute := func(g *game.Game, ap *action.ActionProfile) game.PhaseName {
		getReward(ap, g.ResourceProfile())
		if isFinished(g.ResourceProfile()) {
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
func getReward(ap *action.ActionProfile, rp *resource.ResourceProfile) {
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

func isFinished(rp *resource.ResourceProfile) bool {
	return rp.Player(0).Point() >= 3 || rp.Player(1).Point() >= 3
}
