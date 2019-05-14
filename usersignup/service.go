package usersignup

import (
	"context"
	"time"

	"github.com/salmankhalid187/user-signup/gen/restapi/operations/signup"

	"github.com/salmankhalid187/user-signup/gen/models"
)

// Service handles async log of audit event
type Service interface {
	SignupUser(ctx context.Context, params *signup.SignUpUserParams) (*models.SignupSuccess, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) SignupUser(ctx context.Context, params *signup.SignUpUserParams) (*models.SignupSuccess, error) {

	s.repo.SignupUser(ctx, params)
	t := time.Now()
	signUpSuccess := models.SignupSuccess{
		Name: t.String(),
	}

	return &signUpSuccess, nil
}
