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

	"github.com/clamoriniere/go-quay/models"
)

// NewUpdateOrganizationApplicationParams creates a new UpdateOrganizationApplicationParams object
// with the default values initialized.
func NewUpdateOrganizationApplicationParams() *UpdateOrganizationApplicationParams {
	var ()
	return &UpdateOrganizationApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationApplicationParamsWithTimeout creates a new UpdateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrganizationApplicationParamsWithTimeout(timeout time.Duration) *UpdateOrganizationApplicationParams {
	var ()
	return &UpdateOrganizationApplicationParams{

		timeout: timeout,
	}
}

// NewUpdateOrganizationApplicationParamsWithContext creates a new UpdateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrganizationApplicationParamsWithContext(ctx context.Context) *UpdateOrganizationApplicationParams {
	var ()
	return &UpdateOrganizationApplicationParams{

		Context: ctx,
	}
}

// NewUpdateOrganizationApplicationParamsWithHTTPClient creates a new UpdateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrganizationApplicationParamsWithHTTPClient(client *http.Client) *UpdateOrganizationApplicationParams {
	var ()
	return &UpdateOrganizationApplicationParams{
		HTTPClient: client,
	}
}

/*UpdateOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the update organization application operation typically these are written to a http.Request
*/
type UpdateOrganizationApplicationParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.UpdateApp
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

// WithTimeout adds the timeout to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithTimeout(timeout time.Duration) *UpdateOrganizationApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithContext(ctx context.Context) *UpdateOrganizationApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithHTTPClient(client *http.Client) *UpdateOrganizationApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithBody(body *models.UpdateApp) *UpdateOrganizationApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetBody(body *models.UpdateApp) {
	o.Body = body
}

// WithClientID adds the clientID to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithClientID(clientID string) *UpdateOrganizationApplicationParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithOrgname adds the orgname to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithOrgname(orgname string) *UpdateOrganizationApplicationParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the update organization application params
func (o *UpdateOrganizationApplicationParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
