package strategy

type Strategy interface {
	GetPassLineAmount() int
	GetOddsAmount(point int) int
	GetBuyAmount(point int) int
	GetDontPassAmount() int
}
