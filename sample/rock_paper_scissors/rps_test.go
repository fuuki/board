package rock_paper_scissors

import (
	"testing"

	"github.com/fuuki/board/board"
)

func TestRockPaperScissorsGame(t *testing.T) {
	tests := []struct {
		name string
		aps  []*jAction
		want *JankenBoardProfile
	}{
		{
			name: "",
			aps: func() []*jAction {
				a0 := &JankenActionProfile{
					Hand: ROCK,
				}
				a1 := &JankenActionProfile{
					Hand: PPDER,
				}
				ap := board.NewActionProfile[*JankenActionProfile](2)
				ap.SetPlayerAction(board.Player(0), a0)
				ap.SetPlayerAction(board.Player(1), a1)
				return []*jAction{ap, ap, ap}
			}(),

			want: func() *JankenBoardProfile {
				rp := NewJankenBoardProfile(2)
				rp.Player(1).AddPoint(3)
				return rp
			}(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := rockPaperScissorsGame()
			g.Start()
			for _, ap := range tt.aps {
				g.Next(ap)
			}
			if !g.BoardProfile().Equal(tt.want) {
				t.Errorf("RockPaperScissorsGame() = %v, want %v", g.BoardProfile(), tt.want)
			}
		})
	}
}
