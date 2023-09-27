package domain

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
