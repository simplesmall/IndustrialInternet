package response

import "net/http"

type ResponseStructure struct {
	Code int `json:"code"`
	Msg string	`json:"msg"`
	Data interface{} `json:"data"`
}

func (res ResponseStructure) NotFound() (result ResponseStructure) {
	return ResponseStructure{
		http.StatusNotFound,
		"Sorry,Not found",
		"",
	}
}



