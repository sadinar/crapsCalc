package strategy

type DontPass struct {
	dontPassAmount int
}

func NewDontPass(dontPassAmount int) *DontPass {
	return &DontPass{dontPassAmount}
}

func (d DontPass) GetPassLineAmount() int {
	return 0
}

func (d DontPass) GetOddsAmount(point, maxOddsMultiplier int) int {
	return 0
}

func (d DontPass) GetBuyAmount(point int) int {
	return 0
}

func (d DontPass) GetDontPassAmount() int {
	return d.dontPassAmount
}

func (d DontPass) GetDontComeAmount() int {
	return 0
}

type DontComeDontPass struct {
	betAmount int
}

func NewDontComeDontPass(dontPassAmount int) *DontComeDontPass {
	return &DontComeDontPass{dontPassAmount}
}

func (d DontComeDontPass) GetPassLineAmount() int {
	return 0
}

func (d DontComeDontPass) GetOddsAmount(point, maxOddsMultiplier int) int {
	return 0
}

func (d DontComeDontPass) GetBuyAmount(point int) int {
	return 0
}

func (d DontComeDontPass) GetDontPassAmount() int {
	return d.betAmount
}

func (d DontComeDontPass) GetDontComeAmount() int {
	return d.betAmount
}
