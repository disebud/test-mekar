package utils

// Response Struct
type Response struct {
	Date    string      `json:"date"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func GenerateResponse(date string, status int, message string, result interface{}) Response {
	var responseInfo Response
	responseInfo.Date = date
	responseInfo.Status = status
	responseInfo.Message = message
	responseInfo.Result = result
	return responseInfo
}
