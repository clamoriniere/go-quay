// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewListTagImagesParams creates a new ListTagImagesParams object
// with the default values initialized.
func NewListTagImagesParams() *ListTagImagesParams {
	var ()
	return &ListTagImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTagImagesParamsWithTimeout creates a new ListTagImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTagImagesParamsWithTimeout(timeout time.Duration) *ListTagImagesParams {
	var ()
	return &ListTagImagesParams{

		timeout: timeout,
	}
}

// NewListTagImagesParamsWithContext creates a new ListTagImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTagImagesParamsWithContext(ctx context.Context) *ListTagImagesParams {
	var ()
	return &ListTagImagesParams{

		Context: ctx,
	}
}

// NewListTagImagesParamsWithHTTPClient creates a new ListTagImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTagImagesParamsWithHTTPClient(client *http.Client) *ListTagImagesParams {
	var ()
	return &ListTagImagesParams{
		HTTPClient: client,
	}
}

/*ListTagImagesParams contains all the parameters to send to the API endpoint
for the list tag images operation typically these are written to a http.Request
*/
type ListTagImagesParams struct {

	/*Owned
	  If specified, only images wholely owned by this tag are returned.

	*/
	Owned *bool
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tag images params
func (o *ListTagImagesParams) WithTimeout(timeout time.Duration) *ListTagImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tag images params
func (o *ListTagImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tag images params
func (o *ListTagImagesParams) WithContext(ctx context.Context) *ListTagImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tag images params
func (o *ListTagImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tag images params
func (o *ListTagImagesParams) WithHTTPClient(client *http.Client) *ListTagImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tag images params
func (o *ListTagImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwned adds the owned to the list tag images params
func (o *ListTagImagesParams) WithOwned(owned *bool) *ListTagImagesParams {
	o.SetOwned(owned)
	return o
}

// SetOwned adds the owned to the list tag images params
func (o *ListTagImagesParams) SetOwned(owned *bool) {
	o.Owned = owned
}

// WithRepository adds the repository to the list tag images params
func (o *ListTagImagesParams) WithRepository(repository string) *ListTagImagesParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the list tag images params
func (o *ListTagImagesParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTag adds the tag to the list tag images params
func (o *ListTagImagesParams) WithTag(tag string) *ListTagImagesParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the list tag images params
func (o *ListTagImagesParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ListTagImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Owned != nil {

		// query param owned
		var qrOwned bool
		if o.Owned != nil {
			qrOwned = *o.Owned
		}
		qOwned := swag.FormatBool(qrOwned)
		if qOwned != "" {
			if err := r.SetQueryParam("owned", qOwned); err != nil {
				return err
			}
		}

	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
