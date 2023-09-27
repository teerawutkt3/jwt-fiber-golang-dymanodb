package common

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
