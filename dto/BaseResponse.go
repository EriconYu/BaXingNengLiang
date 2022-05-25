package dto

type BaseResponse struct {
	Msg    string      `json:"msg,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
	Res    interface{} `json:"res,omitempty"`
}
