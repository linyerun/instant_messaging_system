package middlewares

import (
	"github.com/gin-gonic/gin"
	"instant_messaging_system/define"
	"instant_messaging_system/utils"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id, email, err := utils.ParseToken(token)
		if err != nil {
			// 只是阻止之后的middleware+handler的调用
			c.Abort()
			c.JSON(200, define.NewResponseResult(402, "认证不通过, 需要重新登录", nil))
		}
		c.Set("user_id", id)
		c.Set("user_email", email)

		c.Next() // 继续执行middlewares + 核心handler
	}
}
