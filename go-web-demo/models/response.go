package models

type Response struct {
	Msg  string `json:"msg"`
	Code int8   `json:"code"`
	Data string `json:"data"`
}
