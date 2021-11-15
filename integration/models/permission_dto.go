// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PermissionDto PermissionDto PermissionDto permission dto
//
// swagger:model PermissionDto
type PermissionDto struct {

	// created on
	CreatedOn string `json:"createdOn,omitempty"`

	// Id
	ID string `json:"id,omitempty"`

	// last updated
	LastUpdated string `json:"lastUpdated,omitempty"`

	// on user Id
	OnUserID string `json:"onUserId,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// permission
	Permission PermissionType `json:"permission,omitempty"`
}

// Validate validates this permission dto
func (m *PermissionDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionDto) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		}
		return err
	}

	return nil
}

// ContextValidate validate this permission dto based on the context it is used
func (m *PermissionDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionDto) contextValidatePermission(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Permission.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PermissionDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionDto) UnmarshalBinary(b []byte) error {
	var res PermissionDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
