package action

import (
	"fmt"

	"github.com/fuuki/board/player"
)

type ActionInputer interface {
	Input(apd *ActionProfileDefinition) *ActionProfile
}

type InteractiveActionInputer struct {
}

var _ ActionInputer = (*InteractiveActionInputer)(nil)

func (i *InteractiveActionInputer) Input(apd *ActionProfileDefinition) *ActionProfile {
	ap := apd.NewActionProfile()
	for {
		if err := registerAction(ap); err != nil {
			fmt.Println(err)
		}
		if ap.IsSelectCompleted() {
			break
		}
	}
	return ap
}

func registerAction(ap *ActionProfile) error {
	ap.ShowAllChoices()
	p, a := getInput()
	err := ap.Select(player.Player(p), Action(a))
	return err
}

func getInput() (int, int) {
	var p int
	var a int
	for {
		n, _ := fmt.Scanln(&p, &a)
		if n == 2 {
			break
		}
		fmt.Println("無効な入力です。もう一度入力してください。")
	}
	return p, a
}

type AutoActionInputer struct {
	next func() *ActionProfile
}

var _ ActionInputer = (*AutoActionInputer)(nil)

func NewAutoActionInputer(aps []*ActionProfile) *AutoActionInputer {
	iterator := iterator(aps)
	return &AutoActionInputer{
		next: iterator,
	}
}

func (a *AutoActionInputer) Input(_ *ActionProfileDefinition) *ActionProfile {
	return a.next()
}

func iterator[T any](arr []T) func() T {
	index := 0
	return func() T {
		item := arr[index]
		index++
		// If index is out of range, reset it to 0.
		if index >= len(arr) {
			index = 0
		}
		return item
	}
}
