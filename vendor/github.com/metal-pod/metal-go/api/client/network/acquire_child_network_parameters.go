// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewAcquireChildNetworkParams creates a new AcquireChildNetworkParams object
// with the default values initialized.
func NewAcquireChildNetworkParams() *AcquireChildNetworkParams {
	var ()
	return &AcquireChildNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcquireChildNetworkParamsWithTimeout creates a new AcquireChildNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcquireChildNetworkParamsWithTimeout(timeout time.Duration) *AcquireChildNetworkParams {
	var ()
	return &AcquireChildNetworkParams{

		timeout: timeout,
	}
}

// NewAcquireChildNetworkParamsWithContext creates a new AcquireChildNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcquireChildNetworkParamsWithContext(ctx context.Context) *AcquireChildNetworkParams {
	var ()
	return &AcquireChildNetworkParams{

		Context: ctx,
	}
}

// NewAcquireChildNetworkParamsWithHTTPClient creates a new AcquireChildNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcquireChildNetworkParamsWithHTTPClient(client *http.Client) *AcquireChildNetworkParams {
	var ()
	return &AcquireChildNetworkParams{
		HTTPClient: client,
	}
}

/*AcquireChildNetworkParams contains all the parameters to send to the API endpoint
for the acquire child network operation typically these are written to a http.Request
*/
type AcquireChildNetworkParams struct {

	/*Body*/
	Body *models.V1NetworkAcquireRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the acquire child network params
func (o *AcquireChildNetworkParams) WithTimeout(timeout time.Duration) *AcquireChildNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the acquire child network params
func (o *AcquireChildNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the acquire child network params
func (o *AcquireChildNetworkParams) WithContext(ctx context.Context) *AcquireChildNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the acquire child network params
func (o *AcquireChildNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the acquire child network params
func (o *AcquireChildNetworkParams) WithHTTPClient(client *http.Client) *AcquireChildNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the acquire child network params
func (o *AcquireChildNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the acquire child network params
func (o *AcquireChildNetworkParams) WithBody(body *models.V1NetworkAcquireRequest) *AcquireChildNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the acquire child network params
func (o *AcquireChildNetworkParams) SetBody(body *models.V1NetworkAcquireRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AcquireChildNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
