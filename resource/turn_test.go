package resource

import (
	"reflect"
	"testing"

	"github.com/fuuki/board/logic"
)

func TestTurn_Next(t *testing.T) {
	type fields struct {
		order   []logic.Player
		current int
	}
	tests := []struct {
		name   string
		fields fields
		want   logic.Player
	}{
		{
			name: "次のプレイヤーを取得できる",
			fields: fields{
				order: []logic.Player{
					3, 2, 1, 0,
				},
				current: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := &Turn{
				order:   tt.fields.order,
				current: tt.fields.current,
			}
			if got := tr.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Turn.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
