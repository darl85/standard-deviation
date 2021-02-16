package numbers

import (
	"context"
	"standard-deviation/src/random_api"
	"sync"
)

func CollectNumberSets(
	requests int,
	numberOfIntegers int,
	randomApiClient random_api.RandomApiClientInterface,
) (
	[][]int,
	random_api.ClientErrorInterface,
) {
	var apiResponseWaitGroup sync.WaitGroup
	var numbersSetsCollection [][]int
	var apiError random_api.ClientErrorInterface

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
