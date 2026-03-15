package house

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayComeOutWin(t *testing.T) {
	assert.Equal(t, 5, Casino{}.PayComeOutWin(5))
}

func TestPayOddsWin(t *testing.T) {
	casino := Casino{}
	assert.Equal(t, 10, casino.PayOddsWin(5, 4))
	assert.Equal(t, 10, casino.PayOddsWin(5, 10))

	assert.Equal(t, 9, casino.PayOddsWin(6, 5))
	assert.Equal(t, 9, casino.PayOddsWin(6, 9))

	assert.Equal(t, 6, casino.PayOddsWin(5, 6))
	assert.Equal(t, 6, casino.PayOddsWin(5, 8))

	assert.Equal(t, 15, casino.PayOddsWin(5, 3))
	assert.Equal(t, 15, casino.PayOddsWin(5, 11))

	assert.Equal(t, 30, casino.PayOddsWin(5, 2))
	assert.Equal(t, 30, casino.PayOddsWin(5, 12))
}

func TestPayBuyWin(t *testing.T) {
	casino := Casino{}
	assert.Equal(t, 9, casino.PayBuyWin(5, 4))
	assert.Equal(t, 9, casino.PayBuyWin(5, 10))

	assert.Equal(t, 8, casino.PayBuyWin(6, 5))
	assert.Equal(t, 8, casino.PayBuyWin(6, 9))

	assert.Equal(t, 5, casino.PayBuyWin(5, 6))
	assert.Equal(t, 5, casino.PayBuyWin(5, 8))

	assert.Equal(t, 49, casino.PayBuyWin(25, 4))
	assert.Equal(t, 149, casino.PayBuyWin(25, 2))
}
