package player

import (
	"crapsSimulator/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p := NewPlayer(strategy.NewBuyAllStrategy(45), 13)
	assert.Equal(t, 13, p.bank)
	assert.NotEmpty(t, p.strategy)
}

func TestMoneyMovement(t *testing.T) {
	p := NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 400)
	assert.Equal(t, 400, p.GetBank())

	p.ReceiveMoney(100)
	assert.Equal(t, 500, p.GetBank())
}

func TestBuyBetMovement(t *testing.T) {
	p := NewPlayer(strategy.NewBuyExtremesStrategy(50, true, false), 1000)

	p.OfferBuyBets([]int{2, 3, 11, 12})
	assert.Equal(t, 800, p.GetBank())
	assert.Equal(t, 50, p.GetBuyBet(2))
	assert.Equal(t, 50, p.GetBuyBet(3))
	assert.Equal(t, 50, p.GetBuyBet(11))
	assert.Equal(t, 50, p.GetBuyBet(12))

	p.ReturnBuyBet(2)
	assert.Equal(t, 850, p.GetBank())
	assert.Equal(t, 0, p.GetBuyBet(2))

	p.ReturnBuyBet(11)
	assert.Equal(t, 900, p.GetBank())
	assert.Equal(t, 0, p.GetBuyBet(11))

	p.RemoveAllBuyBets()
	assert.Equal(t, 900, p.GetBank())
	assert.Equal(t, 0, p.GetBuyBet(3))
	assert.Equal(t, 0, p.GetBuyBet(12))
}

func TestPassLineMovement(t *testing.T) {
	p := NewPlayer(strategy.NewComePassStrategy(15, map[int]int{}), 750)

	p.OfferPassLineBet()
	assert.Equal(t, 735, p.GetBank())
	assert.Equal(t, 15, p.GetPassLineBet())

	p.OfferPassLineBet()
	assert.Equal(t, 735, p.GetBank())
	assert.Equal(t, 15, p.GetPassLineBet())

	p.ReturnPassLineBet()
	assert.Equal(t, 750, p.GetBank())
	assert.Equal(t, 0, p.GetPassLineBet())

	p.OfferPassLineBet()
	p.RemovePassLineBet()
	assert.Equal(t, 735, p.GetBank())
	assert.Equal(t, 0, p.GetPassLineBet())
}

func TestOddsMovement(t *testing.T) {
	multipliers := map[int]int{
		2:  1,
		3:  2,
		4:  3,
		5:  4,
		6:  5,
		8:  5,
		9:  4,
		10: 3,
		11: 2,
		12: 1,
	}
	p := NewPlayer(strategy.NewComePassStrategy(15, multipliers), 750)

	p.OfferOddsBet(2)
	assert.Equal(t, 735, p.GetBank())
	assert.Equal(t, 15, p.GetOddsBet(2))

	p.OfferOddsBet(3)
	assert.Equal(t, 705, p.GetBank())
	assert.Equal(t, 30, p.GetOddsBet(3))

	p.OfferOddsBet(10)
	assert.Equal(t, 660, p.GetBank())
	assert.Equal(t, 45, p.GetOddsBet(10))

	p.OfferOddsBet(6)
	assert.Equal(t, 585, p.GetBank())
	assert.Equal(t, 75, p.GetOddsBet(6))

	p.ReturnOddsBet(6)
	assert.Equal(t, 660, p.GetBank())
	assert.Equal(t, 0, p.GetOddsBet(6))

	p.RemoveAllOddsBets()
	assert.Equal(t, 660, p.GetBank())
	assert.Equal(t, 0, p.GetOddsBet(10))
	assert.Equal(t, 0, p.GetOddsBet(3))
	assert.Equal(t, 0, p.GetOddsBet(2))
}

func TestComeLineMovement(t *testing.T) {
	p := NewPlayer(strategy.NewComePassStrategy(100, map[int]int{}), 1350)

	p.OfferComeLineBet()
	assert.Equal(t, 1250, p.GetBank())
	assert.Equal(t, 100, p.GetComeLineBet())

	p.OfferComeLineBet()
	assert.Equal(t, 1250, p.GetBank())
	assert.Equal(t, 100, p.GetComeLineBet())

	p.RemoveComeLineBet()
	assert.Equal(t, 1250, p.GetBank())
	assert.Equal(t, 0, p.GetComeLineBet())

	p.OfferComeLineBet()
	assert.Equal(t, 1150, p.GetBank())
	assert.Equal(t, 100, p.GetComeLineBet())

	p.ReturnComeLineBet()
	assert.Equal(t, 1250, p.GetBank())
	assert.Equal(t, 0, p.GetComeLineBet())
}

func TestComeMovement(t *testing.T) {
	p := NewPlayer(strategy.NewComePassStrategy(72, map[int]int{}), 2800)

	p.SetComeBet(72, 6)
	assert.Equal(t, 2800, p.GetBank())
	assert.Equal(t, 72, p.GetComeBet(6))

	p.ReturnComeBet(6)
	assert.Equal(t, 2872, p.GetBank())
	assert.Equal(t, 0, p.GetComeBet(6))

	p.SetComeBet(43, 9)
	assert.Equal(t, 2872, p.GetBank())
	assert.Equal(t, 43, p.GetComeBet(9))

	p.RemoveAllComeBets()
	assert.Equal(t, 2872, p.GetBank())
	assert.Equal(t, 0, p.GetComeBet(9))
}
