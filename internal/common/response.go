package common

import "net/http"

type successRes struct {
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data"`
	Paging     interface{} `json:"paging,omitempty"`
	Filter     interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{StatusCode: http.StatusOK, Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}
