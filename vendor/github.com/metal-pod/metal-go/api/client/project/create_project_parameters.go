// Code generated by go-swagger; DO NOT EDIT.

package project

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

	models "github.com/metal-pod/metal-go/api/models"
)

// NewCreateProjectParams creates a new CreateProjectParams object
// with the default values initialized.
func NewCreateProjectParams() *CreateProjectParams {
	var ()
	return &CreateProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectParamsWithTimeout creates a new CreateProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectParamsWithTimeout(timeout time.Duration) *CreateProjectParams {
	var ()
	return &CreateProjectParams{

		timeout: timeout,
	}
}

// NewCreateProjectParamsWithContext creates a new CreateProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectParamsWithContext(ctx context.Context) *CreateProjectParams {
	var ()
	return &CreateProjectParams{

		Context: ctx,
	}
}

// NewCreateProjectParamsWithHTTPClient creates a new CreateProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectParamsWithHTTPClient(client *http.Client) *CreateProjectParams {
	var ()
	return &CreateProjectParams{
		HTTPClient: client,
	}
}

/*CreateProjectParams contains all the parameters to send to the API endpoint
for the create project operation typically these are written to a http.Request
*/
type CreateProjectParams struct {

	/*Body*/
	Body *models.V1ProjectCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create project params
func (o *CreateProjectParams) WithTimeout(timeout time.Duration) *CreateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project params
func (o *CreateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project params
func (o *CreateProjectParams) WithContext(ctx context.Context) *CreateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project params
func (o *CreateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project params
func (o *CreateProjectParams) WithHTTPClient(client *http.Client) *CreateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project params
func (o *CreateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create project params
func (o *CreateProjectParams) WithBody(body *models.V1ProjectCreateRequest) *CreateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create project params
func (o *CreateProjectParams) SetBody(body *models.V1ProjectCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
