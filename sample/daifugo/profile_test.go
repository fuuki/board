package daifugo

import (
	"errors"
	"testing"

	"github.com/fuuki/board/action"
	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

func Test_daifugoActionRequest_IsValid(t *testing.T) {
	type args struct {
		ap jAction
	}
	tests := []struct {
		name    string
		ar      *daifugoActionRequest
		args    args
		wantErr error
	}{
		{
			name: "プレイヤー1の手番で、プレイヤー1が所持しているカードを選択した場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoPlayerAction{
						nil,
						{Select: []resource.CardID{"s1"}},
						nil,
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "プレイヤー1の手番で、プレイヤー1が所持していないカードを選択した場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoPlayerAction{
						nil,
						{Select: []resource.CardID{"s2"}},
						nil,
					},
				),
			},
			wantErr: action.ErrInvalidAction,
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
					[]*daifugoPlayerAction{
						nil,
						{Select: []resource.CardID{}},
						nil,
					},
				),
			},
			wantErr: action.ErrMustTakeAction,
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
					[]*daifugoPlayerAction{
						nil,
						nil,
						nil,
					},
				),
			},
			wantErr: action.ErrMustTakeAction, // 不定
		},
		{
			name: "プレイヤー1の手番で、プレイヤー2がカードを選択した場合",
			ar: func() *daifugoActionRequest {
				ar := &daifugoActionRequest{}
				ar.SetPlayer(board.Player(1))
				return ar
			}(),
			args: args{
				ap: *board.NewActionProfileWithAction(
					[]*daifugoPlayerAction{
						nil,
						{Select: []resource.CardID{"s1"}},
						{Select: []resource.CardID{"s2"}},
					},
				),
			},
			wantErr: action.ErrMustNotTakeAction,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ar := tt.ar
			got := ar.IsValid(tt.args.ap)
			if !errors.Is(got, tt.wantErr) {
				t.Errorf("daifugoActionRequest.IsValid() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
