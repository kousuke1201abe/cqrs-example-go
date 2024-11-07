package query

import "context"

type CustomerQueryService interface {
	QueryAll(context.Context) ([]Customer, error)
}

type Customer struct {
	ID string
}
