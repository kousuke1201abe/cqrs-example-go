package database

import (
	"context"

	"github.com/giftee/cqrs-example-go/application/usecase/query"
)

type CustomerQueryService struct {
	Conn DBConn
}

type customerQueryData struct {
	customerData `db:"customer"`
}

func (qs CustomerQueryService) QueryAll(ctx context.Context) ([]query.Customer, error) {
	var rows []customerQueryData

	const statement = `SELECT id AS "customer.id" FROM customers`

	if err := qs.Conn.db.SelectContext(ctx, &rows, statement); err != nil {
		return nil, err
	}

	customers := make([]query.Customer, 0, len(rows))
	for _, row := range rows {
		customers = append(customers, query.Customer{ID: row.ID})
	}

	return customers, nil
}
