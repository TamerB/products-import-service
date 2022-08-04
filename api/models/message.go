// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Message Informational message
//
// swagger:model Message
type Message struct {

	// The message code to use when communicating about this problem
	// Example: 8001
	Number int64 `json:"Number,omitempty"`

	// A description of this message, which may include formatting specific to this context
	// Example: Location closed
	Text string `json:"Text,omitempty"`

	// Type of message, e.g. warning/info
	// Example: Info
	Type string `json:"Type,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this message based on context it is used
func (m *Message) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
