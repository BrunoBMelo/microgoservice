package offer

import (
	"net/http"

	"github.com/BrunoBMelo/handlerhttp"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/BrunoBMelo/appconfig"
)

func GetMapRoutes(cfg appconfig.Config) []handlerhttp.MapRoute {

	return []handlerhttp.MapRoute{
		{
			HttpMethod:   http.MethodGet,
			RelativePath: "/consortium/offers/:id",
			HandlerFunc:  ConsortiumOffer,
			IoC: func() interface{} {
				return Di{
					DB: NewFromRepository(dynamodb.NewFromConfig(*cfg.AwsConfig)),
				}
			},
		},
	}
}
