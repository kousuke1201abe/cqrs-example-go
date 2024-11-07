// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type ApplyPromotionInput struct {
	CustomerID  string `json:"customerId"`
	PromotionID string `json:"promotionId"`
}

type ApplyPromotionPayload struct {
	PromotionID string `json:"promotionId"`
}

type Coupon struct {
	DiscountAmount int       `json:"discountAmount"`
	ExpiresAt      time.Time `json:"expiresAt"`
	ID             string    `json:"id"`
	Invalidated    bool      `json:"invalidated"`
	Redeemed       bool      `json:"redeemed"`
}

type Customer struct {
	ID string `json:"id"`
}

type GrantCouponInput struct {
	CustomerID     string `json:"customerId"`
	DiscountAmount int    `json:"discountAmount"`
}

type GrantCouponPayload struct {
	CouponID string `json:"couponId"`
}

type InvalidateCouponInput struct {
	CouponID string `json:"couponId"`
}

type InvalidateCouponPayload struct {
	CouponID string `json:"couponId"`
}

type Mutation struct {
}

type Promotion struct {
	AppliedCustomerNumber int    `json:"appliedCustomerNumber"`
	DiscountAmount        int    `json:"discountAmount"`
	SlotRemainingAmount   int    `json:"slotRemainingAmount"`
	Published             bool   `json:"published"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
}

type PublishPromotionInput struct {
	PromotionID string `json:"promotionId"`
}

type PublishPromotionPayload struct {
	PromotionID string `json:"promotionId"`
}

type Query struct {
}

type SubmitPromotionInput struct {
	DiscountAmount      int    `json:"discountAmount"`
	SlotRemainingAmount int    `json:"slotRemainingAmount"`
	Name                string `json:"name"`
}

type SubmitPromotionPayload struct {
	PromotionID string `json:"promotionId"`
}