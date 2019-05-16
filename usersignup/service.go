package usersignup

import (
	"context"

	"github.com/salmankhalid187/user-signup/gen/models"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations/signup"
)

// Service handles async log of audit event
type Service interface {
	// SignupUser(ctx context.Context, params *signup.SignUpUserParams) (*models.SignupSuccess, error)
	CreateUser(ctx context.Context, params *signup.CreateUserParams) (*models.SignupSuccess, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// func (s *service) SignupUser(ctx context.Context, params *signup.SignUpUserParams) (*models.SignupSuccess, error) {

// 	// s.repo.SignupUser(ctx, params)
// 	t := time.Now()
// 	signUpSuccess := models.SignupSuccess{
// 		Name: t.String(),
// 	}

// 	return &signUpSuccess, nil
// }

func (s *service) CreateUser(ctx context.Context, params *signup.CreateUserParams) (*models.SignupSuccess, error) {
	s.repo.SignupUser(ctx, params)
	signUpSuccess := models.SignupSuccess{
		Name: "User created successfully",
	}

	return &signUpSuccess, nil
}
