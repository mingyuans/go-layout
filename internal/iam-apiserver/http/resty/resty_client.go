package http_resty

import (
	"github.com/go-resty/resty/v2"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/http"
	http_options "github.com/mingyuans/go-layout/internal/iam-apiserver/http/options"
)

type restyFactory struct {
	wx *wxClientImpl
}

func (r restyFactory) WXClient() http.WXClient {
	return r.wx
}

func NewFactory(wxOption http_options.WXClientOption) http.Factory {
	factory := &restyFactory{}

	restyClient := resty.New()

	factory.wx = &wxClientImpl{
		client: restyClient,
		option: wxOption,
	}

	return factory
}
