package main

import (
	"crapsSimulator/runtime"
	"fmt"
	"slices"

	"gonum.org/v1/gonum/stat"
)

func main() {
	numRuns := 1000000
	workerCount := 20
	results := make([]int, numRuns)

	mgr := runtime.NewManager(numRuns, workerCount, runtime.RegularComePass)
	mgr.SimulateGames(results)
	printStats(results)
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
