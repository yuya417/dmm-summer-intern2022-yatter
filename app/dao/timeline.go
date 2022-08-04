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
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

// Create timeline repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

// AllStatuses : Statusを全て取得
func (r *timeline) AllStatuses(ctx context.Context, max_id int, since_id int, limit int) ([]*object.Status, error) {
    rows, err := r.db.QueryxContext(ctx, "select s.id, s.content, s.create_at, a.username as 'account.username', a.create_at as 'account.create_at' from status as s inner join account as a on s.account_id = a.id where s.id <= ? and s.id >= ? limit ?", max_id, since_id, limit)
    var timelines []*object.Status
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }

        return nil, fmt.Errorf("%w", err)
    }
    for rows.Next() {
        status := new(object.Status)
        err = rows.StructScan(&status)
        if err != nil {
            return nil, err
        }
        timelines = append(timelines, status)

    }
    return timelines, nil
}