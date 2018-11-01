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

// NewGetDriverVersionByIDParams creates a new GetDriverVersionByIDParams object
// with the default values initialized.
func NewGetDriverVersionByIDParams() *GetDriverVersionByIDParams {
	var ()
	return &GetDriverVersionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDriverVersionByIDParamsWithTimeout creates a new GetDriverVersionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDriverVersionByIDParamsWithTimeout(timeout time.Duration) *GetDriverVersionByIDParams {
	var ()
	return &GetDriverVersionByIDParams{

		timeout: timeout,
	}
}

// NewGetDriverVersionByIDParamsWithContext creates a new GetDriverVersionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDriverVersionByIDParamsWithContext(ctx context.Context) *GetDriverVersionByIDParams {
	var ()
	return &GetDriverVersionByIDParams{

		Context: ctx,
	}
}

// NewGetDriverVersionByIDParamsWithHTTPClient creates a new GetDriverVersionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDriverVersionByIDParamsWithHTTPClient(client *http.Client) *GetDriverVersionByIDParams {
	var ()
	return &GetDriverVersionByIDParams{
		HTTPClient: client,
	}
}

/*GetDriverVersionByIDParams contains all the parameters to send to the API endpoint
for the get driver version by Id operation typically these are written to a http.Request
*/
type GetDriverVersionByIDParams struct {

	/*PackageID*/
	PackageID string
	/*VersionID*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get driver version by Id params
func (o *GetDriverVersionByIDParams) WithTimeout(timeout time.Duration) *GetDriverVersionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get driver version by Id params
func (o *GetDriverVersionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get driver version by Id params
func (o *GetDriverVersionByIDParams) WithContext(ctx context.Context) *GetDriverVersionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get driver version by Id params
func (o *GetDriverVersionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get driver version by Id params
func (o *GetDriverVersionByIDParams) WithHTTPClient(client *http.Client) *GetDriverVersionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get driver version by Id params
func (o *GetDriverVersionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the get driver version by Id params
func (o *GetDriverVersionByIDParams) WithPackageID(packageID string) *GetDriverVersionByIDParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get driver version by Id params
func (o *GetDriverVersionByIDParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithVersionID adds the versionID to the get driver version by Id params
func (o *GetDriverVersionByIDParams) WithVersionID(versionID string) *GetDriverVersionByIDParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get driver version by Id params
func (o *GetDriverVersionByIDParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDriverVersionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
