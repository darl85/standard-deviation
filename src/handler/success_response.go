package handler

import (
	"encoding/json"
	"net/http"
)

type StandardDeviation struct {
	StdDev int `json:"StdDev"`
	Data   []int  `json:"Data"`
}

