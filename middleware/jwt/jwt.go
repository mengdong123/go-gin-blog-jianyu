package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/e"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/util"
	"net/http"
	"time"
)

// gin中间件的使用说明：https://juejin.cn/post/7026709702293585950
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if claims.ExpiresAt < time.Now().Unix() {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
