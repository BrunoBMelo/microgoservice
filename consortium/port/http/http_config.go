package porthttp

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, maps []MapRoute) {
	for _, mr := range maps {
		if mr.HttpMethod == "GET" {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}
	}
}
