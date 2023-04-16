package board

type BoardProfile interface {
	Players() []Player
	Show() string
}

// BoardProfileBase is a base struct for BoardProfile.
type BoardProfileBase struct {
	players []Player
}

// NewBoardProfileBase creates a new BoardProfileBase.
func NewBoardProfileBase(playerCount uint) *BoardProfileBase {
	players := make([]Player, playerCount)
	for i := uint(0); i < playerCount; i++ {
		players[i] = Player(i)
	}
	return &BoardProfileBase{
		players: players,
	}
}

func (b *BoardProfileBase) Players() []Player {
	return b.players
}
