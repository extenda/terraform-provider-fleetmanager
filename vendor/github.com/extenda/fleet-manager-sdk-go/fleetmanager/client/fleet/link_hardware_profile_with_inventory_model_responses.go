// Code generated by go-swagger; DO NOT EDIT.

package fleet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// LinkHardwareProfileWithInventoryModelReader is a Reader for the LinkHardwareProfileWithInventoryModel structure.
type LinkHardwareProfileWithInventoryModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LinkHardwareProfileWithInventoryModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewLinkHardwareProfileWithInventoryModelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLinkHardwareProfileWithInventoryModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLinkHardwareProfileWithInventoryModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLinkHardwareProfileWithInventoryModelNoContent creates a LinkHardwareProfileWithInventoryModelNoContent with default headers values
func NewLinkHardwareProfileWithInventoryModelNoContent() *LinkHardwareProfileWithInventoryModelNoContent {
	return &LinkHardwareProfileWithInventoryModelNoContent{}
}

/*LinkHardwareProfileWithInventoryModelNoContent handles this case with default header values.

No Content
*/
type LinkHardwareProfileWithInventoryModelNoContent struct {
}

func (o *LinkHardwareProfileWithInventoryModelNoContent) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] linkHardwareProfileWithInventoryModelNoContent ", 204)
}

func (o *LinkHardwareProfileWithInventoryModelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLinkHardwareProfileWithInventoryModelBadRequest creates a LinkHardwareProfileWithInventoryModelBadRequest with default headers values
func NewLinkHardwareProfileWithInventoryModelBadRequest() *LinkHardwareProfileWithInventoryModelBadRequest {
	return &LinkHardwareProfileWithInventoryModelBadRequest{}
}

/*LinkHardwareProfileWithInventoryModelBadRequest handles this case with default header values.

Bad Request
*/
type LinkHardwareProfileWithInventoryModelBadRequest struct {
	Payload *models.Error
}

func (o *LinkHardwareProfileWithInventoryModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] linkHardwareProfileWithInventoryModelBadRequest  %+v", 400, o.Payload)
}

func (o *LinkHardwareProfileWithInventoryModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLinkHardwareProfileWithInventoryModelNotFound creates a LinkHardwareProfileWithInventoryModelNotFound with default headers values
func NewLinkHardwareProfileWithInventoryModelNotFound() *LinkHardwareProfileWithInventoryModelNotFound {
	return &LinkHardwareProfileWithInventoryModelNotFound{}
}

/*LinkHardwareProfileWithInventoryModelNotFound handles this case with default header values.

Not Found
*/
type LinkHardwareProfileWithInventoryModelNotFound struct {
}

func (o *LinkHardwareProfileWithInventoryModelNotFound) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] linkHardwareProfileWithInventoryModelNotFound ", 404)
}

func (o *LinkHardwareProfileWithInventoryModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
