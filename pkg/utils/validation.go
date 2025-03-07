package utils

import "github.com/gin-gonic/gin"

// ExtractStructFromValidator extracts a struct of type V from the Gin context.
// It retrieves the value associated with the key "parser" and attempts to cast it to *V.
// If the key does not exist, it returns nil.
//
// V: The type of the struct to be extracted.
// c: The Gin context from which to extract the struct.
//
// Returns a pointer to the extracted struct of type V, or nil if the key does not exist.
func ExtractStructFromValidator[V any](c *gin.Context) *V {
	value, ok := c.Get("parser")
	v, _ := value.(*V)
	if !ok {
		return v
	}
	return v
}
