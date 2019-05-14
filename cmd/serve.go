package cmd

import (
	"github.com/salmankhalid187/user-signup/gen/restapi"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations"
)

// Start UserSignUpApisAPI the user sign up apis API
func Start(api *operations.UserSignUpApisAPI, portFlag int) error {
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = portFlag

	return server.Serve()
}
