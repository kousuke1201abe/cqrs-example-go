package command

import (
	"context"
	"time"

	"github.com/giftee/cqrs-example-go/application/domain/model/promotion"
)

type SubmitPromotionCommand struct {
	Repository promotion.Repository
}

func (c SubmitPromotionCommand) Exec(ctx context.Context, name promotion.Name, discount promotion.Discount, slot promotion.Slot) (promotion.Submitted, error) {
	current := time.Now()

	event, err := promotion.Submit(name, discount, slot, current)
	if err != nil {
		return promotion.Submitted{}, err
	}

	if err := c.Repository.PersistEvent(ctx, event); err != nil {
		return promotion.Submitted{}, err
	}

	return event, nil
}
