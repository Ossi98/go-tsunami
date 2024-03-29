package main

import (
	"Ossi98/go-tsunami/http"
	"Ossi98/go-tsunami/http/health"
	"Ossi98/go-tsunami/http/scanner"
	"Ossi98/go-tsunami/internal/cmd"
	"Ossi98/go-tsunami/internal/config"
	"Ossi98/go-tsunami/internal/utils/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	env := os.Getenv("OD_ENV")
	if env == "" {
		env = "dev"
	}

	logger.NewLogger(env, "logs")

	// Viper Config
	c := config.NewConfig(env)

	// Gin Engine and Config
	var e *gin.Engine

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
		e = gin.New()
		e.Use(gin.Recovery())
		e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			log.Infof(fmt.Sprintf("%s - \"%s %s %s %d %s %s %s",
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			))
			return ""
		}))
	} else {
		e = gin.Default()
	}

	// use of a middleware to allow all origins
	e.Use(cors.Default())

	// Block reverse proxy process
	err := e.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	// Init Controllers and Router
	routes(e, c)

	// Run Server
	if e.Run(fmt.Sprintf(":%v", c.GetStringMap("api")["port"])) != nil {
		log.Fatalln("can not start the gin server")
	}
}

func routes(e *gin.Engine, c *viper.Viper) {
	//services
	ps := cmd.NewProcessScan(c)

	// Controller
	hc := health.NewHealthCtrl()
	sc := scanner.NewScannerCtrl(ps, c)

	// Router
	router := http.NewRouter(e, c)

	//Routes
	router.GET("/health", hc.Index)
	router.POST("/scanner/start", sc.StartScan)
	router.GET("/scanner/:id", sc.ReadScanFile)

}
