package game

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/resource"
)

type Game struct {
	apd       *action.ActionProfileDefinition
	rp        *resource.ResourceProfile
	getResult func(*action.ActionProfile, *resource.ResourceProfile)
}

func NewGame(
	apd *action.ActionProfileDefinition,
	rp *resource.ResourceProfile,
	getResult func(*action.ActionProfile, *resource.ResourceProfile),
) *Game {
	return &Game{
		apd:       apd,
		rp:        rp,
		getResult: getResult,
	}
}

func (g *Game) Play(
	inputer action.ActionInputer,
) {
	ap := inputer.Input(g.apd)
	g.getResult(ap, g.rp)
	g.rp.Show()
}

func (g *Game) ActionProfileDefinition() *action.ActionProfileDefinition {
	return g.apd
}

func (g *Game) ResourceProfile() *resource.ResourceProfile {
	return g.rp
}
