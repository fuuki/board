package game

import (
	"fmt"

	"github.com/fuuki/board/action"
	"github.com/fuuki/board/board"
	"github.com/fuuki/board/result"
)

type CurrentPhase struct {
	phaseName PhaseName
	apd       *action.ActionProfileDefinition
}

type Game[BP board.BoardProfile] struct {
	// definition of games.
	initialPhase PhaseName
	phaseMap     []*Phase[BP]
	boardProfile BP
	resultFn     func(*Game[BP]) *result.Result

	// dynamic information of games.
	current *CurrentPhase
}

func NewGame[BP board.BoardProfile](
	initialPhase PhaseName,
	phases []*Phase[BP],
	boardProfile BP,
	resultFn func(*Game[BP]) *result.Result,
) *Game[BP] {
	return &Game[BP]{
		initialPhase: initialPhase,
		phaseMap:     phases,
		boardProfile: boardProfile,
		resultFn:     resultFn,
	}
}

func (g *Game[BP]) Play(inputer action.ActionInputer) {
	var ap *action.ActionProfile
	for {
		cnt, apd := g.Next(ap)
		if !cnt {
			break
		}
		ap = inputer.Input(apd)
	}
	result := g.resultFn(g)
	fmt.Printf("%+v", result)
}

// Start returns the initial action profile definition.
func (g *Game[BP]) Start() (bool, *action.ActionProfileDefinition) {
	return g.Next(nil)
}

// Next returns the next action profile definition.
// bool is true if the game continues.
func (g *Game[BP]) Next(ap *action.ActionProfile) (bool, *action.ActionProfileDefinition) {
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
	apd := nextPhase.prepare(g)
	g.current = &CurrentPhase{
		phaseName: next,
		apd:       apd,
	}
	return true, apd
}

func (g *Game[BP]) getPhase(phaseName PhaseName) *Phase[BP] {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

func (g *Game[BP]) BoardProfile() BP {
	return g.boardProfile
}
