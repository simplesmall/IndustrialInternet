package response

import "net/http"

type ResponseStructure struct {
	Code int `json:"code"`
	Msg string	`json:"msg"`
	Data interface{} `json:"data"`
}
//未找到响应体
func (res ResponseStructure) NotFound() (result ResponseStructure) {
	return ResponseStructure{
		http.StatusNotFound,
		"Sorry,Not found",
		"",
	}
}
//带data 正常响应体
func (res ResponseStructure) OKResult(data interface{})(result ResponseStructure){
	result.Code = http.StatusOK
	result.Msg = "Normal......"
	result.Data = data
	return result
}
//带 msg data 正常响应体
func (res ResponseStructure) OKResultWithMsg(msg string,data interface{})(result ResponseStructure){
	result.Code = http.StatusOK
	result.Msg = msg
	result.Data = data
	return result
}



