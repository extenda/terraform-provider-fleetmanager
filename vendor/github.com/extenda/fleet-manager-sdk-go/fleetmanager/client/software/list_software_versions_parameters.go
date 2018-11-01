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

// NewListSoftwareVersionsParams creates a new ListSoftwareVersionsParams object
// with the default values initialized.
func NewListSoftwareVersionsParams() *ListSoftwareVersionsParams {
	var ()
	return &ListSoftwareVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSoftwareVersionsParamsWithTimeout creates a new ListSoftwareVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSoftwareVersionsParamsWithTimeout(timeout time.Duration) *ListSoftwareVersionsParams {
	var ()
	return &ListSoftwareVersionsParams{

		timeout: timeout,
	}
}

// NewListSoftwareVersionsParamsWithContext creates a new ListSoftwareVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSoftwareVersionsParamsWithContext(ctx context.Context) *ListSoftwareVersionsParams {
	var ()
	return &ListSoftwareVersionsParams{

		Context: ctx,
	}
}

// NewListSoftwareVersionsParamsWithHTTPClient creates a new ListSoftwareVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSoftwareVersionsParamsWithHTTPClient(client *http.Client) *ListSoftwareVersionsParams {
	var ()
	return &ListSoftwareVersionsParams{
		HTTPClient: client,
	}
}

/*ListSoftwareVersionsParams contains all the parameters to send to the API endpoint
for the list software versions operation typically these are written to a http.Request
*/
type ListSoftwareVersionsParams struct {

	/*NextToken*/
	NextToken *string
	/*PackageID*/
	PackageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list software versions params
func (o *ListSoftwareVersionsParams) WithTimeout(timeout time.Duration) *ListSoftwareVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list software versions params
func (o *ListSoftwareVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list software versions params
func (o *ListSoftwareVersionsParams) WithContext(ctx context.Context) *ListSoftwareVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list software versions params
func (o *ListSoftwareVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list software versions params
func (o *ListSoftwareVersionsParams) WithHTTPClient(client *http.Client) *ListSoftwareVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list software versions params
func (o *ListSoftwareVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextToken adds the nextToken to the list software versions params
func (o *ListSoftwareVersionsParams) WithNextToken(nextToken *string) *ListSoftwareVersionsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list software versions params
func (o *ListSoftwareVersionsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithPackageID adds the packageID to the list software versions params
func (o *ListSoftwareVersionsParams) WithPackageID(packageID string) *ListSoftwareVersionsParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the list software versions params
func (o *ListSoftwareVersionsParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSoftwareVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NextToken != nil {

		// query param nextToken
		var qrNextToken string
		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {
			if err := r.SetQueryParam("nextToken", qNextToken); err != nil {
				return err
			}
		}

	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
