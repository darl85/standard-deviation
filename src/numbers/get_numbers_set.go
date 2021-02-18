package numbers

import (
	"context"
	"net/http"
	"standard-deviation/src/random_api"
	"sync"
	"time"
)

const singleQueryTimeout = 800 * time.Millisecond

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

type clientResponse struct{
	clientResult []int
	clientError  error
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
	defer apiResponseWaitGroup.Done()

	clientResponseChannel := make(chan clientResponse)

	go func() {
		result, clientError := randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
		clientResponseChannel <- clientResponse{result, clientError}
	}()

	select {
	case clientResponse := <-clientResponseChannel:
		clientError := clientResponse.clientError
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
		} else {
			*numbersSetsCollection = append(*numbersSetsCollection, clientResponse.clientResult)
		}
	case <-singleQueryCtx.Done():
		*apiError = &CollectingNumbersError{
			code:    http.StatusInternalServerError,
			message: singleQueryCtx.Err().Error(),
		}
		cancel()
	}
}
