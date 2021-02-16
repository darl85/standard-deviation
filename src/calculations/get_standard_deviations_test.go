package calculations

import (
	"github.com/stretchr/testify/assert"
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

	expectedData := []handler.StandardDeviation{
		{
		0,
		firstNumbersCollection,
		},
		{
			0,
			secondNumbersCollection,
		},
		{
			0,
			append(firstNumbersCollection, secondNumbersCollection...),
		},
	}

	assert.Equal(t, expectedData, GetStandardDeviations(data))
}