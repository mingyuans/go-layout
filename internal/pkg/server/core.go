package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
	"net/http"
)

type Meta struct {
	Code    int      `json:"code"`
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func WriteResponse(c *gin.Context, data interface{}, err error) {
	var response = &struct {
		Meta Meta        `json:"meta"`
		Data interface{} `json:"data"`
	}{}

	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), response)
		return
	}

	c.JSON(http.StatusOK, data)
}
