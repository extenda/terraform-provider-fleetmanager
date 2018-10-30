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

// LinkFleetWorkstation2FleetSoftwareProfile link fleet workstation2 fleet software profile
// swagger:model LinkFleetWorkstation2FleetSoftwareProfile
type LinkFleetWorkstation2FleetSoftwareProfile struct {

	// software profile
	// Required: true
	SoftwareProfile *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile `json:"softwareProfile"`
}

// Validate validates this link fleet workstation2 fleet software profile
func (m *LinkFleetWorkstation2FleetSoftwareProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSoftwareProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkFleetWorkstation2FleetSoftwareProfile) validateSoftwareProfile(formats strfmt.Registry) error {

	if err := validate.Required("softwareProfile", "body", m.SoftwareProfile); err != nil {
		return err
	}

	if m.SoftwareProfile != nil {
		if err := m.SoftwareProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkFleetWorkstation2FleetSoftwareProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkFleetWorkstation2FleetSoftwareProfile) UnmarshalBinary(b []byte) error {
	var res LinkFleetWorkstation2FleetSoftwareProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile link fleet workstation2 fleet software profile software profile
// swagger:model LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile
type LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile struct {

	// brand Id
	// Required: true
	BrandID *string `json:"brandId"`

	// software profile Id
	// Required: true
	SoftwareProfileID *string `json:"softwareProfileId"`

	// tenant Id
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this link fleet workstation2 fleet software profile software profile
func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) validateBrandID(formats strfmt.Registry) error {

	if err := validate.Required("softwareProfile"+"."+"brandId", "body", m.BrandID); err != nil {
		return err
	}

	return nil
}

func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) validateSoftwareProfileID(formats strfmt.Registry) error {

	if err := validate.Required("softwareProfile"+"."+"softwareProfileId", "body", m.SoftwareProfileID); err != nil {
		return err
	}

	return nil
}

func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("softwareProfile"+"."+"tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile) UnmarshalBinary(b []byte) error {
	var res LinkFleetWorkstation2FleetSoftwareProfileSoftwareProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
