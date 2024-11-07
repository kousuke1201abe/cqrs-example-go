package coupon

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	IsCouponEvent()
}

type Granted struct {
	AggregateID    uuid.UUID
	CustomerID     uuid.UUID
	DiscountAmount int
	ExpiresAt      time.Time
	GrantedAt      time.Time
}

func (Granted) IsCouponEvent() {}

type Invalidated struct {
	AggregateID   uuid.UUID
	InvalidatedAt time.Time
}

func (Invalidated) IsCouponEvent() {}
