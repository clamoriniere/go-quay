// Code generated by go-swagger; DO NOT EDIT.

package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new repositorynotification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repositorynotification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRepoNotification(params *CreateRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRepoNotificationCreated, error)

	DeleteRepoNotification(params *DeleteRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRepoNotificationNoContent, error)

	GetRepoNotification(params *GetRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoNotificationOK, error)

	ListRepoNotifications(params *ListRepoNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoNotificationsOK, error)

	ResetRepositoryNotificationFailures(params *ResetRepositoryNotificationFailuresParams, authInfo runtime.ClientAuthInfoWriter) (*ResetRepositoryNotificationFailuresCreated, error)

	TestRepoNotification(params *TestRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*TestRepoNotificationCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRepoNotification create repo notification API
*/
func (a *Client) CreateRepoNotification(params *CreateRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRepoNotificationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRepoNotification",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/notification/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRepoNotificationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRepoNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRepoNotification Deletes the specified notification.
*/
func (a *Client) DeleteRepoNotification(params *DeleteRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRepoNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRepoNotification",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRepoNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRepoNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRepoNotification Get information for the specified notification.
*/
func (a *Client) GetRepoNotification(params *GetRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoNotification",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRepoNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRepoNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRepoNotifications List the notifications for the specified repository.
*/
func (a *Client) ListRepoNotifications(params *ListRepoNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepoNotifications",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/notification/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepoNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRepoNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRepoNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetRepositoryNotificationFailures Resets repository notification to 0 failures.
*/
func (a *Client) ResetRepositoryNotificationFailures(params *ResetRepositoryNotificationFailuresParams, authInfo runtime.ClientAuthInfoWriter) (*ResetRepositoryNotificationFailuresCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetRepositoryNotificationFailuresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resetRepositoryNotificationFailures",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetRepositoryNotificationFailuresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetRepositoryNotificationFailuresCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetRepositoryNotificationFailures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TestRepoNotification Queues a test notification for this repository.
*/
func (a *Client) TestRepoNotification(params *TestRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*TestRepoNotificationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testRepoNotification",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestRepoNotificationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testRepoNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
