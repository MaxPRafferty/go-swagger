// Code generated by go-swagger; DO NOT EDIT.

package models2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtUsersArrayFlat users array flat
// swagger:model ExtUsersArrayFlat
type ExtUsersArrayFlat []*ExtUsersArrayFlatItems0

// Validate validates this users array flat
func (m ExtUsersArrayFlat) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ExtUsersArrayFlatItems0 users array flat items0
// swagger:model ExtUsersArrayFlatItems0
type ExtUsersArrayFlatItems0 struct {

	// age
	Age int64 `json:"age,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this users array flat items0
func (m *ExtUsersArrayFlatItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtUsersArrayFlatItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtUsersArrayFlatItems0) UnmarshalBinary(b []byte) error {
	var res ExtUsersArrayFlatItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
