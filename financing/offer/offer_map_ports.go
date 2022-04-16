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
			RelativePath: "/financing/offers/:id",
			HandlerFunc:  GetConsortiumOffer,
			IoC: func() interface{} {
				return Dependency{
					Database: NewFromRepository(dynamodb.NewFromConfig(*cfg.AwsConfig)),
				}
			},
		},
	}
}
