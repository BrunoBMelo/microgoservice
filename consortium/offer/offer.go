package offer

import "context"

type ConsortiumOffer struct {
	CustomerId   string
	Available    float64
	PercentageAA float64
	QuotaMax     int
}

type IDatabase interface {
	GetItem(ctx context.Context, customerId string) (ConsortiumOffer, error)
}

func GetConsortiumOffer(ctx context.Context, db IDatabase, customerId string) (ConsortiumOffer, error) {
	result, err := db.GetItem(ctx, customerId)
	return result, err
}
