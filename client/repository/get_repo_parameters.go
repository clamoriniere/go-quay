// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRepoParams creates a new GetRepoParams object
// with the default values initialized.
func NewGetRepoParams() *GetRepoParams {
	var ()
	return &GetRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoParamsWithTimeout creates a new GetRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepoParamsWithTimeout(timeout time.Duration) *GetRepoParams {
	var ()
	return &GetRepoParams{

		timeout: timeout,
	}
}

// NewGetRepoParamsWithContext creates a new GetRepoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepoParamsWithContext(ctx context.Context) *GetRepoParams {
	var ()
	return &GetRepoParams{

		Context: ctx,
	}
}

// NewGetRepoParamsWithHTTPClient creates a new GetRepoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepoParamsWithHTTPClient(client *http.Client) *GetRepoParams {
	var ()
	return &GetRepoParams{
		HTTPClient: client,
	}
}

/*GetRepoParams contains all the parameters to send to the API endpoint
for the get repo operation typically these are written to a http.Request
*/
type GetRepoParams struct {

	/*IncludeStats
	  Whether to include action statistics

	*/
	IncludeStats *bool
	/*IncludeTags
	  Whether to include repository tags

	*/
	IncludeTags *bool
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repo params
func (o *GetRepoParams) WithTimeout(timeout time.Duration) *GetRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo params
func (o *GetRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo params
func (o *GetRepoParams) WithContext(ctx context.Context) *GetRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo params
func (o *GetRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo params
func (o *GetRepoParams) WithHTTPClient(client *http.Client) *GetRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo params
func (o *GetRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeStats adds the includeStats to the get repo params
func (o *GetRepoParams) WithIncludeStats(includeStats *bool) *GetRepoParams {
	o.SetIncludeStats(includeStats)
	return o
}

// SetIncludeStats adds the includeStats to the get repo params
func (o *GetRepoParams) SetIncludeStats(includeStats *bool) {
	o.IncludeStats = includeStats
}

// WithIncludeTags adds the includeTags to the get repo params
func (o *GetRepoParams) WithIncludeTags(includeTags *bool) *GetRepoParams {
	o.SetIncludeTags(includeTags)
	return o
}

// SetIncludeTags adds the includeTags to the get repo params
func (o *GetRepoParams) SetIncludeTags(includeTags *bool) {
	o.IncludeTags = includeTags
}

// WithRepository adds the repository to the get repo params
func (o *GetRepoParams) WithRepository(repository string) *GetRepoParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get repo params
func (o *GetRepoParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeStats != nil {

		// query param includeStats
		var qrIncludeStats bool
		if o.IncludeStats != nil {
			qrIncludeStats = *o.IncludeStats
		}
		qIncludeStats := swag.FormatBool(qrIncludeStats)
		if qIncludeStats != "" {
			if err := r.SetQueryParam("includeStats", qIncludeStats); err != nil {
				return err
			}
		}

	}

	if o.IncludeTags != nil {

		// query param includeTags
		var qrIncludeTags bool
		if o.IncludeTags != nil {
			qrIncludeTags = *o.IncludeTags
		}
		qIncludeTags := swag.FormatBool(qrIncludeTags)
		if qIncludeTags != "" {
			if err := r.SetQueryParam("includeTags", qIncludeTags); err != nil {
				return err
			}
		}

	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
