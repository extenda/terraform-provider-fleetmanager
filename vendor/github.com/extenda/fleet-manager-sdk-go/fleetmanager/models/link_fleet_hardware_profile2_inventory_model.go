// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LinkFleetHardwareProfile2InventoryModel link fleet hardware profile2 inventory model
// swagger:model LinkFleetHardwareProfile2InventoryModel
type LinkFleetHardwareProfile2InventoryModel struct {

	// layer
	// Required: true
	// Enum: [cashchanger cashdrawer eft fiscal linedisplay msr printer printer2 scale scanner scanner2]
	Layer *string `json:"layer"`

	// model
	// Required: true
	Model *LinkFleetHardwareProfile2InventoryModelModel `json:"model"`

	// port
	// Required: true
	// Enum: [usb com1 com2 com3 com4 com5 com6 com7 com8 com9]
	Port *string `json:"port"`
}

// Validate validates this link fleet hardware profile2 inventory model
func (m *LinkFleetHardwareProfile2InventoryModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var linkFleetHardwareProfile2InventoryModelTypeLayerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cashchanger","cashdrawer","eft","fiscal","linedisplay","msr","printer","printer2","scale","scanner","scanner2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkFleetHardwareProfile2InventoryModelTypeLayerPropEnum = append(linkFleetHardwareProfile2InventoryModelTypeLayerPropEnum, v)
	}
}

const (

	// LinkFleetHardwareProfile2InventoryModelLayerCashchanger captures enum value "cashchanger"
	LinkFleetHardwareProfile2InventoryModelLayerCashchanger string = "cashchanger"

	// LinkFleetHardwareProfile2InventoryModelLayerCashdrawer captures enum value "cashdrawer"
	LinkFleetHardwareProfile2InventoryModelLayerCashdrawer string = "cashdrawer"

	// LinkFleetHardwareProfile2InventoryModelLayerEft captures enum value "eft"
	LinkFleetHardwareProfile2InventoryModelLayerEft string = "eft"

	// LinkFleetHardwareProfile2InventoryModelLayerFiscal captures enum value "fiscal"
	LinkFleetHardwareProfile2InventoryModelLayerFiscal string = "fiscal"

	// LinkFleetHardwareProfile2InventoryModelLayerLinedisplay captures enum value "linedisplay"
	LinkFleetHardwareProfile2InventoryModelLayerLinedisplay string = "linedisplay"

	// LinkFleetHardwareProfile2InventoryModelLayerMsr captures enum value "msr"
	LinkFleetHardwareProfile2InventoryModelLayerMsr string = "msr"

	// LinkFleetHardwareProfile2InventoryModelLayerPrinter captures enum value "printer"
	LinkFleetHardwareProfile2InventoryModelLayerPrinter string = "printer"

	// LinkFleetHardwareProfile2InventoryModelLayerPrinter2 captures enum value "printer2"
	LinkFleetHardwareProfile2InventoryModelLayerPrinter2 string = "printer2"

	// LinkFleetHardwareProfile2InventoryModelLayerScale captures enum value "scale"
	LinkFleetHardwareProfile2InventoryModelLayerScale string = "scale"

	// LinkFleetHardwareProfile2InventoryModelLayerScanner captures enum value "scanner"
	LinkFleetHardwareProfile2InventoryModelLayerScanner string = "scanner"

	// LinkFleetHardwareProfile2InventoryModelLayerScanner2 captures enum value "scanner2"
	LinkFleetHardwareProfile2InventoryModelLayerScanner2 string = "scanner2"
)

// prop value enum
func (m *LinkFleetHardwareProfile2InventoryModel) validateLayerEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, linkFleetHardwareProfile2InventoryModelTypeLayerPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LinkFleetHardwareProfile2InventoryModel) validateLayer(formats strfmt.Registry) error {

	if err := validate.Required("layer", "body", m.Layer); err != nil {
		return err
	}

	// value enum
	if err := m.validateLayerEnum("layer", "body", *m.Layer); err != nil {
		return err
	}

	return nil
}

