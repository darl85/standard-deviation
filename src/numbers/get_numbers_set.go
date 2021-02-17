package numbers

import (
	"context"
	"net/http"
)

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
) ([]int, error) {
	select {
		case <- randomApiContext.Done():
			return nil, &CollectingNumbersError{
				code:    http.StatusRequestTimeout,
				message: "Reqeust timeout exceeded:",
			}
		default:
	}

	return randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
}
