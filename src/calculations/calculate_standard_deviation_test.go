package calculations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProperlyCalculateStandardDeviationForNumbersSet(t *testing.T) {
	data := []int{10, 12, 23, 23, 16, 23, 21, 16}

	assert.Equal(t, 4.898979485566356, StandardDeviationCalculator.calculateStandardDeviation(data))
}