package dice

import "math/rand"

type RollGenerator interface {
	Roll() int
}

type Dice struct{}

func (d Dice) Roll() int {
	return rand.Intn(6) + 1 + rand.Intn(6) + 1
}
