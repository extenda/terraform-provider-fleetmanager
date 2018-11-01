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
)

// NewDeleteInventoryModelJposEntriesParams creates a new DeleteInventoryModelJposEntriesParams object
// with the default values initialized.
func NewDeleteInventoryModelJposEntriesParams() *DeleteInventoryModelJposEntriesParams {
	var ()
	return &DeleteInventoryModelJposEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInventoryModelJposEntriesParamsWithTimeout creates a new DeleteInventoryModelJposEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInventoryModelJposEntriesParamsWithTimeout(timeout time.Duration) *DeleteInventoryModelJposEntriesParams {
	var ()
	return &DeleteInventoryModelJposEntriesParams{

		timeout: timeout,
	}
}

// NewDeleteInventoryModelJposEntriesParamsWithContext creates a new DeleteInventoryModelJposEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInventoryModelJposEntriesParamsWithContext(ctx context.Context) *DeleteInventoryModelJposEntriesParams {
	var ()
	return &DeleteInventoryModelJposEntriesParams{

		Context: ctx,
	}
}

// NewDeleteInventoryModelJposEntriesParamsWithHTTPClient creates a new DeleteInventoryModelJposEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInventoryModelJposEntriesParamsWithHTTPClient(client *http.Client) *DeleteInventoryModelJposEntriesParams {
	var ()
	return &DeleteInventoryModelJposEntriesParams{
		HTTPClient: client,
	}
}

/*DeleteInventoryModelJposEntriesParams contains all the parameters to send to the API endpoint
for the delete inventory model jpos entries operation typically these are written to a http.Request
*/
type DeleteInventoryModelJposEntriesParams struct {

	/*ManufacturerID*/
	ManufacturerID string
	/*ModelID*/
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) WithTimeout(timeout time.Duration) *DeleteInventoryModelJposEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) WithContext(ctx context.Context) *DeleteInventoryModelJposEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) WithHTTPClient(client *http.Client) *DeleteInventoryModelJposEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManufacturerID adds the manufacturerID to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) WithManufacturerID(manufacturerID string) *DeleteInventoryModelJposEntriesParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WithModelID adds the modelID to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) WithModelID(modelID string) *DeleteInventoryModelJposEntriesParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the delete inventory model jpos entries params
func (o *DeleteInventoryModelJposEntriesParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInventoryModelJposEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
