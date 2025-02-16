package utilities

import "github.com/gin-gonic/gin"

func ExtractStructFromValidator[V any](c *gin.Context) *V {
	value, ok := c.Get("parser")
	v, _ := value.(*V)
	if !ok {
		return v
	}
	return v
}
