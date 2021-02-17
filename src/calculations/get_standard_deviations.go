package calculations

import "standard-deviation/src/handler"

func GetStandardDeviations(numberSets [][]int, calculate calculateStandardDeviationInterface) []handler.StandardDeviation {
	var deviations []handler.StandardDeviation
	var numbersSum []int

	for _, numberSet := range numberSets {
		deviations = append(deviations, handler.StandardDeviation{
			StdDev: calculate.calculateStandardDeviation(numberSet),
			Data: numberSet,
		})

		numbersSum = append(numbersSum, numberSet...)
	}

	sumDeviation := handler.StandardDeviation{
		StdDev: calculate.calculateStandardDeviation(numbersSum),
		Data:   numbersSum,
	}

	return append(deviations, sumDeviation)
}