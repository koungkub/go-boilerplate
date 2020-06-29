package worker

import (
	"context"
	"database/sql"
)

type Reviewer interface {
	Get(ctx context.Context, id string)
}

func NewReview(db *sql.DB) Reviewer {
	return &Review{
		DB: db,
	}
}

func (r *Review) Get(ctx context.Context, id string) {

}
