package runtime

import (
	"crapsSimulator/dice"
	"crapsSimulator/odds"
	"crapsSimulator/player"
	"crapsSimulator/strategy"
	"crapsSimulator/table"
	"sync"
)

const RegularComePass = "RegularComePass"
const HorseshoeDigitalComePass = "HorseshoeDigitalComePass"
const CraplessComePass = "CraplessComePass"
const StratosphereComePass = "StratosphereComePass"
const CraplessFarExtremes = "CraplessFarExtremes"
const CraplessExtremes = "CraplessExtremes"
const CraplessLeastExtremes = "LeastExtremes"
const CraplessBuyAll = "BuyAll"
const RegularBuyAll = "RegularBuyAll"
const DoNotPassDoNotCome = "DoNotPassDoNotCome"

type GameRunner struct {
	wg            *sync.WaitGroup
	inputChannel  chan int
	outputChannel chan int
}

func (gr *GameRunner) Start(tableConfiguration string) {
	defer gr.wg.Done()

	for range gr.inputChannel {
		gr.playAtTable(tableConfiguration, gr.outputChannel)
	}
}

func (gr *GameRunner) playAtTable(tableType string, resultChannel chan int) {
	var tbl *table.Table
	switch tableType {
	case RegularComePass:
		tbl = gr.setupRegularComePass()
	case HorseshoeDigitalComePass:
		tbl = gr.setupHorseshoeDigitalComePass()
	case CraplessComePass:
		tbl = gr.setupCraplessComePass()
	case StratosphereComePass:
		tbl = gr.setupStratosphereComePass()
	case CraplessFarExtremes:
		tbl = gr.setupCraplessFarExtremes()
	case CraplessExtremes:
		tbl = gr.setupCraplessExtremes()
	case CraplessLeastExtremes:
		tbl = gr.setupCraplessLeastExtremes()
	case CraplessBuyAll:
		tbl = gr.setupCraplessBuyAll()
	case RegularBuyAll:
		tbl = gr.setupRegularBuyAll()
	case DoNotPassDoNotCome:
		tbl = gr.setupDoNotPassDoNotCome()
	default:
		panic("Unrecognized table type")
	}

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
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupCraplessComePass() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupStratosphereComePass() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
		odds.Get100xMaxOdds(),
	)
}

func (gr *GameRunner) setupHorseshoeDigitalComePass() *table.Table {
	return table.NewRegularTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
		odds.Get2xMaxOdds(),
	)
}

func (gr *GameRunner) setupCraplessFarExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupCraplessExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, false), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupCraplessLeastExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, true), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupCraplessBuyAll() *table.Table {
	return table.NewCraplessTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyAllStrategy(25), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupRegularBuyAll() *table.Table {
	return table.NewRegularTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyAllStrategy(25), 0),
		},
		odds.GetStdMaxOdds(),
	)
}

func (gr *GameRunner) setupDoNotPassDoNotCome() *table.Table {
	return table.NewRegularTable(
		dice.NewSeededDice(),
		[]*player.Gambler{
			player.NewPlayer(strategy.NewDontComeDontPass(15), 0),
		},
		odds.GetStdMaxOdds(),
	)
}
