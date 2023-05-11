package mech

import "github.com/fuuki/board/logic"

// Mech is a game mechanism.
type Mech[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config] struct {
	initialPhase           logic.PhaseName
	phaseMap               []logic.Phase[AD, BP, CF]
	boardProfileDefinition logic.BoardProfileDefinition[BP]
}

// NewMech returns a new game mechanism.
func NewMech[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	initialPhase logic.PhaseName,
	phases []logic.Phase[AD, BP, CF],
	bpd logic.BoardProfileDefinition[BP],
) *Mech[AD, BP, CF] {
	g := &Mech[AD, BP, CF]{
		initialPhase:           initialPhase,
		phaseMap:               phases,
		boardProfileDefinition: bpd,
	}
	return g
}

// GetPhase returns the phase.
// It returns nil if the phase is not found.
func (mc *Mech[AD, BP, CF]) GetPhase(phaseName logic.PhaseName) logic.Phase[AD, BP, CF] {
	for _, p := range mc.phaseMap {
		if p.Name() == phaseName {
			return p
		}
	}
	return nil
}

// GetInitialPhase returns the initial phase.
func (mc *Mech[AD, BP, CF]) GetInitialPhase() logic.Phase[AD, BP, CF] {
	return mc.GetPhase(mc.initialPhase)
}

// GetBoardProfileDefinition returns the board profile definition.
func (mc *Mech[AD, BP, CF]) GetBoardProfileDefinition() logic.BoardProfileDefinition[BP] {
	return mc.boardProfileDefinition
}
