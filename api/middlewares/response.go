package middlewares

type Response_t struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseObject(success bool, data interface{}) *Response_t {
	return &Response_t{
		Success: success,
		Data:    data,
	}
}

func RetOkData(obj interface{}) *Response_t {
	return NewResponseObject(true, obj)
}

func RetOK() *Response_t {
	return NewResponseObject(true, nil)
}

func RetFail(msg string) *Response_t {
	return &Response_t{
		Success: false,
	}
}
