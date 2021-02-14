package random_api

import (
	"github.com/ybbus/jsonrpc/v2"
	"sync"
)

var (
	rpcClientInstance jsonrpc.RPCClient
	once sync.Once
)

type ApiParams struct {
	ApiKey string `json:"apiKey"`
	NumberOfIntegers int `json:"n"`
	Min int `json:"min"`
	Max int `json:"max"`
}

type ApiRandomResult struct {
	Data [5]int
	CompletionTime string
}

type ApiResult struct {
	Random ApiRandomResult
}

func clientInit() jsonrpc.RPCClient {
	once.Do(func () {
		rpcClientInstance = jsonrpc.NewClient("https://api.random.org/json-rpc/2/invoke")
	})

	return rpcClientInstance
}

func getRandomIntegers(numberOfIntegers int, min int, max int) ([5]int, error) {
	rpcClient := clientInit()
	response, _ := rpcClient.Call(
		"generateIntegers",
		&ApiParams{"7033c3d0-1314-4f3e-9cf5-7944e992ac9f", numberOfIntegers, min, max},
	)

	result := &ApiResult{Random: ApiRandomResult{
		Data: [5]int{},
		CompletionTime: "",
	}}

	err := response.GetObject(&result)

	return result.Random.Data, err
}
