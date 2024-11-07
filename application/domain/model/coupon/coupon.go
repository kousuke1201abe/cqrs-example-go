package coupon

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	id          uuid.UUID
	invalidated bool
}

func Grant(customerID uuid.UUID, discount Discount, grantAt time.Time) (Granted, error) {
	expiresAt := grantAt.AddDate(1, 0, 0)

	aggregateID, err := uuid.NewRandom()
	if err != nil {
		return Granted{}, err
	}

	return Granted{AggregateID: aggregateID, CustomerID: customerID, DiscountAmount: discount.amount, ExpiresAt: expiresAt, GrantedAt: grantAt}, nil
}

func (c *Coupon) Invalidate(invalidateAt time.Time) (Invalidated, error) {
	if c.invalidated {
		return Invalidated{}, errors.New("coupon has been already invalidated")
	}

	return Invalidated{AggregateID: c.id, InvalidatedAt: invalidateAt}, nil
}

func Reconstruct(id uuid.UUID, invalidated bool) (*Coupon, error) {
	return &Coupon{id: id, invalidated: invalidated}, nil
}
