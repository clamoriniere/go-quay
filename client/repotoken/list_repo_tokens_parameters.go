// Code generated by go-swagger; DO NOT EDIT.

package repotoken

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

// NewListRepoTokensParams creates a new ListRepoTokensParams object
// with the default values initialized.
func NewListRepoTokensParams() *ListRepoTokensParams {
	var ()
	return &ListRepoTokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRepoTokensParamsWithTimeout creates a new ListRepoTokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRepoTokensParamsWithTimeout(timeout time.Duration) *ListRepoTokensParams {
	var ()
	return &ListRepoTokensParams{

		timeout: timeout,
	}
}

// NewListRepoTokensParamsWithContext creates a new ListRepoTokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRepoTokensParamsWithContext(ctx context.Context) *ListRepoTokensParams {
	var ()
	return &ListRepoTokensParams{

		Context: ctx,
	}
}

// NewListRepoTokensParamsWithHTTPClient creates a new ListRepoTokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRepoTokensParamsWithHTTPClient(client *http.Client) *ListRepoTokensParams {
	var ()
	return &ListRepoTokensParams{
		HTTPClient: client,
	}
}

/*ListRepoTokensParams contains all the parameters to send to the API endpoint
for the list repo tokens operation typically these are written to a http.Request
*/
type ListRepoTokensParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list repo tokens params
func (o *ListRepoTokensParams) WithTimeout(timeout time.Duration) *ListRepoTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list repo tokens params
func (o *ListRepoTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list repo tokens params
func (o *ListRepoTokensParams) WithContext(ctx context.Context) *ListRepoTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list repo tokens params
func (o *ListRepoTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list repo tokens params
func (o *ListRepoTokensParams) WithHTTPClient(client *http.Client) *ListRepoTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list repo tokens params
func (o *ListRepoTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the list repo tokens params
func (o *ListRepoTokensParams) WithRepository(repository string) *ListRepoTokensParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the list repo tokens params
func (o *ListRepoTokensParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *ListRepoTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
