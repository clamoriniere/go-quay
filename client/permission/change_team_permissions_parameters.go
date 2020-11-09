// Code generated by go-swagger; DO NOT EDIT.

package permission

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

// NewChangeTeamPermissionsParams creates a new ChangeTeamPermissionsParams object
// with the default values initialized.
func NewChangeTeamPermissionsParams() *ChangeTeamPermissionsParams {
	var ()
	return &ChangeTeamPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTeamPermissionsParamsWithTimeout creates a new ChangeTeamPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTeamPermissionsParamsWithTimeout(timeout time.Duration) *ChangeTeamPermissionsParams {
	var ()
	return &ChangeTeamPermissionsParams{

		timeout: timeout,
	}
}

// NewChangeTeamPermissionsParamsWithContext creates a new ChangeTeamPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTeamPermissionsParamsWithContext(ctx context.Context) *ChangeTeamPermissionsParams {
	var ()
	return &ChangeTeamPermissionsParams{

		Context: ctx,
	}
}

// NewChangeTeamPermissionsParamsWithHTTPClient creates a new ChangeTeamPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTeamPermissionsParamsWithHTTPClient(client *http.Client) *ChangeTeamPermissionsParams {
	var ()
	return &ChangeTeamPermissionsParams{
		HTTPClient: client,
	}
}

/*ChangeTeamPermissionsParams contains all the parameters to send to the API endpoint
for the change team permissions operation typically these are written to a http.Request
*/
type ChangeTeamPermissionsParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.TeamPermission
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Teamname
	  The name of the team to which the permission applies

	*/
	Teamname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithTimeout(timeout time.Duration) *ChangeTeamPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithContext(ctx context.Context) *ChangeTeamPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithHTTPClient(client *http.Client) *ChangeTeamPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithBody(body *models.TeamPermission) *ChangeTeamPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetBody(body *models.TeamPermission) {
	o.Body = body
}

// WithRepository adds the repository to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithRepository(repository string) *ChangeTeamPermissionsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTeamname adds the teamname to the change team permissions params
func (o *ChangeTeamPermissionsParams) WithTeamname(teamname string) *ChangeTeamPermissionsParams {
	o.SetTeamname(teamname)
	return o
}

// SetTeamname adds the teamname to the change team permissions params
func (o *ChangeTeamPermissionsParams) SetTeamname(teamname string) {
	o.Teamname = teamname
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTeamPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param teamname
	if err := r.SetPathParam("teamname", o.Teamname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
