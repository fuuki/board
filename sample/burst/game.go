package burst

import (
	"github.com/fuuki/board"
)

const (
	DealPhase board.PhaseName = "deal"
	PlayPhase board.PhaseName = "play"
)

type jGame = board.Game[*burstBoardProfile, *burstPlayerAction]
type jPhase = board.Phase[*burstBoardProfile, *burstPlayerAction]
type jAction = board.ActionProfile[*burstPlayerAction]
type jActionReq = board.ActionRequest[*burstPlayerAction]

// burstGame returns a game of rock-paper-scissors.
func burstGame(totalPlayer uint) (*jGame, <-chan board.PeriodCount) {
	p1 := dealPhase()
	p2 := playPhase()

	c := make(chan board.PeriodCount)

	var bpd board.BoardProfileDefinition[*burstBoardProfile] = &burstBoardProfileDefinition{
		totalPlayer: totalPlayer,
	}
	g := board.NewGame(
		totalPlayer,
		DealPhase,
		[]*jPhase{p1, p2},
		bpd,
		board.PhaseChangeChan[*burstBoardProfile, *burstPlayerAction](c),
	)
	return g, c
}
