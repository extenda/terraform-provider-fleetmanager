// Code generated by go-swagger; DO NOT EDIT.

package fleet

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

// NewListSoftwareProfilesForBrandParams creates a new ListSoftwareProfilesForBrandParams object
// with the default values initialized.
func NewListSoftwareProfilesForBrandParams() *ListSoftwareProfilesForBrandParams {
	var ()
	return &ListSoftwareProfilesForBrandParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSoftwareProfilesForBrandParamsWithTimeout creates a new ListSoftwareProfilesForBrandParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSoftwareProfilesForBrandParamsWithTimeout(timeout time.Duration) *ListSoftwareProfilesForBrandParams {
	var ()
	return &ListSoftwareProfilesForBrandParams{

		timeout: timeout,
	}
}

// NewListSoftwareProfilesForBrandParamsWithContext creates a new ListSoftwareProfilesForBrandParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSoftwareProfilesForBrandParamsWithContext(ctx context.Context) *ListSoftwareProfilesForBrandParams {
	var ()
	return &ListSoftwareProfilesForBrandParams{

		Context: ctx,
	}
}

// NewListSoftwareProfilesForBrandParamsWithHTTPClient creates a new ListSoftwareProfilesForBrandParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSoftwareProfilesForBrandParamsWithHTTPClient(client *http.Client) *ListSoftwareProfilesForBrandParams {
	var ()
	return &ListSoftwareProfilesForBrandParams{
		HTTPClient: client,
	}
}

/*ListSoftwareProfilesForBrandParams contains all the parameters to send to the API endpoint
for the list software profiles for brand operation typically these are written to a http.Request
*/
type ListSoftwareProfilesForBrandParams struct {

	/*BrandID*/
	BrandID string
	/*NextToken*/
	NextToken *string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithTimeout(timeout time.Duration) *ListSoftwareProfilesForBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithContext(ctx context.Context) *ListSoftwareProfilesForBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithHTTPClient(client *http.Client) *ListSoftwareProfilesForBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithBrandID(brandID string) *ListSoftwareProfilesForBrandParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithNextToken adds the nextToken to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithNextToken(nextToken *string) *ListSoftwareProfilesForBrandParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithTenantID adds the tenantID to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) WithTenantID(tenantID string) *ListSoftwareProfilesForBrandParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list software profiles for brand params
func (o *ListSoftwareProfilesForBrandParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSoftwareProfilesForBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
		return err
	}

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

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}