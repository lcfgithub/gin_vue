package responce

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Responce(ctx *gin.Context, HttpStatus int, code int, data gin.H, msg string)  {
	ctx.JSON(HttpStatus, gin.H{
		"code": code,
		"data": data,
		"msg": msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string)  {
	Responce(ctx, http.StatusOK, 200, data, msg)
}

func Error(ctx *gin.Context, data gin.H, msg string)  {
	Responce(ctx, http.StatusOK, 400, data, msg)
}