package numbers

import (
	"context"
	"sync"
)

type ClientErrorInterface interface {
	Error() string
	GetCode() int
}

type RandomApiClientInterface interface {
	GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, interface{ClientErrorInterface})
}

func CollectNumberSets(
	requests int,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
) (
	[][]int,
	ClientErrorInterface,
) {
	var apiResponseWaitGroup sync.WaitGroup
	var numbersSetsCollection [][]int
	var apiError ClientErrorInterface

	randomApiContext, cancel := context.WithCancel(context.Background())

	// TODO possibilities to refactor ?
	for i := 0; i < requests; i++ {
		apiResponseWaitGroup.Add(1)
		go func() {
			numbers, err := getNumbersSet(randomApiContext, numberOfIntegers, randomApiClient)
			if err != nil {
				cancel()
				apiError = err
				numbersSetsCollection = nil
			} else {
				numbersSetsCollection = append(numbersSetsCollection, numbers)
			}

			apiResponseWaitGroup.Done()
		}()
	}

	apiResponseWaitGroup.Wait()
	cancel()

	return numbersSetsCollection, apiError
}
