package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/forumtest-api/responce"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				responce.Error(ctx, nil, fmt.Sprint(err))
			}
		}()

		ctx.Next()
	}
}
