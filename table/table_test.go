package table

import (
	"crapsSimulator/house"
	"crapsSimulator/player"
	"crapsSimulator/ruleset"
	"crapsSimulator/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FixedDice struct {
	callCount  int
	returnVals []int
}

func (f *FixedDice) Roll() int {
	returnVal := f.returnVals[f.callCount]
	f.callCount++

	return returnVal
}

func TestShootHandlesComeOutWinsAndLosses(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{7, 11, 3, 11, 7, 2, 12, 12, 3}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(5, ruleset.GetStdOddsMultipliers()), 0),
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 0),
		},
		house: house.Casino{},
	}
	tbl.Shoot()
	assert.Equal(t, 5, tbl.gamblers[0].GetBank())
	assert.Equal(t, 15, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 10, tbl.gamblers[0].GetBank())
	assert.Equal(t, 30, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 5, tbl.gamblers[0].GetBank())
	assert.Equal(t, 15, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 10, tbl.gamblers[0].GetBank())
	assert.Equal(t, 30, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetBank())
	assert.Equal(t, 45, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 10, tbl.gamblers[0].GetBank())
	assert.Equal(t, 30, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 5, tbl.gamblers[0].GetBank())
	assert.Equal(t, 15, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetBank())
	assert.Equal(t, 0, tbl.gamblers[1].GetBank())

	tbl.Shoot()
	assert.Equal(t, -5, tbl.gamblers[0].GetBank())
	assert.Equal(t, -15, tbl.gamblers[1].GetBank())
}

func TestTableSetsNewPointAndOffersBets(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(5, ruleset.GetStdOddsMultipliers()), 0),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.NotEqual(t, ruleset.PointOff, tbl.point)
	assert.Equal(t, 4, tbl.point)
	assert.Equal(t, -20, tbl.gamblers[0].GetBank())
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 5, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(4))
}

type BuyEightStrategy struct{}

func (b BuyEightStrategy) GetPassLineAmount() int {
	return 0
}

func (b BuyEightStrategy) GetOddsAmount(point int) int {
	return 0
}

func (b BuyEightStrategy) GetBuyAmount(point int) int {
	if point == 8 {
		return 100
	}
	return 0
}

func TestBuyPayoutAndOffer(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4, 8}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(BuyEightStrategy{}, 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	tbl.Shoot()
	assert.Equal(t, 615, tbl.gamblers[0].GetBank())
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(8))
	assert.Equal(t, 4, tbl.point)
}

func TestComePayoutAndOffer(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4, 6, 6}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 440, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(6))
	assert.Equal(t, 350, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(6))
	assert.Equal(t, 455, tbl.gamblers[0].GetBank())
}

func TestNewPointPaysOldBets(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4, 8, 10, 4, 8}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(8))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(8))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 440, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(8))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 75, tbl.gamblers[0].GetOddsBet(8))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 350, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(8))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 75, tbl.gamblers[0].GetOddsBet(8))
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 290, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(8))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 75, tbl.gamblers[0].GetOddsBet(8))
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 395, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(8))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 75, tbl.gamblers[0].GetOddsBet(8))
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 410, tbl.gamblers[0].GetBank())
}

func TestBackToBackOddsWins(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4, 5, 5}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 440, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 60, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 365, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 60, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 470, tbl.gamblers[0].GetBank())
}

func TestRoundEndsInSevenRolls(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{4, 5, 10, 7}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 440, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 60, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 365, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 60, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 45, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 305, tbl.gamblers[0].GetBank())

	assert.Equal(t, 0, tbl.GetRoundCount())
	tbl.Shoot()
	assert.Equal(t, 1, tbl.GetRoundCount())
	assert.True(t, tbl.LastRoundEndedOnSeven())
	assert.Equal(t, ruleset.PointOff, tbl.point)
	assert.Equal(t, 0, tbl.gamblers[0].GetPassLineBet())
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(4))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(5))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(10))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(10))
	assert.Equal(t, 320, tbl.gamblers[0].GetBank())
	assert.Equal(t, 1, len(tbl.GetPlayerBanks()))
	assert.Equal(t, tbl.gamblers[0].GetBank(), tbl.GetPlayerBanks()[0])
}

func TestComeLineWin(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{9, 11}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 425, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 440, tbl.gamblers[0].GetBank())
}

func TestComeLineLoss(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{8, 2, 3, 12}},
		ruleset: ruleset.Regular{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 410, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 395, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 380, tbl.gamblers[0].GetBank())
	tbl.Shoot()
	assert.Equal(t, 365, tbl.gamblers[0].GetBank())
}

func TestCraplessPayouts(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{8, 2, 3, 12, 12, 3, 11, 11, 2, 7}},
		ruleset: ruleset.Crapless{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, ruleset.GetStdOddsMultipliers()), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 410, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 380, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 335, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 305, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 410, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 515, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(11))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(11))
	assert.Equal(t, 470, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(11))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(11))
	assert.Equal(t, 575, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 15, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 30, tbl.gamblers[0].GetOddsBet(11))
	assert.Equal(t, 15, tbl.gamblers[0].GetComeBet(11))
	assert.Equal(t, 680, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(3))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(3))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(12))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(12))
	assert.Equal(t, 0, tbl.gamblers[0].GetOddsBet(11))
	assert.Equal(t, 0, tbl.gamblers[0].GetComeBet(11))
	assert.Equal(t, 695, tbl.gamblers[0].GetBank())
}

func TestBuyExtremes(t *testing.T) {
	tbl := Table{
		dice:    &FixedDice{returnVals: []int{8, 2, 3, 12, 12, 3, 11, 11, 2, 7}},
		ruleset: ruleset.Crapless{},
		point:   ruleset.PointOff,
		gamblers: []*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 500),
		},
		house: house.Casino{},
	}

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 500, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 624, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 599, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 773, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 922, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 897, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 897, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 897, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 25, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 1071, tbl.gamblers[0].GetBank())

	tbl.Shoot()
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(2))
	assert.Equal(t, 0, tbl.gamblers[0].GetBuyBet(12))
	assert.Equal(t, 1046, tbl.gamblers[0].GetBank())
}
