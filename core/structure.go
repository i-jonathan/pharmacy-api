package core

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Previous bool        `json:"previous"`
	Next     bool        `json:"next"`
	Page     int         `json:"page"`
	Count    int64       `json:"count"`
	Data     interface{} `json:"data"`
}
