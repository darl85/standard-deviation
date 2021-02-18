package numbers


type RandomApiClientInterface interface {
	GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error)
}

type CollectingNumbersError struct {
	code    int
	message string
}

func (numbersError *CollectingNumbersError) Error() string {
	return numbersError.message
}
func (numbersError *CollectingNumbersError) GetCode() int {
	return numbersError.code
}

type clientResponse struct{
	clientResult [][]int
	clientError  error
}

type singleClientResponse struct{
	clientResult []int
	clientError  error
}
