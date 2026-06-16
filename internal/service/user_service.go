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

func NewUserService(repo *repository.UserRepository) *UserService {
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

	age := calculateAge(user.Dob.Time)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *UserService) CreateUser(
	ctx context.Context,
	req models.CreateUserRequest,
) error {

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return err
	}

	_, err = s.repo.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Name: req.Name,
			Dob: pgtype.Date{
				Time:  dob,
				Valid: true,
			},
		},
	)

	return err
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