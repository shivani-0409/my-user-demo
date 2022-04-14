// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
)

// AddUserCreatedCode is the HTTP code returned for type AddUserCreated
const AddUserCreatedCode int = 201

/*AddUserCreated Created

swagger:response addUserCreated
*/
type AddUserCreated struct {
}

// NewAddUserCreated creates AddUserCreated with default headers values
func NewAddUserCreated() *AddUserCreated {

	return &AddUserCreated{}
}

// WriteResponse to the client
func (o *AddUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddUserBadRequestCode is the HTTP code returned for type AddUserBadRequest
const AddUserBadRequestCode int = 400

/*AddUserBadRequest Name should be atleast 3 characters long

swagger:response addUserBadRequest
*/
type AddUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddUserBadRequest creates AddUserBadRequest with default headers values
func NewAddUserBadRequest() *AddUserBadRequest {

	return &AddUserBadRequest{}
}

// WithPayload adds the payload to the add user bad request response
func (o *AddUserBadRequest) WithPayload(payload *models.Error) *AddUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user bad request response
func (o *AddUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddUserDefault error

swagger:response addUserDefault
*/
type AddUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddUserDefault creates AddUserDefault with default headers values
func NewAddUserDefault(code int) *AddUserDefault {
	if code <= 0 {
		code = 500
	}

	return &AddUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add user default response
func (o *AddUserDefault) WithStatusCode(code int) *AddUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add user default response
func (o *AddUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add user default response
func (o *AddUserDefault) WithPayload(payload *models.Error) *AddUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user default response
func (o *AddUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}