package offer

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository struct {
	dbClient *dynamodb.Client
}

func (db Repository) GetItem(ctx context.Context, customerId string) (OfferModel, error) {

	result, err := db.dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"customerid": &types.AttributeValueMemberS{Value: customerId},
		},
		TableName: aws.String("consortium-offers"),
	})

	if err != nil {
		return OfferModel{}, err
	}

	if len(result.Item) == 0 {
		return OfferModel{}, errors.New("customer-id not found")
	}

	consortiumOffer := OfferModel{}
	err = attributevalue.UnmarshalMap(result.Item, &consortiumOffer)
	if err != nil {
		return OfferModel{}, err
	}

	return consortiumOffer, nil
}

func NewFromRepository(client *dynamodb.Client) Repository {
	return Repository{dbClient: client}
}
