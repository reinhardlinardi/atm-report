package httpjson

type Response struct {
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
