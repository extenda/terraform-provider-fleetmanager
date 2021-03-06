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

// NewLinkSoftwareProfileWithSoftwareVersionParams creates a new LinkSoftwareProfileWithSoftwareVersionParams object
// with the default values initialized.
func NewLinkSoftwareProfileWithSoftwareVersionParams() *LinkSoftwareProfileWithSoftwareVersionParams {
	var ()
	return &LinkSoftwareProfileWithSoftwareVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLinkSoftwareProfileWithSoftwareVersionParamsWithTimeout creates a new LinkSoftwareProfileWithSoftwareVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLinkSoftwareProfileWithSoftwareVersionParamsWithTimeout(timeout time.Duration) *LinkSoftwareProfileWithSoftwareVersionParams {
	var ()
	return &LinkSoftwareProfileWithSoftwareVersionParams{

		timeout: timeout,
	}
}

// NewLinkSoftwareProfileWithSoftwareVersionParamsWithContext creates a new LinkSoftwareProfileWithSoftwareVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewLinkSoftwareProfileWithSoftwareVersionParamsWithContext(ctx context.Context) *LinkSoftwareProfileWithSoftwareVersionParams {
	var ()
	return &LinkSoftwareProfileWithSoftwareVersionParams{

		Context: ctx,
	}
}

// NewLinkSoftwareProfileWithSoftwareVersionParamsWithHTTPClient creates a new LinkSoftwareProfileWithSoftwareVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLinkSoftwareProfileWithSoftwareVersionParamsWithHTTPClient(client *http.Client) *LinkSoftwareProfileWithSoftwareVersionParams {
	var ()
	return &LinkSoftwareProfileWithSoftwareVersionParams{
		HTTPClient: client,
	}
}

/*LinkSoftwareProfileWithSoftwareVersionParams contains all the parameters to send to the API endpoint
for the link software profile with software version operation typically these are written to a http.Request
*/
type LinkSoftwareProfileWithSoftwareVersionParams struct {

	/*Body*/
	Body *models.LinkFleetSoftwareProfile2SoftwareVersion
	/*BrandID*/
	BrandID string
	/*SoftwareProfileID*/
	SoftwareProfileID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithTimeout(timeout time.Duration) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithContext(ctx context.Context) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithHTTPClient(client *http.Client) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithBody(body *models.LinkFleetSoftwareProfile2SoftwareVersion) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetBody(body *models.LinkFleetSoftwareProfile2SoftwareVersion) {
	o.Body = body
}

// WithBrandID adds the brandID to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithBrandID(brandID string) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithSoftwareProfileID adds the softwareProfileID to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithSoftwareProfileID(softwareProfileID string) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetSoftwareProfileID(softwareProfileID)
	return o
}

// SetSoftwareProfileID adds the softwareProfileId to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetSoftwareProfileID(softwareProfileID string) {
	o.SoftwareProfileID = softwareProfileID
}

// WithTenantID adds the tenantID to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WithTenantID(tenantID string) *LinkSoftwareProfileWithSoftwareVersionParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the link software profile with software version params
func (o *LinkSoftwareProfileWithSoftwareVersionParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *LinkSoftwareProfileWithSoftwareVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param softwareProfileId
	if err := r.SetPathParam("softwareProfileId", o.SoftwareProfileID); err != nil {
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
