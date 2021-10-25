package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
	"github.com/mingyuans/go-layout/internal/pkg/code"
)

const (
	// invalidStatusCode means the code isn't setup, please set it.
	invalidStatusCode = 0
	// successMetaCode meas the response is success.It may be different from the HTTP status code.
	// For example, one GET request returns 404 while its meta.code is success.
	successMetaCode = 100000
)

var metaTypes = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	400: "BadRequest",
	401: "Unauthorized",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	409: "Conflict",
	422: "UnprocessableEntity",
	429: "TooManyRequests",
	500: "InternalError",
	502: "InternalError",
	503: "InternalError",
	504: "InternalError",
}

type DetailError struct {
	Detail string `json:"detail"`
}

type Meta struct {
	Code    int           `json:"code"`
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Errors  []DetailError `json:"errors"`
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
	statusCode := b.buildStatusCode(coder.HTTPStatus())
	b.Response.Meta.Code = coder.Code()
	b.Response.Meta.Message = coder.String()
	b.Response.Meta.Type = getMetaType(statusCode)
	b.Response.Meta.Errors = []DetailError{
		{
			Detail: b.err.Error(),
		},
	}
	return statusCode, *b.Response
}

func (b *builder) buildStatusCode(statusCode int) int {
	if b.statusCode != invalidStatusCode {
		return b.statusCode
	}
	return statusCode
}

func (b *builder) buildSuccessResponse() (int, Response) {
	b.err = errors.WithCode(code.Success, "")
	statusCode, response := b.buildErrorResponse()
	// we don't fill meta.errors when the request is success.
	response.Meta.Errors = nil
	return statusCode, response
}

func (b *builder) SendJSON() {
	statusCode, response := b.Build()
	b.context.JSON(statusCode, response)
}

func getMetaType(statusCode int) string {
	metaTypeString, ok := metaTypes[statusCode]
	if !ok {
		metaTypeString = ""
	}
	return metaTypeString
}
