package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthCtrl struct {
}

func (h *healthCtrl) Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "alive",
		"version": gin.Version,
	})
}

func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}
