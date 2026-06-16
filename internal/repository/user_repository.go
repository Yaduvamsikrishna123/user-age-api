package repository

import (
	"context"

	"github.com/yaduvamsi/user-age-api/db/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	params sqlc.CreateUserParams,
) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, params)
}