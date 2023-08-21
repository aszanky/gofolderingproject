package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aszanky/gofolderingproject/internal/model"
	"github.com/opentracing/opentracing-go"
)

func (r *repository) GetUsers(ctx context.Context, username string) (data model.User, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository_GetUser")
	defer span.Finish()

	//Check if user is already exist
	err = r.db.QueryRowxContext(ctx, queryGetUser, username).StructScan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, errors.New("user not found")
		}
		return
	}
	return
}
