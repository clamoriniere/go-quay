// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewGetOrganizationCollaboratorsParams creates a new GetOrganizationCollaboratorsParams object
// with the default values initialized.
func NewGetOrganizationCollaboratorsParams() *GetOrganizationCollaboratorsParams {
	var ()
	return &GetOrganizationCollaboratorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationCollaboratorsParamsWithTimeout creates a new GetOrganizationCollaboratorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationCollaboratorsParamsWithTimeout(timeout time.Duration) *GetOrganizationCollaboratorsParams {
	var ()
	return &GetOrganizationCollaboratorsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationCollaboratorsParamsWithContext creates a new GetOrganizationCollaboratorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationCollaboratorsParamsWithContext(ctx context.Context) *GetOrganizationCollaboratorsParams {
	var ()
	return &GetOrganizationCollaboratorsParams{

		Context: ctx,
	}
}

// NewGetOrganizationCollaboratorsParamsWithHTTPClient creates a new GetOrganizationCollaboratorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationCollaboratorsParamsWithHTTPClient(client *http.Client) *GetOrganizationCollaboratorsParams {
	var ()
	return &GetOrganizationCollaboratorsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationCollaboratorsParams contains all the parameters to send to the API endpoint
for the get organization collaborators operation typically these are written to a http.Request
*/
type GetOrganizationCollaboratorsParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) WithTimeout(timeout time.Duration) *GetOrganizationCollaboratorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) WithContext(ctx context.Context) *GetOrganizationCollaboratorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) WithHTTPClient(client *http.Client) *GetOrganizationCollaboratorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgname adds the orgname to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) WithOrgname(orgname string) *GetOrganizationCollaboratorsParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the get organization collaborators params
func (o *GetOrganizationCollaboratorsParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationCollaboratorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
