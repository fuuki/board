package board_test

import (
	"testing"

	"github.com/fuuki/board"
	"github.com/fuuki/board/logic"
)

const (
	samplePhaseNamePlay logic.PhaseName = "play"
)

type sampleBoardProfile struct {
	points []int
	turn   int
}

type sampleBoardProfileDefinition struct {
	totalPlayer uint
}

// New returns a new board profile.
func (d *sampleBoardProfileDefinition) New() *sampleBoardProfile {
	points := make([]int, d.totalPlayer)
	return &sampleBoardProfile{
		turn:   0,
		points: points,
	}
}

// Clone returns a clone of the board profile.
func (d *sampleBoardProfileDefinition) Clone(bp *sampleBoardProfile) *sampleBoardProfile {
	return &sampleBoardProfile{
		points: append([]int{}, bp.points...),
		turn:   bp.turn,
	}
}

var _ logic.BoardProfileDefinition[*sampleBoardProfile] = &sampleBoardProfileDefinition{}

type sampleActionDefinition struct {
}

type sampleConfig struct{}

func (c *sampleConfig) TotalPlayer() uint {
	return 3
}

type sGame = board.Game[*sampleActionDefinition, *sampleBoardProfile, *sampleConfig]
type sPhase = logic.Phase[*sampleActionDefinition, *sampleBoardProfile, *sampleConfig]
type sAP = logic.ActionProfile[*sampleActionDefinition]
type sAR = logic.ActionRequest[*sampleActionDefinition]
type sBPD = logic.BoardProfileDefinition[*sampleBoardProfile]

// samplePhaseAllPlayerAction is a sample phase that all players can do an action.
type samplePhaseAllPlayerAction struct{}

var _ sPhase = &samplePhaseAllPlayerAction{}

func (s *samplePhaseAllPlayerAction) Name() logic.PhaseName {
	return samplePhaseNamePlay
}

func (s *samplePhaseAllPlayerAction) Prepare(config *sampleConfig, bp *sampleBoardProfile) (*sAR, *sampleBoardProfile, error) {
	apr := newSampleAR(config.TotalPlayer())
	return apr, bp, nil
}

func (s *samplePhaseAllPlayerAction) Execute(config *sampleConfig, bp *sampleBoardProfile, ap *sAP) (logic.PhaseName, *sampleBoardProfile, error) {
	if bp.turn > 3 {
		return "", bp, nil
	}
	return samplePhaseNamePlay, bp, nil
}

func newSampleAR(totalPlayer uint) *sAR {
	r := logic.NewActionRequest[*sampleActionDefinition](totalPlayer)
	for i := 0; i < int(totalPlayer); i++ {
		r.RegisterValidator(logic.Player(i), func(a *sampleActionDefinition) error { return nil })
	}
	return r
}

func TestNewGame(t *testing.T) {
	t.Parallel()

	type args struct {
		totalPlayer  uint
		initialPhase logic.PhaseName
		phases       []sPhase
		bpd          sBPD
	}
	tests := []struct {
		name string
		args args
		want *sGame
	}{
		{
			name: "ゲームが生成できる",
			args: args{
				totalPlayer:  3,
				initialPhase: samplePhaseNamePlay,
				phases: []sPhase{
					&samplePhaseAllPlayerAction{},
				},
				bpd: &sampleBoardProfileDefinition{totalPlayer: 3},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := board.NewGame(tt.args.initialPhase, tt.args.phases, tt.args.bpd)
			if got == nil {
				t.Errorf("NewGame() should return a game")
			}
		})
	}
}
