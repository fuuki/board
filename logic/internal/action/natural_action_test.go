package action

import (
	"testing"
)

func TestNaturalPlayer_AddDice(t *testing.T) {
	type args struct {
		name    string
		numDice int
		numFace int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "10d6",
			args: args{
				name:    "10d6",
				numDice: 10,
				numFace: 6,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			np := NewNaturalPlayer()
			if err := np.AddDice(tt.args.name, tt.args.numDice, tt.args.numFace); (err != nil) != tt.wantErr {
				t.Errorf("NaturalPlayer.AddDice() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := np.GetResults()[tt.args.name]
			if len(got) != tt.args.numDice {
				t.Errorf("NaturalPlayer.AddDice() got = %v, want %v", got, tt.args.numDice)
			}
		})
	}
}

func TestNaturalPlayer_AddShuffle(t *testing.T) {
	type args struct {
		name    string
		numCard int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "5 cards",
			args: args{
				name:    "5 cards",
				numCard: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			np := NewNaturalPlayer()
			if err := np.AddShuffle(tt.args.name, tt.args.numCard); (err != nil) != tt.wantErr {
				t.Errorf("NaturalPlayer.AddShuffle() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := np.GetResults()[tt.args.name]
			if len(got) != tt.args.numCard {
				t.Errorf("NaturalPlayer.AddShuffle() got = %v, want %v", got, tt.args.numCard)
			}
		})
	}
}
