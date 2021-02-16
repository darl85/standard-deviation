package calculations

import (
	"math"
)

func calculateStandardDeviation(numberSet []int) float64 {
	var sum, mean, powSum float64
	setLength := float64(len(numberSet))

	for _, num := range numberSet {
		sum += float64(num)
	}
	mean = sum / setLength

	for _, num := range numberSet {
		powSum += math.Pow(float64(num) - mean, 2)
	}

	return math.Sqrt(powSum / setLength)
}