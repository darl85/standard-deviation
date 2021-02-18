package numbers

import (
	"context"
	"standard-deviation/src/random_api"
	"sync"
	"time"
)

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
	var clientResponse clientResponse

	randomApiContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for i := 0; i < requests; i++ {
		apiResponseWaitGroup.Add(1)
		go getNumbersSet(
			randomApiContext,
			numberOfIntegers,
			randomApiClient,
			&clientResponse,
			&apiResponseWaitGroup,
		)
	}

	apiResponseWaitGroup.Wait()

	if clientResponse.clientError != nil {
		if assertError, ok := clientResponse.clientError.(*random_api.ClientError); ok {
			return nil, &CollectingNumbersError{
				code:    assertError.GetCode(),
				message: assertError.Error(),
			}
		} else {
			return nil, clientResponse.clientError
		}
	}

	return clientResponse.clientResult, nil
}
