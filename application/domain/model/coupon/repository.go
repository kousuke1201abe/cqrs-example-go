package coupon

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Find(context.Context, uuid.UUID) (*Coupon, error)
	PersistEvent(context.Context, Event) error
}
