// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/models"
)

// SegmentCreateOKCode is the HTTP code returned for type SegmentCreateOK
const SegmentCreateOKCode int = 200

/*
SegmentCreateOK Success

swagger:response segmentCreateOK
*/
type SegmentCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewSegmentCreateOK creates SegmentCreateOK with default headers values
func NewSegmentCreateOK() *SegmentCreateOK {

	return &SegmentCreateOK{}
}

// WithPayload adds the payload to the segment create o k response
func (o *SegmentCreateOK) WithPayload(payload *models.SuccessResponse) *SegmentCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the segment create o k response
func (o *SegmentCreateOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SegmentCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SegmentCreateBadRequestCode is the HTTP code returned for type SegmentCreateBadRequest
const SegmentCreateBadRequestCode int = 400

/*
SegmentCreateBadRequest Bad Request

swagger:response segmentCreateBadRequest
*/
type SegmentCreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSegmentCreateBadRequest creates SegmentCreateBadRequest with default headers values
func NewSegmentCreateBadRequest() *SegmentCreateBadRequest {

	return &SegmentCreateBadRequest{}
}

// WithPayload adds the payload to the segment create bad request response
func (o *SegmentCreateBadRequest) WithPayload(payload *models.ErrorResponse) *SegmentCreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the segment create bad request response
func (o *SegmentCreateBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SegmentCreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SegmentCreateInternalServerErrorCode is the HTTP code returned for type SegmentCreateInternalServerError
const SegmentCreateInternalServerErrorCode int = 500

/*
SegmentCreateInternalServerError Internal Server Error

swagger:response segmentCreateInternalServerError
*/
type SegmentCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSegmentCreateInternalServerError creates SegmentCreateInternalServerError with default headers values
func NewSegmentCreateInternalServerError() *SegmentCreateInternalServerError {

	return &SegmentCreateInternalServerError{}
}

// WithPayload adds the payload to the segment create internal server error response
func (o *SegmentCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *SegmentCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the segment create internal server error response
func (o *SegmentCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SegmentCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
