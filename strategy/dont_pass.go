package strategy

type DoNotPass struct {
	dontPassAmount int
}

func NewDoNotPass(dontPassAmount int) *DoNotPass {
	return &DoNotPass{dontPassAmount}
}

func (d DoNotPass) GetPassLineAmount() int {
	return 0
}

func (d DoNotPass) GetOddsAmount(point int) int {
	return 0
}

func (d DoNotPass) GetBuyAmount(point int) int {
	return 0
}

func (d DoNotPass) GetDontPassAmount() int {
	return d.dontPassAmount
}
