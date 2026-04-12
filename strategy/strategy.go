package strategy

type Strategy interface {
	GetPassLineAmount() int
	GetOddsAmount(point, maxOddsMultiplier int) int
	GetBuyAmount(point int) int
	GetDontPassAmount() int
	GetDontComeAmount() int
}
