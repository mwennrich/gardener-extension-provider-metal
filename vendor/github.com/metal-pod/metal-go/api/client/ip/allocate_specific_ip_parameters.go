// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// NewAllocateSpecificIPParams creates a new AllocateSpecificIPParams object
// with the default values initialized.
func NewAllocateSpecificIPParams() *AllocateSpecificIPParams {
	var ()
	return &AllocateSpecificIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateSpecificIPParamsWithTimeout creates a new AllocateSpecificIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllocateSpecificIPParamsWithTimeout(timeout time.Duration) *AllocateSpecificIPParams {
	var ()
	return &AllocateSpecificIPParams{

		timeout: timeout,
	}
}

// NewAllocateSpecificIPParamsWithContext creates a new AllocateSpecificIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllocateSpecificIPParamsWithContext(ctx context.Context) *AllocateSpecificIPParams {
	var ()
	return &AllocateSpecificIPParams{

		Context: ctx,
	}
}

// NewAllocateSpecificIPParamsWithHTTPClient creates a new AllocateSpecificIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllocateSpecificIPParamsWithHTTPClient(client *http.Client) *AllocateSpecificIPParams {
	var ()
	return &AllocateSpecificIPParams{
		HTTPClient: client,
	}
}

/*AllocateSpecificIPParams contains all the parameters to send to the API endpoint
for the allocate specific IP operation typically these are written to a http.Request
*/
type AllocateSpecificIPParams struct {

	/*Body*/
	Body *models.V1IPAllocateRequest
	/*IP
	  ip to try to allocate

	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the allocate specific IP params
func (o *AllocateSpecificIPParams) WithTimeout(timeout time.Duration) *AllocateSpecificIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate specific IP params
func (o *AllocateSpecificIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate specific IP params
func (o *AllocateSpecificIPParams) WithContext(ctx context.Context) *AllocateSpecificIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate specific IP params
func (o *AllocateSpecificIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate specific IP params
func (o *AllocateSpecificIPParams) WithHTTPClient(client *http.Client) *AllocateSpecificIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate specific IP params
func (o *AllocateSpecificIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allocate specific IP params
func (o *AllocateSpecificIPParams) WithBody(body *models.V1IPAllocateRequest) *AllocateSpecificIPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allocate specific IP params
func (o *AllocateSpecificIPParams) SetBody(body *models.V1IPAllocateRequest) {
	o.Body = body
}

// WithIP adds the ip to the allocate specific IP params
func (o *AllocateSpecificIPParams) WithIP(ip string) *AllocateSpecificIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the allocate specific IP params
func (o *AllocateSpecificIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateSpecificIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
