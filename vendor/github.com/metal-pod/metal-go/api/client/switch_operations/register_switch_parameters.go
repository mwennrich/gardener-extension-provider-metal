// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewRegisterSwitchParams creates a new RegisterSwitchParams object
// with the default values initialized.
func NewRegisterSwitchParams() *RegisterSwitchParams {
	var ()
	return &RegisterSwitchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterSwitchParamsWithTimeout creates a new RegisterSwitchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterSwitchParamsWithTimeout(timeout time.Duration) *RegisterSwitchParams {
	var ()
	return &RegisterSwitchParams{

		timeout: timeout,
	}
}

// NewRegisterSwitchParamsWithContext creates a new RegisterSwitchParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterSwitchParamsWithContext(ctx context.Context) *RegisterSwitchParams {
	var ()
	return &RegisterSwitchParams{

		Context: ctx,
	}
}

// NewRegisterSwitchParamsWithHTTPClient creates a new RegisterSwitchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterSwitchParamsWithHTTPClient(client *http.Client) *RegisterSwitchParams {
	var ()
	return &RegisterSwitchParams{
		HTTPClient: client,
	}
}

/*RegisterSwitchParams contains all the parameters to send to the API endpoint
for the register switch operation typically these are written to a http.Request
*/
type RegisterSwitchParams struct {

	/*Body*/
	Body *models.V1SwitchRegisterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register switch params
func (o *RegisterSwitchParams) WithTimeout(timeout time.Duration) *RegisterSwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register switch params
func (o *RegisterSwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register switch params
func (o *RegisterSwitchParams) WithContext(ctx context.Context) *RegisterSwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register switch params
func (o *RegisterSwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register switch params
func (o *RegisterSwitchParams) WithHTTPClient(client *http.Client) *RegisterSwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register switch params
func (o *RegisterSwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register switch params
func (o *RegisterSwitchParams) WithBody(body *models.V1SwitchRegisterRequest) *RegisterSwitchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register switch params
func (o *RegisterSwitchParams) SetBody(body *models.V1SwitchRegisterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterSwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
