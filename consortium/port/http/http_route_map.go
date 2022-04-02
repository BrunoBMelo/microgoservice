package porthttp

import (
	"github.com/brunobmelo/consortium/adapter"
	"github.com/gin-gonic/gin"
)

type MapRoute struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  func(a interface{}) gin.HandlerFunc
	IoC          func() interface{}
}

var mapRoutes = []MapRoute{
	{
		HttpMethod:   "GET",
		RelativePath: "/consortium/offers/:id",
		HandlerFunc:  adapter.GetConsortiumOffer,
		IoC: func() interface{} {
			return adapter.Di{
				DB: nil,
			}
		},
	},
}

func GetMapRoutes() []MapRoute {
	return mapRoutes
}
