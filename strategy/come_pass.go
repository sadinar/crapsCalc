package strategy

type ComePassStrategy struct {
	passLineAmount int
}

func (s *ComePassStrategy) GetDontComeAmount() int {
	return 0
}

func (s *ComePassStrategy) GetDontPassAmount() int {
	return 0
}

func NewComePassStrategy(passLineAmount int) *ComePassStrategy {
	return &ComePassStrategy{
		passLineAmount: passLineAmount,
	}
}

func (s *ComePassStrategy) GetPassLineAmount() int {
	return s.passLineAmount
}

func (s *ComePassStrategy) GetOddsAmount(point, maxOddsMultiplier int) int {
	return s.passLineAmount * maxOddsMultiplier
}

func (s *ComePassStrategy) GetBuyAmount(point int) int {
	return 0
}
