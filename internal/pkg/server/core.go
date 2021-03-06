package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/mingyuans/go-layout/internal/pkg/code"
	"github.com/mingyuans/go-layout/pkg/log"
)

const (
	// invalidStatusCode means the code isn't setup, please set it.
	invalidStatusCode = 0
)

type DetailError struct {
	Detail string `json:"detail"`
}

type Meta struct {
	// Business error code. Please check the code with our docs.
	//
	// Example: 100000
	Code int `json:"code"`
	// The type of error.
	//
	// Example: NotFound
	Type string `json:"type"`
	// The detail message of the error.
	//
	// Example: The user existed.
	Message string `json:"message"`
	// The other messages. But most of this will be empty.
	//
	// Example: [""]
	Errors []string `json:"errors"`
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
	if !errors.IsCode(b.err, code.Success) {
		log.Errorf("%#+v", b.err)
	}
	coder := errors.ParseCoder(b.err)
	statusCode := b.buildStatusCode(coder.HTTPStatus())
	b.Response.Meta.Code = coder.Code()
	b.Response.Meta.Message = coder.String()
	b.Response.Meta.Type = code.GetMetaType(statusCode)
	// Ignore this field at first.
	b.Response.Meta.Errors = nil
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
