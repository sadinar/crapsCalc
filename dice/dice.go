package dice

import (
	cryptorand "crypto/rand"
	"hash/maphash"
	"math/big"
	"math/rand/v2"
	"strconv"
	"time"
)

type RollGenerator interface {
	Roll() int
}

type SeededDice struct {
	randGenerator *rand.Rand
}

func NewSeededDice() *SeededDice {
	var h maphash.Hash
	_, _ = h.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
	seed := h.Sum64()

	return &SeededDice{
		randGenerator: rand.New(rand.NewPCG(seed, (seed<<1)|1)),
	}
}

func (sd *SeededDice) Roll() int {
	return sd.randGenerator.IntN(6) + 1 + sd.randGenerator.IntN(6) + 1
}

type RandomDice struct{}

func (rd RandomDice) Roll() int {
	randNum, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(6))
	die1 := int(randNum.Int64()) + 1

	randNum, _ = cryptorand.Int(cryptorand.Reader, big.NewInt(6))
	die2 := int(randNum.Int64()) + 1

	return die1 + die2
}
