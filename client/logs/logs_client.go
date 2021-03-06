// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListOrgLogs(params *ListOrgLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrgLogsOK, error)

	ListRepoLogs(params *ListRepoLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoLogsOK, error)

	ListUserLogs(params *ListUserLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserLogsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListOrgLogs List the logs for the specified organization.
*/
func (a *Client) ListOrgLogs(params *ListOrgLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrgLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrgLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listOrgLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrgLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrgLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOrgLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRepoLogs List the logs for the specified repository.
*/
func (a *Client) ListRepoLogs(params *ListRepoLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepoLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepoLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRepoLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRepoLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserLogs List the logs for the current user.
*/
func (a *Client) ListUserLogs(params *ListUserLogsParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/user/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
