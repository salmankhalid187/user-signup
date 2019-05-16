package currenttime

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations/currenttime"
	"github.com/salmankhalid187/user-signup/swagger"
)

// Configure setups handlers on api with Service
func Configure(api *operations.UserSignUpApisAPI, service Service) {

	api.CurrenttimeGetCurrentTimeHandler = currenttime.GetCurrentTimeHandlerFunc(func(params currenttime.GetCurrentTimeParams) middleware.Responder {

		result, err := service.GetTime(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.SignupErrorHandler("GetTime", err)
		}
		return currenttime.NewGetCurrentTimeOK().WithPayload(result)
	})

}
