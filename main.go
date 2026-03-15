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
	for i := 0; i < 1000000; i++ {
		result := playAtTable()
		results = append(results, result)
	}

	//fmt.Print("[")
	//for _, result := range results {
	//	fmt.Printf(",%d", result)
	//}
	//fmt.Println("]")
	mean, stdDeviation := findMeanAndStdDeviation(results)
	median := findMedian(results)
	fmt.Printf(
		"Median: %.6f\nMean: %.6f\nStdDeviation: %.6f\n",
		median,
		mean,
		stdDeviation,
	)
}

func playAtTable() int {
	tbl := setupRegularComePass()
	//tbl := setupCraplessComePass()
	//tbl := setupCraplessFarExtremes()
	//tbl := setupCraplessExtremes()
	//tbl := setupLeastExtremes()

	for {
		if tbl.LastRoundEndedOnSeven() {
			break
		}
		tbl.Shoot()
	}

	return tbl.GetPlayerBanks()[0]
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
		dice.Dice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
	)
}

func setupCraplessComePass() *table.Table {
	return table.NewCraplessTable(
		dice.Dice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewComePassStrategy(15), 0),
		},
	)
}

func setupCraplessFarExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.Dice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, false, false), 0),
		},
	)
}

func setupCraplessExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.Dice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, false), 0),
		},
	)
}

func setupLeastExtremes() *table.Table {
	return table.NewCraplessTable(
		dice.Dice{},
		[]*player.Gambler{
			player.NewPlayer(strategy.NewBuyExtremesStrategy(25, true, true), 0),
		},
	)
}
