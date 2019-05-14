package usersignup

import (
	"context"
)

type Repository interface {
	SignupUser(ctx context.Context) (bool, error)
}

type repository struct {
	temp bool
}

// NewRepository pass db reference coming from orm here
func NewRepository() Repository {
	return &repository{
		temp: true,
	}
}

func (r *repository) SignupUser(ctx context.Context) (bool, error) {
	return true, nil
}
