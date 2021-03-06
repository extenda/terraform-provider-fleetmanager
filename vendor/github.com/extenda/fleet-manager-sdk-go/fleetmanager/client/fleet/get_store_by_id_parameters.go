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

// NewGetStoreByIDParams creates a new GetStoreByIDParams object
// with the default values initialized.
func NewGetStoreByIDParams() *GetStoreByIDParams {
	var ()
	return &GetStoreByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoreByIDParamsWithTimeout creates a new GetStoreByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoreByIDParamsWithTimeout(timeout time.Duration) *GetStoreByIDParams {
	var ()
	return &GetStoreByIDParams{

		timeout: timeout,
	}
}

// NewGetStoreByIDParamsWithContext creates a new GetStoreByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStoreByIDParamsWithContext(ctx context.Context) *GetStoreByIDParams {
	var ()
	return &GetStoreByIDParams{

		Context: ctx,
	}
}

// NewGetStoreByIDParamsWithHTTPClient creates a new GetStoreByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStoreByIDParamsWithHTTPClient(client *http.Client) *GetStoreByIDParams {
	var ()
	return &GetStoreByIDParams{
		HTTPClient: client,
	}
}

/*GetStoreByIDParams contains all the parameters to send to the API endpoint
for the get store by Id operation typically these are written to a http.Request
*/
type GetStoreByIDParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*StoreID*/
	StoreID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get store by Id params
func (o *GetStoreByIDParams) WithTimeout(timeout time.Duration) *GetStoreByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get store by Id params
func (o *GetStoreByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get store by Id params
func (o *GetStoreByIDParams) WithContext(ctx context.Context) *GetStoreByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get store by Id params
func (o *GetStoreByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get store by Id params
func (o *GetStoreByIDParams) WithHTTPClient(client *http.Client) *GetStoreByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get store by Id params
func (o *GetStoreByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get store by Id params
func (o *GetStoreByIDParams) WithBrandID(brandID string) *GetStoreByIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get store by Id params
func (o *GetStoreByIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the get store by Id params
func (o *GetStoreByIDParams) WithCountryID(countryID string) *GetStoreByIDParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the get store by Id params
func (o *GetStoreByIDParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the get store by Id params
func (o *GetStoreByIDParams) WithStoreID(storeID string) *GetStoreByIDParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get store by Id params
func (o *GetStoreByIDParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the get store by Id params
func (o *GetStoreByIDParams) WithTenantID(tenantID string) *GetStoreByIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get store by Id params
func (o *GetStoreByIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoreByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
