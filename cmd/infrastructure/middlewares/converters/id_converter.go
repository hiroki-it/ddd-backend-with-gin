package converters

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"errors": err})
			return
		}

		c.Set("id", id)
		c.Next()
	}
}
