package numbers

import (
	"context"
	"standard-deviation/src/random_api"
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

	// TODO possibilities to refactor ?
	for i := 0; i < requests; i++ {
		apiResponseWaitGroup.Add(1)
		go func() {
			numbers, err := getNumbersSet(randomApiContext, numberOfIntegers, randomApiClient)
			if err != nil {
				cancel()

				if assertError, ok := err.(*random_api.ClientError); ok {
					apiError = &CollectingNumbersError{
						code:    assertError.GetCode(),
						message: assertError.Error(),
					}
				} else {
					apiError = err
				}

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
