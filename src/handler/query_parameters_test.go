package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestHandleQueryParametersExtraction(t *testing.T) {
	expectedRequestsParam := "2"
	expectedLengthParam := "5"
	request := prepareRequest(expectedRequestsParam, expectedLengthParam)

	requests, length, _ := HandleQueryParameters(request)

	convertedRequests, _ := strconv.Atoi(expectedRequestsParam)
	convertedLength, _ := strconv.Atoi(expectedLengthParam)

	assert.Equal(t, convertedRequests, requests)
	assert.Equal(t, convertedLength, length)
}

func TestReturningErrorOnConversion(t *testing.T) {
	request := prepareRequest("str", "str")

	_, _, handlerError := HandleQueryParameters(request)

	assertErr, _ := handlerError.(*ErrorResponse)

	assert.Equal(t, http.StatusBadRequest, assertErr.GetCode())
	assert.Equal(t, "Parameters are invalid, pass integers for both params", assertErr.Error())
}

func TestReturningErrorOnNegativeParamValue(t *testing.T) {
	request := prepareRequest("-2", "8")

	_, _, handlerError := HandleQueryParameters(request)
	assertErr, _ := handlerError.(*ErrorResponse)

	assert.Equal(t, http.StatusBadRequest, assertErr.GetCode())
	assert.Equal(t, "Parameters cannot be negative", assertErr.Error())
}

func prepareRequest(requestsParam string, lengthParam string) *http.Request {
	return &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "www.google.com",
			Path:   "/search",
			RawQuery: "requests=" + requestsParam + "&length=" + lengthParam,
		},
		ProtoMajor:       1,
		ProtoMinor:       1,
		TransferEncoding: []string{"chunked"},
	}
}