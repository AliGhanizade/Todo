package share

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}, ctx *gin.Context)  {
ctx.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}


type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string, ctx *gin.Context)  {
	ctx.JSON(code, Error{
		Code:    code,
		Message: message,
	})
}
