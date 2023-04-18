package host

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fuuki/board/board"
)

// TerminalHost is a host of the game.
// This host is for playing the game in terminal.
type TerminalHost[BP board.BoardProfile, PD board.PlayerActionDefinition] struct {
	g *board.Game[BP, PD]
}

// NewTerminalHost returns a new host.
func NewTerminalHost[BP board.BoardProfile, PD board.PlayerActionDefinition](
	g *board.Game[BP, PD],
) *TerminalHost[BP, PD] {
	return &TerminalHost[BP, PD]{
		g: g,
	}
}

// Play starts the game.
func (gh *TerminalHost[BP, PD]) Play() {
	gh.g.Start()
	for {
		p, act, err := gh.entryInput()
		if err != nil {
			log.Default().Println(err)
			continue
		}
		if err := gh.g.RegisterAction(p, act); err != nil {
			log.Default().Println(err)
			continue
		}
		if gh.g.IsOver() {
			break
		}
	}
	// result := gh.g.resultFn(gh.g)
	// fmt.Printf("%+v", result)
}

// entryInput returns the input of player and action.
func (gh *TerminalHost[BP, PD]) entryInput() (board.Player, PD, error) {
	getInput := func() (int, string) {
		var p int
		var str string
		for {
			n, _ := fmt.Scanln(&p, &str)
			if n == 2 {
				break
			}
			fmt.Println("無効な入力です。もう一度入力してください。")
		}
		return p, str
	}

	fmt.Println("プレイヤー番号とアクションを入力してください。 ex: 0 {\"Hand\":1}")

	act := new(PD) // ここがポインタかどうかで動かないかも？
	p, str := getInput()
	if err := json.Unmarshal([]byte(str), act); err != nil {
		return 0, *new(PD), err
	}
	return board.Player(p), *act, nil
}
