package numbers

import (
	"context"
	"sync"
	"time"
)

type RandomApiClientInterface interface {
	GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error)
}

func CollectNumberSets(
	requests int,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
	timeout time.Duration,
) (
	[][]int,
	error,
) {
	var apiResponseWaitGroup sync.WaitGroup
	var numbersSetsCollection [][]int
	var apiError error

	randomApiContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for i := 0; i < requests; i++ {
		apiResponseWaitGroup.Add(1)
 		go getNumbersSet(
			randomApiContext,
			numberOfIntegers,
			randomApiClient,
			cancel,
			&numbersSetsCollection,
			&apiError,
			&apiResponseWaitGroup,
		)
	}

	apiResponseWaitGroup.Wait()

	return numbersSetsCollection, apiError
}
