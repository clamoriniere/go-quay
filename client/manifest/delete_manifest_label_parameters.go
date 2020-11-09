// Code generated by go-swagger; DO NOT EDIT.

package manifest

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

// NewDeleteManifestLabelParams creates a new DeleteManifestLabelParams object
// with the default values initialized.
func NewDeleteManifestLabelParams() *DeleteManifestLabelParams {
	var ()
	return &DeleteManifestLabelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteManifestLabelParamsWithTimeout creates a new DeleteManifestLabelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteManifestLabelParamsWithTimeout(timeout time.Duration) *DeleteManifestLabelParams {
	var ()
	return &DeleteManifestLabelParams{

		timeout: timeout,
	}
}

// NewDeleteManifestLabelParamsWithContext creates a new DeleteManifestLabelParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteManifestLabelParamsWithContext(ctx context.Context) *DeleteManifestLabelParams {
	var ()
	return &DeleteManifestLabelParams{

		Context: ctx,
	}
}

// NewDeleteManifestLabelParamsWithHTTPClient creates a new DeleteManifestLabelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteManifestLabelParamsWithHTTPClient(client *http.Client) *DeleteManifestLabelParams {
	var ()
	return &DeleteManifestLabelParams{
		HTTPClient: client,
	}
}

/*DeleteManifestLabelParams contains all the parameters to send to the API endpoint
for the delete manifest label operation typically these are written to a http.Request
*/
type DeleteManifestLabelParams struct {

	/*Labelid
	  The ID of the label

	*/
	Labelid string
	/*Manifestref
	  The digest of the manifest

	*/
	Manifestref string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete manifest label params
func (o *DeleteManifestLabelParams) WithTimeout(timeout time.Duration) *DeleteManifestLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete manifest label params
func (o *DeleteManifestLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete manifest label params
func (o *DeleteManifestLabelParams) WithContext(ctx context.Context) *DeleteManifestLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete manifest label params
func (o *DeleteManifestLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete manifest label params
func (o *DeleteManifestLabelParams) WithHTTPClient(client *http.Client) *DeleteManifestLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete manifest label params
func (o *DeleteManifestLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabelid adds the labelid to the delete manifest label params
func (o *DeleteManifestLabelParams) WithLabelid(labelid string) *DeleteManifestLabelParams {
	o.SetLabelid(labelid)
	return o
}

// SetLabelid adds the labelid to the delete manifest label params
func (o *DeleteManifestLabelParams) SetLabelid(labelid string) {
	o.Labelid = labelid
}

// WithManifestref adds the manifestref to the delete manifest label params
func (o *DeleteManifestLabelParams) WithManifestref(manifestref string) *DeleteManifestLabelParams {
	o.SetManifestref(manifestref)
	return o
}

// SetManifestref adds the manifestref to the delete manifest label params
func (o *DeleteManifestLabelParams) SetManifestref(manifestref string) {
	o.Manifestref = manifestref
}

// WithRepository adds the repository to the delete manifest label params
func (o *DeleteManifestLabelParams) WithRepository(repository string) *DeleteManifestLabelParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the delete manifest label params
func (o *DeleteManifestLabelParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteManifestLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param labelid
	if err := r.SetPathParam("labelid", o.Labelid); err != nil {
		return err
	}

	// path param manifestref
	if err := r.SetPathParam("manifestref", o.Manifestref); err != nil {
		return err
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