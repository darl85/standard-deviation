package numbers

import (
	"context"
)

func getNumbersSet(
	randomApiContext context.Context,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
) ([]int, error) {
	select {
		case <- randomApiContext.Done():
			return nil, nil
		default:
	}

	return randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
}
