// Code generated by go-swagger; DO NOT EDIT.

package subsystems

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

// NewGetSubsystemsIDParams creates a new GetSubsystemsIDParams object
// with the default values initialized.
func NewGetSubsystemsIDParams() *GetSubsystemsIDParams {
	var ()
	return &GetSubsystemsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubsystemsIDParamsWithTimeout creates a new GetSubsystemsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubsystemsIDParamsWithTimeout(timeout time.Duration) *GetSubsystemsIDParams {
	var ()
	return &GetSubsystemsIDParams{

		timeout: timeout,
	}
}

// NewGetSubsystemsIDParamsWithContext creates a new GetSubsystemsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubsystemsIDParamsWithContext(ctx context.Context) *GetSubsystemsIDParams {
	var ()
	return &GetSubsystemsIDParams{

		Context: ctx,
	}
}

// NewGetSubsystemsIDParamsWithHTTPClient creates a new GetSubsystemsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubsystemsIDParamsWithHTTPClient(client *http.Client) *GetSubsystemsIDParams {
	var ()
	return &GetSubsystemsIDParams{
		HTTPClient: client,
	}
}

/*GetSubsystemsIDParams contains all the parameters to send to the API endpoint
for the get subsystems ID operation typically these are written to a http.Request
*/
type GetSubsystemsIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subsystems ID params
func (o *GetSubsystemsIDParams) WithTimeout(timeout time.Duration) *GetSubsystemsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subsystems ID params
func (o *GetSubsystemsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subsystems ID params
func (o *GetSubsystemsIDParams) WithContext(ctx context.Context) *GetSubsystemsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subsystems ID params
func (o *GetSubsystemsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subsystems ID params
func (o *GetSubsystemsIDParams) WithHTTPClient(client *http.Client) *GetSubsystemsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subsystems ID params
func (o *GetSubsystemsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get subsystems ID params
func (o *GetSubsystemsIDParams) WithID(id int64) *GetSubsystemsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get subsystems ID params
func (o *GetSubsystemsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubsystemsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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