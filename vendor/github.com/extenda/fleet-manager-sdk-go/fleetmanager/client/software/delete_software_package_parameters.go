// Code generated by go-swagger; DO NOT EDIT.

package software

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

// NewDeleteSoftwarePackageParams creates a new DeleteSoftwarePackageParams object
// with the default values initialized.
func NewDeleteSoftwarePackageParams() *DeleteSoftwarePackageParams {
	var ()
	return &DeleteSoftwarePackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSoftwarePackageParamsWithTimeout creates a new DeleteSoftwarePackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSoftwarePackageParamsWithTimeout(timeout time.Duration) *DeleteSoftwarePackageParams {
	var ()
	return &DeleteSoftwarePackageParams{

		timeout: timeout,
	}
}

// NewDeleteSoftwarePackageParamsWithContext creates a new DeleteSoftwarePackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSoftwarePackageParamsWithContext(ctx context.Context) *DeleteSoftwarePackageParams {
	var ()
	return &DeleteSoftwarePackageParams{

		Context: ctx,
	}
}

// NewDeleteSoftwarePackageParamsWithHTTPClient creates a new DeleteSoftwarePackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSoftwarePackageParamsWithHTTPClient(client *http.Client) *DeleteSoftwarePackageParams {
	var ()
	return &DeleteSoftwarePackageParams{
		HTTPClient: client,
	}
}

/*DeleteSoftwarePackageParams contains all the parameters to send to the API endpoint
for the delete software package operation typically these are written to a http.Request
*/
type DeleteSoftwarePackageParams struct {

	/*PackageID*/
	PackageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete software package params
func (o *DeleteSoftwarePackageParams) WithTimeout(timeout time.Duration) *DeleteSoftwarePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete software package params
func (o *DeleteSoftwarePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete software package params
func (o *DeleteSoftwarePackageParams) WithContext(ctx context.Context) *DeleteSoftwarePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete software package params
func (o *DeleteSoftwarePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete software package params
func (o *DeleteSoftwarePackageParams) WithHTTPClient(client *http.Client) *DeleteSoftwarePackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete software package params
func (o *DeleteSoftwarePackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the delete software package params
func (o *DeleteSoftwarePackageParams) WithPackageID(packageID string) *DeleteSoftwarePackageParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the delete software package params
func (o *DeleteSoftwarePackageParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSoftwarePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
