// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FleetStores fleet stores
// swagger:model FleetStores
type FleetStores struct {

	// next token
	NextToken string `json:"nextToken,omitempty"`

	// stores
	// Required: true
	Stores []*FleetStoresStoresItems0 `json:"stores"`
}

// Validate validates this fleet stores
func (m *FleetStores) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetStores) validateStores(formats strfmt.Registry) error {

	if err := validate.Required("stores", "body", m.Stores); err != nil {
		return err
	}

	for i := 0; i < len(m.Stores); i++ {
		if swag.IsZero(m.Stores[i]) { // not required
			continue
		}

		if m.Stores[i] != nil {
			if err := m.Stores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetStores) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetStores) UnmarshalBinary(b []byte) error {
	var res FleetStores
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FleetStoresStoresItems0 fleet stores stores items0
// swagger:model FleetStoresStoresItems0
type FleetStoresStoresItems0 struct {

	// href
	// Required: true
	Href *string `json:"href"`
}

// Validate validates this fleet stores stores items0
func (m *FleetStoresStoresItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetStoresStoresItems0) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetStoresStoresItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetStoresStoresItems0) UnmarshalBinary(b []byte) error {
	var res FleetStoresStoresItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
