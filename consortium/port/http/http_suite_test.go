package porthttp_test

import (
	"context"
	"net/http"
	"net/http/httptest"
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

var _ = Describe("Check if the route is working normally", Label("PortHttp"), func() {

	var r *gin.Engine

	BeforeEach(func() {
		r = gin.Default()
		porthttp.ConfigureRoutes(r, getMapsRoutes())
	})

	When("Make a call to the endpoint: /consortium/offers/:id", func() {
		Context("and send an customerId that doesn't exist", func() {
			It("Should return a message with statusCode:400 (BadRequest)", func() {

				messageExpec := "{\"message\":\"customer not found\"}"
				path := "/consortium/offers/idTest"
				req, _ := http.NewRequest(http.MethodGet, path, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Expect(w).To(HaveHTTPBody(MatchJSON(messageExpec)), "the body should contain the value: %s", messageExpec)
				Expect(w).To(HaveHTTPStatus(http.StatusBadRequest))
			})
		})
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

type mockDb struct{}

func (mc mockDb) GetItem(ctx context.Context, customerId string) (offer.ConsortiumOffer, error) {
	return offer.ConsortiumOffer{}, nil
}
