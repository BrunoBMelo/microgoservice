package main

import (
	"github.com/BrunoBMelo/appconfig"
	"github.com/BrunoBMelo/handlerhttp"
	"github.com/brunobmelo/financing/offer"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var cfg appconfig.Config

func init() {

	cfg = appconfig.LoadConfig()
	r = handlerhttp.New()
	handlerhttp.ConfigureMapRoute(func() []handlerhttp.MapRoute {
		return offer.GetMapRoutes(cfg)
	})

}

func main() {

	r.Run(":" + cfg.PortApp)

}
