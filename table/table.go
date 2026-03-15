package table

import (
	"crapsSimulator/dice"
	"crapsSimulator/house"
	"crapsSimulator/player"
	"crapsSimulator/ruleset"
	"fmt"
)

type Table struct {
	dice              dice.RollGenerator
	ruleset           ruleset.Ruleset
	point             int
	gamblers          []*player.Gambler
	house             house.House
	roundCounter      int
	sevenOutLastRound bool
}

func NewRegularTable(dice dice.RollGenerator, gamblers []*player.Gambler) *Table {
	return &Table{
		dice:         dice,
		ruleset:      ruleset.Regular{},
		point:        ruleset.PointOff,
		gamblers:     gamblers,
		house:        house.Casino{},
		roundCounter: 0,
	}
}

func NewCraplessTable(dice dice.RollGenerator, gamblers []*player.Gambler) *Table {
	return &Table{
		dice:         dice,
		ruleset:      ruleset.Crapless{},
		point:        ruleset.PointOff,
		gamblers:     gamblers,
		house:        house.Casino{},
		roundCounter: 0,
	}
}

func (t *Table) Shoot() {
	t.sevenOutLastRound = false
	t.offerLineBets()
	roll := t.dice.Roll()
	//fmt.Println("roll of: ", roll)

	if t.ruleset.IsComeOutRollWin(roll, t.point) {
		t.handleComeOutWin()
		return
	}

	if t.ruleset.IsComeOutRollLoss(roll, t.point) {
		for _, person := range t.gamblers {
			person.RemovePassLineBet()
		}
		return
	}

	if t.ruleset.IsNewPointSet(roll, t.point) {
		for _, person := range t.gamblers {
			if person.GetComeBet(roll) > 0 {
				person.ReceiveMoney(
					t.house.PayComeOutWin(person.GetComeBet(roll)),
				)
				person.ReturnComeBet(roll)
				person.ReturnOddsBet(roll)
			}

			person.OfferOddsBet(roll)
			person.OfferBuyBets(t.ruleset.GetAllowedBuyPoints())
		}

		t.point = roll
		return
	}

	if t.ruleset.IsPointHit(roll, t.point) {
		t.handlePointHit(roll)

		//fmt.Println("Point hit! Round ends in win!")
		t.roundCounter++
		//t.printPlayerBanks()
		return
	}

	if t.ruleset.HasPointEndedInCraps(roll, t.point) {
		for _, person := range t.gamblers {
			if person.GetComeLineBet() > 0 {
				person.ReceiveMoney(
					t.house.PayComeOutWin(person.GetComeLineBet()),
				)
				person.ReturnComeLineBet()
			}

			person.RemovePassLineBet()
			person.RemoveAllComeBets()
			person.RemoveAllOddsBets()
			person.RemoveAllBuyBets()
		}

		t.point = ruleset.PointOff

		//fmt.Println("Seven rolled! Round ends in a loss. Whomp whomp.")
		t.roundCounter++
		t.sevenOutLastRound = true
		//t.printPlayerBanks()
		return
	}

	if t.ruleset.IsPossiblePoint(roll) {
		t.handleOffPointRoll(roll)
		return
	}

	if t.ruleset.IsComeLineWin(roll) {
		for _, person := range t.gamblers {
			if person.GetComeLineBet() > 0 {
				person.ReceiveMoney(
					t.house.PayComeOutWin(person.GetComeLineBet()),
				)
			}
			person.ReturnComeLineBet()
		}
		return
	}

	if t.ruleset.IsComeLineLoss(roll) {
		for _, person := range t.gamblers {
			if person.GetComeLineBet() > 0 {
				person.RemoveComeLineBet()
			}
		}
	}
}

func (t *Table) GetRoundCount() int {
	return t.roundCounter
}

func (t *Table) LastRoundEndedOnSeven() bool {
	return t.sevenOutLastRound
}

func (t *Table) GetPlayerBanks() []int {
	banks := make([]int, 0)
	for _, person := range t.gamblers {
		banks = append(banks, person.GetBank())
	}

	return banks
}

func (t *Table) offerLineBets() {
	if t.point == ruleset.PointOff {
		for _, gambler := range t.gamblers {
			gambler.OfferPassLineBet()
		}
		return
	}

	for _, gambler := range t.gamblers {
		gambler.OfferComeLineBet()
	}
}

func (t *Table) handleComeOutWin() {
	for _, person := range t.gamblers {
		if person.GetPassLineBet() > 0 {
			person.ReceiveMoney(
				t.house.PayComeOutWin(person.GetPassLineBet()),
			)
			person.ReturnPassLineBet()
		}
	}
}

func (t *Table) handleOffPointRoll(roll int) {
	for _, person := range t.gamblers {
		t.rewardBuyBet(person, roll)
		t.rewardOddsBet(person, roll)

		if person.GetComeBet(roll) > 0 {
			person.ReceiveMoney(
				t.house.PayComeOutWin(person.GetComeBet(roll)),
			)
			person.ReturnComeBet(roll)
		}

		t.moveComeLineBetUpAndOfferOdds(person, roll)
		person.OfferBuyBets(t.ruleset.GetAllowedBuyPoints())
	}
}

func (t *Table) handlePointHit(roll int) {
	for _, person := range t.gamblers {
		if person.GetPassLineBet() > 0 {
			person.ReceiveMoney(
				t.house.PayComeOutWin(person.GetPassLineBet()),
			)
			person.ReturnPassLineBet()
		}

		t.rewardOddsBet(person, roll)
		t.rewardBuyBet(person, roll)
		t.moveComeLineBetUpAndOfferOdds(person, roll)
	}

	t.point = ruleset.PointOff
	t.offerLineBets()
}

func (t *Table) moveComeLineBetUpAndOfferOdds(person *player.Gambler, roll int) {
	if person.GetComeLineBet() > 0 {
		person.SetComeBet(person.GetComeLineBet(), roll)
		person.RemoveComeLineBet()
		person.OfferOddsBet(roll)
	}
}

func (t *Table) rewardBuyBet(person *player.Gambler, roll int) {
	if person.GetBuyBet(roll) > 0 {
		person.ReceiveMoney(
			t.house.PayBuyWin(person.GetBuyBet(roll), roll),
		)
		person.ReturnBuyBet(roll)
	}
}

func (t *Table) rewardOddsBet(person *player.Gambler, roll int) {
	if person.GetOddsBet(roll) > 0 {
		person.ReceiveMoney(
			t.house.PayOddsWin(person.GetOddsBet(roll), roll),
		)
		person.ReturnOddsBet(roll)
	}
}

func (t *Table) printPlayerBanks() {
	for _, person := range t.gamblers {
		fmt.Println(person.GetBank())
	}
}
