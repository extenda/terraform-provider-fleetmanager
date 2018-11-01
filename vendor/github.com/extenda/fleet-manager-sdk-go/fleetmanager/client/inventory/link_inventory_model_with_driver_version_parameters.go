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

// NewLinkInventoryModelWithDriverVersionParams creates a new LinkInventoryModelWithDriverVersionParams object
// with the default values initialized.
func NewLinkInventoryModelWithDriverVersionParams() *LinkInventoryModelWithDriverVersionParams {
	var ()
	return &LinkInventoryModelWithDriverVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLinkInventoryModelWithDriverVersionParamsWithTimeout creates a new LinkInventoryModelWithDriverVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLinkInventoryModelWithDriverVersionParamsWithTimeout(timeout time.Duration) *LinkInventoryModelWithDriverVersionParams {
	var ()
	return &LinkInventoryModelWithDriverVersionParams{

		timeout: timeout,
	}
}

// NewLinkInventoryModelWithDriverVersionParamsWithContext creates a new LinkInventoryModelWithDriverVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewLinkInventoryModelWithDriverVersionParamsWithContext(ctx context.Context) *LinkInventoryModelWithDriverVersionParams {
	var ()
	return &LinkInventoryModelWithDriverVersionParams{

		Context: ctx,
	}
}

// NewLinkInventoryModelWithDriverVersionParamsWithHTTPClient creates a new LinkInventoryModelWithDriverVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLinkInventoryModelWithDriverVersionParamsWithHTTPClient(client *http.Client) *LinkInventoryModelWithDriverVersionParams {
	var ()
	return &LinkInventoryModelWithDriverVersionParams{
		HTTPClient: client,
	}
}

/*LinkInventoryModelWithDriverVersionParams contains all the parameters to send to the API endpoint
for the link inventory model with driver version operation typically these are written to a http.Request
*/
type LinkInventoryModelWithDriverVersionParams struct {

	/*Body*/
	Body *models.LinkInventoryModel2DriverVersion
	/*ManufacturerID*/
	ManufacturerID string
	/*ModelID*/
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithTimeout(timeout time.Duration) *LinkInventoryModelWithDriverVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithContext(ctx context.Context) *LinkInventoryModelWithDriverVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithHTTPClient(client *http.Client) *LinkInventoryModelWithDriverVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithBody(body *models.LinkInventoryModel2DriverVersion) *LinkInventoryModelWithDriverVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetBody(body *models.LinkInventoryModel2DriverVersion) {
	o.Body = body
}

// WithManufacturerID adds the manufacturerID to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithManufacturerID(manufacturerID string) *LinkInventoryModelWithDriverVersionParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WithModelID adds the modelID to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) WithModelID(modelID string) *LinkInventoryModelWithDriverVersionParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the link inventory model with driver version params
func (o *LinkInventoryModelWithDriverVersionParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *LinkInventoryModelWithDriverVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
