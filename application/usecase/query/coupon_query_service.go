package query

import (
	"context"
	"time"
)

type CouponQueryService interface {
	QueryByCustomerID(context.Context, string) ([]Coupon, error)
}

type Coupon struct {
	ID             string
	DiscountAmount int
	ExpiredAt      time.Time
	Invalidated    bool
	Redeemed       bool
}
