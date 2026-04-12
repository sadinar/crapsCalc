package strategy

import (
	"crapsSimulator/odds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPassLineAmount(t *testing.T) {
	ba := NewBuyAllStrategy(5)
	be := NewBuyExtremesStrategy(5, false, false)
	be2 := NewBuyExtremesStrategy(5, true, false)
	be3 := NewBuyExtremesStrategy(5, true, true)
	cp := NewComePassMaxOddsStrategy(5)

	assert.Empty(t, ba.GetPassLineAmount())
	assert.Empty(t, be.GetPassLineAmount())
	assert.Empty(t, be2.GetPassLineAmount())
	assert.Empty(t, be3.GetPassLineAmount())
	assert.Equal(t, 5, cp.GetPassLineAmount())
}

func TestGetOddsAmount(t *testing.T) {
	ba := NewBuyAllStrategy(5)
	be := NewBuyExtremesStrategy(5, false, false)
	be2 := NewBuyExtremesStrategy(5, true, false)
	be3 := NewBuyExtremesStrategy(5, true, true)
	cp := NewComePassMaxOddsStrategy(5)

	stdMaxOdds := odds.GetStdMaxOdds()
	assert.Empty(t, ba.GetOddsAmount(2, stdMaxOdds[2]))
	assert.Empty(t, be.GetOddsAmount(3, stdMaxOdds[3]))
	assert.Empty(t, be2.GetOddsAmount(4, stdMaxOdds[4]))
	assert.Empty(t, be3.GetOddsAmount(5, stdMaxOdds[5]))
	assert.Equal(t, 5*5, cp.GetOddsAmount(6, stdMaxOdds[6]))
}

func TestGetBuyAmount(t *testing.T) {
	ba := NewBuyAllStrategy(5)
	be := NewBuyExtremesStrategy(5, false, false)
	be2 := NewBuyExtremesStrategy(5, true, false)
	be3 := NewBuyExtremesStrategy(5, true, true)
	cp := NewComePassMaxOddsStrategy(5)

	assert.Equal(t, 5, ba.GetBuyAmount(2))
	assert.Equal(t, 5, ba.GetBuyAmount(3))
	assert.Equal(t, 5, ba.GetBuyAmount(4))
	assert.Equal(t, 5, ba.GetBuyAmount(5))
	assert.Equal(t, 5, ba.GetBuyAmount(6))
	assert.Equal(t, 5, ba.GetBuyAmount(8))
	assert.Equal(t, 5, ba.GetBuyAmount(9))
	assert.Equal(t, 5, ba.GetBuyAmount(10))
	assert.Equal(t, 5, ba.GetBuyAmount(11))
	assert.Equal(t, 5, ba.GetBuyAmount(12))

	assert.Equal(t, 5, be.GetBuyAmount(2))
	assert.Empty(t, be.GetBuyAmount(3))
	assert.Empty(t, be.GetBuyAmount(4))
	assert.Empty(t, be.GetBuyAmount(5))
	assert.Empty(t, be.GetBuyAmount(6))
	assert.Empty(t, be.GetBuyAmount(8))
	assert.Empty(t, be.GetBuyAmount(9))
	assert.Empty(t, be.GetBuyAmount(10))
	assert.Empty(t, be.GetBuyAmount(11))
	assert.Equal(t, 5, be.GetBuyAmount(12))

	assert.Equal(t, 5, be2.GetBuyAmount(2))
	assert.Equal(t, 5, be2.GetBuyAmount(3))
	assert.Empty(t, be2.GetBuyAmount(4))
	assert.Empty(t, be2.GetBuyAmount(5))
	assert.Empty(t, be2.GetBuyAmount(6))
	assert.Empty(t, be2.GetBuyAmount(8))
	assert.Empty(t, be2.GetBuyAmount(9))
	assert.Empty(t, be2.GetBuyAmount(10))
	assert.Equal(t, 5, be2.GetBuyAmount(11))
	assert.Equal(t, 5, be2.GetBuyAmount(12))

	assert.Equal(t, 5, be3.GetBuyAmount(2))
	assert.Equal(t, 5, be3.GetBuyAmount(3))
	assert.Equal(t, 5, be3.GetBuyAmount(4))
	assert.Empty(t, be3.GetBuyAmount(5))
	assert.Empty(t, be3.GetBuyAmount(6))
	assert.Empty(t, be3.GetBuyAmount(8))
	assert.Empty(t, be3.GetBuyAmount(9))
	assert.Equal(t, 5, be3.GetBuyAmount(10))
	assert.Equal(t, 5, be3.GetBuyAmount(11))
	assert.Equal(t, 5, be3.GetBuyAmount(12))

	assert.Empty(t, cp.GetBuyAmount(6))
}

func TestNoPass(t *testing.T) {
	np := DontComeDontPass{}
	assert.Empty(t, np.GetPassLineAmount())

	assert.Empty(t, np.GetOddsAmount(4, 0))
	assert.Empty(t, np.GetOddsAmount(5, 0))
	assert.Empty(t, np.GetOddsAmount(6, 0))
	assert.Empty(t, np.GetOddsAmount(8, 0))
	assert.Empty(t, np.GetOddsAmount(9, 0))
	assert.Empty(t, np.GetOddsAmount(10, 0))

	assert.Empty(t, np.GetBuyAmount(4))
	assert.Empty(t, np.GetBuyAmount(5))
	assert.Empty(t, np.GetBuyAmount(6))
	assert.Empty(t, np.GetBuyAmount(8))
	assert.Empty(t, np.GetBuyAmount(9))
	assert.Empty(t, np.GetBuyAmount(10))

	np.betAmount = 250
	assert.Equal(t, 250, np.GetDontPassAmount())

	ba := NewBuyAllStrategy(345)
	assert.Empty(t, ba.GetDontPassAmount())

	be := BuyExtremes{betAmount: 8765}
	assert.Empty(t, be.GetDontPassAmount())

	cp := ComePassMaxOddsStrategy{passLineAmount: 8567}
	assert.Empty(t, cp.GetDontPassAmount())
}
