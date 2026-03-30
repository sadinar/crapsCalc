package strategy

type ComePassStrategy struct {
	passLineAmount     int
	maxOddsMultipliers map[int]int
}

func (s *ComePassStrategy) GetDontComeAmount() int {
	return 0
}

func (s *ComePassStrategy) GetDontPassAmount() int {
	return 0
}

func NewComePassStrategy(passLineAmount int, multipliers map[int]int) *ComePassStrategy {
	return &ComePassStrategy{
		passLineAmount:     passLineAmount,
		maxOddsMultipliers: multipliers,
	}
}

func (s *ComePassStrategy) GetPassLineAmount() int {
	return s.passLineAmount
}

func (s *ComePassStrategy) GetOddsAmount(point int) int {
	switch point {
	case 2:
		fallthrough
	case 12:
		return s.passLineAmount * s.maxOddsMultipliers[12]
	case 3:
		fallthrough
	case 11:
		return s.passLineAmount * s.maxOddsMultipliers[11]
	case 4:
		fallthrough
	case 10:
		return s.passLineAmount * s.maxOddsMultipliers[10]
	case 5:
		fallthrough
	case 9:
		return s.passLineAmount * s.maxOddsMultipliers[9]
	case 6:
		fallthrough
	case 8:
		return s.passLineAmount * s.maxOddsMultipliers[8]
	default:
		panic("invalid point")
	}
}

func (s *ComePassStrategy) GetBuyAmount(point int) int {
	return 0
}
