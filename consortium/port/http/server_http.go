package porthttp

import "github.com/gin-gonic/gin"

func StartServer() {

	r := gin.Default()

	maps := GetMapRoutes()

	ConfigureRoutes(r, maps)

	r.Run(":8080")
}
