package table

import (
	"crapsSimulator/dice"
	"crapsSimulator/house"
	"crapsSimulator/player"
	"crapsSimulator/ruleset"
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
	roll := t.dice.Roll()

	if t.point == ruleset.PointOff {
		for _, gambler := range t.gamblers {
			gambler.OfferPassLineBet()
			gambler.OfferDontPassBet()
		}

		t.handlePointOffRoll(roll)
		return
	}

	for _, person := range t.gamblers {
		person.OfferComeLineBet()
		person.OfferDontComeLineBet()
		person.OfferBuyBets(t.ruleset.GetAllowedBuyPoints())
	}
	t.handlePointOnRoll(roll)
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

func (t *Table) handlePointOffRoll(roll int) {
	t.handlePointOffPassActivities(roll)
	t.handlePointOffDontPassActivities(roll)

	if t.ruleset.IsNewPointSet(roll, t.point) {
		t.point = roll
	}
}

func (t *Table) handlePointOffPassActivities(roll int) {
	if t.ruleset.IsComeOutRollWin(roll, t.point) {
		t.handleComeOutWin()
	}

	if t.ruleset.IsComeOutRollLoss(roll, t.point) {
		for _, person := range t.gamblers {
			person.RemovePassLineBet()
		}
	}

	if t.ruleset.IsNewPointSet(roll, t.point) {
		t.handlePassNewPointActivity(roll)
	}
}

func (t *Table) handlePointOffDontPassActivities(roll int) {
	if t.ruleset.IsDontPassWin(roll) {
		t.handleDontPassComeOutWin()
	}

	if t.ruleset.IsDontPassTie(roll) {
		t.handleDontPassComeOutTie()
	}

	if t.ruleset.IsDontPassLoss(roll) {
		t.handleDontPassComeOutLoss()
	}
}

func (t *Table) handlePointOnRoll(roll int) {
	t.handlePointOnPassActivities(roll)
	t.handlePointOnDontPassActivities(roll)
	t.handlePointBoxActivity(roll)

	if t.ruleset.IsPointHit(roll, t.point) {
		t.point = ruleset.PointOff
		t.roundCounter++
		return
	}

	if t.ruleset.HasPointEndedInCraps(roll, t.point) {
		t.point = ruleset.PointOff
		t.roundCounter++
		t.sevenOutLastRound = true
		return
	}
}

func (t *Table) handlePointBoxActivity(roll int) {
	if t.ruleset.IsPointHit(roll, t.point) {
		for _, person := range t.gamblers {
			t.rewardBuyBet(person, roll)
		}
		return
	}

	if t.ruleset.HasPointEndedInCraps(roll, t.point) {
		for _, person := range t.gamblers {
			person.RemoveAllBuyBets()
		}
		return
	}

	if t.ruleset.IsPointBoxNumber(roll) {
		for _, person := range t.gamblers {
			t.rewardBuyBet(person, roll)
		}
	}
}

func (t *Table) handlePointOnPassActivities(roll int) {
	if t.ruleset.IsPointHit(roll, t.point) {
		for _, person := range t.gamblers {
			if person.GetPassLineBet() > 0 {
				person.ReceiveMoney(
					t.house.PayComeOutWin(person.GetPassLineBet()),
				)
				person.ReturnPassLineBet()
			}

			t.rewardOddsBet(person, roll)
			t.moveComeLineBetUpAndOfferOdds(person, roll)
		}
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
		}
		return
	}

	if t.ruleset.IsPointBoxNumber(roll) {
		for _, person := range t.gamblers {
			t.rewardOddsBet(person, roll)

			if person.GetComeBet(roll) > 0 {
				person.ReceiveMoney(
					t.house.PayComeOutWin(person.GetComeBet(roll)),
				)
				person.ReturnComeBet(roll)
			}
			t.moveComeLineBetUpAndOfferOdds(person, roll)
		}
		return
	}

	if t.ruleset.IsComeLineWin(roll) {
		t.processComeLineWin()
		return
	}

	if t.ruleset.IsComeLineLoss(roll) {
		t.processComeLineLoss()
	}
}

