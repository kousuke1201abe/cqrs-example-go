package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/giftee/cqrs-example-go/application/domain/model/coupon"
	"github.com/google/uuid"
)

type CouponRepository struct {
	Conn DBConn
}

func (r CouponRepository) PersistEvent(ctx context.Context, event coupon.Event) error {
	switch e := event.(type) {
	case coupon.Granted:
		return r.persistCouponGranted(ctx, e)
	case coupon.Invalidated:
		return r.persistCouponInvalidated(ctx, e)
	default:
		return errors.New("invalid event")
	}
}

func (r CouponRepository) persistCouponGranted(ctx context.Context, event coupon.Granted) error {
	data := r.serializeGranted(event)

	if _, err := r.Conn.db.NamedExecContext(ctx, `INSERT INTO coupons (id, customer_id, discount_amount, expires_at, granted_at) VALUES (:id, :customer_id, :discount_amount, :expires_at, :granted_at);`, data.couponData); err != nil {
		return err
	}

	return nil
}

func (r CouponRepository) persistCouponInvalidated(ctx context.Context, event coupon.Invalidated) error {
	data := r.serializeInvalidated(event)

	if _, err := r.Conn.db.NamedExecContext(ctx, `INSERT INTO coupon_invalidations (coupon_id, invalidated_at) VALUES (:coupon_id, :invalidated_at);`, data.couponInvalidationData); err != nil {
		return err
	}

	return nil
}

func (r CouponRepository) Find(ctx context.Context, id uuid.UUID) (*coupon.Coupon, error) {
	var data couponAggregateData

	if err := r.Conn.db.GetContext(ctx, &data.couponData, `SELECT id FROM coupons WHERE id = ?`, id.String()); err != nil {
		return nil, err
	}

	if err := r.Conn.db.GetContext(ctx, &data.invalidated, `SELECT EXISTS( SELECT * FROM coupon_invalidations WHERE coupon_id = ?)`, id.String()); err != nil {
		return nil, err
	}

	return r.deserialize(data)
}

type couponGrantedData struct {
	couponData
}

func (r CouponRepository) serializeGranted(event coupon.Granted) couponGrantedData {
	return couponGrantedData{
		couponData: couponData{
			ID:             event.AggregateID.String(),
			DiscountAmount: event.DiscountAmount,
			ExpiresAt:      event.ExpiresAt,
			CustomerID:     event.CustomerID.String(),
			GrantedAt:      event.GrantedAt,
		},
	}
}

type couponInvalidatedData struct {
	couponInvalidationData
}

func (r CouponRepository) serializeInvalidated(event coupon.Invalidated) couponInvalidatedData {
	return couponInvalidatedData{
		couponInvalidationData: couponInvalidationData{
			CouponID:      sql.NullString{String: event.AggregateID.String(), Valid: true},
			InvalidatedAt: sql.NullTime{Time: event.InvalidatedAt, Valid: true},
		},
	}
}

type couponAggregateData struct {
	couponData
	invalidated bool
}

func (r CouponRepository) deserialize(data couponAggregateData) (*coupon.Coupon, error) {
	aggregateID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, err
	}

	return coupon.Reconstruct(aggregateID, data.invalidated)
}