func (m *LinkFleetHardwareProfile2InventoryModel) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

var linkFleetHardwareProfile2InventoryModelTypePortPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["usb","com1","com2","com3","com4","com5","com6","com7","com8","com9"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkFleetHardwareProfile2InventoryModelTypePortPropEnum = append(linkFleetHardwareProfile2InventoryModelTypePortPropEnum, v)
	}
}

const (

	// LinkFleetHardwareProfile2InventoryModelPortUsb captures enum value "usb"
	LinkFleetHardwareProfile2InventoryModelPortUsb string = "usb"

	// LinkFleetHardwareProfile2InventoryModelPortCom1 captures enum value "com1"
	LinkFleetHardwareProfile2InventoryModelPortCom1 string = "com1"

	// LinkFleetHardwareProfile2InventoryModelPortCom2 captures enum value "com2"
	LinkFleetHardwareProfile2InventoryModelPortCom2 string = "com2"

	// LinkFleetHardwareProfile2InventoryModelPortCom3 captures enum value "com3"
	LinkFleetHardwareProfile2InventoryModelPortCom3 string = "com3"

	// LinkFleetHardwareProfile2InventoryModelPortCom4 captures enum value "com4"
	LinkFleetHardwareProfile2InventoryModelPortCom4 string = "com4"

	// LinkFleetHardwareProfile2InventoryModelPortCom5 captures enum value "com5"
	LinkFleetHardwareProfile2InventoryModelPortCom5 string = "com5"

	// LinkFleetHardwareProfile2InventoryModelPortCom6 captures enum value "com6"
	LinkFleetHardwareProfile2InventoryModelPortCom6 string = "com6"

	// LinkFleetHardwareProfile2InventoryModelPortCom7 captures enum value "com7"
	LinkFleetHardwareProfile2InventoryModelPortCom7 string = "com7"

	// LinkFleetHardwareProfile2InventoryModelPortCom8 captures enum value "com8"
	LinkFleetHardwareProfile2InventoryModelPortCom8 string = "com8"

	// LinkFleetHardwareProfile2InventoryModelPortCom9 captures enum value "com9"
	LinkFleetHardwareProfile2InventoryModelPortCom9 string = "com9"
)

// prop value enum
func (m *LinkFleetHardwareProfile2InventoryModel) validatePortEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, linkFleetHardwareProfile2InventoryModelTypePortPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LinkFleetHardwareProfile2InventoryModel) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	// value enum
	if err := m.validatePortEnum("port", "body", *m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkFleetHardwareProfile2InventoryModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkFleetHardwareProfile2InventoryModel) UnmarshalBinary(b []byte) error {
	var res LinkFleetHardwareProfile2InventoryModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinkFleetHardwareProfile2InventoryModelModel link fleet hardware profile2 inventory model model
// swagger:model LinkFleetHardwareProfile2InventoryModelModel
type LinkFleetHardwareProfile2InventoryModelModel struct {

	// manufacturer Id
	// Required: true
	ManufacturerID *string `json:"manufacturerId"`

	// model Id
	// Required: true
	ModelID *string `json:"modelId"`
}

// Validate validates this link fleet hardware profile2 inventory model model
func (m *LinkFleetHardwareProfile2InventoryModelModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManufacturerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkFleetHardwareProfile2InventoryModelModel) validateManufacturerID(formats strfmt.Registry) error {

	if err := validate.Required("model"+"."+"manufacturerId", "body", m.ManufacturerID); err != nil {
		return err
	}

	return nil
}

func (m *LinkFleetHardwareProfile2InventoryModelModel) validateModelID(formats strfmt.Registry) error {

	if err := validate.Required("model"+"."+"modelId", "body", m.ModelID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkFleetHardwareProfile2InventoryModelModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkFleetHardwareProfile2InventoryModelModel) UnmarshalBinary(b []byte) error {
	var res LinkFleetHardwareProfile2InventoryModelModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
