package random_api

import (
	"github.com/ybbus/jsonrpc/v2"
	"net/http"
)

var (
	// TODO handle params via .env
	rpcClientInstance = jsonrpc.NewClient("https://api.random.org/json-rpc/2/invoke")
	RandomApiClient = &randomApiClient{
		rpcClientInstance,
		"7033c3d0-1314-4f3e-9cf5-7944e992ac9f",
	}
)

type randomApiClient struct{
	rpcClient interface{CallableClientInterface}
	apiKey string
}

type RandomApiClientInterface interface {
	GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, interface{ clientErrorInterface })
}

type CallableClientInterface interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
}

func (client *randomApiClient) GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, interface{ clientErrorInterface }) {
	rpcClient := client.rpcClient
	response, apiError := rpcClient.Call(
		"generateIntegers",
		&apiParams{client.apiKey, numberOfIntegers, min, max},
	)

	if apiError != nil {
		return nil, &clientError{
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
		return nil, &clientError{
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
		return nil, &clientError{
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

type clientError struct {
	code    int
	message string
}

type clientErrorInterface interface {
	Error() string
	GetCode() int
}

func (error *clientError) Error() string {
	return error.message
}
func (error *clientError) GetCode() int {
	return error.code
}
