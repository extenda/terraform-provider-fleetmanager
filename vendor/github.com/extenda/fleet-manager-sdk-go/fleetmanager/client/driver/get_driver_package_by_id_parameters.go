// Code generated by go-swagger; DO NOT EDIT.

package driver

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

// NewGetDriverPackageByIDParams creates a new GetDriverPackageByIDParams object
// with the default values initialized.
func NewGetDriverPackageByIDParams() *GetDriverPackageByIDParams {
	var ()
	return &GetDriverPackageByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDriverPackageByIDParamsWithTimeout creates a new GetDriverPackageByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDriverPackageByIDParamsWithTimeout(timeout time.Duration) *GetDriverPackageByIDParams {
	var ()
	return &GetDriverPackageByIDParams{

		timeout: timeout,
	}
}

// NewGetDriverPackageByIDParamsWithContext creates a new GetDriverPackageByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDriverPackageByIDParamsWithContext(ctx context.Context) *GetDriverPackageByIDParams {
	var ()
	return &GetDriverPackageByIDParams{

		Context: ctx,
	}
}

// NewGetDriverPackageByIDParamsWithHTTPClient creates a new GetDriverPackageByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDriverPackageByIDParamsWithHTTPClient(client *http.Client) *GetDriverPackageByIDParams {
	var ()
	return &GetDriverPackageByIDParams{
		HTTPClient: client,
	}
}

/*GetDriverPackageByIDParams contains all the parameters to send to the API endpoint
for the get driver package by Id operation typically these are written to a http.Request
*/
type GetDriverPackageByIDParams struct {

	/*PackageID*/
	PackageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get driver package by Id params
func (o *GetDriverPackageByIDParams) WithTimeout(timeout time.Duration) *GetDriverPackageByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get driver package by Id params
func (o *GetDriverPackageByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get driver package by Id params
func (o *GetDriverPackageByIDParams) WithContext(ctx context.Context) *GetDriverPackageByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get driver package by Id params
func (o *GetDriverPackageByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get driver package by Id params
func (o *GetDriverPackageByIDParams) WithHTTPClient(client *http.Client) *GetDriverPackageByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get driver package by Id params
func (o *GetDriverPackageByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the get driver package by Id params
func (o *GetDriverPackageByIDParams) WithPackageID(packageID string) *GetDriverPackageByIDParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get driver package by Id params
func (o *GetDriverPackageByIDParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDriverPackageByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
