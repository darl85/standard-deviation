package random_api

import "github.com/ybbus/jsonrpc/v2"

type CallableClientInterface interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
}

type randomApiClient struct {
	rpcClient interface{CallableClientInterface}
	apiKey string
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
	Code    int
	Message string
}

func (clientError *ClientError) Error() string {
	return clientError.Message
}
func (clientError *ClientError) GetCode() int {
	return clientError.Code
}
