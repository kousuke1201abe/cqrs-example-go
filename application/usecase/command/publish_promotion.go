package command

import (
	"context"
	"time"

	"github.com/giftee/cqrs-example-go/application/domain/model/promotion"
	"github.com/google/uuid"
)

type PublishPromotionCommand struct {
	Repository promotion.Repository
}

func (c PublishPromotionCommand) Exec(ctx context.Context, promotionID uuid.UUID) (promotion.Published, error) {
	current := time.Now()

	promo, err := c.Repository.Find(ctx, promotionID)
	if err != nil {
		return promotion.Published{}, err
	}

	event, err := promo.Publish(current)
	if err != nil {
		return promotion.Published{}, err
	}

	if err := c.Repository.PersistEvent(ctx, event); err != nil {
		return promotion.Published{}, err
	}

	return event, nil
}
