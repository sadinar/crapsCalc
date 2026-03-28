package runtime

import (
	"crapsSimulator/dice"
	"crapsSimulator/player"
	"crapsSimulator/strategy"
	"crapsSimulator/table"
	"sync"
)

type GameRunner struct {
	wg            *sync.WaitGroup
	inputChannel  chan int
	outputChannel chan int
}

func (gr *GameRunner) Start() {
	defer gr.wg.Done()

	for range gr.inputChannel {
		gr.playAtTable(gr.outputChannel)
	}
}

func (gr *GameRunner) playAtTable(resultChannel chan int) {
	tbl := gr.setupRegularComePass()
	//tbl := gr.setupHorseshoeDigitalComePass()
	//tbl := gr.setupCraplessComePass()
	//tbl := gr.setupStratosphereComePass()
	//tbl := gr.setupCraplessFarExtremes()
	//tbl := gr.setupCraplessExtremes()
	//tbl := gr.setupLeastExtremes()
	//tbl := gr.setupBuyAll()

	for {
		if tbl.LastRoundEndedOnSeven() {
			break
		}
		tbl.Shoot()
	}

	resultChannel <- tbl.GetPlayerBanks()[0]
}

func (gr *GameRunner) setupRegularComePass() *table.Table {
	return table.NewRegularTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.GetStdOddsMultipliers()), 0),
		},
	)
}

func (gr *GameRunner) setupCraplessComePass() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.GetStdOddsMultipliers()), 0),
		},
	)
}

func (gr *GameRunner) setupStratosphereComePass() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.Get100xMultipliers()), 0),
		},
	)
}

func (gr *GameRunner) setupHorseshoeDigitalComePass() *table.Table {
	return table.NewRegularTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.Get2xMultipliers()), 0),
		},
	)
}

func (gr *GameRunner) setupCraplessFarExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 0),
		},
	)
}

func (gr *GameRunner) setupCraplessExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, false), 0),
		},
	)
}

func (gr *GameRunner) setupLeastExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, true), 0),
		},
	)
}

func (gr *GameRunner) setupBuyAll() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyAllStrategy(25), 0),
		},
	)
}
