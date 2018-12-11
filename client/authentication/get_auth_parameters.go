// Code generated by go-swagger; DO NOT EDIT.

package authentication

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
)

// NewGetAuthParams creates a new GetAuthParams object
// with the default values initialized.
func NewGetAuthParams() *GetAuthParams {
	var ()
	return &GetAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthParamsWithTimeout creates a new GetAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthParamsWithTimeout(timeout time.Duration) *GetAuthParams {
	var ()
	return &GetAuthParams{

		timeout: timeout,
	}
}

// NewGetAuthParamsWithContext creates a new GetAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthParamsWithContext(ctx context.Context) *GetAuthParams {
	var ()
	return &GetAuthParams{

		Context: ctx,
	}
}

// NewGetAuthParamsWithHTTPClient creates a new GetAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthParamsWithHTTPClient(client *http.Client) *GetAuthParams {
	var ()
	return &GetAuthParams{
		HTTPClient: client,
	}
}

/*GetAuthParams contains all the parameters to send to the API endpoint
for the get auth operation typically these are written to a http.Request
*/
type GetAuthParams struct {

	/*Grant*/
	Grant string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth params
func (o *GetAuthParams) WithTimeout(timeout time.Duration) *GetAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth params
func (o *GetAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth params
func (o *GetAuthParams) WithContext(ctx context.Context) *GetAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth params
func (o *GetAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth params
func (o *GetAuthParams) WithHTTPClient(client *http.Client) *GetAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth params
func (o *GetAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrant adds the grant to the get auth params
func (o *GetAuthParams) WithGrant(grant string) *GetAuthParams {
	o.SetGrant(grant)
	return o
}

// SetGrant adds the grant to the get auth params
func (o *GetAuthParams) SetGrant(grant string) {
	o.Grant = grant
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param grant
	qrGrant := o.Grant
	qGrant := qrGrant
	if qGrant != "" {
		if err := r.SetQueryParam("grant", qGrant); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}