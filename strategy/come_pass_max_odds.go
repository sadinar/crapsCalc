package strategy

type ComePassMaxOddsStrategy struct {
	passLineAmount int
}

func (s *ComePassMaxOddsStrategy) GetDontComeAmount() int {
	return 0
}

func (s *ComePassMaxOddsStrategy) GetDontPassAmount() int {
	return 0
}

func NewComePassMaxOddsStrategy(passLineAmount int) *ComePassMaxOddsStrategy {
	return &ComePassMaxOddsStrategy{
		passLineAmount: passLineAmount,
	}
}

func (s *ComePassMaxOddsStrategy) GetPassLineAmount() int {
	return s.passLineAmount
}

func (s *ComePassMaxOddsStrategy) GetOddsAmount(point, maxOddsMultiplier int) int {
	return s.passLineAmount * maxOddsMultiplier
}

func (s *ComePassMaxOddsStrategy) GetBuyAmount(point int) int {
	return 0
}
