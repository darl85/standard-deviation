package numbers

import (
	"context"
	"net/http"
	"standard-deviation/src/random_api"
	"sync"
	"time"
)

const singleQueryTimeout = 600 * time.Millisecond

type CollectingNumbersError struct {
	code    int
	message string
}

func (timeoutError *CollectingNumbersError) Error() string {
	return timeoutError.message
}
func (timeoutError *CollectingNumbersError) GetCode() int {
	return timeoutError.code
}

func getNumbersSet(
	randomApiContext context.Context,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
	cancel context.CancelFunc,
	numbersSetsCollection *[][]int,
	apiError *error,
	apiResponseWaitGroup *sync.WaitGroup,
) {
	singleQueryCtx, singleQueryCancel := context.WithTimeout(randomApiContext, singleQueryTimeout)
	defer singleQueryCancel()

	singleResultChannel := make(chan []int)
	errorChannel := make(chan error)

	go func(singleResultChannel chan []int, errorChannel chan error) {
		result, clientError := randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
		singleResultChannel <- result
		errorChannel <- clientError
	}(singleResultChannel, errorChannel)

	select {
	case singleNumberSet := <-singleResultChannel:
		*numbersSetsCollection = append(*numbersSetsCollection, singleNumberSet)
	case clientError := <-errorChannel:
		if clientError != nil {
			if assertError, ok := clientError.(*random_api.ClientError); ok {
				*apiError = &CollectingNumbersError{
					code:    assertError.GetCode(),
					message: assertError.Error(),
				}
			} else {
				*apiError = clientError
			}

			*numbersSetsCollection = nil
			cancel()
		}
	case <-singleQueryCtx.Done():
		*apiError = &CollectingNumbersError{
			code:    http.StatusInternalServerError,
			message: singleQueryCtx.Err().Error(),
		}
	}

	apiResponseWaitGroup.Done()
}
