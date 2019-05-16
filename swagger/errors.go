package swagger

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/salmankhalid187/user-signup/gen/models"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations/signup"
	"github.com/sirupsen/logrus"
)

type codedResponse interface {
	Code() string
}

// ErrorResponse wraps the error in the api standard models.ErrorResponse object
func ErrorResponse(err error) *models.ErrorResponse {
	cd := ""
	if e, ok := err.(codedResponse); ok {
		cd = e.Code()
	}

	e := models.ErrorResponse{
		Code:    cd,
		Message: err.Error(),
	}
	return &e
}

func SignupErrorHandler(label string, err error) middleware.Responder {

	logrus.WithError(err).Error(label)
	return signup.NewCreateUserBadRequest().WithPayload(ErrorResponse(err))
}
