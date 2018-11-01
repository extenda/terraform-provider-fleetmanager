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

// UpdateBrandReader is a Reader for the UpdateBrand structure.
type UpdateBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateBrandNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBrandNoContent creates a UpdateBrandNoContent with default headers values
func NewUpdateBrandNoContent() *UpdateBrandNoContent {
	return &UpdateBrandNoContent{}
}

/*UpdateBrandNoContent handles this case with default header values.

No Content
*/
type UpdateBrandNoContent struct {
}

func (o *UpdateBrandNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}][%d] updateBrandNoContent ", 204)
}

func (o *UpdateBrandNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBrandBadRequest creates a UpdateBrandBadRequest with default headers values
func NewUpdateBrandBadRequest() *UpdateBrandBadRequest {
	return &UpdateBrandBadRequest{}
}

/*UpdateBrandBadRequest handles this case with default header values.

Bad Request
*/
type UpdateBrandBadRequest struct {
	Payload *models.Error
}

func (o *UpdateBrandBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}][%d] updateBrandBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
