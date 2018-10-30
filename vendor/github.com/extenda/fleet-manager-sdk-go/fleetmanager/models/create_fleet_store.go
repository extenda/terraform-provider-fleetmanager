// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateFleetStore create fleet store
// swagger:model CreateFleetStore
type CreateFleetStore struct {

	// city
	// Required: true
	City *string `json:"city"`

	// contact person
	// Required: true
	ContactPerson *FleetStoreContactPerson `json:"contactPerson"`

	// name
	// Required: true
	Name *string `json:"name"`

	// post code
	// Required: true
	PostCode *string `json:"postCode"`

	// street
	// Required: true
	Street *string `json:"street"`
}

// Validate validates this create fleet store
func (m *CreateFleetStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFleetStore) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *CreateFleetStore) validateContactPerson(formats strfmt.Registry) error {

	if err := validate.Required("contactPerson", "body", m.ContactPerson); err != nil {
		return err
	}

	if m.ContactPerson != nil {
		if err := m.ContactPerson.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactPerson")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFleetStore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateFleetStore) validatePostCode(formats strfmt.Registry) error {

	if err := validate.Required("postCode", "body", m.PostCode); err != nil {
		return err
	}

	return nil
}

func (m *CreateFleetStore) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFleetStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFleetStore) UnmarshalBinary(b []byte) error {
	var res CreateFleetStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
