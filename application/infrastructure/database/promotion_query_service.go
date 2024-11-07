package database

import (
	"context"

	"github.com/giftee/cqrs-example-go/application/usecase/query"
)

type PromotionQueryService struct {
	Conn DBConn
}

type promotionQueryData struct {
	promotionData            `db:"promotion"`
	promotionPublicationData `db:"promotion_publication"`
	AppliedCustomerNumber    int `db:"applied_customer_number"`
}

func (qs PromotionQueryService) QueryAll(ctx context.Context) ([]query.Promotion, error) {
	var rows []promotionQueryData

	const statement = `
		SELECT
			id AS "promotion.id",
			name AS "promotion.name",
			discount_amount AS "promotion.discount_amount",
			slot_remaining_amount AS "promotion.slot_remaining_amount",
			count(promotion_applications.customer_id) AS "applied_customer_number",
			promotion_publications.promotion_id AS "promotion_publication.promotion_id"
		FROM promotions
		LEFT JOIN promotion_applications ON promotions.id = promotion_applications.promotion_id
		LEFT JOIN promotion_publications ON promotions.id = promotion_publications.promotion_id
		GROUP BY promotions.id
	`

	if err := qs.Conn.db.SelectContext(ctx, &rows, statement); err != nil {
		return nil, err
	}

	promotions := make([]query.Promotion, 0, len(rows))
	for _, row := range rows {
		var published bool
		if row.promotionPublicationData.PromotionID.Valid {
			published = true
		}

		promotions = append(promotions, query.Promotion{
			ID:                    row.ID,
			Name:                  row.Name,
			Published:             published,
			DiscountAmount:        row.DiscountAmount,
			SlotRemainingAmount:   row.SlotRemainingAmount,
			AppliedCustomerNumber: row.AppliedCustomerNumber,
		})
	}

	return promotions, nil
}
