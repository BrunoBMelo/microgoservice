package offer

import "context"

type ConsortiumOffer struct {
	CustomerId string `json:"customerid"`
	Available  string `json:"available"`
	Tax        string `json:"tax"`
	Quota      int    `json:"quota"`
}

type IDatabase interface {
	GetItem(ctx context.Context, customerId string) (ConsortiumOffer, error)
}

func GetConsortiumOffer(ctx context.Context, db IDatabase, customerId string) (ConsortiumOffer, error) {
	result, err := db.GetItem(ctx, customerId)
	return result, err
}
