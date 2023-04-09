package rock_paper_scissors_test

import (
	"testing"

	"github.com/fuuki/board/action"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/resource"
	"github.com/fuuki/board/sample/rock_paper_scissors"
)

func TestRockPaperScissorsGame(t *testing.T) {
	tests := []struct {
		name string
		aps  func(g *game.Game) []*action.ActionProfile
		want func(g *game.Game) *resource.ResourceProfile
	}{
		{
			name: "",
			aps: func(g *game.Game) []*action.ActionProfile {
				ap := g.ActionProfileDefinition().NewActionProfile()
				ap.Select(0, 0)
				ap.Select(1, 1)
				return []*action.ActionProfile{ap}
			},
			want: func(g *game.Game) *resource.ResourceProfile {
				rp := resource.NewResourceProfile()
				rp.AddResource(0, resource.NewResource())
				rp.AddResource(1, resource.NewResource())
				rp.Player(1).AddPoint(1)
				return rp
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := rock_paper_scissors.RockPaperScissorsGame()
			aps := tt.aps(g)
			g.Play(action.NewAutoActionInputer(aps))

			if !g.ResourceProfile().Equal(tt.want(g)) {
				t.Errorf("RockPaperScissorsGame() = %v, want %v", g.ResourceProfile(), tt.want(g))
			}
		})
	}
}
