package database

import (
	"database/sql"
	"time"
)

type customerData struct {
	ID string `db:"id"`
}

type promotionData struct {
	ID                  string    `db:"id"`
	Name                string    `db:"name"`
	DiscountAmount      int       `db:"discount_amount"`
	SlotRemainingAmount int       `db:"slot_remaining_amount"`
	SubmittedAt         time.Time `db:"submitted_at"`
	ModifiedAt          time.Time `db:"modified_at"`
}

type promotionPublicationData struct {
	PromotionID sql.NullString `db:"promotion_id"`
	PublishedAt time.Time      `db:"published_at"`
}

type promotionApplicationData struct {
	PromotionID string    `db:"promotion_id"`
	CustomerID  string    `db:"customer_id"`
	AppliedAt   time.Time `db:"applied_at"`
}

type couponData struct {
	ID             string    `db:"id"`
	DiscountAmount int       `db:"discount_amount"`
	ExpiresAt      time.Time `db:"expires_at"`
	CustomerID     string    `db:"customer_id"`
	GrantedAt      time.Time `db:"granted_at"`
}

type couponInvalidationData struct {
	CouponID      sql.NullString `db:"coupon_id"`
	InvalidatedAt sql.NullTime   `db:"invalidated_at"`
}

type couponRedemptionData struct {
	CouponID   sql.NullString `db:"coupon_id"`
	RedeemedAt sql.NullTime   `db:"redeemed_at"`
}
