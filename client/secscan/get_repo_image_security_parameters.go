// Code generated by go-swagger; DO NOT EDIT.

package secscan

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

// NewGetRepoImageSecurityParams creates a new GetRepoImageSecurityParams object
// with the default values initialized.
func NewGetRepoImageSecurityParams() *GetRepoImageSecurityParams {
	var ()
	return &GetRepoImageSecurityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoImageSecurityParamsWithTimeout creates a new GetRepoImageSecurityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepoImageSecurityParamsWithTimeout(timeout time.Duration) *GetRepoImageSecurityParams {
	var ()
	return &GetRepoImageSecurityParams{

		timeout: timeout,
	}
}

// NewGetRepoImageSecurityParamsWithContext creates a new GetRepoImageSecurityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepoImageSecurityParamsWithContext(ctx context.Context) *GetRepoImageSecurityParams {
	var ()
	return &GetRepoImageSecurityParams{

		Context: ctx,
	}
}

// NewGetRepoImageSecurityParamsWithHTTPClient creates a new GetRepoImageSecurityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepoImageSecurityParamsWithHTTPClient(client *http.Client) *GetRepoImageSecurityParams {
	var ()
	return &GetRepoImageSecurityParams{
		HTTPClient: client,
	}
}

/*GetRepoImageSecurityParams contains all the parameters to send to the API endpoint
for the get repo image security operation typically these are written to a http.Request
*/
type GetRepoImageSecurityParams struct {

	/*Imageid
	  The image ID

	*/
	Imageid string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Vulnerabilities
	  Include vulnerabilities information

	*/
	Vulnerabilities *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repo image security params
func (o *GetRepoImageSecurityParams) WithTimeout(timeout time.Duration) *GetRepoImageSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo image security params
func (o *GetRepoImageSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo image security params
func (o *GetRepoImageSecurityParams) WithContext(ctx context.Context) *GetRepoImageSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo image security params
func (o *GetRepoImageSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo image security params
func (o *GetRepoImageSecurityParams) WithHTTPClient(client *http.Client) *GetRepoImageSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo image security params
func (o *GetRepoImageSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageid adds the imageid to the get repo image security params
func (o *GetRepoImageSecurityParams) WithImageid(imageid string) *GetRepoImageSecurityParams {
	o.SetImageid(imageid)
	return o
}

// SetImageid adds the imageid to the get repo image security params
func (o *GetRepoImageSecurityParams) SetImageid(imageid string) {
	o.Imageid = imageid
}

// WithRepository adds the repository to the get repo image security params
func (o *GetRepoImageSecurityParams) WithRepository(repository string) *GetRepoImageSecurityParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get repo image security params
func (o *GetRepoImageSecurityParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithVulnerabilities adds the vulnerabilities to the get repo image security params
func (o *GetRepoImageSecurityParams) WithVulnerabilities(vulnerabilities *bool) *GetRepoImageSecurityParams {
	o.SetVulnerabilities(vulnerabilities)
	return o
}

// SetVulnerabilities adds the vulnerabilities to the get repo image security params
func (o *GetRepoImageSecurityParams) SetVulnerabilities(vulnerabilities *bool) {
	o.Vulnerabilities = vulnerabilities
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoImageSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageid
	if err := r.SetPathParam("imageid", o.Imageid); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if o.Vulnerabilities != nil {

		// query param vulnerabilities
		var qrVulnerabilities bool
		if o.Vulnerabilities != nil {
			qrVulnerabilities = *o.Vulnerabilities
		}
		qVulnerabilities := swag.FormatBool(qrVulnerabilities)
		if qVulnerabilities != "" {
			if err := r.SetQueryParam("vulnerabilities", qVulnerabilities); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
