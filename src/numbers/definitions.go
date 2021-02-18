package numbers


type RandomApiClientInterface interface {
	GetRandomIntegers(numberOfIntegers int, min int, max int) ([]int, error)
}

type CollectingNumbersError struct {
	code    int
	message string
}

func (timeoutError *CollectingNumbersError) Error() string {
	return timeoutError.message
}
func (timeoutError *CollectingNumbersError) GetCode() int {
	return timeoutError.code
}

type clientResponse struct{
	clientResult [][]int
	clientError  error
}

type singleClientResponse struct{
	clientResult []int
	clientError  error
}
