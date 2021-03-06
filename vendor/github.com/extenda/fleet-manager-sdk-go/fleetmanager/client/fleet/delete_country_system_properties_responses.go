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

// DeleteCountrySystemPropertiesReader is a Reader for the DeleteCountrySystemProperties structure.
type DeleteCountrySystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCountrySystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCountrySystemPropertiesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteCountrySystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCountrySystemPropertiesNoContent creates a DeleteCountrySystemPropertiesNoContent with default headers values
func NewDeleteCountrySystemPropertiesNoContent() *DeleteCountrySystemPropertiesNoContent {
	return &DeleteCountrySystemPropertiesNoContent{}
}

/*DeleteCountrySystemPropertiesNoContent handles this case with default header values.

No Content
*/
type DeleteCountrySystemPropertiesNoContent struct {
}

func (o *DeleteCountrySystemPropertiesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] deleteCountrySystemPropertiesNoContent ", 204)
}

func (o *DeleteCountrySystemPropertiesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCountrySystemPropertiesBadRequest creates a DeleteCountrySystemPropertiesBadRequest with default headers values
func NewDeleteCountrySystemPropertiesBadRequest() *DeleteCountrySystemPropertiesBadRequest {
	return &DeleteCountrySystemPropertiesBadRequest{}
}

/*DeleteCountrySystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type DeleteCountrySystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *DeleteCountrySystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] deleteCountrySystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCountrySystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
