package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test/forumtest-api/common"
	"test/forumtest-api/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		//vcalidate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足0"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			ctx.Abort()
			return
		}

		DB := common.GetDB()
		var user model.User
		DB.First(&user, claims.UserId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			ctx.Abort()
			return
		}
		//写入用户信息到上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
