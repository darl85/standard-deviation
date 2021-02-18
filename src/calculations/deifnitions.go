package calculations

type standardDeviationCalculator struct{}

type calculateStandardDeviationInterface interface {
	calculateStandardDeviation(numberSet []int) float64
}
