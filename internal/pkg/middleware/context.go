// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mingyuans/go-layout/pkg/log"
)

const (
	KeyRequestID string = "reqId"
)

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Filled required fields to new logger instance.
		// Then, we can use `log.L(c)` to get logger instance.
		logger := log.WithValues(KeyRequestID, c.GetString(XRequestIDKey))
		c.Set(log.ContextKey, logger)
		c.Next()
	}
}
