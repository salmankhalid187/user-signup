// Code generated by go-swagger; DO NOT EDIT.

package signup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/salmankhalid187/user-signup/gen/models"
)

// CreateUserOKCode is the HTTP code returned for type CreateUserOK
const CreateUserOKCode int = 200

/*CreateUserOK Success

swagger:response createUserOK
*/
type CreateUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.SignupSuccess `json:"body,omitempty"`
}

// NewCreateUserOK creates CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {

	return &CreateUserOK{}
}

// WithPayload adds the payload to the create user o k response
func (o *CreateUserOK) WithPayload(payload *models.SignupSuccess) *CreateUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user o k response
func (o *CreateUserOK) SetPayload(payload *models.SignupSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserBadRequestCode is the HTTP code returned for type CreateUserBadRequest
const CreateUserBadRequestCode int = 400

/*CreateUserBadRequest Invalid request

swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {

	return &CreateUserBadRequest{}
}

// WithPayload adds the payload to the create user bad request response
func (o *CreateUserBadRequest) WithPayload(payload *models.ErrorResponse) *CreateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user bad request response
func (o *CreateUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
