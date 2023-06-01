package burst

import (
	"github.com/fuuki/board"
	"github.com/fuuki/board/logic"
)

const (
	DealPhase logic.PhaseName = "deal"
	PlayPhase logic.PhaseName = "play"
)

type bTable = board.Table[*burstPlayerAction, *burstBoardProfile, *burstConfig]
type bPhase = logic.Phase[*burstPlayerAction, *burstBoardProfile, *burstConfig]
type bAction = logic.ActionProfile[*burstPlayerAction]
type bActionReq = logic.ActionRequest[*burstPlayerAction]

// burstGame returns a game of rock-paper-scissors.
func burstGame(totalPlayer uint) (*bTable, <-chan *board.Event) {

	var bpd logic.BoardProfileDefinition[*burstBoardProfile] = &burstBoardProfileDefinition{
		totalPlayer: totalPlayer,
	}
	g := board.NewGame(
		DealPhase,
		[]bPhase{&dealPhase{}, &playPhase{}},
		bpd,
	)
	config := &burstConfig{
		totalPlayer: totalPlayer,
	}
	t, c := g.NewTable(config)
	return t, c
}
