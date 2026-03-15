package strategy

type ComePassStrategy struct {
	passLineAmount int
}

func NewComePassStrategy(passLineAmount int) *ComePassStrategy {
	return &ComePassStrategy{passLineAmount: passLineAmount}
}

func (s *ComePassStrategy) GetPassLineAmount() int {
	return s.passLineAmount
}

func (s *ComePassStrategy) GetOddsAmount(point int) int {
	switch point {
	case 2:
		fallthrough
	case 12:
		return s.passLineAmount
	case 3:
		fallthrough
	case 11:
		return 2 * s.passLineAmount
	case 4:
		fallthrough
	case 10:
		return 3 * s.passLineAmount
	case 5:
		fallthrough
	case 9:
		return 4 * s.passLineAmount
	case 6:
		fallthrough
	case 8:
		return 5 * s.passLineAmount
	default:
		panic("invalid point")
	}
}

func (s *ComePassStrategy) GetBuyAmount(point int) int {
	return 0
}
