package action

import (
	"errors"
	"testing"

	"github.com/fuuki/board/board"
)

type pd struct{}

func TestTurnValid(t *testing.T) {
	type args struct {
		ap            board.ActionProfile[*pd]
		currentPlayer board.Player
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "プレイヤーアクションが空の場合",
			args: args{
				currentPlayer: 1,
				ap:            board.ActionProfile[*pd]{},
			},
			wantErr: ErrMustTakeAction,
		},
		{
			name: "誰も選択していない場合",
			args: args{
				currentPlayer: 1,
				ap: *board.NewActionProfileWithAction(
					[]*pd{
						nil,
						nil,
						nil,
					},
				),
			},
			wantErr: ErrMustTakeAction,
		},
		{
			name: "プレイヤー1の手番で、プレイヤー1が選択している場合",
			args: args{
				currentPlayer: 1,
				ap: *board.NewActionProfileWithAction(
					[]*pd{
						nil,
						{},
						nil,
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "プレイヤー1の手番で、プレイヤー2が選択している場合",
			args: args{
				currentPlayer: 1,
				ap: *board.NewActionProfileWithAction(
					[]*pd{
						nil,
						{},
						{},
					},
				),
			},
			wantErr: ErrMustNotTakeAction,
		}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := TurnValid(tt.args.ap, tt.args.currentPlayer); !errors.Is(err, tt.wantErr) {
				t.Errorf("TurnValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
