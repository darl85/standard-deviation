package calculations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"standard-deviation/src/handler"
	"testing"
)

func TestReturnCalculationsForEachSetAndSum(t *testing.T) {
	firstNumbersCollection := []int{1,4,5,10,12}
	secondNumbersCollection := []int{15,14,25,55,66}
	data := [][]int{
		firstNumbersCollection,
		secondNumbersCollection,
	}

	sumNumbersCollection := append(firstNumbersCollection, secondNumbersCollection...)
	expectedStdDev := 1.0000
	expectedData := []handler.StandardDeviation{
		{
			expectedStdDev,
			firstNumbersCollection,
		},
		{
			expectedStdDev,
			secondNumbersCollection,
		},
		{
			expectedStdDev,
			sumNumbersCollection,
		},
	}

	calcMock := new(calculateStandardDeviationMock)
	calcMock.On("calculateStandardDeviation", firstNumbersCollection).Return(int(expectedStdDev))
	calcMock.On("calculateStandardDeviation", secondNumbersCollection).Return(int(expectedStdDev))
	calcMock.On("calculateStandardDeviation", sumNumbersCollection).Return(int(expectedStdDev))

	assert.Equal(t, expectedData, GetStandardDeviations(data, calcMock))
}

type calculateStandardDeviationMock struct {
	mock.Mock
}

func (m *calculateStandardDeviationMock) calculateStandardDeviation(numberSet []int) float64 {
	args := m.Called(numberSet)
	return float64(args.Int(0))
}
