package service

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func ResponseOK(statusmsg string) Response {
	response := Response{
		StatusCode: 0,
		StatusMsg:  statusmsg,
	}
	return response
}
func ResponseERR(statusmsg string) Response {
	response := Response{
		StatusCode: 1,
		StatusMsg:  statusmsg,
	}
	return response
}
