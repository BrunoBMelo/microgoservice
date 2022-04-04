package porthttp

import (
	"github.com/brunobmelo/consortium/appconfig"
	"github.com/gin-gonic/gin"
)

func RunServer() {

	r := gin.Default()

	cfg := appconfig.LoadConfig()

	maps := GetMapRoutes(cfg)

	ConfigureRoutes(r, maps)

	r.Run(":8081")
}
