package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BadRequestType struct {
	HTTPCode int         `json:"httpCode"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Healthz is a handler for health check
func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func JSONResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    code,
		Message: "ok",
		Data:    data,
	})
}

func NoRouteResult(c *gin.Context) {
	c.JSON(404, gin.H{
		"path":    c.Request.URL.Path,
		"message": "Not Found",
	})
}
