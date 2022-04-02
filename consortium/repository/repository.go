package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brunobmelo/consortium/offer"
)

type Repository struct {
	dbClient *dynamodb.Client
}

func (db Repository) GetItem(ctx context.Context, customerId string) (offer.ConsortiumOffer, error) {
	return offer.ConsortiumOffer{}, nil
}

func New(client *dynamodb.Client) Repository {
	return Repository{dbClient: client}
}
