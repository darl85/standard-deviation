package numbers

import (
	"context"
)

func getNumbersSet(
	randomApiContext context.Context,
	numberOfIntegers int,
	randomApiClient RandomApiClientInterface,
) ([]int, ClientErrorInterface) {
	select {
		case <- randomApiContext.Done():
			return nil, nil
		default:
	}

	return randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
}
