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

func RetOK() string {
	return "OK"
}

func RetFail(msg string) string {
	return msg
}
