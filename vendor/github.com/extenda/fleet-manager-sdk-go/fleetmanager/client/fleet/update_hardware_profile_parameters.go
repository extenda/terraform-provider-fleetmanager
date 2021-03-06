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

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewUpdateHardwareProfileParams creates a new UpdateHardwareProfileParams object
// with the default values initialized.
func NewUpdateHardwareProfileParams() *UpdateHardwareProfileParams {
	var ()
	return &UpdateHardwareProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHardwareProfileParamsWithTimeout creates a new UpdateHardwareProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHardwareProfileParamsWithTimeout(timeout time.Duration) *UpdateHardwareProfileParams {
	var ()
	return &UpdateHardwareProfileParams{

		timeout: timeout,
	}
}

// NewUpdateHardwareProfileParamsWithContext creates a new UpdateHardwareProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHardwareProfileParamsWithContext(ctx context.Context) *UpdateHardwareProfileParams {
	var ()
	return &UpdateHardwareProfileParams{

		Context: ctx,
	}
}

// NewUpdateHardwareProfileParamsWithHTTPClient creates a new UpdateHardwareProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHardwareProfileParamsWithHTTPClient(client *http.Client) *UpdateHardwareProfileParams {
	var ()
	return &UpdateHardwareProfileParams{
		HTTPClient: client,
	}
}

/*UpdateHardwareProfileParams contains all the parameters to send to the API endpoint
for the update hardware profile operation typically these are written to a http.Request
*/
type UpdateHardwareProfileParams struct {

	/*Body*/
	Body *models.UpdateFleetHardwareProfile
	/*BrandID*/
	BrandID string
	/*HardwareProfileID*/
	HardwareProfileID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithTimeout(timeout time.Duration) *UpdateHardwareProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithContext(ctx context.Context) *UpdateHardwareProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithHTTPClient(client *http.Client) *UpdateHardwareProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithBody(body *models.UpdateFleetHardwareProfile) *UpdateHardwareProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetBody(body *models.UpdateFleetHardwareProfile) {
	o.Body = body
}

// WithBrandID adds the brandID to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithBrandID(brandID string) *UpdateHardwareProfileParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithHardwareProfileID adds the hardwareProfileID to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithHardwareProfileID(hardwareProfileID string) *UpdateHardwareProfileParams {
	o.SetHardwareProfileID(hardwareProfileID)
	return o
}

// SetHardwareProfileID adds the hardwareProfileId to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetHardwareProfileID(hardwareProfileID string) {
	o.HardwareProfileID = hardwareProfileID
}

// WithTenantID adds the tenantID to the update hardware profile params
func (o *UpdateHardwareProfileParams) WithTenantID(tenantID string) *UpdateHardwareProfileParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update hardware profile params
func (o *UpdateHardwareProfileParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHardwareProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param hardwareProfileId
	if err := r.SetPathParam("hardwareProfileId", o.HardwareProfileID); err != nil {
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
