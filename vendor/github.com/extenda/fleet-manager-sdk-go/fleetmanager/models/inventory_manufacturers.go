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

// InventoryManufacturers inventory manufacturers
// swagger:model InventoryManufacturers
type InventoryManufacturers struct {

	// manufacturers
	// Required: true
	Manufacturers []*InventoryManufacturersManufacturersItems0 `json:"manufacturers"`

	// next token
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this inventory manufacturers
func (m *InventoryManufacturers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManufacturers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryManufacturers) validateManufacturers(formats strfmt.Registry) error {

	if err := validate.Required("manufacturers", "body", m.Manufacturers); err != nil {
		return err
	}

	for i := 0; i < len(m.Manufacturers); i++ {
		if swag.IsZero(m.Manufacturers[i]) { // not required
			continue
		}

		if m.Manufacturers[i] != nil {
			if err := m.Manufacturers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manufacturers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryManufacturers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryManufacturers) UnmarshalBinary(b []byte) error {
	var res InventoryManufacturers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryManufacturersManufacturersItems0 inventory manufacturers manufacturers items0
// swagger:model InventoryManufacturersManufacturersItems0
type InventoryManufacturersManufacturersItems0 struct {

	// href
	// Required: true
	Href *string `json:"href"`
}

// Validate validates this inventory manufacturers manufacturers items0
func (m *InventoryManufacturersManufacturersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryManufacturersManufacturersItems0) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryManufacturersManufacturersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryManufacturersManufacturersItems0) UnmarshalBinary(b []byte) error {
	var res InventoryManufacturersManufacturersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}