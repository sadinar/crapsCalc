package strategy

type BuyExtremes struct {
	betAmount          int
	includeThreeEleven bool
	includeFourTen     bool
}

func (b BuyExtremes) GetDontPassAmount() int {
	return 0
}

func NewBuyExtremesStrategy(betAmount int, includeThreeEleven, includeFourTen bool) *BuyExtremes {
	return &BuyExtremes{
		betAmount:          betAmount,
		includeThreeEleven: includeThreeEleven,
		includeFourTen:     includeFourTen,
	}
}

func (b BuyExtremes) GetPassLineAmount() int {
	return 0
}

func (b BuyExtremes) GetOddsAmount(point int) int {
	return 0
}

func (b BuyExtremes) GetBuyAmount(point int) int {
	if point == 2 || point == 12 {
		return b.betAmount
	}

	if b.includeThreeEleven {
		if point == 3 || point == 11 {
			return b.betAmount
		}
	}

	if b.includeFourTen {
		if point == 4 || point == 10 {
			return b.betAmount
		}
	}

	return 0
}
