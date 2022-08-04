package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindByStatusID(ctx context.Context, id string) (*object.Status, error)
	CreateStatus(ctx context.Context, account *object.Account, status *object.Status) (*object.Status, error)
}
