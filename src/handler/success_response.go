package handler

import (
	"encoding/json"
	"net/http"
)

type StandardDeviation struct {
	StdDev float64 `json:"StdDev"`
	Data   []int  `json:"Data"`
}

func HandleSuccessResponse(writer http.ResponseWriter, deviations []StandardDeviation) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(deviations)
}
