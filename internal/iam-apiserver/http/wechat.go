package http

import "context"

type SendMessageParam struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type SendMessageResponse struct {
}

type WXClient interface {
	SendMessage(ctx context.Context, param SendMessageParam) (SendMessageResponse, error)
}
