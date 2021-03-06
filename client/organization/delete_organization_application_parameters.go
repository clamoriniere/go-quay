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

// NewDeleteOrganizationApplicationParams creates a new DeleteOrganizationApplicationParams object
// with the default values initialized.
func NewDeleteOrganizationApplicationParams() *DeleteOrganizationApplicationParams {
	var ()
	return &DeleteOrganizationApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationApplicationParamsWithTimeout creates a new DeleteOrganizationApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrganizationApplicationParamsWithTimeout(timeout time.Duration) *DeleteOrganizationApplicationParams {
	var ()
	return &DeleteOrganizationApplicationParams{

		timeout: timeout,
	}
}

// NewDeleteOrganizationApplicationParamsWithContext creates a new DeleteOrganizationApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrganizationApplicationParamsWithContext(ctx context.Context) *DeleteOrganizationApplicationParams {
	var ()
	return &DeleteOrganizationApplicationParams{

		Context: ctx,
	}
}

// NewDeleteOrganizationApplicationParamsWithHTTPClient creates a new DeleteOrganizationApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrganizationApplicationParamsWithHTTPClient(client *http.Client) *DeleteOrganizationApplicationParams {
	var ()
	return &DeleteOrganizationApplicationParams{
		HTTPClient: client,
	}
}

/*DeleteOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the delete organization application operation typically these are written to a http.Request
*/
type DeleteOrganizationApplicationParams struct {

	/*ClientID
	  The OAuth client ID

	*/
	ClientID string
	/*Orgname
	  The name of the organization

	*/
	Orgname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithTimeout(timeout time.Duration) *DeleteOrganizationApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization application params
func (o *DeleteOrganizationApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithContext(ctx context.Context) *DeleteOrganizationApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization application params
func (o *DeleteOrganizationApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithHTTPClient(client *http.Client) *DeleteOrganizationApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization application params
func (o *DeleteOrganizationApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithClientID(clientID string) *DeleteOrganizationApplicationParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete organization application params
func (o *DeleteOrganizationApplicationParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithOrgname adds the orgname to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithOrgname(orgname string) *DeleteOrganizationApplicationParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the delete organization application params
func (o *DeleteOrganizationApplicationParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
