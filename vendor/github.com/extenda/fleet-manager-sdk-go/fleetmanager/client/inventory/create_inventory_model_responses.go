// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// CreateInventoryModelReader is a Reader for the CreateInventoryModel structure.
type CreateInventoryModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInventoryModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateInventoryModelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateInventoryModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInventoryModelCreated creates a CreateInventoryModelCreated with default headers values
func NewCreateInventoryModelCreated() *CreateInventoryModelCreated {
	return &CreateInventoryModelCreated{}
}

/*CreateInventoryModelCreated handles this case with default header values.

Created
*/
type CreateInventoryModelCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *CreateInventoryModelCreated) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model][%d] createInventoryModelCreated ", 201)
}

func (o *CreateInventoryModelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewCreateInventoryModelBadRequest creates a CreateInventoryModelBadRequest with default headers values
func NewCreateInventoryModelBadRequest() *CreateInventoryModelBadRequest {
	return &CreateInventoryModelBadRequest{}
}

/*CreateInventoryModelBadRequest handles this case with default header values.

Bad Request
*/
type CreateInventoryModelBadRequest struct {
	Payload *models.Error
}

func (o *CreateInventoryModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model][%d] createInventoryModelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInventoryModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
