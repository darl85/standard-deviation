package random_api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v2"
	"net/http"
	"testing"
)

func TestReturnCorrectnessCallingRpcClientWithSuccess(t *testing.T) {
	apiCallParams := &apiParams{"some api key", 3, 1,3}
	expectedData := []int{1,2,3}

	clientMock := prepareMockWithSuccessResponse(apiCallParams, expectedData)

	clientInstance := randomApiClient{
		clientMock,
		apiCallParams.ApiKey,
	}

	result, _ := clientInstance.GetRandomIntegers(
		apiCallParams.NumberOfIntegers,
		apiCallParams.Min,
		apiCallParams.Max,
	)

	assert.Equal(t, expectedData, result)
}

func TestApiErrorFromClientCall(t *testing.T) {
	apiError := errors.New("some api error")
	apiCallParams := &apiParams{"some api key", 3, 1,3}
	expectedError := &ClientError{
		code:    http.StatusInternalServerError,
		message: apiError.Error(),
	}

	clientMock := prepareMockWithApiError(apiCallParams, apiError)

	clientInstance := randomApiClient{
		clientMock,
		apiCallParams.ApiKey,
	}

	_, apiError = clientInstance.GetRandomIntegers(
		apiCallParams.NumberOfIntegers,
		apiCallParams.Min,
		apiCallParams.Max,
	)

	assert.Equal(t, expectedError, apiError)
}

func TestRpcResponseApiErrorFromClientCall(t *testing.T) {
	apiCallParams := &apiParams{"some api key", 3, 1,3}

	var testData = []struct{
		rpcApiError *jsonrpc.RPCError
		expectedClientError *ClientError
	}{
		{
			rpcApiError: &jsonrpc.RPCError{
				Code:    666,
				Message: "some internal error",
				Data:    nil,
			},
			expectedClientError : &ClientError{
				code:    666,
				message: "some internal error",
			},
		},
		{
			rpcApiError: &jsonrpc.RPCError{
				Code:    -666,
				Message: "some internal error sourced in negative error code",
				Data:    nil,
			},
			expectedClientError : &ClientError{
				code:    500,
				message: "some internal error sourced in negative error code",
			},
		},
	}

	for _, data := range testData {
		clientMock := prepareMockWithResponseRpcError(apiCallParams, data.rpcApiError)

		clientInstance := randomApiClient{
			clientMock,
			apiCallParams.ApiKey,
		}

		_, apiError := clientInstance.GetRandomIntegers(
			apiCallParams.NumberOfIntegers,
			apiCallParams.Min,
			apiCallParams.Max,
		)

		assert.Equal(t, data.expectedClientError, apiError)
	}
}

func prepareMockWithSuccessResponse(apiParams *apiParams, expectedData []int) *Client {
	rpcResponse := jsonrpc.RPCResponse{
		JSONRPC: "",
		Result: &apiResult{Random: apiRandomResult{
			Data: expectedData,
			CompletionTime: "",
		}},
		Error:   nil,
		ID:      0,
	}

	clientMock := new(Client)
	clientMock.On("Call", "generateIntegers", apiParams).Return(
		&rpcResponse,
		nil,
	).Once()

	return clientMock
}

func prepareMockWithApiError(apiParams *apiParams, apiError error) *Client {
	clientMock := new(Client)
	clientMock.On("Call", "generateIntegers", apiParams).Return(
		nil,
		apiError,
	).Once()

	return clientMock
}

func prepareMockWithResponseRpcError(apiParams *apiParams, apiError *jsonrpc.RPCError) *Client {
	rpcResponse := jsonrpc.RPCResponse{
		JSONRPC: "",
		Result:  nil,
		Error:   apiError,
		ID:      0,
	}

	clientMock := new(Client)
	clientMock.On("Call", "generateIntegers", apiParams).Return(
		&rpcResponse,
		nil,
	).Once()

	return clientMock
}

type Client struct {
	mock.Mock
}

func (c *Client) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	args := c.Called(method, params[0])

	var rpcResponse *jsonrpc.RPCResponse
	if args.Get(0) == nil {
		rpcResponse = nil
	} else {
		rpcResponse = args.Get(0).(*jsonrpc.RPCResponse)
	}

	return rpcResponse, args.Error(1)
}