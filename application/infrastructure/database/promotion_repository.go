package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/giftee/cqrs-example-go/application/domain/model/promotion"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PromotionRepository struct {
	Conn DBConn
}

func (r PromotionRepository) PersistEvent(ctx context.Context, event promotion.Event) error {
	switch e := event.(type) {
	case promotion.Submitted:
		return r.persistPromotionSubmitted(ctx, e)
	case promotion.Published:
		return r.persistPromotionPublished(ctx, e)
	case promotion.Applied:
		return r.persistPromotionApplied(ctx, e)
	default:
		return errors.New("invalid event")
	}
}

func (r PromotionRepository) persistPromotionSubmitted(ctx context.Context, event promotion.Submitted) error {
	data := r.serializeSubmitted(event)

	if _, err := r.Conn.db.NamedExecContext(ctx, `INSERT INTO promotions (id, name, discount_amount, slot_remaining_amount, submitted_at, modified_at) VALUES (:id, :name, :discount_amount, :slot_remaining_amount, :submitted_at, :modified_at);`, data.promotionData); err != nil {
		return err
	}

	return nil
}

func (r PromotionRepository) persistPromotionPublished(ctx context.Context, event promotion.Published) error {
	data := r.serializePublished(event)

	if _, err := r.Conn.db.NamedExecContext(ctx, `INSERT INTO promotion_publications (promotion_id, published_at) VALUES (:promotion_id, :published_at);`, data.promotionPublicationData); err != nil {
		return err
	}

	return nil
}

func (r PromotionRepository) persistPromotionApplied(ctx context.Context, event promotion.Applied) error {
	data := r.serializeApplied(event)

	r.Conn.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		if _, err := tx.NamedExec(`UPDATE promotions SET slot_remaining_amount = :slot_remaining_amount where id = :id AND slot_remaining_amount = :slot_remaining_amount + 1;`, data.promotionData); err != nil {
			return err
		}

		if _, err := tx.NamedExec(`INSERT INTO promotion_applications (promotion_id, customer_id, applied_at) VALUES (:promotion_id, :customer_id, :applied_at);`, data.promotionApplicationData); err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (r PromotionRepository) Find(ctx context.Context, id uuid.UUID) (*promotion.Promotion, error) {
	var data promotionAggregateData

	if err := r.Conn.db.GetContext(ctx, &data, `
		SELECT id AS "promotion.id", slot_remaining_amount AS "promotion.slot_remaining_amount", promotion_publications.promotion_id AS "promotion_publication.promotion_id"
		FROM promotions
		LEFT JOIN promotion_publications ON promotions.id = promotion_publications.promotion_id
		WHERE promotions.id = ?
	`, id.String()); err != nil {
		return nil, err
	}

	if err := r.Conn.db.SelectContext(ctx, &data.appliedCustomers, `
		SELECT id
		FROM customers
		LEFT JOIN promotion_applications ON promotion_applications.customer_id = customers.id
		WHERE promotion_applications.promotion_id = ?
	`, id.String()); err != nil {
		return nil, err
	}

	return r.deserialize(data)
}

type promotionSubmittedData struct {
	promotionData
}

func (r PromotionRepository) serializeSubmitted(event promotion.Submitted) promotionSubmittedData {
	return promotionSubmittedData{
		promotionData: promotionData{
			ID:                  event.AggregateID.String(),
			Name:                event.Name,
			DiscountAmount:      event.DiscountAmount,
			SlotRemainingAmount: event.SlotRemainingAmount,
			SubmittedAt:         event.SubmittedAt,
			ModifiedAt:          event.SubmittedAt,
		},
	}
}

type promotionPublishedData struct {
	promotionPublicationData
}

func (r PromotionRepository) serializePublished(event promotion.Published) promotionPublishedData {
	return promotionPublishedData{
		promotionPublicationData: promotionPublicationData{
			PromotionID: sql.NullString{String: event.AggregateID.String(), Valid: true},
			PublishedAt: event.PublishedAt,
		},
	}
}

type promotionAppliedData struct {
	promotionData
	promotionApplicationData
}

func (r PromotionRepository) serializeApplied(event promotion.Applied) promotionAppliedData {
	return promotionAppliedData{
		promotionData: promotionData{
			ID:                  event.AggregateID.String(),
			SlotRemainingAmount: event.SlotRemainingAmount,
		},
		promotionApplicationData: promotionApplicationData{
			PromotionID: event.AggregateID.String(),
			CustomerID:  event.CustomerID.String(),
			AppliedAt:   event.AppliedAt,
		},
	}
}

type promotionAggregateData struct {
	promotionData            `db:"promotion"`
	promotionPublicationData `db:"promotion_publication"`
	appliedCustomers         []customerData
}

func (r PromotionRepository) deserialize(data promotionAggregateData) (*promotion.Promotion, error) {
	aggregateID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, err
	}

	slot, err := promotion.NewSlot(data.SlotRemainingAmount)
	if err != nil {
		return nil, err
	}

	var published bool
	if data.promotionPublicationData.PromotionID.Valid {
		published = true
	}

	appliedCustomerIDs := make([]uuid.UUID, 0, len(data.appliedCustomers))
	for _, appliedCustomer := range data.appliedCustomers {
		appliedCustomerID, err := uuid.Parse(appliedCustomer.ID)
		if err != nil {
			return nil, err
		}

		appliedCustomerIDs = append(appliedCustomerIDs, appliedCustomerID)
	}

	return promotion.Reconstruct(aggregateID, published, slot, appliedCustomerIDs)
}
