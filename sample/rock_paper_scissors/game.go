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

func Game() {
	g := RockPaperScissorsGame()
	g.Play(&action.InteractiveActionInputer{})
}

func RockPaperScissorsGame() *game.Game {
	// Define action profile
	def := action.NewActionDefinition()
	def.AddChoice([]action.ActionChoiceOption{ROCK, PAPER, SCISSORS})
	apd := action.NewActionProfileDefinition()
	apd.AddActionDefinition(0, def)
	apd.AddActionDefinition(1, def)

	// Define resource profile
	rp := resource.NewResourceProfile()
	rp.AddResource(0, resource.NewResource())
	rp.AddResource(1, resource.NewResource())

	g := game.NewGame(apd, rp, getResult)
	return g
}

func getResult(ap *action.ActionProfile, rp *resource.ResourceProfile) {
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
