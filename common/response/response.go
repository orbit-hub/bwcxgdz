package response

type ErrResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	//Data       interface{} `json:"data"`
}
