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

// NewUploadDriverBinaryParams creates a new UploadDriverBinaryParams object
// with the default values initialized.
func NewUploadDriverBinaryParams() *UploadDriverBinaryParams {
	var ()
	return &UploadDriverBinaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDriverBinaryParamsWithTimeout creates a new UploadDriverBinaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadDriverBinaryParamsWithTimeout(timeout time.Duration) *UploadDriverBinaryParams {
	var ()
	return &UploadDriverBinaryParams{

		timeout: timeout,
	}
}

// NewUploadDriverBinaryParamsWithContext creates a new UploadDriverBinaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadDriverBinaryParamsWithContext(ctx context.Context) *UploadDriverBinaryParams {
	var ()
	return &UploadDriverBinaryParams{

		Context: ctx,
	}
}

// NewUploadDriverBinaryParamsWithHTTPClient creates a new UploadDriverBinaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadDriverBinaryParamsWithHTTPClient(client *http.Client) *UploadDriverBinaryParams {
	var ()
	return &UploadDriverBinaryParams{
		HTTPClient: client,
	}
}

/*UploadDriverBinaryParams contains all the parameters to send to the API endpoint
for the upload driver binary operation typically these are written to a http.Request
*/
type UploadDriverBinaryParams struct {

	/*PackageID*/
	PackageID string
	/*VersionID*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload driver binary params
func (o *UploadDriverBinaryParams) WithTimeout(timeout time.Duration) *UploadDriverBinaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload driver binary params
func (o *UploadDriverBinaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload driver binary params
func (o *UploadDriverBinaryParams) WithContext(ctx context.Context) *UploadDriverBinaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload driver binary params
func (o *UploadDriverBinaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload driver binary params
func (o *UploadDriverBinaryParams) WithHTTPClient(client *http.Client) *UploadDriverBinaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload driver binary params
func (o *UploadDriverBinaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the upload driver binary params
func (o *UploadDriverBinaryParams) WithPackageID(packageID string) *UploadDriverBinaryParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the upload driver binary params
func (o *UploadDriverBinaryParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithVersionID adds the versionID to the upload driver binary params
func (o *UploadDriverBinaryParams) WithVersionID(versionID string) *UploadDriverBinaryParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the upload driver binary params
func (o *UploadDriverBinaryParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDriverBinaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
