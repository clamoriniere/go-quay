// Code generated by go-swagger; DO NOT EDIT.

package robot

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

// NewDeleteUserRobotParams creates a new DeleteUserRobotParams object
// with the default values initialized.
func NewDeleteUserRobotParams() *DeleteUserRobotParams {
	var ()
	return &DeleteUserRobotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserRobotParamsWithTimeout creates a new DeleteUserRobotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserRobotParamsWithTimeout(timeout time.Duration) *DeleteUserRobotParams {
	var ()
	return &DeleteUserRobotParams{

		timeout: timeout,
	}
}

// NewDeleteUserRobotParamsWithContext creates a new DeleteUserRobotParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserRobotParamsWithContext(ctx context.Context) *DeleteUserRobotParams {
	var ()
	return &DeleteUserRobotParams{

		Context: ctx,
	}
}

// NewDeleteUserRobotParamsWithHTTPClient creates a new DeleteUserRobotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserRobotParamsWithHTTPClient(client *http.Client) *DeleteUserRobotParams {
	var ()
	return &DeleteUserRobotParams{
		HTTPClient: client,
	}
}

/*DeleteUserRobotParams contains all the parameters to send to the API endpoint
for the delete user robot operation typically these are written to a http.Request
*/
type DeleteUserRobotParams struct {

	/*RobotShortname
	  The short name for the robot, without any user or organization prefix

	*/
	RobotShortname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user robot params
func (o *DeleteUserRobotParams) WithTimeout(timeout time.Duration) *DeleteUserRobotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user robot params
func (o *DeleteUserRobotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user robot params
func (o *DeleteUserRobotParams) WithContext(ctx context.Context) *DeleteUserRobotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user robot params
func (o *DeleteUserRobotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user robot params
func (o *DeleteUserRobotParams) WithHTTPClient(client *http.Client) *DeleteUserRobotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user robot params
func (o *DeleteUserRobotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRobotShortname adds the robotShortname to the delete user robot params
func (o *DeleteUserRobotParams) WithRobotShortname(robotShortname string) *DeleteUserRobotParams {
	o.SetRobotShortname(robotShortname)
	return o
}

// SetRobotShortname adds the robotShortname to the delete user robot params
func (o *DeleteUserRobotParams) SetRobotShortname(robotShortname string) {
	o.RobotShortname = robotShortname
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserRobotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
