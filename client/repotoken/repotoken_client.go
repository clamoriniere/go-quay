// Code generated by go-swagger; DO NOT EDIT.

package repotoken

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new repotoken API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repotoken API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeToken(params *ChangeTokenParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeTokenOK, error)

	CreateToken(params *CreateTokenParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTokenCreated, error)

	DeleteToken(params *DeleteTokenParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTokenNoContent, error)

	GetTokens(params *GetTokensParams, authInfo runtime.ClientAuthInfoWriter) (*GetTokensOK, error)

	ListRepoTokens(params *ListRepoTokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoTokensOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeToken Update the permissions for the specified repository token.
*/
func (a *Client) ChangeToken(params *ChangeTokenParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeToken",
		Method:             "PUT",
		PathPattern:        "/api/v1/repository/{repository}/tokens/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateToken Create a new repository token.
*/
func (a *Client) CreateToken(params *CreateTokenParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTokenCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createToken",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/tokens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTokenCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteToken Delete the repository token.
*/
func (a *Client) DeleteToken(params *DeleteTokenParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTokenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteToken",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}/tokens/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTokenNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTokens Fetch the specified repository token information.
*/
func (a *Client) GetTokens(params *GetTokensParams, authInfo runtime.ClientAuthInfoWriter) (*GetTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTokens",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/tokens/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRepoTokens List the tokens for the specified repository.
*/
func (a *Client) ListRepoTokens(params *ListRepoTokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoTokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepoTokens",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/tokens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepoTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRepoTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRepoTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
