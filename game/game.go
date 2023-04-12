package game

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/resource"
)

type CurrentPhase struct {
	phaseName PhaseName
	apd       *action.ActionProfileDefinition
}

type Game struct {
	initialPhase PhaseName
	phaseMap     []*Phase
	rp           *resource.ResourceProfile
	current      *CurrentPhase
}

func NewGame(
	initialPhase PhaseName,
	phases []*Phase,
	rp *resource.ResourceProfile,
) *Game {
	return &Game{
		initialPhase: initialPhase,
		phaseMap:     phases,
		rp:           rp,
	}
}

func (g *Game) Play(inputer action.ActionInputer) {
	var ap *action.ActionProfile
	for {
		cnt, apd := g.Next(ap)
		if !cnt {
			break
		}
		ap = inputer.Input(apd)
	}
}

// Start returns the initial action profile definition.
func (g *Game) Start() (bool, *action.ActionProfileDefinition) {
	return g.Next(nil)
}

// Next returns the next action profile definition.
// bool is true if the game continues.
func (g *Game) Next(ap *action.ActionProfile) (bool, *action.ActionProfileDefinition) {
	var next PhaseName
	if g.current == nil {
		next = g.initialPhase
	} else {
		phase := g.getPhase(g.current.phaseName)
		next = phase.execute(g, ap)
	}

	g.rp.Show()
	if next == "" {
		return false, nil
	}

	nextPhase := g.getPhase(next)
	apd := nextPhase.prepare(g)
	g.current = &CurrentPhase{
		phaseName: next,
		apd:       apd,
	}
	return true, apd
}

func (g *Game) getPhase(phaseName PhaseName) *Phase {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

func (g *Game) ResourceProfile() *resource.ResourceProfile {
	return g.rp
}
