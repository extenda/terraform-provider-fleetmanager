// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetSoftwarePackageByIDReader is a Reader for the GetSoftwarePackageByID structure.
type GetSoftwarePackageByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSoftwarePackageByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSoftwarePackageByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSoftwarePackageByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSoftwarePackageByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSoftwarePackageByIDOK creates a GetSoftwarePackageByIDOK with default headers values
func NewGetSoftwarePackageByIDOK() *GetSoftwarePackageByIDOK {
	return &GetSoftwarePackageByIDOK{}
}

/*GetSoftwarePackageByIDOK handles this case with default header values.

OK
*/
type GetSoftwarePackageByIDOK struct {
	Payload *models.SoftwarePackage
}

func (o *GetSoftwarePackageByIDOK) Error() string {
	return fmt.Sprintf("[GET /software/package/{packageId}][%d] getSoftwarePackageByIdOK  %+v", 200, o.Payload)
}

func (o *GetSoftwarePackageByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwarePackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSoftwarePackageByIDBadRequest creates a GetSoftwarePackageByIDBadRequest with default headers values
func NewGetSoftwarePackageByIDBadRequest() *GetSoftwarePackageByIDBadRequest {
	return &GetSoftwarePackageByIDBadRequest{}
}

/*GetSoftwarePackageByIDBadRequest handles this case with default header values.

Bad Request
*/
type GetSoftwarePackageByIDBadRequest struct {
	Payload *models.Error
}

func (o *GetSoftwarePackageByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /software/package/{packageId}][%d] getSoftwarePackageByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSoftwarePackageByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSoftwarePackageByIDNotFound creates a GetSoftwarePackageByIDNotFound with default headers values
func NewGetSoftwarePackageByIDNotFound() *GetSoftwarePackageByIDNotFound {
	return &GetSoftwarePackageByIDNotFound{}
}

/*GetSoftwarePackageByIDNotFound handles this case with default header values.

Not Found
*/
type GetSoftwarePackageByIDNotFound struct {
}

func (o *GetSoftwarePackageByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /software/package/{packageId}][%d] getSoftwarePackageByIdNotFound ", 404)
}

func (o *GetSoftwarePackageByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}