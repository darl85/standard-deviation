package random_api

import (
	"github.com/ybbus/jsonrpc/v2"
	"net/http"
	"os"
)

var (
	rpcClientInstance = jsonrpc.NewClient(os.Getenv("RANDOM_API_ADDRESS"))
	RandomApiClient = &randomApiClient{
		rpcClientInstance,
		os.Getenv("RANDOM_API_KEY"),
	}
)

func (client *randomApiClient) GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error) {
	rpcClient := client.rpcClient
	response, apiError := rpcClient.Call(
		"generateIntegers",
		&apiParams{client.apiKey, numberOfIntegers, min, max},
	)

	if apiError != nil {
		return nil, &ClientError{
			Code:    http.StatusInternalServerError,
			Message: apiError.Error(),
		}
	}
	if response != nil && response.Error != nil {
		code := http.StatusInternalServerError
		if response.Error.Code > 0 {
			code = response.Error.Code
		}

		return nil, &ClientError{
			Code:    code,
			Message: response.Error.Message,
		}
	}

	result := &apiResult{Random: apiRandomResult{
		Data: []int{},
		CompletionTime: "",
	}}

	gettingObjectError := response.GetObject(&result)
	if gettingObjectError != nil {
		return nil, &ClientError{
			Code:    http.StatusInternalServerError,
			Message: gettingObjectError.Error(),
		}
	}

	return result.Random.Data, nil
}
