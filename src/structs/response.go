package structs

// Response represents a standard API response structure
type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// APIResponse creates a new Response instance
func APIResponse(code int, status string, data interface{}) *Response {
	return &Response{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
