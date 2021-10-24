package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
	"net/http"
)

const (
	// invalidStatusCode means the code isn't setup, please set it.
	invalidStatusCode = 0
	// successMetaCode meas the response is success.It may be different from the HTTP status code.
	// For example, one GET request returns 404 while its meta.code is success.
	successMetaCode = 100000
)

type Meta struct {
	Code    int      `json:"code"`
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type builder struct {
	err        error
	context    *gin.Context
	statusCode int
	*Response
}

func NewRestfulResponseBuilder(c *gin.Context) *builder {
	return &builder{
		statusCode: invalidStatusCode,
		context:    c,
		Response: &Response{
			Data: nil,
			Meta: Meta{
				Code: invalidStatusCode,
			},
		},
	}
}

func (b *builder) Meta(meta Meta) *builder {
	b.Response.Meta = meta
	return b
}

func (b *builder) Data(data interface{}) *builder {
	b.Response.Data = data
	return b
}

func (b *builder) Error(err error) *builder {
	b.err = err
	return b
}

func (b *builder) StatusCode(statusCode int) *builder {
	b.statusCode = statusCode
	return b
}

func (b *builder) Build() (int, Response) {
	if b.err != nil {
		return b.buildErrorResponse()
	}
	return b.buildSuccessResponse()
}

func (b *builder) buildErrorResponse() (int, Response) {
	log.Errorf("%#+v", b.err)
	coder := errors.ParseCoder(b.err)
	b.Response.Meta.Code = coder.Code()
	statusCode := b.buildStatusCode(coder.HTTPStatus())
	return statusCode, *b.Response
}

func (b *builder) buildStatusCode(statusCode int) int {
	if b.statusCode != invalidStatusCode {
		return b.statusCode
	}
	return statusCode
}

func (b *builder) buildSuccessResponse() (int, Response) {
	var statusCode = http.StatusOK
	method := b.context.Request.Method
	switch {
	case method == http.MethodGet && b.Response.Data == nil:
		{
			statusCode = http.StatusNotFound
		}
	default:

	}

	statusCode = b.buildStatusCode(statusCode)
	if b.Response.Meta.Code == invalidStatusCode {
		b.Response.Meta.Code = 100000
	}
	return statusCode, *b.Response
}

func (b *builder) SendJSON() {
	statusCode, response := b.Build()
	b.context.JSON(statusCode, response)
}
