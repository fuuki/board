package card

// BoardProfile is a profile of board.
type BoardProfile struct {
	central *CentralBoard
	player  []*PlayerBoard
}

// NewBoardProfile returns a new BoardProfile.
func NewBoardProfile(playerNum int) *BoardProfile {
	return &BoardProfile{
		central: NewCentralBoard(),
		player:  make([]*PlayerBoard, playerNum),
	}
}

type CentralBoard struct {
	Deck        *CardLine
	PlayArea    *CardLine
	DiscardArea *CardLine
}

type PlayerBoard struct {
	HandArea *CardLine
}

func NewCentralBoard() *CentralBoard {
	return &CentralBoard{}
}

func NewPlayerBoard() *PlayerBoard {
	return &PlayerBoard{}
}
