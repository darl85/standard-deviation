package numbers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

	result, _ := CollectNumberSets(numOfRequests, numOfIntegers, randomApiMock)

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
		&ClientError{},
	).Once()

	result, apiError := CollectNumberSets(numOfRequests, numOfIntegers, randomApiMock)

	assert.Nil(t, result)
	assert.NotNil(t, apiError)
}


type ClientError struct {}
func (clientErr *ClientError) Error () string {
	return "some error"
}
func (clientErr *ClientError) GetCode () int {
	return 666
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
