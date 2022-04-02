package porthttp_test

import (
	"context"
	"testing"

	"github.com/brunobmelo/consortium/adapter"
	"github.com/brunobmelo/consortium/offer"
	porthttp "github.com/brunobmelo/consortium/port/http"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPortHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Http Suite")
}

var r *gin.Engine

var _ = Describe("", func() {

	BeforeEach(func() {
		startServer()
	})
})

func getMapsRoutes() []porthttp.MapRoute {
	return []porthttp.MapRoute{
		{
			HttpMethod:   "GET",
			RelativePath: "/consortium/offers/:id",
			HandlerFunc:  adapter.GetConsortiumOffer,
			IoC: func() interface{} {
				return adapter.Di{
					DB: mockDb{},
				}
			},
		},
	}
}

func startServer() {
	r = gin.Default()
	porthttp.ConfigureRoutes(r, getMapsRoutes())
	r.Run(":3000")
}

type mockDb struct{}

func (mc mockDb) GetItem(ctx context.Context, customerId string) (offer.ConsortiumOffer, error) {
	return offer.ConsortiumOffer{}, nil
}
