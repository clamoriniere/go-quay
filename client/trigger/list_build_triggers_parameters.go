// Code generated by go-swagger; DO NOT EDIT.

package trigger

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
)

// NewListBuildTriggersParams creates a new ListBuildTriggersParams object
// with the default values initialized.
func NewListBuildTriggersParams() *ListBuildTriggersParams {
	var ()
	return &ListBuildTriggersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBuildTriggersParamsWithTimeout creates a new ListBuildTriggersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBuildTriggersParamsWithTimeout(timeout time.Duration) *ListBuildTriggersParams {
	var ()
	return &ListBuildTriggersParams{

		timeout: timeout,
	}
}

// NewListBuildTriggersParamsWithContext creates a new ListBuildTriggersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBuildTriggersParamsWithContext(ctx context.Context) *ListBuildTriggersParams {
	var ()
	return &ListBuildTriggersParams{

		Context: ctx,
	}
}

// NewListBuildTriggersParamsWithHTTPClient creates a new ListBuildTriggersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBuildTriggersParamsWithHTTPClient(client *http.Client) *ListBuildTriggersParams {
	var ()
	return &ListBuildTriggersParams{
		HTTPClient: client,
	}
}

/*ListBuildTriggersParams contains all the parameters to send to the API endpoint
for the list build triggers operation typically these are written to a http.Request
*/
type ListBuildTriggersParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list build triggers params
func (o *ListBuildTriggersParams) WithTimeout(timeout time.Duration) *ListBuildTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list build triggers params
func (o *ListBuildTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list build triggers params
func (o *ListBuildTriggersParams) WithContext(ctx context.Context) *ListBuildTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list build triggers params
func (o *ListBuildTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list build triggers params
func (o *ListBuildTriggersParams) WithHTTPClient(client *http.Client) *ListBuildTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list build triggers params
func (o *ListBuildTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the list build triggers params
func (o *ListBuildTriggersParams) WithRepository(repository string) *ListBuildTriggersParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the list build triggers params
func (o *ListBuildTriggersParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *ListBuildTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
