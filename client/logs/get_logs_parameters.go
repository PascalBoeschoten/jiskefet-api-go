// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewGetLogsParams creates a new GetLogsParams object
// with the default values initialized.
func NewGetLogsParams() *GetLogsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetLogsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsParamsWithTimeout creates a new GetLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogsParamsWithTimeout(timeout time.Duration) *GetLogsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetLogsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetLogsParamsWithContext creates a new GetLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogsParamsWithContext(ctx context.Context) *GetLogsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetLogsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetLogsParamsWithHTTPClient creates a new GetLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogsParamsWithHTTPClient(client *http.Client) *GetLogsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetLogsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetLogsParams contains all the parameters to send to the API endpoint
for the get logs operation typically these are written to a http.Request
*/
type GetLogsParams struct {

	/*EndCreationTime
	  The upper bound of the creation time filter range.

	*/
	EndCreationTime *string
	/*LogID
	  The id of the log.

	*/
	LogID *string
	/*OrderBy
	  On which field to order on.

	*/
	OrderBy *string
	/*OrderDirection
	  The order direction, either ascending or descending.

	*/
	OrderDirection *string
	/*Origin
	  The origin/creator of the log.

	*/
	Origin *string
	/*PageNumber
	  The current page, i.e. the offset in the result set based on pageSize.

	*/
	PageNumber *string
	/*PageSize
	  The maximum amount of logs in each result.

	*/
	PageSize *string
	/*Searchterm
	  Search for a term that exists in the title field of a log.

	*/
	Searchterm *string
	/*StartCreationTime
	  The lower bound of the creation time filter range.

	*/
	StartCreationTime *string
	/*Subtype
	  The subtype of the log.

	*/
	Subtype *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logs params
func (o *GetLogsParams) WithTimeout(timeout time.Duration) *GetLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs params
func (o *GetLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs params
func (o *GetLogsParams) WithContext(ctx context.Context) *GetLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs params
func (o *GetLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) WithHTTPClient(client *http.Client) *GetLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndCreationTime adds the endCreationTime to the get logs params
func (o *GetLogsParams) WithEndCreationTime(endCreationTime *string) *GetLogsParams {
	o.SetEndCreationTime(endCreationTime)
	return o
}

// SetEndCreationTime adds the endCreationTime to the get logs params
func (o *GetLogsParams) SetEndCreationTime(endCreationTime *string) {
	o.EndCreationTime = endCreationTime
}

// WithLogID adds the logID to the get logs params
func (o *GetLogsParams) WithLogID(logID *string) *GetLogsParams {
	o.SetLogID(logID)
	return o
}

// SetLogID adds the logId to the get logs params
func (o *GetLogsParams) SetLogID(logID *string) {
	o.LogID = logID
}

// WithOrderBy adds the orderBy to the get logs params
func (o *GetLogsParams) WithOrderBy(orderBy *string) *GetLogsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get logs params
func (o *GetLogsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the get logs params
func (o *GetLogsParams) WithOrderDirection(orderDirection *string) *GetLogsParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the get logs params
func (o *GetLogsParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithOrigin adds the origin to the get logs params
func (o *GetLogsParams) WithOrigin(origin *string) *GetLogsParams {
	o.SetOrigin(origin)
	return o
}

// SetOrigin adds the origin to the get logs params
func (o *GetLogsParams) SetOrigin(origin *string) {
	o.Origin = origin
}

// WithPageNumber adds the pageNumber to the get logs params
func (o *GetLogsParams) WithPageNumber(pageNumber *string) *GetLogsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get logs params
func (o *GetLogsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get logs params
func (o *GetLogsParams) WithPageSize(pageSize *string) *GetLogsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get logs params
func (o *GetLogsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithSearchterm adds the searchterm to the get logs params
func (o *GetLogsParams) WithSearchterm(searchterm *string) *GetLogsParams {
	o.SetSearchterm(searchterm)
	return o
}

// SetSearchterm adds the searchterm to the get logs params
func (o *GetLogsParams) SetSearchterm(searchterm *string) {
	o.Searchterm = searchterm
}

// WithStartCreationTime adds the startCreationTime to the get logs params
func (o *GetLogsParams) WithStartCreationTime(startCreationTime *string) *GetLogsParams {
	o.SetStartCreationTime(startCreationTime)
	return o
}

// SetStartCreationTime adds the startCreationTime to the get logs params
func (o *GetLogsParams) SetStartCreationTime(startCreationTime *string) {
	o.StartCreationTime = startCreationTime
}

// WithSubtype adds the subtype to the get logs params
func (o *GetLogsParams) WithSubtype(subtype *string) *GetLogsParams {
	o.SetSubtype(subtype)
	return o
}

// SetSubtype adds the subtype to the get logs params
func (o *GetLogsParams) SetSubtype(subtype *string) {
	o.Subtype = subtype
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndCreationTime != nil {

		// query param endCreationTime
		var qrEndCreationTime string
		if o.EndCreationTime != nil {
			qrEndCreationTime = *o.EndCreationTime
		}
		qEndCreationTime := qrEndCreationTime
		if qEndCreationTime != "" {
			if err := r.SetQueryParam("endCreationTime", qEndCreationTime); err != nil {
				return err
			}
		}

	}

	if o.LogID != nil {

		// query param logId
		var qrLogID string
		if o.LogID != nil {
			qrLogID = *o.LogID
		}
		qLogID := qrLogID
		if qLogID != "" {
			if err := r.SetQueryParam("logId", qLogID); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.OrderDirection != nil {

		// query param orderDirection
		var qrOrderDirection string
		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {
			if err := r.SetQueryParam("orderDirection", qOrderDirection); err != nil {
				return err
			}
		}

	}

	if o.Origin != nil {

		// query param origin
		var qrOrigin string
		if o.Origin != nil {
			qrOrigin = *o.Origin
		}
		qOrigin := qrOrigin
		if qOrigin != "" {
			if err := r.SetQueryParam("origin", qOrigin); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Searchterm != nil {

		// query param searchterm
		var qrSearchterm string
		if o.Searchterm != nil {
			qrSearchterm = *o.Searchterm
		}
		qSearchterm := qrSearchterm
		if qSearchterm != "" {
			if err := r.SetQueryParam("searchterm", qSearchterm); err != nil {
				return err
			}
		}

	}

	if o.StartCreationTime != nil {

		// query param startCreationTime
		var qrStartCreationTime string
		if o.StartCreationTime != nil {
			qrStartCreationTime = *o.StartCreationTime
		}
		qStartCreationTime := qrStartCreationTime
		if qStartCreationTime != "" {
			if err := r.SetQueryParam("startCreationTime", qStartCreationTime); err != nil {
				return err
			}
		}

	}

	if o.Subtype != nil {

		// query param subtype
		var qrSubtype string
		if o.Subtype != nil {
			qrSubtype = *o.Subtype
		}
		qSubtype := qrSubtype
		if qSubtype != "" {
			if err := r.SetQueryParam("subtype", qSubtype); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}