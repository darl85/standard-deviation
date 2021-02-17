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

type randomApiClient struct{
	rpcClient interface{CallableClientInterface}
	apiKey string
}

type CallableClientInterface interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
}

func (client *randomApiClient) GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error) {
	rpcClient := client.rpcClient
	response, apiError := rpcClient.Call(
		"generateIntegers",
		&apiParams{client.apiKey, numberOfIntegers, min, max},
	)

	if apiError != nil {
		return nil, &ClientError{
			code:    http.StatusInternalServerError,
			message: apiError.Error(),
		}
	}
	if response != nil && response.Error != nil {
		var code int
		if response.Error.Code < 0 {
			// TODO dubious - https://api.random.org/json-rpc/2/error-codes can be mapped more strictly
			code = http.StatusInternalServerError
		} else {
			code = response.Error.Code
		}
		return nil, &ClientError{
			code:    code,
			message: response.Error.Message,
		}
	}

	result := &apiResult{Random: apiRandomResult{
		Data: []int{},
		CompletionTime: "",
	}}

	gettingObjectError := response.GetObject(&result)
	if gettingObjectError != nil {
		return nil, &ClientError{
			code:    http.StatusInternalServerError,
			message: gettingObjectError.Error(),
		}
	}

	return result.Random.Data, nil
}

type apiParams struct {
	ApiKey string `json:"apiKey"`
	NumberOfIntegers int `json:"n"`
	Min int `json:"min"`
	Max int `json:"max"`
}

type apiRandomResult struct {
	Data []int
	CompletionTime string
}

type apiResult struct {
	Random apiRandomResult
}

type ClientError struct {
	code    int
	message string
}

func (clientError *ClientError) Error() string {
	return clientError.message
}
func (clientError *ClientError) GetCode() int {
	return clientError.code
}