func (t *Table) handlePointOnDontPassActivities(roll int) {
	if t.ruleset.IsPointHit(roll, t.point) {
		t.handleDontPassLoss()
		return
	}

	if t.ruleset.HasPointEndedInCraps(roll, t.point) {
		t.handleDontPassWin()

		for _, person := range t.gamblers {
			if person.GetDontComeLineBet() > 0 {
				person.RemoveDontComeLineBet()
			}

			for _, point := range t.ruleset.GetAllowedDontComePoints() {
				if person.GetDontComeBet(point) > 0 {
					person.ReceiveMoney(t.house.PayNoPassWin(person.GetDontComeBet(point)))
					person.ReturnDontComeBet(point)
				}
			}
		}
		return
	}

	if t.ruleset.IsPointBoxNumber(roll) {
		for _, person := range t.gamblers {
			if person.GetDontComeBet(roll) > 0 {
				person.RemoveDontComeBet(roll)
			}
			t.moveDontComeLineBetUp(person, roll)
		}
		return
	}

	if t.ruleset.IsDontPassWin(roll) {
		for _, person := range t.gamblers {
			if person.GetDontComeLineBet() > 0 {
				person.ReceiveMoney(t.house.PayNoPassWin(person.GetDontComeLineBet()))
				person.ReturnDontComeLineBet()
			}
		}
		return
	}

	if t.ruleset.IsDontPassTie(roll) {
		for _, person := range t.gamblers {
			if person.GetDontComeLineBet() > 0 {
				person.ReturnDontComeLineBet()
			}
		}
		return
	}

	if t.ruleset.IsDontPassLoss(roll) {
		for _, person := range t.gamblers {
			if person.GetDontComeLineBet() > 0 {
				person.RemoveDontComeLineBet()
			}
		}
	}
}

func (t *Table) handleDontPassWin() {
	for _, person := range t.gamblers {
		if person.GetDontPassBet() > 0 {
			person.ReceiveMoney(t.house.PayNoPassWin(person.GetDontPassBet()))
			person.ReturnDontPassBet()
		}
	}
}

func (t *Table) handleDontPassLoss() {
	for _, person := range t.gamblers {
		if person.GetDontPassBet() > 0 {
			person.RemoveDontPassBet()
		}
	}
}

func (t *Table) handleDontPassComeOutWin() {
	for _, person := range t.gamblers {
		if person.GetDontPassBet() > 0 {
			person.ReceiveMoney(t.house.PayNoPassWin(person.GetDontPassBet()))
			person.ReturnDontPassBet()
		}
	}
}

func (t *Table) handleDontPassComeOutTie() {
	for _, person := range t.gamblers {
		if person.GetDontPassBet() > 0 {
			person.ReturnDontPassBet()
		}
	}
}

func (t *Table) handleDontPassComeOutLoss() {
	for _, person := range t.gamblers {
		if person.GetDontPassBet() > 0 {
			person.RemoveDontPassBet()
		}
	}
}

func (t *Table) handlePassNewPointActivity(roll int) {
	for _, person := range t.gamblers {
		if person.GetComeBet(roll) > 0 {
			person.ReceiveMoney(
				t.house.PayComeOutWin(person.GetComeBet(roll)),
			)
			person.ReturnComeBet(roll)
			person.ReturnOddsBet(roll)
		}

		person.OfferOddsBet(roll)
	}
}

func (t *Table) processComeLineWin() {
	for _, person := range t.gamblers {
		if person.GetComeLineBet() > 0 {
			person.ReceiveMoney(
				t.house.PayComeOutWin(person.GetComeLineBet()),
			)
		}
		person.ReturnComeLineBet()
	}
}

func (t *Table) processComeLineLoss() {
	for _, person := range t.gamblers {
		if person.GetComeLineBet() > 0 {
			person.RemoveComeLineBet()
		}
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

func (t *Table) moveDontComeLineBetUp(person *player.Gambler, roll int) {
	if person.GetDontComeLineBet() == 0 {
		return
	}

	person.SetDontComeBet(person.GetDontComeLineBet(), roll)
	person.RemoveDontComeLineBet()
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
