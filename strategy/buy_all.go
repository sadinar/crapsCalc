package strategy

type BuyAllStrategy struct {
	buyAmount int
}

func NewBuyAllStrategy(amount int) *BuyAllStrategy {
	return &BuyAllStrategy{
		buyAmount: amount,
	}
}

func (b BuyAllStrategy) GetPassLineAmount() int {
	return 0
}

func (b BuyAllStrategy) GetOddsAmount(point int) int {
	return 0
}

func (b BuyAllStrategy) GetBuyAmount(point int) int {
	return b.buyAmount
}
