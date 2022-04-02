package porthttp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brunobmelo/consortium/adapter"
	"github.com/brunobmelo/consortium/repository"
	"github.com/gin-gonic/gin"
)

type MapRoute struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  func(a interface{}) gin.HandlerFunc
	IoC          func() interface{}
}

func GetMapRoutes() []MapRoute {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return []MapRoute{
		{
			HttpMethod:   "GET",
			RelativePath: "/consortium/offers/:id",
			HandlerFunc:  adapter.GetConsortiumOffer,
			IoC: func() interface{} {
				return adapter.Di{
					DB: repository.New(dynamodb.NewFromConfig(cfg)),
				}
			},
		},
	}
}
