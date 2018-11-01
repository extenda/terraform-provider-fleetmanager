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

// NewDeleteDriverBinaryParams creates a new DeleteDriverBinaryParams object
// with the default values initialized.
func NewDeleteDriverBinaryParams() *DeleteDriverBinaryParams {
	var ()
	return &DeleteDriverBinaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDriverBinaryParamsWithTimeout creates a new DeleteDriverBinaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDriverBinaryParamsWithTimeout(timeout time.Duration) *DeleteDriverBinaryParams {
	var ()
	return &DeleteDriverBinaryParams{

		timeout: timeout,
	}
}

// NewDeleteDriverBinaryParamsWithContext creates a new DeleteDriverBinaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDriverBinaryParamsWithContext(ctx context.Context) *DeleteDriverBinaryParams {
	var ()
	return &DeleteDriverBinaryParams{

		Context: ctx,
	}
}

// NewDeleteDriverBinaryParamsWithHTTPClient creates a new DeleteDriverBinaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDriverBinaryParamsWithHTTPClient(client *http.Client) *DeleteDriverBinaryParams {
	var ()
	return &DeleteDriverBinaryParams{
		HTTPClient: client,
	}
}

/*DeleteDriverBinaryParams contains all the parameters to send to the API endpoint
for the delete driver binary operation typically these are written to a http.Request
*/
type DeleteDriverBinaryParams struct {

	/*PackageID*/
	PackageID string
	/*VersionID*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete driver binary params
func (o *DeleteDriverBinaryParams) WithTimeout(timeout time.Duration) *DeleteDriverBinaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete driver binary params
func (o *DeleteDriverBinaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete driver binary params
func (o *DeleteDriverBinaryParams) WithContext(ctx context.Context) *DeleteDriverBinaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete driver binary params
func (o *DeleteDriverBinaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete driver binary params
func (o *DeleteDriverBinaryParams) WithHTTPClient(client *http.Client) *DeleteDriverBinaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete driver binary params
func (o *DeleteDriverBinaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the delete driver binary params
func (o *DeleteDriverBinaryParams) WithPackageID(packageID string) *DeleteDriverBinaryParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the delete driver binary params
func (o *DeleteDriverBinaryParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithVersionID adds the versionID to the delete driver binary params
func (o *DeleteDriverBinaryParams) WithVersionID(versionID string) *DeleteDriverBinaryParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete driver binary params
func (o *DeleteDriverBinaryParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDriverBinaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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