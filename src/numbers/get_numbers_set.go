package numbers

import (
	"context"
	"net/http"
	"sync"
	"time"
)

const singleQueryTimeout = 800 * time.Millisecond

func getNumbersSet(
	randomApiContext context.Context,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
	clientResponseResult *clientResponse,
	apiResponseWaitGroup *sync.WaitGroup,
) {
	singleQueryCtx, singleQueryCancel := context.WithTimeout(randomApiContext, singleQueryTimeout)
	defer singleQueryCancel()
	defer apiResponseWaitGroup.Done()

	clientResponseChannel := make(chan singleClientResponse)

	go func() {
		result, clientError := randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
		clientResponseChannel <- singleClientResponse{result, clientError}
	}()

	select {
	case clientResponse := <-clientResponseChannel:
		clientResponseResult.clientResult = append(clientResponseResult.clientResult, clientResponse.clientResult)
		clientResponseResult.clientError = clientResponse.clientError
	case <-singleQueryCtx.Done():
		*clientResponseResult = clientResponse{
			nil,
			&CollectingNumbersError{
				code:    http.StatusInternalServerError,
				message: "Request timeout exceeded",
			},
		}
	}
}
