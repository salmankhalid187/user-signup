package currenttime

import (
	"context"
	"time"

	"github.com/salmankhalid187/user-signup/gen/models"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations/currenttime"
)

// Service handles async log of audit event
type Service interface {
	GetTime(ctx context.Context, params *currenttime.GetCurrentTimeParams) (*models.CurrenttimeSuccess, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetTime(ctx context.Context, params *currenttime.GetCurrentTimeParams) (*models.CurrenttimeSuccess, error) {

	t := time.Now()
	currenttimeSuccess := models.CurrenttimeSuccess{
		Time: t.String(),
	}

	return &currenttimeSuccess, nil
}
