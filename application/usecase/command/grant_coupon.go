package command

import (
	"context"
	"time"

	"github.com/giftee/cqrs-example-go/application/domain/model/coupon"
	"github.com/google/uuid"
)

type GrantCouponCommand struct {
	Repository coupon.Repository
}

func (c GrantCouponCommand) Exec(ctx context.Context, customerID uuid.UUID, discount coupon.Discount) (coupon.Granted, error) {
	current := time.Now()

	event, err := coupon.Grant(customerID, discount, current)
	if err != nil {
		return coupon.Granted{}, err
	}

	if err := c.Repository.PersistEvent(ctx, event); err != nil {
		return coupon.Granted{}, err
	}

	return event, nil
}
