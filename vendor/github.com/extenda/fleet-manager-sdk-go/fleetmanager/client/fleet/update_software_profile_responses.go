// Code generated by go-swagger; DO NOT EDIT.

package fleet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// UpdateSoftwareProfileReader is a Reader for the UpdateSoftwareProfile structure.
type UpdateSoftwareProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSoftwareProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSoftwareProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSoftwareProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSoftwareProfileNoContent creates a UpdateSoftwareProfileNoContent with default headers values
func NewUpdateSoftwareProfileNoContent() *UpdateSoftwareProfileNoContent {
	return &UpdateSoftwareProfileNoContent{}
}

/*UpdateSoftwareProfileNoContent handles this case with default header values.

No Content
*/
type UpdateSoftwareProfileNoContent struct {
}

func (o *UpdateSoftwareProfileNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}][%d] updateSoftwareProfileNoContent ", 204)
}

func (o *UpdateSoftwareProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSoftwareProfileBadRequest creates a UpdateSoftwareProfileBadRequest with default headers values
func NewUpdateSoftwareProfileBadRequest() *UpdateSoftwareProfileBadRequest {
	return &UpdateSoftwareProfileBadRequest{}
}

/*UpdateSoftwareProfileBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSoftwareProfileBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSoftwareProfileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}][%d] updateSoftwareProfileBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSoftwareProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
