// Code generated by go-swagger; DO NOT EDIT.

package signing

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

// NewGetRepoSignaturesParams creates a new GetRepoSignaturesParams object
// with the default values initialized.
func NewGetRepoSignaturesParams() *GetRepoSignaturesParams {
	var ()
	return &GetRepoSignaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoSignaturesParamsWithTimeout creates a new GetRepoSignaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepoSignaturesParamsWithTimeout(timeout time.Duration) *GetRepoSignaturesParams {
	var ()
	return &GetRepoSignaturesParams{

		timeout: timeout,
	}
}

// NewGetRepoSignaturesParamsWithContext creates a new GetRepoSignaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepoSignaturesParamsWithContext(ctx context.Context) *GetRepoSignaturesParams {
	var ()
	return &GetRepoSignaturesParams{

		Context: ctx,
	}
}

// NewGetRepoSignaturesParamsWithHTTPClient creates a new GetRepoSignaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepoSignaturesParamsWithHTTPClient(client *http.Client) *GetRepoSignaturesParams {
	var ()
	return &GetRepoSignaturesParams{
		HTTPClient: client,
	}
}

/*GetRepoSignaturesParams contains all the parameters to send to the API endpoint
for the get repo signatures operation typically these are written to a http.Request
*/
type GetRepoSignaturesParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repo signatures params
func (o *GetRepoSignaturesParams) WithTimeout(timeout time.Duration) *GetRepoSignaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo signatures params
func (o *GetRepoSignaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo signatures params
func (o *GetRepoSignaturesParams) WithContext(ctx context.Context) *GetRepoSignaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo signatures params
func (o *GetRepoSignaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo signatures params
func (o *GetRepoSignaturesParams) WithHTTPClient(client *http.Client) *GetRepoSignaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo signatures params
func (o *GetRepoSignaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the get repo signatures params
func (o *GetRepoSignaturesParams) WithRepository(repository string) *GetRepoSignaturesParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get repo signatures params
func (o *GetRepoSignaturesParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoSignaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
