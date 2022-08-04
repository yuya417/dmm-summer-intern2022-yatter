package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindBystatus : ユーザ名からユーザを取得
// func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
// 	entity := new(object.Account)
// 	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}

// 		return nil, fmt.Errorf("%w", err)
// 	}

// 	return entity, nil
// }

func (r *status) CreateStatus(ctx context.Context, account *object.Account, status *object.Status) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, "insert into status (account_id, content) value (?, ?)", account.ID, status.Content)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	status_id, err := result.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

    if err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", status_id).StructScan(status); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }

        return nil, fmt.Errorf("%w", err)
    }
    return status, nil
}
