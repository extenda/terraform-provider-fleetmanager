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

// UploadWorkstationSystemPropertiesReader is a Reader for the UploadWorkstationSystemProperties structure.
type UploadWorkstationSystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadWorkstationSystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadWorkstationSystemPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadWorkstationSystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadWorkstationSystemPropertiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadWorkstationSystemPropertiesOK creates a UploadWorkstationSystemPropertiesOK with default headers values
func NewUploadWorkstationSystemPropertiesOK() *UploadWorkstationSystemPropertiesOK {
	return &UploadWorkstationSystemPropertiesOK{}
}

/*UploadWorkstationSystemPropertiesOK handles this case with default header values.

OK
*/
type UploadWorkstationSystemPropertiesOK struct {
	Payload *models.S3PresignedPost
}

func (o *UploadWorkstationSystemPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/workstation/{workstationId}/systemproperties][%d] uploadWorkstationSystemPropertiesOK  %+v", 200, o.Payload)
}

func (o *UploadWorkstationSystemPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadWorkstationSystemPropertiesBadRequest creates a UploadWorkstationSystemPropertiesBadRequest with default headers values
func NewUploadWorkstationSystemPropertiesBadRequest() *UploadWorkstationSystemPropertiesBadRequest {
	return &UploadWorkstationSystemPropertiesBadRequest{}
}

/*UploadWorkstationSystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type UploadWorkstationSystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *UploadWorkstationSystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/workstation/{workstationId}/systemproperties][%d] uploadWorkstationSystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadWorkstationSystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadWorkstationSystemPropertiesNotFound creates a UploadWorkstationSystemPropertiesNotFound with default headers values
func NewUploadWorkstationSystemPropertiesNotFound() *UploadWorkstationSystemPropertiesNotFound {
	return &UploadWorkstationSystemPropertiesNotFound{}
}

/*UploadWorkstationSystemPropertiesNotFound handles this case with default header values.

Not Found
*/
type UploadWorkstationSystemPropertiesNotFound struct {
}

func (o *UploadWorkstationSystemPropertiesNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/workstation/{workstationId}/systemproperties][%d] uploadWorkstationSystemPropertiesNotFound ", 404)
}

func (o *UploadWorkstationSystemPropertiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
