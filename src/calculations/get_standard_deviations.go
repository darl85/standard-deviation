package calculations

import "standard-deviation/src/handler"

func GetStandardDeviations(numberSets [][]int) []handler.StandardDeviation {
	var deviations []handler.StandardDeviation
	var numbersSum []int

	for _, numberSet := range numberSets {
		deviations = append(deviations, handler.StandardDeviation{
			StdDev: 0,
			Data: numberSet,
		})

		numbersSum = append(numbersSum, numberSet...)
	}

	sumDeviation := handler.StandardDeviation{
		StdDev: 0,
		Data:   numbersSum,
	}

	return append(deviations, sumDeviation)
}