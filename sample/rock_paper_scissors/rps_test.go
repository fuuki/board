package rock_paper_scissors

import (
	"testing"

	"github.com/fuuki/board/action"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/resource"
)

func TestRockPaperScissorsGame(t *testing.T) {
	tests := []struct {
		name string
		aps  func(g *game.Game[*resource.ResourceProfile]) []*action.ActionProfile
		want func(g *game.Game[*resource.ResourceProfile]) *resource.ResourceProfile
	}{
		{
			name: "",
			aps: func(g *game.Game[*resource.ResourceProfile]) []*action.ActionProfile {
				ap := profileDef().NewActionProfile()
				ap.Select(0, 0)
				ap.Select(1, 1)
				return []*action.ActionProfile{ap, ap, ap}
			},

			want: func(g *game.Game[*resource.ResourceProfile]) *resource.ResourceProfile {
				rp := resource.NewResourceProfile()
				rp.AddResource(0, resource.NewResource())
				rp.AddResource(1, resource.NewResource())
				rp.Player(1).AddPoint(3)
				return rp
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := rockPaperScissorsGame()
			aps := tt.aps(g)
			g.Play(action.NewAutoActionInputer(aps))

			if !g.BoardProfile().Equal(tt.want(g)) {
				t.Errorf("RockPaperScissorsGame() = %v, want %v", g.BoardProfile(), tt.want(g))
			}
		})
	}
}
