package main

import (
	"crapsSimulator/runtime"
	"fmt"
	"slices"

	"gonum.org/v1/gonum/stat"
)

func main() {
	numRuns := 1000
	workerCount := 20
	results := make([][]int, numRuns)

	mgr := runtime.NewManager(numRuns, workerCount, runtime.RegularComePass)
	mgr.SimulateGames(results)

	justBanks := make([]int, 0)
	topGames := make([][]int, 0)
	for _, result := range results {
		justBanks = append(justBanks, result[0])

		if len(topGames) < 10 {
			topGames = append(topGames, result)
			topGames = sortTopGames(topGames)
			continue
		}

		if result[0] > topGames[9][0] {
			topGames[9] = result
			topGames = sortTopGames(topGames)
		}
	}

	printBestGames(topGames)
	fmt.Println()
	printStats(justBanks)
}

func printBestGames(bestGames [][]int) {
	fmt.Printf("Top %d highest winning games:\n", len(bestGames))
	for _, game := range bestGames {
		fmt.Printf("Change in bank: %d    Rolls at the table: %d\n", game[0], game[1])
	}
}

func sortTopGames(allGames [][]int) [][]int {
	orderedGames := make([][]int, len(allGames))
	remainingGames := allGames

	for i := 0; i < len(allGames); i++ {
		bestIndex := 0
		for x := 0; x < len(remainingGames); x++ {
			if remainingGames[x][0] > remainingGames[bestIndex][0] {
				bestIndex = x
			}
		}
		orderedGames[i] = remainingGames[bestIndex]

		newRemainingGames := make([][]int, 0)
		for x := 0; x < len(remainingGames); x++ {
			if x != bestIndex {
				newRemainingGames = append(newRemainingGames, remainingGames[x])
			}
		}
		remainingGames = newRemainingGames
	}

	return orderedGames
}

func printStats(results []int) {
	mean, stdDeviation := findMeanAndStdDeviation(results)
	fmt.Printf(
		"Median: %.6f\nMean: %.6f\nStdDeviation: %.6f\nTotal: %d\n",
		findMedian(results),
		mean,
		stdDeviation,
		sumResults(results),
	)
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
