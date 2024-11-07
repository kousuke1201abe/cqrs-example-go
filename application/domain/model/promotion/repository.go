package promotion

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Find(context.Context, uuid.UUID) (*Promotion, error)
	PersistEvent(context.Context, Event) error
}
