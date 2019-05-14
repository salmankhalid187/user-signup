// Code generated by go-swagger; DO NOT EDIT.

package signup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/salmankhalid187/user-signup/gen/models"
)

// SignUpUserOKCode is the HTTP code returned for type SignUpUserOK
const SignUpUserOKCode int = 200

/*SignUpUserOK Success

swagger:response signUpUserOK
*/
type SignUpUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.SignupSuccess `json:"body,omitempty"`
}

// NewSignUpUserOK creates SignUpUserOK with default headers values
func NewSignUpUserOK() *SignUpUserOK {

	return &SignUpUserOK{}
}

// WithPayload adds the payload to the sign up user o k response
func (o *SignUpUserOK) WithPayload(payload *models.SignupSuccess) *SignUpUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign up user o k response
func (o *SignUpUserOK) SetPayload(payload *models.SignupSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignUpUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SignUpUserBadRequestCode is the HTTP code returned for type SignUpUserBadRequest
const SignUpUserBadRequestCode int = 400

/*SignUpUserBadRequest Invalid request

swagger:response signUpUserBadRequest
*/
type SignUpUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSignUpUserBadRequest creates SignUpUserBadRequest with default headers values
func NewSignUpUserBadRequest() *SignUpUserBadRequest {

	return &SignUpUserBadRequest{}
}

// WithPayload adds the payload to the sign up user bad request response
func (o *SignUpUserBadRequest) WithPayload(payload *models.ErrorResponse) *SignUpUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign up user bad request response
func (o *SignUpUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignUpUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}