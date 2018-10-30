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

// NewListWorkstationsParams creates a new ListWorkstationsParams object
// with the default values initialized.
func NewListWorkstationsParams() *ListWorkstationsParams {
	var ()
	return &ListWorkstationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListWorkstationsParamsWithTimeout creates a new ListWorkstationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListWorkstationsParamsWithTimeout(timeout time.Duration) *ListWorkstationsParams {
	var ()
	return &ListWorkstationsParams{

		timeout: timeout,
	}
}

// NewListWorkstationsParamsWithContext creates a new ListWorkstationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListWorkstationsParamsWithContext(ctx context.Context) *ListWorkstationsParams {
	var ()
	return &ListWorkstationsParams{

		Context: ctx,
	}
}

// NewListWorkstationsParamsWithHTTPClient creates a new ListWorkstationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListWorkstationsParamsWithHTTPClient(client *http.Client) *ListWorkstationsParams {
	var ()
	return &ListWorkstationsParams{
		HTTPClient: client,
	}
}

/*ListWorkstationsParams contains all the parameters to send to the API endpoint
for the list workstations operation typically these are written to a http.Request
*/
type ListWorkstationsParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*NextToken*/
	NextToken *string
	/*StoreID*/
	StoreID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list workstations params
func (o *ListWorkstationsParams) WithTimeout(timeout time.Duration) *ListWorkstationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list workstations params
func (o *ListWorkstationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list workstations params
func (o *ListWorkstationsParams) WithContext(ctx context.Context) *ListWorkstationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list workstations params
func (o *ListWorkstationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list workstations params
func (o *ListWorkstationsParams) WithHTTPClient(client *http.Client) *ListWorkstationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list workstations params
func (o *ListWorkstationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the list workstations params
func (o *ListWorkstationsParams) WithBrandID(brandID string) *ListWorkstationsParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the list workstations params
func (o *ListWorkstationsParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the list workstations params
func (o *ListWorkstationsParams) WithCountryID(countryID string) *ListWorkstationsParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the list workstations params
func (o *ListWorkstationsParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithNextToken adds the nextToken to the list workstations params
func (o *ListWorkstationsParams) WithNextToken(nextToken *string) *ListWorkstationsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list workstations params
func (o *ListWorkstationsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithStoreID adds the storeID to the list workstations params
func (o *ListWorkstationsParams) WithStoreID(storeID string) *ListWorkstationsParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the list workstations params
func (o *ListWorkstationsParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the list workstations params
func (o *ListWorkstationsParams) WithTenantID(tenantID string) *ListWorkstationsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list workstations params
func (o *ListWorkstationsParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *ListWorkstationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
		return err
	}

	// path param countryId
	if err := r.SetPathParam("countryId", o.CountryID); err != nil {
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

	// path param storeId
	if err := r.SetPathParam("storeId", o.StoreID); err != nil {
		return err
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