package daifugo

import (
	"testing"

	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

func Test_daifugoActionRequest_IsValid(t *testing.T) {
	type args struct {
		ap jAction
	}
	tests := []struct {
		name string
		ar   *daifugoActionRequest
		args args
		want bool
	}{
		{
			name: "プレイヤー1の手番で、プレイヤー1がカードを選択した場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoActionProfile{
						nil,
						{Select: []resource.CardID{"s1"}},
						nil,
					},
				),
			},
			want: true,
		},
		{
			name: "プレイヤー1の手番で、プレイヤー1が何もカードを選択していない場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoActionProfile{
						nil,
						{Select: []resource.CardID{}},
						nil,
					},
				),
			},
			want: false,
		},
		{
			name: "プレイヤー1の手番で、誰も何もカードを選択していない場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoActionProfile{
						nil,
						nil,
						nil,
					},
				),
			},
			want: false,
		},
		{
			name: "プレイヤー2の手番で、プレイヤー1がカードを選択した場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(2))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoActionProfile{
						nil,
						{Select: []resource.CardID{"s1"}},
						nil,
					},
				),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ar := tt.ar
			if got := ar.IsValid(tt.args.ap); got != tt.want {
				t.Errorf("daifugoActionRequest.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
