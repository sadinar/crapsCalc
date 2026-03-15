package house

import "math"

type House interface {
	PayComeOutWin(bet int) int
	PayOddsWin(bet, point int) int
	PayBuyWin(bet, point int) int
}

type Casino struct{}

func (c Casino) PayComeOutWin(bet int) int {
	return bet
}

func (c Casino) PayOddsWin(bet, point int) int {
	switch point {
	case 2:
		fallthrough
	case 12:
		return bet * 6
	case 3:
		fallthrough
	case 11:
		return bet * 3
	case 4:
		fallthrough
	case 10:
		return bet * 2
	case 5:
		fallthrough
	case 9:
		return bet / 2 * 3
	case 6:
		fallthrough
	case 8:
		return bet / 5 * 6
	}

	panic("not a point value!")
}

func (c Casino) PayBuyWin(bet, point int) int {
	baseWinning := c.PayOddsWin(bet, point)
	commission := int(math.Round(0.05 * float64(bet)))
	if commission == 0 {
		return baseWinning - 1
	}

	return baseWinning - commission
}
