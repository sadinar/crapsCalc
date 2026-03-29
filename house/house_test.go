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

func TestPayNoPassWin(t *testing.T) {
	casino := Casino{}
	assert.Equal(t, 1, casino.PayNoPassWin(1))
	assert.Equal(t, 5, casino.PayNoPassWin(5))
	assert.Equal(t, 234, casino.PayNoPassWin(234))
}

func TestPayLayOdds(t *testing.T) {
	casino := Casino{}

	assert.Equal(t, 5, casino.PayLayOddsWin(10, 4))
	assert.Equal(t, 15, casino.PayLayOddsWin(30, 10))

	assert.Equal(t, 2, casino.PayLayOddsWin(3, 5))
	assert.Equal(t, 6, casino.PayLayOddsWin(9, 9))

	assert.Equal(t, 5, casino.PayLayOddsWin(6, 6))
	assert.Equal(t, 30, casino.PayLayOddsWin(36, 8))
}
