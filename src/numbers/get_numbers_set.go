package numbers

import (
	"context"
	"standard-deviation/src/random_api"
)

func getNumbersSet(
	randomApiContext context.Context,
	numberOfIntegers int,
	randomApiClient random_api.RandomApiClientInterface,
) ([]int, random_api.ClientErrorInterface) {
	select {
		case <- randomApiContext.Done():
			return nil, nil
		default:
	}

	return randomApiClient.GetRandomIntegers(numberOfIntegers, 1, 100)
}
