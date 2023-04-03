package common

import "net/http"

type successRes struct {
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *successRes {
	return &successRes{StatusCode: http.StatusOK, Data: data}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}
