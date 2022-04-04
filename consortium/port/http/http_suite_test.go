package porthttp_test

import (
	"context"
	"errors"
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
		Context("and send an customerId to get the offer", func() {
			It("Should return a message with statusCode:400 (BadRequest)", func() {

				messageExpec := "{\"message\":\"customer not found\"}"
				path := "/consortium/offers/idTest"
				req, _ := http.NewRequest(http.MethodGet, path, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Expect(w).To(HaveHTTPBody(MatchJSON(messageExpec)), "the body should contain the value: %s", messageExpec)
				Expect(w).To(HaveHTTPStatus(http.StatusBadRequest), "the status should be 400 - Bad Request")
			})

			It("Should return statusCode: 200 (Successed)", func() {

				jsonExpec := "{\"CustomerId\":\"39819584-50b3-45ee-a4e9-ad4d3607b167\",\"Available\": 13000,\"PercentageAA\": 0.02,\"QuotaMax\": 36}"
				path := "/consortium/offers/39819584-50b3-45ee-a4e9-ad4d3607b167"
				req, _ := http.NewRequest(http.MethodGet, path, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Expect(w).To(HaveHTTPBody(MatchJSON(jsonExpec)), "the body should contain the value: %s", jsonExpec)
				Expect(w).To(HaveHTTPStatus(http.StatusOK), "the status should be 200 - Bad Request")
			})

			It("Should return statusCode: 500 (Internal Server Error)", func() {

				path := "/consortium/offers/398195t4-50b3-45ee-a4e9-ad4d3607b167"
				req, _ := http.NewRequest(http.MethodGet, path, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Expect(w).To(HaveHTTPStatus(http.StatusInternalServerError), "the status should be 500 - Internal Server Error")
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

	if customerId == "39819584-50b3-45ee-a4e9-ad4d3607b167" {
		return offer.ConsortiumOffer{
			CustomerId:   "39819584-50b3-45ee-a4e9-ad4d3607b167",
			Available:    13000.00,
			PercentageAA: 0.02,
			QuotaMax:     36,
		}, nil
	}

	if customerId == "398195t4-50b3-45ee-a4e9-ad4d3607b167" {
		return offer.ConsortiumOffer{}, errors.New("error")
	}

	return offer.ConsortiumOffer{}, nil
}
