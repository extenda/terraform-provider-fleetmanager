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

// ListSoftwareProfilesForBrandReader is a Reader for the ListSoftwareProfilesForBrand structure.
type ListSoftwareProfilesForBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSoftwareProfilesForBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSoftwareProfilesForBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSoftwareProfilesForBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSoftwareProfilesForBrandOK creates a ListSoftwareProfilesForBrandOK with default headers values
func NewListSoftwareProfilesForBrandOK() *ListSoftwareProfilesForBrandOK {
	return &ListSoftwareProfilesForBrandOK{}
}

/*ListSoftwareProfilesForBrandOK handles this case with default header values.

OK
*/
type ListSoftwareProfilesForBrandOK struct {
	Payload *models.FleetSoftwareProfiles
}

func (o *ListSoftwareProfilesForBrandOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile][%d] listSoftwareProfilesForBrandOK  %+v", 200, o.Payload)
}

func (o *ListSoftwareProfilesForBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetSoftwareProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSoftwareProfilesForBrandBadRequest creates a ListSoftwareProfilesForBrandBadRequest with default headers values
func NewListSoftwareProfilesForBrandBadRequest() *ListSoftwareProfilesForBrandBadRequest {
	return &ListSoftwareProfilesForBrandBadRequest{}
}

/*ListSoftwareProfilesForBrandBadRequest handles this case with default header values.

Bad Request
*/
type ListSoftwareProfilesForBrandBadRequest struct {
	Payload *models.Error
}

func (o *ListSoftwareProfilesForBrandBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile][%d] listSoftwareProfilesForBrandBadRequest  %+v", 400, o.Payload)
}

func (o *ListSoftwareProfilesForBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
