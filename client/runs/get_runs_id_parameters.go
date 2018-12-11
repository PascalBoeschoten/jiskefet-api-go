// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRunsIDParams creates a new GetRunsIDParams object
// with the default values initialized.
func NewGetRunsIDParams() *GetRunsIDParams {
	var ()
	return &GetRunsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunsIDParamsWithTimeout creates a new GetRunsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunsIDParamsWithTimeout(timeout time.Duration) *GetRunsIDParams {
	var ()
	return &GetRunsIDParams{

		timeout: timeout,
	}
}

// NewGetRunsIDParamsWithContext creates a new GetRunsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunsIDParamsWithContext(ctx context.Context) *GetRunsIDParams {
	var ()
	return &GetRunsIDParams{

		Context: ctx,
	}
}

// NewGetRunsIDParamsWithHTTPClient creates a new GetRunsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunsIDParamsWithHTTPClient(client *http.Client) *GetRunsIDParams {
	var ()
	return &GetRunsIDParams{
		HTTPClient: client,
	}
}

/*GetRunsIDParams contains all the parameters to send to the API endpoint
for the get runs ID operation typically these are written to a http.Request
*/
type GetRunsIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runs ID params
func (o *GetRunsIDParams) WithTimeout(timeout time.Duration) *GetRunsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runs ID params
func (o *GetRunsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runs ID params
func (o *GetRunsIDParams) WithContext(ctx context.Context) *GetRunsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runs ID params
func (o *GetRunsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runs ID params
func (o *GetRunsIDParams) WithHTTPClient(client *http.Client) *GetRunsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runs ID params
func (o *GetRunsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get runs ID params
func (o *GetRunsIDParams) WithID(id int64) *GetRunsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runs ID params
func (o *GetRunsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}