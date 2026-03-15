package player

import strategy2 "crapsSimulator/strategy"

type Gambler struct {
	buyBets     map[int]int
	passLineBet int
	comeLineBet int
	comeBets    map[int]int
	oddsBets    map[int]int
	bank        int
	strategy    strategy2.Strategy
}

func NewPlayer(strategy strategy2.Strategy, startingBank int) *Gambler {
	return &Gambler{
		strategy: strategy,
		bank:     startingBank,
		comeBets: make(map[int]int),
		oddsBets: make(map[int]int),
		buyBets:  make(map[int]int),
	}
}

func (p *Gambler) GetBank() int {
	return p.bank
}

func (p *Gambler) ReceiveMoney(amount int) {
	p.bank += amount
}

func (p *Gambler) GetPassLineBet() int {
	return p.passLineBet
}

func (p *Gambler) ReturnPassLineBet() {
	p.bank += p.passLineBet
	p.passLineBet = 0
}

func (p *Gambler) RemovePassLineBet() {
	p.passLineBet = 0
}

func (p *Gambler) OfferPassLineBet() {
	if p.passLineBet > 0 {
		return
	}

	p.passLineBet = p.strategy.GetPassLineAmount()
	p.bank -= p.passLineBet
}

func (p *Gambler) OfferComeLineBet() {
	if p.comeLineBet > 0 {
		return
	}

	p.comeLineBet = p.strategy.GetPassLineAmount()
	p.bank -= p.comeLineBet
}

func (p *Gambler) GetComeLineBet() int {
	return p.comeLineBet
}

func (p *Gambler) RemoveComeLineBet() {
	p.comeLineBet = 0
}

func (p *Gambler) SetComeBet(bet, point int) {
	p.comeBets[point] = bet
}

func (p *Gambler) GetComeBet(point int) int {
	return p.comeBets[point]
}

func (p *Gambler) ReturnComeBet(point int) {
	p.bank += p.comeBets[point]
	p.comeBets[point] = 0
}

func (p *Gambler) OfferOddsBet(point int) {
	p.oddsBets[point] = p.strategy.GetOddsAmount(point)
	p.bank -= p.oddsBets[point]
}

func (p *Gambler) GetOddsBet(point int) int {
	return p.oddsBets[point]
}

func (p *Gambler) ReturnOddsBet(roll int) {
	p.bank += p.oddsBets[roll]
	p.oddsBets[roll] = 0
}

func (p *Gambler) OfferBuyBets(allowedPoints []int) {
	for _, point := range allowedPoints {
		if p.buyBets[point] == 0 {
			p.buyBets[point] = p.strategy.GetBuyAmount(point)
			p.bank -= p.buyBets[point]
		}
	}
}

func (p *Gambler) GetBuyBet(point int) int {
	return p.buyBets[point]
}

func (p *Gambler) ReturnBuyBet(point int) {
	p.bank += p.buyBets[point]
	p.buyBets[point] = 0
}

func (p *Gambler) RemoveAllOddsBets() {
	for i, _ := range p.oddsBets {
		p.oddsBets[i] = 0
	}
}

func (p *Gambler) RemoveAllBuyBets() {
	for i, _ := range p.buyBets {
		p.buyBets[i] = 0
	}
}

func (p *Gambler) RemoveAllComeBets() {
	for i, _ := range p.comeBets {
		p.comeBets[i] = 0
	}
}

func (p *Gambler) ReturnComeLineBet() {
	p.bank += p.comeLineBet
	p.comeLineBet = 0
}
