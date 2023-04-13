package board

import (
	"fmt"
)

type CurrentPhase struct {
	phaseName PhaseName
}

type Game[BP BoardProfile, AP PlayerActionDefinition] struct {
	// definition of games.
	initialPhase PhaseName
	phaseMap     []*Phase[BP, AP]
	boardProfile BP
	resultFn     func(*Game[BP, AP]) *Result

	// dynamic information of games.
	current *CurrentPhase
}

func NewGame[BP BoardProfile, AP PlayerActionDefinition](
	initialPhase PhaseName,
	phases []*Phase[BP, AP],
	boardProfile BP,
	resultFn func(*Game[BP, AP]) *Result,
) *Game[BP, AP] {
	return &Game[BP, AP]{
		initialPhase: initialPhase,
		phaseMap:     phases,
		boardProfile: boardProfile,
		resultFn:     resultFn,
	}
}

func (g *Game[BP, AP]) Play(inputer ActionInputer[AP]) {
	var ap *ActionProfile[AP]
	for {
		cnt, apr := g.Next(ap)
		if !cnt {
			break
		}
		ap = inputer.Input(apr)
	}
	result := g.resultFn(g)
	fmt.Printf("%+v", result)
}

// Start returns the initial action profile definition.
// bool is true if the game continues.
func (g *Game[BP, AP]) Start() (bool, *ActionRequest[AP]) {
	return g.Next(nil)
}

// Next returns the next action profile definition.
// bool is true if the game continues.
func (g *Game[BP, AP]) Next(ap *ActionProfile[AP]) (bool, *ActionRequest[AP]) {
	var next PhaseName
	if g.current == nil {
		next = g.initialPhase
	} else {
		phase := g.getPhase(g.current.phaseName)
		next = phase.execute(g, ap)
	}

	fmt.Printf("== BoardProfile ==\n%s\n", g.BoardProfile().Show())

	if next == "" {
		return false, nil
	}

	nextPhase := g.getPhase(next)
	apr := nextPhase.prepare(g)
	g.current = &CurrentPhase{
		phaseName: next,
	}
	return true, apr
}

func (g *Game[BP, AP]) getPhase(phaseName PhaseName) *Phase[BP, AP] {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

func (g *Game[BP, AP]) BoardProfile() BP {
	return g.boardProfile
}
