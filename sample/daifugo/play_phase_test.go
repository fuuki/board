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
	if len(expected.PlayerHands) != len(actual.PlayerHands) {
		t.Errorf("expected playerHands length is %d, but got %d", len(expected.PlayerHands), len(actual.PlayerHands))
	}
	for p, expectedHand := range expected.PlayerHands {
		actualHand := actual.PlayerHands[p]
		if !expectedHand.Equals(actualHand) {
			t.Errorf("expected playerHands is %v, but got %v", expectedHand, actualHand)
		}
	}
	if !expected.PlayArea.Equals(actual.PlayArea) {
		t.Errorf("expected PlayArea is %v, but got %v", expected.PlayArea, actual.PlayArea)
	}
	if !reflect.DeepEqual(expected.Turn, actual.Turn) {
		t.Errorf("expected turn is %v, but got %v", expected.Turn, actual.Turn)
	}
}

func Test_playPhaseExecute(t *testing.T) {
	gm, _ := daifugoGame(2)
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
				g: gm,
				bp: &daifugoBoardProfile{
					Turn: resource.NewTurn([]board.Player{0, 1}, 0),
					PlayerHands: []*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{CardID: "s1", Suit: SuitSpade, Rank: 1},
							{CardID: "s2", Suit: SuitSpade, Rank: 2},
							{CardID: "s3", Suit: SuitSpade, Rank: 3},
						}),
						1: resource.NewCardLine([]*Card{
							{CardID: "d1", Suit: SuitDiamond, Rank: 1},
							{CardID: "d2", Suit: SuitDiamond, Rank: 2},
							{CardID: "d3", Suit: SuitDiamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{CardID: "c2", Suit: SuitClub, Rank: 2},
					})},
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
				Turn: resource.NewTurn([]board.Player{0, 1}, 1),
				PlayerHands: []*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{
						{CardID: "s1", Suit: SuitSpade, Rank: 1},
						{CardID: "s3", Suit: SuitSpade, Rank: 3},
					}),
					1: resource.NewCardLine([]*Card{
						{CardID: "d1", Suit: SuitDiamond, Rank: 1},
						{CardID: "d2", Suit: SuitDiamond, Rank: 2},
						{CardID: "d3", Suit: SuitDiamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{CardID: "s2", Suit: SuitDiamond, Rank: 2},
				}),
			},
		},
		{
			name: "プレイヤー1がパスし、プレイヤー2の番になる",
			args: args{
				g: gm,
				bp: &daifugoBoardProfile{
					Turn: resource.NewTurn([]board.Player{0, 1, 2}, 0),
					PlayerHands: []*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{CardID: "s1", Suit: SuitSpade, Rank: 1},
							{CardID: "s2", Suit: SuitSpade, Rank: 2},
							{CardID: "s3", Suit: SuitSpade, Rank: 3},
						}),
						1: resource.NewCardLine([]*Card{
							{CardID: "d1", Suit: SuitDiamond, Rank: 1},
							{CardID: "d2", Suit: SuitDiamond, Rank: 2},
							{CardID: "d3", Suit: SuitDiamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{CardID: "c2", Suit: SuitClub, Rank: 2},
					}),
				},
				ap: func() *jAction {
					p := &daifugoPlayerAction{
						Pass: true,
					}
					ap := board.NewActionProfile[*daifugoPlayerAction](3)
					ap.SetPlayerAction(0, p)
					return ap
				}(),
			},
			want: PlayPhase,
			want1: &daifugoBoardProfile{
				Turn: resource.NewTurn([]board.Player{1, 2}, 0),
				PlayerHands: []*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{
						{CardID: "s1", Suit: SuitSpade, Rank: 1},
						{CardID: "s2", Suit: SuitSpade, Rank: 2},
						{CardID: "s3", Suit: SuitSpade, Rank: 3},
					}),
					1: resource.NewCardLine([]*Card{
						{CardID: "d1", Suit: SuitDiamond, Rank: 1},
						{CardID: "d2", Suit: SuitDiamond, Rank: 2},
						{CardID: "d3", Suit: SuitDiamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{CardID: "c2", Suit: SuitClub, Rank: 2},
				}),
			},
		},
		{
			name: "プレイヤーが s1 を出し、ラウンドが終わる",
			args: args{
				g: gm,
				bp: &daifugoBoardProfile{
					Turn: resource.NewTurn([]board.Player{0, 1}, 0),
					PlayerHands: []*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{CardID: "s1", Suit: SuitSpade, Rank: 1},
						}),
						1: resource.NewCardLine([]*Card{
							{CardID: "d1", Suit: SuitDiamond, Rank: 1},
							{CardID: "d2", Suit: SuitDiamond, Rank: 2},
							{CardID: "d3", Suit: SuitDiamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{CardID: "c2", Suit: SuitClub, Rank: 2},
					}),
				},
				ap: func() *jAction {
					p := &daifugoPlayerAction{
						Select: []resource.CardID{"s1"},
					}
					ap := board.NewActionProfile[*daifugoPlayerAction](2)
					ap.SetPlayerAction(0, p)
					return ap
				}(),
			},
			want: "",
			want1: &daifugoBoardProfile{
				Turn: resource.NewTurn([]board.Player{0, 1}, 0),
				PlayerHands: []*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{}),
					1: resource.NewCardLine([]*Card{
						{CardID: "d1", Suit: SuitDiamond, Rank: 1},
						{CardID: "d2", Suit: SuitDiamond, Rank: 2},
						{CardID: "d3", Suit: SuitDiamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{CardID: "s1", Suit: SuitSpade, Rank: 1},
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
