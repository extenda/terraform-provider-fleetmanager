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

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// NewUpdateWorkstationParams creates a new UpdateWorkstationParams object
// with the default values initialized.
func NewUpdateWorkstationParams() *UpdateWorkstationParams {
	var ()
	return &UpdateWorkstationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkstationParamsWithTimeout creates a new UpdateWorkstationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWorkstationParamsWithTimeout(timeout time.Duration) *UpdateWorkstationParams {
	var ()
	return &UpdateWorkstationParams{

		timeout: timeout,
	}
}

// NewUpdateWorkstationParamsWithContext creates a new UpdateWorkstationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWorkstationParamsWithContext(ctx context.Context) *UpdateWorkstationParams {
	var ()
	return &UpdateWorkstationParams{

		Context: ctx,
	}
}

// NewUpdateWorkstationParamsWithHTTPClient creates a new UpdateWorkstationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWorkstationParamsWithHTTPClient(client *http.Client) *UpdateWorkstationParams {
	var ()
	return &UpdateWorkstationParams{
		HTTPClient: client,
	}
}

/*UpdateWorkstationParams contains all the parameters to send to the API endpoint
for the update workstation operation typically these are written to a http.Request
*/
type UpdateWorkstationParams struct {

	/*Body*/
	Body *models.UpdateFleetWorkstation
	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*StoreID*/
	StoreID string
	/*TenantID*/
	TenantID string
	/*WorkstationID*/
	WorkstationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update workstation params
func (o *UpdateWorkstationParams) WithTimeout(timeout time.Duration) *UpdateWorkstationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workstation params
func (o *UpdateWorkstationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workstation params
func (o *UpdateWorkstationParams) WithContext(ctx context.Context) *UpdateWorkstationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workstation params
func (o *UpdateWorkstationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workstation params
func (o *UpdateWorkstationParams) WithHTTPClient(client *http.Client) *UpdateWorkstationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workstation params
func (o *UpdateWorkstationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update workstation params
func (o *UpdateWorkstationParams) WithBody(body *models.UpdateFleetWorkstation) *UpdateWorkstationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update workstation params
func (o *UpdateWorkstationParams) SetBody(body *models.UpdateFleetWorkstation) {
	o.Body = body
}

// WithBrandID adds the brandID to the update workstation params
func (o *UpdateWorkstationParams) WithBrandID(brandID string) *UpdateWorkstationParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the update workstation params
func (o *UpdateWorkstationParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the update workstation params
func (o *UpdateWorkstationParams) WithCountryID(countryID string) *UpdateWorkstationParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the update workstation params
func (o *UpdateWorkstationParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the update workstation params
func (o *UpdateWorkstationParams) WithStoreID(storeID string) *UpdateWorkstationParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the update workstation params
func (o *UpdateWorkstationParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the update workstation params
func (o *UpdateWorkstationParams) WithTenantID(tenantID string) *UpdateWorkstationParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update workstation params
func (o *UpdateWorkstationParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the update workstation params
func (o *UpdateWorkstationParams) WithWorkstationID(workstationID string) *UpdateWorkstationParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the update workstation params
func (o *UpdateWorkstationParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkstationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param workstationId
	if err := r.SetPathParam("workstationId", o.WorkstationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}