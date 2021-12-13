package http_resty

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/marmotedu/errors"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/http"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/http/options"
	"github.com/mingyuans/go-layout/internal/pkg/code"
)

type wxClientImpl struct {
	client *resty.Client
	option http_options.WXClientOption
}

func (w wxClientImpl) SendMessage(ctx context.Context, param http.SendMessageParam) (http.SendMessageResponse, error) {
	var url = w.option.Address

	response := &http.SendMessageResponse{}

	_, err := w.client.R().
		SetResult(response).
		Post(url)

	if err != nil {
		err = errors.WrapC(err, code.ErrUnknown, "")
	}

	return *response, err
}
