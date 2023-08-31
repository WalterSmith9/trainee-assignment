// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserSegmentsGetResponse user segments get response
//
// swagger:model UserSegmentsGetResponse
type UserSegmentsGetResponse struct {

	// Response code
	// Example: 1
	Code int32 `json:"code,omitempty"`

	// Array of Segments User is in
	Data []string `json:"data"`

	// Response message
	// Example: success
	Message string `json:"message,omitempty"`
}

// Validate validates this user segments get response
func (m *UserSegmentsGetResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user segments get response based on context it is used
func (m *UserSegmentsGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserSegmentsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSegmentsGetResponse) UnmarshalBinary(b []byte) error {
	var res UserSegmentsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
