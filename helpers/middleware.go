package helpers

import (
	_ "fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header["Token"][0]
		data,err := ValidateToken(token)
		if err != nil {
			c.Abort()
			return
		}
		_,ok := data["authorized"].(bool)
		if !ok {
			c.JSON(400,gin.H{
				"message": "Not Allowed",
			})
			c.Abort()
			return
		}

		id := data["id"].(float64)
		email := data["email"].(string)

		c.Set("email", email)
		c.Set("id", id)
		

		// before request

		c.Next()

	}
}