package response

import "github.com/gin-gonic/gin"

func ResponseError(c *gin.Context, status int, err error, errMessage string) {
	c.Error(err)
	c.JSON(status, gin.H{
		"error": errMessage,
	})
}
