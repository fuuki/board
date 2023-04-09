package action

import (
	"errors"
	"fmt"

	"github.com/fuuki/board/player"
)

type ActionChoiceOption interface {
	fmt.Stringer
}

type ActionDefinition struct {
	choices []ActionChoiceOption
}

func NewActionDefinition() *ActionDefinition {
	return &ActionDefinition{}
}

func (ad *ActionDefinition) AddChoice(choices []ActionChoiceOption) {
	ad.choices = choices
}

type ActionProfileDefinition struct {
	actionDefinitions map[player.Player]*ActionDefinition
}

func NewActionProfileDefinition() *ActionProfileDefinition {
	return &ActionProfileDefinition{
		actionDefinitions: make(map[player.Player]*ActionDefinition),
	}
}

func (apd *ActionProfileDefinition) AddActionDefinition(player player.Player, actionDefinition *ActionDefinition) {
	apd.actionDefinitions[player] = actionDefinition
}

func (apd *ActionProfileDefinition) IsValidAction(player player.Player, action Action) bool {
	if ad, ok := apd.actionDefinitions[player]; ok {
		return 0 <= action && int(action) < len(ad.choices)
	}
	return false
}

func (apd *ActionProfileDefinition) PlayerCount() int {
	return len(apd.actionDefinitions)
}

func (apd *ActionProfileDefinition) NewActionProfile() *ActionProfile {
	return &ActionProfile{
		def:     apd,
		selects: make(map[player.Player]Action),
	}
}

type Action int

type ActionProfile struct {
	def     *ActionProfileDefinition
	selects map[player.Player]Action
}

func (ap *ActionProfile) IsSelectCompleted() bool {
	return len(ap.selects) == ap.def.PlayerCount()
}

func (ap *ActionProfile) Select(player player.Player, action Action) error {
	if ap.IsSelectCompleted() {
		return errors.New("select is completed")
	}
	if !ap.def.IsValidAction(player, action) {
		return errors.New("invalid action")
	}
	ap.selects[player] = action
	return nil
}

func (ap *ActionProfile) GetAction(player player.Player) (ActionChoiceOption, error) {
	if a, ok := ap.selects[player]; ok {
		return ap.def.actionDefinitions[player].choices[a], nil
	}
	return nil, errors.New("not found")
}

// ShowAllChoices shows the action choices of all players.
func (ap *ActionProfile) ShowAllChoices() {
	for player := range ap.def.actionDefinitions {
		fmt.Printf("[player %d]\n", player)
		ap.ShowChoices(player)
	}
}

// Show the action choices of the specified player.
func (ap *ActionProfile) ShowChoices(player player.Player) {
	if ad, ok := ap.def.actionDefinitions[player]; ok {
		selected, ok := ap.selects[player]
		for i, c := range ad.choices {
			if ok && i == int(selected) {
				fmt.Printf("%d: %s (selected)\n", i, c.String())
			} else {
				fmt.Printf("%d: %s\n", i, c.String())
			}
		}
	}
}
