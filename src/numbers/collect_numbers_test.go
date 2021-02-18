package numbers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"standard-deviation/src/random_api"
	"testing"
	"time"
)

func TestNumberSetsCollecting(t *testing.T) {
	numOfRequests := 2
	numOfIntegers := 3
	randomApiMock := new(randomApiClientMock)
	responseNumberSets := [][]int{
		{3,4,5},
		{1,2,3},
	}
	randomApiMock.On("GetRandomIntegers", numOfIntegers, 1, 100).Return(responseNumberSets[0], nil).Once()
	randomApiMock.On("GetRandomIntegers", numOfIntegers, 1, 100).Return(responseNumberSets[1], nil).Once()

	result, _ := CollectNumberSets(numOfRequests, numOfIntegers, randomApiMock, time.Second*30)

	assert.Equal(t, responseNumberSets, result)
}

func TestNumberSetsCollectingFailFromClient(t *testing.T) {
	numOfRequests := 2
	numOfIntegers := 3
	randomApiMock := new(randomApiClientMock)
	responseNumberSets := [][]int{
		{3,4,5},
	}
	randomApiMock.On("GetRandomIntegers", numOfIntegers, 1, 100).Return(
		responseNumberSets[0],
		nil,
	).Once()
	randomApiMock.On("GetRandomIntegers", numOfIntegers, 1, 100).Return(
		nil,
		&random_api.ClientError{
			Code: 123,
			Message: "Some msg",
		},
	).Once()

	result, apiError := CollectNumberSets(numOfRequests, numOfIntegers, randomApiMock, time.Second*30)

	assert.Nil(t, result)
	assert.NotNil(t, apiError)
}

// TODO silly test, check how to deal with time related tests
func TestNumberSetsCollectingFailOnTimeout(t *testing.T) {
	numOfRequests := 2
	numOfIntegers := 3

	result, apiError := CollectNumberSets(numOfRequests, numOfIntegers, &fakedClient{}, time.Microsecond*1)

	assert.Nil(t, result)
	assert.NotNil(t, apiError)
	assert.Equal(t, "Request timeout exceeded", apiError.Error())
}

type fakedClient struct {}

func (client *fakedClient) GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error) {
	return nil, nil
}

type randomApiClientMock struct {
	mock.Mock
}

func (m *randomApiClientMock) GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error) {
	var numbersSet []int
	args := m.Called(numberOfIntegers, min, max)

	if args.Get(0) == nil {
		numbersSet = nil
	} else {
		numbersSet = args.Get(0).([]int)
	}

	return numbersSet, args.Error(1)
}
