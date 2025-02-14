// Code generated by go-swagger; DO NOT EDIT.

package partition

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

// NewCreatePartitionParams creates a new CreatePartitionParams object
// with the default values initialized.
func NewCreatePartitionParams() *CreatePartitionParams {
	var ()
	return &CreatePartitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePartitionParamsWithTimeout creates a new CreatePartitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePartitionParamsWithTimeout(timeout time.Duration) *CreatePartitionParams {
	var ()
	return &CreatePartitionParams{

		timeout: timeout,
	}
}

// NewCreatePartitionParamsWithContext creates a new CreatePartitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePartitionParamsWithContext(ctx context.Context) *CreatePartitionParams {
	var ()
	return &CreatePartitionParams{

		Context: ctx,
	}
}

// NewCreatePartitionParamsWithHTTPClient creates a new CreatePartitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePartitionParamsWithHTTPClient(client *http.Client) *CreatePartitionParams {
	var ()
	return &CreatePartitionParams{
		HTTPClient: client,
	}
}

/*CreatePartitionParams contains all the parameters to send to the API endpoint
for the create partition operation typically these are written to a http.Request
*/
type CreatePartitionParams struct {

	/*Body*/
	Body *models.V1PartitionCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create partition params
func (o *CreatePartitionParams) WithTimeout(timeout time.Duration) *CreatePartitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create partition params
func (o *CreatePartitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create partition params
func (o *CreatePartitionParams) WithContext(ctx context.Context) *CreatePartitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create partition params
func (o *CreatePartitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create partition params
func (o *CreatePartitionParams) WithHTTPClient(client *http.Client) *CreatePartitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create partition params
func (o *CreatePartitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create partition params
func (o *CreatePartitionParams) WithBody(body *models.V1PartitionCreateRequest) *CreatePartitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create partition params
func (o *CreatePartitionParams) SetBody(body *models.V1PartitionCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePartitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
