package database

import (
	"context"

	"github.com/giftee/cqrs-example-go/application/usecase/query"
)

type CouponQueryService struct {
	Conn DBConn
}

type couponQueryData struct {
	couponData             `db:"coupon"`
	couponInvalidationData `db:"coupon_invalidation"`
	couponRedemptionData   `db:"coupon_redemption"`
}

func (qs CouponQueryService) QueryByCustomerID(ctx context.Context, customerID string) ([]query.Coupon, error) {
	var rows []couponQueryData

	const statement = `
		SELECT
			id AS "coupon.id",
			discount_amount AS "coupon.discount_amount",
			expires_at AS "coupon.expires_at",
			coupon_invalidations.coupon_id AS "coupon_invalidation.coupon_id",
			coupon_redemptions.coupon_id AS "coupon_redemption.coupon_id"
		FROM coupons
		LEFT JOIN coupon_invalidations ON coupons.id = coupon_invalidations.coupon_id
		LEFT JOIN coupon_redemptions ON coupons.id = coupon_redemptions.coupon_id
		WHERE customer_id = ?
	`

	if err := qs.Conn.db.SelectContext(ctx, &rows, statement, customerID); err != nil {
		return nil, err
	}

	coupons := make([]query.Coupon, 0, len(rows))
	for _, row := range rows {
		coupons = append(coupons, query.Coupon{
			ID:             row.ID,
			DiscountAmount: row.DiscountAmount,
			ExpiredAt:      row.ExpiresAt,
			Invalidated:    row.couponInvalidationData.CouponID.Valid,
			Redeemed:       row.couponRedemptionData.CouponID.Valid,
		})
	}

	return coupons, nil
}
