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

func (r *UserRepository) ListUsers(
	ctx context.Context,
) ([]sqlc.User, error) {

	return r.queries.ListUsers(ctx)
}
func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int32,
) error {

	return r.queries.DeleteUser(
		ctx,
		id,
	)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	params sqlc.UpdateUserParams,
) (sqlc.User, error) {

	return r.queries.UpdateUser(
		ctx,
		params,
	)
}