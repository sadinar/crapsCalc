package runtime

import (
	"fmt"
	"sync"

	"golang.org/x/text/message"
)

type Manager struct {
	numberOfRuns  int
	workerCount   int
	taskChannel   chan int
	resultChannel chan int
	workerWg      *sync.WaitGroup
	collectionWg  *sync.WaitGroup
}

func NewManager(numberOfRuns int, workerCount int) *Manager {
	return &Manager{
		numberOfRuns:  numberOfRuns,
		workerCount:   workerCount,
		taskChannel:   make(chan int, workerCount),
		resultChannel: make(chan int, workerCount),
		workerWg:      &sync.WaitGroup{},
		collectionWg:  &sync.WaitGroup{},
	}
}

func (rm *Manager) SimulateGames(results []int) []int {
	go rm.fillTaskChannel()
	rm.runWorkers()
	return rm.waitForResult(results)
}

func (rm *Manager) fillTaskChannel() {
	p := message.NewPrinter(message.MatchLanguage("en"))

	for i := 0; i < rm.numberOfRuns; i++ {
		rm.taskChannel <- 1
		if i%1000000 == 0 {
			if i == 0 {
				fmt.Println("starting simulation...")
			} else {
				_, _ = p.Printf("simulating game number %d\n", i)
			}
		}
	}

	close(rm.taskChannel)
}

func (rm *Manager) runWorkers() {
	for i := 0; i < rm.workerCount; i++ {
		rm.workerWg.Add(1)
		gr := GameRunner{
			wg:            rm.workerWg,
			inputChannel:  rm.taskChannel,
			outputChannel: rm.resultChannel,
		}
		go gr.Start()
	}
}

func (rm *Manager) waitForResult(results []int) []int {
	rm.collectionWg.Add(1)

	go rm.collectResults(results)
	rm.workerWg.Wait()
	close(rm.resultChannel)

	rm.collectionWg.Wait()

	return results
}

func (rm *Manager) collectResults(results []int) {
	defer rm.collectionWg.Done()
	index := 0
	for result := range rm.resultChannel {
		results[index] = result
		index++
	}
}
