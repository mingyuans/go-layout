// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import (
	"net/http"

	"github.com/marmotedu/errors"
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

// ErrCode implements `github.com/marmotedu/errors`.Coder interface.
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

var _ errors.Coder = &ErrCode{}

// Code returns the integer code of ErrCode.
func (coder ErrCode) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference returns the reference document.
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}

	return coder.HTTP
}

// nolint: unparam,deadcode
func register(code int, httpStatus int, message string, refs ...string) {
	_, found := metaTypes[httpStatus]
	if !found {
		panic("http code not be defined in code.metaTypes``")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	errors.MustRegister(coder)
}

func GetMetaType(statusCode int) string {
	metaTypeString, ok := metaTypes[statusCode]
	if !ok {
		metaTypeString = ""
	}
	return metaTypeString
}
