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

// CreateHardwareProfileForBrandReader is a Reader for the CreateHardwareProfileForBrand structure.
type CreateHardwareProfileForBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHardwareProfileForBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateHardwareProfileForBrandCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateHardwareProfileForBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateHardwareProfileForBrandCreated creates a CreateHardwareProfileForBrandCreated with default headers values
func NewCreateHardwareProfileForBrandCreated() *CreateHardwareProfileForBrandCreated {
	return &CreateHardwareProfileForBrandCreated{}
}

/*CreateHardwareProfileForBrandCreated handles this case with default header values.

Created
*/
type CreateHardwareProfileForBrandCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *CreateHardwareProfileForBrandCreated) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile][%d] createHardwareProfileForBrandCreated ", 201)
}

func (o *CreateHardwareProfileForBrandCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewCreateHardwareProfileForBrandBadRequest creates a CreateHardwareProfileForBrandBadRequest with default headers values
func NewCreateHardwareProfileForBrandBadRequest() *CreateHardwareProfileForBrandBadRequest {
	return &CreateHardwareProfileForBrandBadRequest{}
}

/*CreateHardwareProfileForBrandBadRequest handles this case with default header values.

Bad Request
*/
type CreateHardwareProfileForBrandBadRequest struct {
	Payload *models.Error
}

func (o *CreateHardwareProfileForBrandBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile][%d] createHardwareProfileForBrandBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHardwareProfileForBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}