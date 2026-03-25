package main

import (
	"crapsSimulator/dice"
	"crapsSimulator/player"
	"crapsSimulator/strategy"
	"crapsSimulator/table"
	"fmt"
	"slices"

	"gonum.org/v1/gonum/stat"
)

func main() {
	results := make([]int, 0)
	for i := 0; i < 100000000; i++ {
		result := playAtTable()
		results = append(results, result)
	}

	//fmt.Print("[")
	//for _, result := range results {
	//	fmt.Printf(",%d", result)
	//}
	//fmt.Println("]")
	mean, stdDeviation := findMeanAndStdDeviation(results)
	fmt.Printf(
		"Median: %.6f\nMean: %.6f\nStdDeviation: %.6f\nTotal: %d\n",
		findMedian(results),
		mean,
		stdDeviation,
		sumResults(results),
	)
}

func playAtTable() int {
	//tbl := setupRegularComePass()
	//tbl := setupHorseshoeDigitalComePass()
	//tbl := setupCraplessComePass()
	tbl := setupStratosphereComePass()
	//tbl := setupCraplessFarExtremes()
	//tbl := setupCraplessExtremes()
	//tbl := setupLeastExtremes()
	//tbl := setupBuyAll()

	for {
		if tbl.LastRoundEndedOnSeven() {
			break
		}
		tbl.Shoot()
	}

	return tbl.GetPlayerBanks()[0]
}

func sumResults(results []int) int {
	sum := 0
	for _, result := range results {
		sum += result
	}

	return sum
}

func findMeanAndStdDeviation(data []int) (float64, float64) {
	floatData := make([]float64, 0)
	for _, dataPoint := range data {
		floatData = append(floatData, float64(dataPoint))
	}

	return stat.MeanStdDev(floatData, nil)
}

func findMedian(data []int) float64 {
	slices.Sort(data)

	if len(data)%2 != 0 {
		index := (len(data) - 1) / 2
		return float64(data[index])
	}

	lastIndex := len(data) / 2
	firstMedian := data[lastIndex-1]
	lastMedian := data[lastIndex]

	return (float64(firstMedian) + float64(lastMedian)) / float64(2)
}

func setupRegularComePass() *table.Table {
	return table.NewRegularTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.GetStdOddsMultipliers()), 0),
		},
	)
}

func setupCraplessComePass() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.GetStdOddsMultipliers()), 0),
		},
	)
}

func setupStratosphereComePass() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.Get100xMultipliers()), 0),
		},
	)
}

func setupHorseshoeDigitalComePass() *table.Table {
	return table.NewRegularTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15, strategy.Get2xMultipliers()), 0),
		},
	)
}

func setupCraplessFarExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 0),
		},
	)
}

func setupCraplessExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, false), 0),
		},
	)
}

func setupLeastExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, true), 0),
		},
	)
}

func setupBuyAll() *table.Table {
	return table.NewCraplessTable(
		dice.SeededDice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyAllStrategy(25), 0),
		},
	)
}
