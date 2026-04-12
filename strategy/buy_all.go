package strategy

type BuyAllStrategy struct {
	buyAmount int
}

func (b BuyAllStrategy) GetDontComeAmount() int {
	return 0
}

func (b BuyAllStrategy) GetDontPassAmount() int {
	return 0
}

func NewBuyAllStrategy(amount int) *BuyAllStrategy {
	return &BuyAllStrategy{
		buyAmount: amount,
	}
}

func (b BuyAllStrategy) GetPassLineAmount() int {
	return 0
}

func (b BuyAllStrategy) GetOddsAmount(point, maxOddsMultiplier int) int {
	return 0
}

func (b BuyAllStrategy) GetBuyAmount(point int) int {
	return b.buyAmount
}
