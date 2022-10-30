package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter(engine *gin.Engine, config *viper.Viper) *gin.RouterGroup {
	v := engine.Group(fmt.Sprintf("%s", config.GetStringMap("api")["version"]))

	return v
}
