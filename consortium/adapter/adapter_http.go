package adapter

import (
	"net/http"

	"github.com/brunobmelo/consortium/offer"
	"github.com/gin-gonic/gin"
)

type OfferOuput struct {
	Available    float64 `json:"value"`
	PercentageAA float64 `json:"tax"`
	QuotaMax     int     `json:"quota_max"`
}

type OfferInput struct {
	CustomerId string `json:"id" uri:"id" binding:"required"`
}

type Di struct {
	DB offer.IDatabase
}

func GetConsortiumOffer(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		env, err := TransformTo[Di](v)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var reqInput OfferInput

		if err := c.ShouldBindUri(&reqInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := offer.GetConsortiumOffer(c.Request.Context(), env.DB, reqInput.CustomerId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if result == (offer.ConsortiumOffer{}) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "customer not found"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
