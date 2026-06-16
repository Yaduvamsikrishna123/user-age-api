package service

import (
	"context"
	"time"

	"github.com/yaduvamsi/user-age-api/internal/models"
	"github.com/yaduvamsi/user-age-api/internal/repository"
	"github.com/yaduvamsi/user-age-api/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"

)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func calculateAge(dob time.Time) int {

	now := time.Now()

	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int32,
) (*models.UserResponse, error) {

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  calculateAge(user.Dob.Time),
	}, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	response := make(
		[]models.UserResponse,
		0,
		len(users),
	)

	for _, user := range users {

		response = append(
			response,
			models.UserResponse{
				ID:   user.ID,
				Name: user.Name,
				DOB:  user.Dob.Time.Format("2006-01-02"),
				Age:  calculateAge(user.Dob.Time),
			},
		)
	}

	return response, nil
}

func (s *UserService) CreateUser(
	ctx context.Context,
	req models.CreateUserRequest,
) (*models.UserResponse, error) {

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Name: req.Name,
			Dob: pgtype.Date{
				Time:  dob,
				Valid: true,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  calculateAge(user.Dob.Time),
	}, nil
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {

	return s.repo.DeleteUser(
		ctx,
		id,
	)
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	req models.UpdateUserRequest,
) (*models.UserResponse, error) {

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.UpdateUser(
		ctx,
		sqlc.UpdateUserParams{
			ID: id,
			Name: req.Name,
			Dob: pgtype.Date{
				Time:  dob,
				Valid: true,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  calculateAge(user.Dob.Time),
	}, nil
}
func (s *UserService) ListUsersPaginated(
	ctx context.Context,
	page int,
	limit int,
) ([]models.UserResponse, error) {

	offset := (page - 1) * limit

	users, err := s.repo.ListUsersPaginated(
		ctx,
		int32(limit),
		int32(offset),
	)

	if err != nil {
		return nil, err
	}

	response := make(
		[]models.UserResponse,
		0,
		len(users),
	)

	for _, user := range users {

		response = append(
			response,
			models.UserResponse{
				ID:   user.ID,
				Name: user.Name,
				DOB:  user.Dob.Time.Format("2006-01-02"),
				Age:  calculateAge(user.Dob.Time),
			},
		)
	}

	return response, nil
}