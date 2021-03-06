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

// FleetStore fleet store
// swagger:model FleetStore
type FleetStore struct {

	// city
	// Required: true
	City *string `json:"city"`

	// contact person
	// Required: true
	ContactPerson *FleetStoreContactPerson `json:"contactPerson"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// post code
	// Required: true
	PostCode *string `json:"postCode"`

	// street
	// Required: true
	Street *string `json:"street"`

	// system properties Url
	SystemPropertiesURL string `json:"systemPropertiesUrl,omitempty"`
}

// Validate validates this fleet store
func (m *FleetStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *FleetStore) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *FleetStore) validateContactPerson(formats strfmt.Registry) error {

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

func (m *FleetStore) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FleetStore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FleetStore) validatePostCode(formats strfmt.Registry) error {

	if err := validate.Required("postCode", "body", m.PostCode); err != nil {
		return err
	}

	return nil
}

func (m *FleetStore) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetStore) UnmarshalBinary(b []byte) error {
	var res FleetStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
