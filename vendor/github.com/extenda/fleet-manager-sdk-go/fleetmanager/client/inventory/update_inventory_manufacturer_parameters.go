// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewUpdateInventoryManufacturerParams creates a new UpdateInventoryManufacturerParams object
// with the default values initialized.
func NewUpdateInventoryManufacturerParams() *UpdateInventoryManufacturerParams {
	var ()
	return &UpdateInventoryManufacturerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInventoryManufacturerParamsWithTimeout creates a new UpdateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInventoryManufacturerParamsWithTimeout(timeout time.Duration) *UpdateInventoryManufacturerParams {
	var ()
	return &UpdateInventoryManufacturerParams{

		timeout: timeout,
	}
}

// NewUpdateInventoryManufacturerParamsWithContext creates a new UpdateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInventoryManufacturerParamsWithContext(ctx context.Context) *UpdateInventoryManufacturerParams {
	var ()
	return &UpdateInventoryManufacturerParams{

		Context: ctx,
	}
}

// NewUpdateInventoryManufacturerParamsWithHTTPClient creates a new UpdateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInventoryManufacturerParamsWithHTTPClient(client *http.Client) *UpdateInventoryManufacturerParams {
	var ()
	return &UpdateInventoryManufacturerParams{
		HTTPClient: client,
	}
}

/*UpdateInventoryManufacturerParams contains all the parameters to send to the API endpoint
for the update inventory manufacturer operation typically these are written to a http.Request
*/
type UpdateInventoryManufacturerParams struct {

	/*Body*/
	Body *models.UpdateInventoryManufacturer
	/*ManufacturerID*/
	ManufacturerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) WithTimeout(timeout time.Duration) *UpdateInventoryManufacturerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) WithContext(ctx context.Context) *UpdateInventoryManufacturerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) WithHTTPClient(client *http.Client) *UpdateInventoryManufacturerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) WithBody(body *models.UpdateInventoryManufacturer) *UpdateInventoryManufacturerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) SetBody(body *models.UpdateInventoryManufacturer) {
	o.Body = body
}

// WithManufacturerID adds the manufacturerID to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) WithManufacturerID(manufacturerID string) *UpdateInventoryManufacturerParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the update inventory manufacturer params
func (o *UpdateInventoryManufacturerParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInventoryManufacturerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param manufacturerId
	if err := r.SetPathParam("manufacturerId", o.ManufacturerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
