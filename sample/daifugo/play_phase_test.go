package daifugo

import (
	"reflect"
	"testing"

	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

// AssertBoardProfile checks if two daifugoBoardProfile are equa
func AssertBoardProfile(t *testing.T, expected, actual *daifugoBoardProfile) {
	t.Helper()
	if len(expected.playerHands) != len(actual.playerHands) {
		t.Errorf("expected playerHands length is %d, but got %d", len(expected.playerHands), len(actual.playerHands))
	}
	for p, expectedHand := range expected.playerHands {
		actualHand := actual.playerHands[p]
		if !expectedHand.Equals(actualHand) {
			t.Errorf("expected playerHands is %v, but got %v", expectedHand, actualHand)
		}
	}
	if !expected.PlayArea.Equals(actual.PlayArea) {
		t.Errorf("expected PlayArea is %v, but got %v", expected.PlayArea, actual.PlayArea)
	}
}

func Test_playPhaseExecute(t *testing.T) {
	type args struct {
		g  *jGame
		bp *daifugoBoardProfile
		ap *jAction
	}
	tests := []struct {
		name  string
		args  args
		want  board.PhaseName
		want1 *daifugoBoardProfile
	}{
		{
			name: "プレイヤーが s1 を出すと、プレイエリアに s1 が追加される",
			args: args{
				g: daifugoGame(),
				bp: &daifugoBoardProfile{
					turn: resource.NewSimpleTurn(2),
					playerHands: map[board.Player]*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{id: "s1", Suit: Spade, Rank: 1},
							{id: "s2", Suit: Spade, Rank: 2},
							{id: "s3", Suit: Spade, Rank: 3},
							{id: "s4", Suit: Spade, Rank: 4},
						}),
						1: resource.NewCardLine([]*Card{
							{id: "d1", Suit: Diamond, Rank: 1},
							{id: "d2", Suit: Diamond, Rank: 2},
							{id: "d3", Suit: Diamond, Rank: 3},
							{id: "d4", Suit: Diamond, Rank: 4},
						}),
					},
					PlayArea: &resource.CardLine[*Card]{},
				},
				ap: func() *jAction {
					p := &daifugoPlayerAction{
						Select: []resource.CardID{"s2"},
					}
					ap := board.NewActionProfile[*daifugoPlayerAction](2)
					ap.SetPlayerAction(0, p)
					return ap
				}(),
			},
			want: PlayPhase,
			want1: &daifugoBoardProfile{
				turn: resource.NewSimpleTurn(2),
				playerHands: map[board.Player]*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{
						{id: "s1", Suit: Spade, Rank: 1},
						{id: "s3", Suit: Spade, Rank: 3},
						{id: "s4", Suit: Spade, Rank: 4},
					}),
					1: resource.NewCardLine([]*Card{
						{id: "d1", Suit: Diamond, Rank: 1},
						{id: "d2", Suit: Diamond, Rank: 2},
						{id: "d3", Suit: Diamond, Rank: 3},
						{id: "d4", Suit: Diamond, Rank: 4},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{id: "s2", Suit: Diamond, Rank: 2},
				}),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := playPhaseExecute(tt.args.g, tt.args.bp, tt.args.ap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playPhaseExecute() got = %v, want %v", got, tt.want)
			}
			AssertBoardProfile(t, tt.want1, got1)
		})
	}
}
