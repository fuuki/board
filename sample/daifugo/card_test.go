package daifugo

import (
	"reflect"
	"testing"
)

func TestSortedCards_Power(t *testing.T) {
	type fields struct {
		list []*Card
	}
	tests := []struct {
		name   string
		fields fields
		want   *LinePower
	}{
		{
			name: "階段",
			fields: fields{
				list: []*Card{
					{id: "h1", Rank: 1, Suit: SuitHeart},
					{id: "h2", Rank: 2, Suit: SuitHeart},
					{id: "h3", Rank: 3, Suit: SuitHeart},
				},
			},
			want: &LinePower{
				LineType: LineTypeStairs,
				Count:    3,
				Power:    1,
			},
		},
		{
			name: "シングル",
			fields: fields{
				list: []*Card{
					{id: "h3", Rank: 3, Suit: SuitHeart},
				},
			},
			want: &LinePower{
				LineType: LineTypeSameRank,
				Count:    1,
				Power:    3,
			},
		},
		{
			name: "ペア",
			fields: fields{
				list: []*Card{
					{id: "h2", Rank: 2, Suit: SuitHeart},
					{id: "s2", Rank: 2, Suit: SuitSpade},
				},
			},
			want: &LinePower{
				LineType: LineTypeSameRank,
				Count:    2,
				Power:    2,
			},
		},
		{
			name: "[不正]スートが異なる階段",
			fields: fields{
				list: []*Card{
					{id: "h2", Rank: 2, Suit: SuitHeart},
					{id: "s3", Rank: 3, Suit: SuitSpade},
					{id: "s4", Rank: 4, Suit: SuitSpade},
				},
			},
			want: nil,
		},
		{
			name: "[不正]2枚の階段",
			fields: fields{
				list: []*Card{
					{id: "s3", Rank: 3, Suit: SuitSpade},
					{id: "s4", Rank: 4, Suit: SuitSpade},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &SortedCards{
				list: tt.fields.list,
			}
			if got := s.Power(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedCards.Power() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
