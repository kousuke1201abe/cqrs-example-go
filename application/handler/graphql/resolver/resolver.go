package resolver

import (
	"github.com/giftee/cqrs-example-go/application/usecase/command"
	"github.com/giftee/cqrs-example-go/application/usecase/query"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SubmitPromotionCommand  command.SubmitPromotionCommand
	PublishPromotionCommand command.PublishPromotionCommand
	ApplyPromotionCommand   command.ApplyPromotionCommand
	GrantCouponCommand      command.GrantCouponCommand
	InvalidateCouponCommand command.InvalidateCouponCommand
	CouponQueryService      query.CouponQueryService
	PromotionQueryService   query.PromotionQueryService
	CustomerQueryService    query.CustomerQueryService
}
