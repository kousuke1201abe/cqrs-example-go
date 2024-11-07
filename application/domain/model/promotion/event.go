package promotion

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	IsPromotionEvent()
}

type Submitted struct {
	AggregateID         uuid.UUID
	Name                string
	DiscountAmount      int
	SlotRemainingAmount int
	SubmittedAt         time.Time
}

func (Submitted) IsPromotionEvent() {}

type Published struct {
	AggregateID uuid.UUID
	PublishedAt time.Time
}

func (Published) IsPromotionEvent() {}

type Applied struct {
	AggregateID         uuid.UUID
	CustomerID          uuid.UUID
	SlotRemainingAmount int
	AppliedAt           time.Time
}

func (Applied) IsPromotionEvent() {}
