package board_test

import (
	"reflect"
	"testing"

	"github.com/fuuki/board"
)

const (
	samplePhaseNamePlay board.PhaseName = "play"
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

var _ board.BoardProfileDefinition[*sampleBoardProfile] = &sampleBoardProfileDefinition{}

type sampleActionDefinition struct {
}
type sGame = board.Game[*sampleBoardProfile, *sampleActionDefinition]
type sPhase = board.Phase[*sampleBoardProfile, *sampleActionDefinition]
type sAP = board.ActionProfile[*sampleActionDefinition]
type sAR = board.ActionRequest[*sampleActionDefinition]
type sBPD = board.BoardProfileDefinition[*sampleBoardProfile]

// samplePhaseAllPlayerAction returns a sample phase
func samplePhaseAllPlayerAction() *sPhase {
	prepare := func(st *board.Status, bp *sampleBoardProfile) (*sAR, *sampleBoardProfile) {
		apr := newSampleAR(st.TotalPlayer())
		return apr, bp
	}
	execute := func(st *board.Status, bp *sampleBoardProfile, ap *sAP) (board.PhaseName, *sampleBoardProfile) {
		if bp.turn > 3 {
			return "", bp
		}
		return samplePhaseNamePlay, bp
	}
	p := board.NewPhase(samplePhaseNamePlay, prepare, execute)
	return p
}

func newSampleAR(totalPlayer uint) *sAR {
	r := board.NewActionRequest[*sampleActionDefinition](totalPlayer)
	for i := 0; i < int(totalPlayer); i++ {
		r.RegisterValidator(board.Player(i), func(a *sampleActionDefinition) error { return nil })
	}
	return r
}

func TestNewGame(t *testing.T) {
	t.Parallel()

	type args struct {
		totalPlayer  uint
		initialPhase board.PhaseName
		phases       []*sPhase
		bpd          sBPD
		options      []board.GameOption[*sampleBoardProfile, *sampleActionDefinition]
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
				phases: []*sPhase{
					samplePhaseAllPlayerAction(),
				},
				bpd:     &sampleBoardProfileDefinition{totalPlayer: 3},
				options: []board.GameOption[*sampleBoardProfile, *sampleActionDefinition]{
					// PhaseChangeChan[*sampleBoardProfile, *sampleActionDefinition](make(chan<- PeriodCount)),
				},
			},
			want: &sGame{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := board.NewGame(tt.args.totalPlayer, tt.args.initialPhase, tt.args.phases, tt.args.bpd, tt.args.options...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
