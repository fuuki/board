package burst

import (
	"github.com/fuuki/board"
	"github.com/fuuki/board/logic"
)

const (
	DealPhase logic.PhaseName = "deal"
	PlayPhase logic.PhaseName = "play"
)

type jTable = board.Table[*burstPlayerAction, *burstBoardProfile, *burstConfig]
type jPhase = logic.Phase[*burstPlayerAction, *burstBoardProfile, *burstConfig]
type jAction = logic.ActionProfile[*burstPlayerAction]
type jActionReq = logic.ActionRequest[*burstPlayerAction]

// burstGame returns a game of rock-paper-scissors.
func burstGame(totalPlayer uint) (*jTable, <-chan int) {

	var bpd logic.BoardProfileDefinition[*burstBoardProfile] = &burstBoardProfileDefinition{
		totalPlayer: totalPlayer,
	}
	g := board.NewGame(
		DealPhase,
		[]jPhase{&dealPhase{}, &playPhase{}},
		bpd,
	)
	config := &burstConfig{
		totalPlayer: totalPlayer,
	}
	t, c := g.NewTable(config)
	return t, c
}
