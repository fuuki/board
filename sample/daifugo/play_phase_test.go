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
	if !reflect.DeepEqual(expected.turn, actual.turn) {
		t.Errorf("expected turn is %v, but got %v", expected.turn, actual.turn)
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
				bp: &daifugoBoardProfile{
					BoardProfileBase: board.NewBoardProfileBase(2),
					turn:             resource.NewTurn([]board.Player{0, 1}, 0),
					playerHands: map[board.Player]*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{id: "s1", Suit: Spade, Rank: 1},
							{id: "s2", Suit: Spade, Rank: 2},
							{id: "s3", Suit: Spade, Rank: 3},
						}),
						1: resource.NewCardLine([]*Card{
							{id: "d1", Suit: Diamond, Rank: 1},
							{id: "d2", Suit: Diamond, Rank: 2},
							{id: "d3", Suit: Diamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{id: "c2", Suit: Club, Rank: 2},
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
				turn: resource.NewTurn([]board.Player{0, 1}, 1),
				playerHands: map[board.Player]*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{
						{id: "s1", Suit: Spade, Rank: 1},
						{id: "s3", Suit: Spade, Rank: 3},
					}),
					1: resource.NewCardLine([]*Card{
						{id: "d1", Suit: Diamond, Rank: 1},
						{id: "d2", Suit: Diamond, Rank: 2},
						{id: "d3", Suit: Diamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{id: "s2", Suit: Diamond, Rank: 2},
				}),
			},
		},
		{
			name: "プレイヤー1がパスし、プレイヤー2の番になる",
			args: args{
				bp: &daifugoBoardProfile{
					BoardProfileBase: board.NewBoardProfileBase(3),
					turn:             resource.NewTurn([]board.Player{0, 1, 2}, 0),
					playerHands: map[board.Player]*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{id: "s1", Suit: Spade, Rank: 1},
							{id: "s2", Suit: Spade, Rank: 2},
							{id: "s3", Suit: Spade, Rank: 3},
						}),
						1: resource.NewCardLine([]*Card{
							{id: "d1", Suit: Diamond, Rank: 1},
							{id: "d2", Suit: Diamond, Rank: 2},
							{id: "d3", Suit: Diamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{id: "c2", Suit: Club, Rank: 2},
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
				turn: resource.NewTurn([]board.Player{1, 2}, 0),
				playerHands: map[board.Player]*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{
						{id: "s1", Suit: Spade, Rank: 1},
						{id: "s2", Suit: Spade, Rank: 2},
						{id: "s3", Suit: Spade, Rank: 3},
					}),
					1: resource.NewCardLine([]*Card{
						{id: "d1", Suit: Diamond, Rank: 1},
						{id: "d2", Suit: Diamond, Rank: 2},
						{id: "d3", Suit: Diamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{id: "c2", Suit: Club, Rank: 2},
				}),
			},
		},
		{
			name: "プレイヤーが s1 を出し、ラウンドが終わる",
			args: args{
				bp: &daifugoBoardProfile{
					BoardProfileBase: board.NewBoardProfileBase(2),
					turn:             resource.NewTurn([]board.Player{0, 1}, 0),
					playerHands: map[board.Player]*resource.CardLine[*Card]{
						0: resource.NewCardLine([]*Card{
							{id: "s1", Suit: Spade, Rank: 1},
						}),
						1: resource.NewCardLine([]*Card{
							{id: "d1", Suit: Diamond, Rank: 1},
							{id: "d2", Suit: Diamond, Rank: 2},
							{id: "d3", Suit: Diamond, Rank: 3},
						}),
					},
					PlayArea: resource.NewCardLine([]*Card{
						{id: "c2", Suit: Club, Rank: 2},
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
				turn: resource.NewTurn([]board.Player{0, 1}, 0),
				playerHands: map[board.Player]*resource.CardLine[*Card]{
					0: resource.NewCardLine([]*Card{}),
					1: resource.NewCardLine([]*Card{
						{id: "d1", Suit: Diamond, Rank: 1},
						{id: "d2", Suit: Diamond, Rank: 2},
						{id: "d3", Suit: Diamond, Rank: 3},
					}),
				},
				PlayArea: resource.NewCardLine([]*Card{
					{id: "s1", Suit: Spade, Rank: 1},
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
