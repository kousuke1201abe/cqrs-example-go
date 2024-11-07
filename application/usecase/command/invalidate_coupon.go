package command

import (
	"context"
	"time"

	"github.com/giftee/cqrs-example-go/application/domain/model/coupon"
	"github.com/google/uuid"
)

type InvalidateCouponCommand struct {
	Repository coupon.Repository
}

func (c InvalidateCouponCommand) Exec(ctx context.Context, couponID uuid.UUID) (coupon.Invalidated, error) {
	current := time.Now()

	cpn, err := c.Repository.Find(ctx, couponID)
	if err != nil {
		return coupon.Invalidated{}, err
	}

	event, err := cpn.Invalidate(current)
	if err != nil {
		return coupon.Invalidated{}, err
	}

	if err := c.Repository.PersistEvent(ctx, event); err != nil {
		return coupon.Invalidated{}, err
	}

	return event, nil
}
