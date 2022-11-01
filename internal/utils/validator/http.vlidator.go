package validator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HttpValidationError(c *gin.Context, err error) {
	field := strings.Split(err.Error(), "\n")
	c.JSON(http.StatusBadRequest, map[string]any{
		"error":  true,
		"status": 400,
		"fields": field,
	})
}
