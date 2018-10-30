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

// UploadBrandSystemPropertiesReader is a Reader for the UploadBrandSystemProperties structure.
type UploadBrandSystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadBrandSystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadBrandSystemPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadBrandSystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadBrandSystemPropertiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadBrandSystemPropertiesOK creates a UploadBrandSystemPropertiesOK with default headers values
func NewUploadBrandSystemPropertiesOK() *UploadBrandSystemPropertiesOK {
	return &UploadBrandSystemPropertiesOK{}
}

/*UploadBrandSystemPropertiesOK handles this case with default header values.

OK
*/
type UploadBrandSystemPropertiesOK struct {
	Payload *models.S3PresignedPost
}

func (o *UploadBrandSystemPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/systemproperties][%d] uploadBrandSystemPropertiesOK  %+v", 200, o.Payload)
}

func (o *UploadBrandSystemPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadBrandSystemPropertiesBadRequest creates a UploadBrandSystemPropertiesBadRequest with default headers values
func NewUploadBrandSystemPropertiesBadRequest() *UploadBrandSystemPropertiesBadRequest {
	return &UploadBrandSystemPropertiesBadRequest{}
}

/*UploadBrandSystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type UploadBrandSystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *UploadBrandSystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/systemproperties][%d] uploadBrandSystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadBrandSystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadBrandSystemPropertiesNotFound creates a UploadBrandSystemPropertiesNotFound with default headers values
func NewUploadBrandSystemPropertiesNotFound() *UploadBrandSystemPropertiesNotFound {
	return &UploadBrandSystemPropertiesNotFound{}
}

/*UploadBrandSystemPropertiesNotFound handles this case with default header values.

Not Found
*/
type UploadBrandSystemPropertiesNotFound struct {
}

func (o *UploadBrandSystemPropertiesNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/systemproperties][%d] uploadBrandSystemPropertiesNotFound ", 404)
}

func (o *UploadBrandSystemPropertiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
