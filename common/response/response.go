package response

type Response struct {
	Code	int 		`json:"code" example:"200"` // 代码
	Data	interface{} `json:"data"`// 数据集
	Msg 	string 		`json:"msg"` // 消息
}

var response *Response

func (r *Response) Error(code int,msg string) *Response {
	response = &Response{}
	response.Code = code
	response.Msg = msg
	return response
}

func (r *Response) Success(data interface{}) *Response {
	response = &Response{}
	response.Code = 0
	response.Msg = "Success"
	response.Data = data
	return response
}

func Fail(code int,msg string)  *Response {
	return response.Error(code,msg)
}

func OK(data interface{})  *Response {
	return response.Success(data)
}