package usersignup

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations/signup"
	"github.com/salmankhalid187/user-signup/swagger"
)

// Configure setups handlers on api with Service
func Configure(api *operations.UserSignUpApisAPI, service Service) {

	// api.SignupSignUpUserHandler = signup.SignUpUserHandlerFunc(func(params signup.SignUpUserParams) middleware.Responder {
	// 	result, err := service.SignupUser(params.HTTPRequest.Context(), &params)
	// 	if err != nil {
	// 		return swagger.SignupErrorHandler("SignUp", err)
	// 	}
	// 	return signup.NewSignUpUserOK().WithPayload(result)
	// })

	api.SignupCreateUserHandler = signup.CreateUserHandlerFunc(func(params signup.CreateUserParams) middleware.Responder {

		result, err := service.CreateUser(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.SignupErrorHandler("CreateUser", err)
		}
		return signup.NewCreateUserOK().WithPayload(result)
	})

}
