// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenDto TokenDto TokenDto TokenDto TokenDto TokenDto TokenDto token dto
//
// swagger:model TokenDto
type TokenDto struct {

	// Id
	ID string `json:"id,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this token dto
func (m *TokenDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this token dto based on context it is used
func (m *TokenDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenDto) UnmarshalBinary(b []byte) error {
	var res TokenDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
