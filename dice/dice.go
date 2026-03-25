package dice

import (
	cryptorand "crypto/rand"
	"math/big"
	stdrand "math/rand"
)

type RollGenerator interface {
	Roll() int
}

type SeededDice struct{}

func (sd SeededDice) Roll() int {
	return stdrand.Intn(6) + 1 + stdrand.Intn(6) + 1
}

type RandomDice struct{}

func (rd RandomDice) Roll() int {
	randNum, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(6))
	die1 := int(randNum.Int64()) + 1

	randNum, _ = cryptorand.Int(cryptorand.Reader, big.NewInt(6))
	die2 := int(randNum.Int64()) + 1

	return die1 + die2
}
