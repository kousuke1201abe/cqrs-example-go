package promotion

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Promotion struct {
	id                 uuid.UUID
	published          bool
	slot               Slot
	appliedCustomerIDs []uuid.UUID
}

func Submit(name Name, discount Discount, slot Slot, submitAt time.Time) (Submitted, error) {
	aggregateID, err := uuid.NewRandom()
	if err != nil {
		return Submitted{}, err
	}

	return Submitted{AggregateID: aggregateID, Name: name.value, DiscountAmount: discount.amount, SlotRemainingAmount: slot.remainingAmount, SubmittedAt: submitAt}, nil
}

func (p *Promotion) Publish(publishAt time.Time) (Published, error) {
	if p.published {
		return Published{}, errors.New("promotion is already published")
	}

	return Published{AggregateID: p.id, PublishedAt: publishAt}, nil
}

func (p *Promotion) Apply(customerID uuid.UUID, applyAt time.Time) (Applied, error) {
	if !p.published {
		return Applied{}, errors.New("promotion is not published")
	}

	if p.slot.isFull() {
		return Applied{}, errors.New("slot is full")
	}

	for _, appliedCustomerID := range p.appliedCustomerIDs {
		if appliedCustomerID == customerID {
			return Applied{}, errors.New("promotion has already been applied")
		}
	}

	return Applied{AggregateID: p.id, CustomerID: customerID, SlotRemainingAmount: p.slot.remainingAmount - 1, AppliedAt: applyAt}, nil
}

func Reconstruct(id uuid.UUID, published bool, slot Slot, appliedCustomerIDs []uuid.UUID) (*Promotion, error) {
	return &Promotion{id: id, published: published, slot: slot, appliedCustomerIDs: appliedCustomerIDs}, nil
}
