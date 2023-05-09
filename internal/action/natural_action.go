package action

import (
	"fmt"
	"math/rand"
)

var ErrInvalidParameter = fmt.Errorf("invalid parameter")

// NaturalPlayer returns the nature's action.
type NaturalPlayer struct {
	actions map[string]naturalAction
}

// NewNaturalPlayer returns a new nature's action.
func NewNaturalPlayer() *NaturalPlayer {
	return &NaturalPlayer{
		actions: make(map[string]naturalAction),
	}
}

// AddDice adds dice to the nature's action.
// The result is a slice of int that represents the dice's face.
// numDice is the number of dice.
// numFace is the number of face of the dice.
// For example, if numDice is 2 and numFace is 6 (it means 2d6), the result may be [4, 2].
func (np *NaturalPlayer) AddDice(name string, numDice int, numFace int) error {
	if numDice <= 0 {
		return fmt.Errorf("%w: numDice must be greater than 0", ErrInvalidParameter)
	}
	if numFace <= 0 {
		return fmt.Errorf("%w: numFace must be greater than 0", ErrInvalidParameter)
	}

	np.actions[name] = &dice{
		numDice: numDice,
		numFace: numFace,
	}
	return nil
}

// AddShuffle adds shuffle to the nature's action.
// The result is a slice of int that represents the index of the card.
// numCard is the number of card.
// For example, if numCard is 6, the result may be [5,1,4,0,2,3].
func (np *NaturalPlayer) AddShuffle(name string, numCard int) error {
	if numCard <= 0 {
		return fmt.Errorf("%w: numCard must be greater than 0", ErrInvalidParameter)
	}

	np.actions[name] = &shuffle{
		numCard: numCard,
	}
	return nil
}

// GetResults returns the nature's actions.
func (np *NaturalPlayer) GetResults() map[string][]int {
	results := make(map[string][]int)
	for name, action := range np.actions {
		results[name] = action.action()
	}
	return results
}

type naturalAction interface {
	action() []int
}

type dice struct {
	numDice int
	numFace int
}

func (d *dice) action() []int {
	result := make([]int, d.numDice)
	for i := 0; i < d.numDice; i++ {
		result[i] = rand.Intn(d.numFace) + 1
	}
	return result
}

type shuffle struct {
	numCard int
}

func (s *shuffle) action() []int {
	result := make([]int, s.numCard)
	for i := 0; i < s.numCard; i++ {
		result[i] = i
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
