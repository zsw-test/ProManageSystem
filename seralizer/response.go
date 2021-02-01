package seralizer

type Response struct {
	Code   string      `json:"code"`
	Result string      `json:"result"`
	Data   interface{} `json:"data,omitempty"`
}
