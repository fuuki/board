package board

import (
	"encoding/json"
	"fmt"
)

type ActionInputer[AP PlayerActionDefinition] interface {
	Input(ActionRequest[AP]) *ActionProfile[AP]
}

type InteractiveActionInputer[AP PlayerActionDefinition] struct {
}

// var _ ActionInputer = (*InteractiveActionInputer)(nil)

func (a *InteractiveActionInputer[AP]) Input(req ActionRequest[AP]) *ActionProfile[AP] {

	ap := NewActionProfile[AP](2)
	for {
		fmt.Println("プレイヤー番号とアクションを入力してください。 ex: 0 {\"Hand\":1}")
		if err := a.entryInput(ap); err != nil {
			fmt.Println(err)
		}
		if req.IsValid(*ap) == nil {
			break
		}
	}
	return ap
}

func (a *InteractiveActionInputer[AP]) entryInput(ap *ActionProfile[AP]) error {
	p, str := getInput()
	if err := a.registerAction(ap, p, str); err != nil {
		return err
	}
	return nil
}

func (a *InteractiveActionInputer[AP]) registerAction(ap *ActionProfile[AP], p int, str string) error {
	if p < 0 || 1 < p {
		return fmt.Errorf("プレイヤー番号は0か1を入力してください。")
	}

	act := new(AP) // ここがポインタかどうかで動かないかも？
	if err := json.Unmarshal([]byte(str), act); err != nil {
		return err
	}
	ap.SetPlayerAction(Player(p), *act)
	return nil
}

func getInput() (int, string) {
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

type AutoActionInputer[AP PlayerActionDefinition] struct {
	next func() *ActionProfile[AP]
}

// var _ ActionInputer = (*AutoActionInputer)(nil)

func NewAutoActionInputer[AP PlayerActionDefinition](apr []*ActionProfile[AP]) *AutoActionInputer[AP] {
	iterator := iterator(apr)
	return &AutoActionInputer[AP]{
		next: iterator,
	}
}

func (a *AutoActionInputer[AP]) Input(_ ActionRequest[AP]) *ActionProfile[AP] {
	return a.next()
}

func iterator[T any](arr []T) func() T {
	index := 0
	return func() T {
		item := arr[index]
		index++
		// If index is out of range, panic occurs.
		return item
	}
}
