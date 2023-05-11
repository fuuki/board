package board

import (
	"github.com/fuuki/board/internal/mech"
	"github.com/fuuki/board/logic"
)

type Game[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config] struct {
	// definition of games.
	mech *mech.Mech[AD, BP, CF]
}

// NewGame returns a new game.
func NewGame[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	initialPhase logic.PhaseName,
	phases []logic.Phase[AD, BP, CF],
	bpd logic.BoardProfileDefinition[BP],
) *Game[AD, BP, CF] {
	mech := mech.NewMech(initialPhase, phases, bpd)
	g := &Game[AD, BP, CF]{
		mech: mech,
	}
	return g
}

// NewTable returns a new table.
func (g *Game[AD, BP, CF]) NewTable(config CF) (*Table[AD, BP, CF], <-chan int) {
	phaseChangeChan := make(chan int)
	table := &Table[AD, BP, CF]{
		mech:            g.mech,
		config:          config,
		phaseChangeChan: phaseChangeChan,
	}
	return table, phaseChangeChan
}
