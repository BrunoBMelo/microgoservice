package adapter

import (
	"context"
	"testing"

	"github.com/brunobmelo/consortium/offer"
)

var testTable = []struct {
	From   interface{}
	To     Di
	TestId int
}{
	{
		TestId: 1,
		From: Di{
			DB: mockStruct{},
		},
		To: Di{
			DB: mockStruct{},
		},
	},
}

func TestAdapterTransformationFunc(t *testing.T) {

	func() {
		for _, testTable := range testTable {

			got, err := TransformTo[Di](testTable.From)

			if err != nil {
				t.Error("Failed: The convert interface in struct fail.")
			}
			if got != testTable.To {
				t.Error("Failed: The value return is different of the expected.")
			}
			if got.DB == nil {
				t.Error("Failed: The DB value loaded should to have value.")
			}
			result, err := got.DB.GetItem(context.TODO(), "")
			if err != nil {
				t.Error("Failed: The err mock return not sholud has value.")
			}
			if result != (offer.ConsortiumOffer{}) {
				t.Error("Failed: The result mock return not sholud has value.")
			}
		}
	}()

	func() {
		_, err := TransformTo[Di]("foo")

		if err == nil {
			t.Error("Failed: Is expect an err when pass a arg different the T.")
		}
	}()

}

type mockStruct struct{}

func (c mockStruct) GetItem(ctx context.Context, customerId string) (offer.ConsortiumOffer, error) {
	return offer.ConsortiumOffer{}, nil
}
