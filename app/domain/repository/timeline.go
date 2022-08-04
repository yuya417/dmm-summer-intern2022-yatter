package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	AllStatuses(ctx context.Context, max_id int, since_id int, limit int) ([]*object.Status, error)
}
