package offer

import (
	"context"
	"testing"
)

func TestOffer(t *testing.T) {

	for _, test := range testTable {

		mock := MockDb{}
		got, err := GetConsortiumOffer(context.Background(), mock, test.CustomerId)

		if err != nil {
			t.Errorf("Test Failed: The value return was err and it isnt expect -%s - testId: %d", err.Error(), test.TestId)
		}

		if got == (ConsortiumOffer{}) {
			t.Errorf("Test Failed: The value returned is nil and it isnt expected - %v - testId: %d", got, test.TestId)
		}
	}
}

var testTable = []struct {
	TestId     int
	CustomerId string
	Entity     ConsortiumOffer
}{
	{
		TestId:     1,
		CustomerId: "sxhgsedoasnashudhen",
		Entity: ConsortiumOffer{
			CustomerId:   "sxhgsedoasnashudhen",
			Available:    "1300.00",
			Tax: "2",
			Quota:     36,
		},
	},
}

type MockDb struct{}

func (d MockDb) GetItem(ctx context.Context, customerId string) (ConsortiumOffer, error) {
	return ConsortiumOffer{
		CustomerId:   "sxhgsedoasnashudhen",
		Available:    "1300.00",
		Tax: "2",
		Quota:     36,
	}, nil
}
