package command

import (
	"context"
	"time"

	"github.com/giftee/cqrs-example-go/application/domain/model/promotion"
	"github.com/google/uuid"
)

type ApplyPromotionCommand struct {
	Repository promotion.Repository
}

func (c ApplyPromotionCommand) Exec(ctx context.Context, promotionID uuid.UUID, customerID uuid.UUID) (promotion.Applied, error) {
	current := time.Now()

	promo, err := c.Repository.Find(ctx, promotionID)
	if err != nil {
		return promotion.Applied{}, err
	}

	event, err := promo.Apply(customerID, current)
	if err != nil {
		return promotion.Applied{}, err
	}

	if err := c.Repository.PersistEvent(ctx, event); err != nil {
		return promotion.Applied{}, err
	}

	return event, nil
}
