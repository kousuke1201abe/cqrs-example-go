package query

import (
	"context"
)

type PromotionQueryService interface {
	QueryAll(context.Context) ([]Promotion, error)
}

type Promotion struct {
	ID                    string
	Name                  string
	Published             bool
	DiscountAmount        int
	SlotRemainingAmount   int
	AppliedCustomerNumber int
}
