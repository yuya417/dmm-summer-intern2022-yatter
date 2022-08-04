package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch account which has specified username
	//FindByUsername(ctx context.Context, username string) (*object.Account, error)
	CreateStatus(ctx context.Context, account *object.Account, status *object.Status) (*object.Status, error)
}
