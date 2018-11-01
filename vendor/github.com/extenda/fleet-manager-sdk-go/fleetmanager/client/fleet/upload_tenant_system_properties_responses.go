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

// UploadTenantSystemPropertiesReader is a Reader for the UploadTenantSystemProperties structure.
type UploadTenantSystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadTenantSystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadTenantSystemPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadTenantSystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadTenantSystemPropertiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadTenantSystemPropertiesOK creates a UploadTenantSystemPropertiesOK with default headers values
func NewUploadTenantSystemPropertiesOK() *UploadTenantSystemPropertiesOK {
	return &UploadTenantSystemPropertiesOK{}
}

/*UploadTenantSystemPropertiesOK handles this case with default header values.

OK
*/
type UploadTenantSystemPropertiesOK struct {
	Payload *models.S3PresignedPost
}

func (o *UploadTenantSystemPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemproperties][%d] uploadTenantSystemPropertiesOK  %+v", 200, o.Payload)
}

func (o *UploadTenantSystemPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTenantSystemPropertiesBadRequest creates a UploadTenantSystemPropertiesBadRequest with default headers values
func NewUploadTenantSystemPropertiesBadRequest() *UploadTenantSystemPropertiesBadRequest {
	return &UploadTenantSystemPropertiesBadRequest{}
}

/*UploadTenantSystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type UploadTenantSystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *UploadTenantSystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemproperties][%d] uploadTenantSystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadTenantSystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTenantSystemPropertiesNotFound creates a UploadTenantSystemPropertiesNotFound with default headers values
func NewUploadTenantSystemPropertiesNotFound() *UploadTenantSystemPropertiesNotFound {
	return &UploadTenantSystemPropertiesNotFound{}
}

/*UploadTenantSystemPropertiesNotFound handles this case with default header values.

Not Found
*/
type UploadTenantSystemPropertiesNotFound struct {
}

func (o *UploadTenantSystemPropertiesNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemproperties][%d] uploadTenantSystemPropertiesNotFound ", 404)
}

func (o *UploadTenantSystemPropertiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
