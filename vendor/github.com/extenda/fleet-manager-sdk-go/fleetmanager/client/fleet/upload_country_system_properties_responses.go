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

// UploadCountrySystemPropertiesReader is a Reader for the UploadCountrySystemProperties structure.
type UploadCountrySystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadCountrySystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadCountrySystemPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadCountrySystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadCountrySystemPropertiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadCountrySystemPropertiesOK creates a UploadCountrySystemPropertiesOK with default headers values
func NewUploadCountrySystemPropertiesOK() *UploadCountrySystemPropertiesOK {
	return &UploadCountrySystemPropertiesOK{}
}

/*UploadCountrySystemPropertiesOK handles this case with default header values.

OK
*/
type UploadCountrySystemPropertiesOK struct {
	Payload *models.S3PresignedPost
}

func (o *UploadCountrySystemPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] uploadCountrySystemPropertiesOK  %+v", 200, o.Payload)
}

func (o *UploadCountrySystemPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadCountrySystemPropertiesBadRequest creates a UploadCountrySystemPropertiesBadRequest with default headers values
func NewUploadCountrySystemPropertiesBadRequest() *UploadCountrySystemPropertiesBadRequest {
	return &UploadCountrySystemPropertiesBadRequest{}
}

/*UploadCountrySystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type UploadCountrySystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *UploadCountrySystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] uploadCountrySystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadCountrySystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadCountrySystemPropertiesNotFound creates a UploadCountrySystemPropertiesNotFound with default headers values
func NewUploadCountrySystemPropertiesNotFound() *UploadCountrySystemPropertiesNotFound {
	return &UploadCountrySystemPropertiesNotFound{}
}

/*UploadCountrySystemPropertiesNotFound handles this case with default header values.

Not Found
*/
type UploadCountrySystemPropertiesNotFound struct {
}

func (o *UploadCountrySystemPropertiesNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] uploadCountrySystemPropertiesNotFound ", 404)
}

func (o *UploadCountrySystemPropertiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
