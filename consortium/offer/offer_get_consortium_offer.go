package offer

import (
	"context"
	"net/http"

	"github.com/BrunoBMelo/converter"
	"github.com/gin-gonic/gin"
)

func GetConsortiumOffer(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		di, err := converter.To[Dependency](v)
		if err != nil {
			setResponseJSON(c, http.StatusInternalServerError, ResponseMessage{"message": err.Error()})
			return
		}

		var input OfferInput
		if err = c.ShouldBindUri(&input); err != nil {
			setResponseJSON(c, http.StatusBadRequest, ResponseMessage{"message": err.Error()})
			return
		}

		var outputResult OfferModel
		outputResult, err = di.Database.GetItem(c.Request.Context(), input.CustomerId)
		if err != nil {
			setResponseJSON(c, http.StatusInternalServerError, ResponseMessage{"message": err.Error()})
			return
		}

		if outputResult == (OfferModel{}) {
			setResponseJSON(c, http.StatusBadRequest, ResponseMessage{"message": "customer not found"})
			return
		}

		setResponseJSON(c, http.StatusOK, outputResult)
	}
}

func setResponseJSON(c *gin.Context, statusCode int, obj interface{}) {
	c.JSON(statusCode, obj)
}

type ResponseMessage map[string]interface{}

type OfferOutput struct {
	Available    float64 `json:"value"`
	PercentageAA float64 `json:"tax"`
	QuotaMax     int     `json:"quota_max"`
}

type OfferInput struct {
	CustomerId string `json:"id" uri:"id" binding:"required"`
}

type Dependency struct {
	Database IDatabase
}

type OfferModel struct {
	CustomerId string `json:"customerid"`
	Available  string `json:"available"`
	Tax        string `json:"tax"`
	Quota      int    `json:"quota"`
}

type IDatabase interface {
	GetItem(ctx context.Context, customerId string) (OfferModel, error)
}
