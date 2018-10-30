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

// NewUploadCountrySystemPropertiesParams creates a new UploadCountrySystemPropertiesParams object
// with the default values initialized.
func NewUploadCountrySystemPropertiesParams() *UploadCountrySystemPropertiesParams {
	var ()
	return &UploadCountrySystemPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadCountrySystemPropertiesParamsWithTimeout creates a new UploadCountrySystemPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadCountrySystemPropertiesParamsWithTimeout(timeout time.Duration) *UploadCountrySystemPropertiesParams {
	var ()
	return &UploadCountrySystemPropertiesParams{

		timeout: timeout,
	}
}

// NewUploadCountrySystemPropertiesParamsWithContext creates a new UploadCountrySystemPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadCountrySystemPropertiesParamsWithContext(ctx context.Context) *UploadCountrySystemPropertiesParams {
	var ()
	return &UploadCountrySystemPropertiesParams{

		Context: ctx,
	}
}

// NewUploadCountrySystemPropertiesParamsWithHTTPClient creates a new UploadCountrySystemPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadCountrySystemPropertiesParamsWithHTTPClient(client *http.Client) *UploadCountrySystemPropertiesParams {
	var ()
	return &UploadCountrySystemPropertiesParams{
		HTTPClient: client,
	}
}

/*UploadCountrySystemPropertiesParams contains all the parameters to send to the API endpoint
for the upload country system properties operation typically these are written to a http.Request
*/
type UploadCountrySystemPropertiesParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithTimeout(timeout time.Duration) *UploadCountrySystemPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithContext(ctx context.Context) *UploadCountrySystemPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithHTTPClient(client *http.Client) *UploadCountrySystemPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithBrandID(brandID string) *UploadCountrySystemPropertiesParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithCountryID(countryID string) *UploadCountrySystemPropertiesParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithTenantID adds the tenantID to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) WithTenantID(tenantID string) *UploadCountrySystemPropertiesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the upload country system properties params
func (o *UploadCountrySystemPropertiesParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadCountrySystemPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}