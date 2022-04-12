package offer

import (
	"net/http"

	"github.com/BrunoBMelo/converter"
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
	DB IDatabase
}

func ConsortiumOffer(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		env, err := converter.To[Di](v)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var reqInput OfferInput
		if err := c.ShouldBindUri(&reqInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := GetConsortiumOffer(c.Request.Context(), env.DB, reqInput.CustomerId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if result == (Offer{}) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "customer not found"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
