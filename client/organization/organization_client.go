// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeOrganizationDetails(params *ChangeOrganizationDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeOrganizationDetailsOK, error)

	CreateOrganization(params *CreateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationCreated, error)

	CreateOrganizationApplication(params *CreateOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationApplicationCreated, error)

	DeleteAdminedOrganization(params *DeleteAdminedOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAdminedOrganizationNoContent, error)

	DeleteOrganizationApplication(params *DeleteOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationApplicationNoContent, error)

	GetApplicationInformation(params *GetApplicationInformationParams) (*GetApplicationInformationOK, error)

	GetOrganization(params *GetOrganizationParams) (*GetOrganizationOK, error)

	GetOrganizationApplication(params *GetOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationApplicationOK, error)

	GetOrganizationApplications(params *GetOrganizationApplicationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationApplicationsOK, error)

	GetOrganizationCollaborators(params *GetOrganizationCollaboratorsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationCollaboratorsOK, error)

	GetOrganizationMember(params *GetOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMemberOK, error)

	GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMembersOK, error)

	RemoveOrganizationMember(params *RemoveOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveOrganizationMemberNoContent, error)

	UpdateOrganizationApplication(params *UpdateOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationApplicationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeOrganizationDetails Change the details for the specified organization.
*/
func (a *Client) ChangeOrganizationDetails(params *ChangeOrganizationDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeOrganizationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeOrganizationDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeOrganizationDetails",
		Method:             "PUT",
		PathPattern:        "/api/v1/organization/{orgname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeOrganizationDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeOrganizationDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeOrganizationDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOrganization Create a new organization.
*/
func (a *Client) CreateOrganization(params *CreateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrganization",
		Method:             "POST",
		PathPattern:        "/api/v1/organization/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrganizationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOrganizationApplication Creates a new application under this organization.
*/
func (a *Client) CreateOrganizationApplication(params *CreateOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationApplicationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrganizationApplication",
		Method:             "POST",
		PathPattern:        "/api/v1/organization/{orgname}/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrganizationApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrganizationApplicationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrganizationApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAdminedOrganization Deletes the specified organization.
*/
func (a *Client) DeleteAdminedOrganization(params *DeleteAdminedOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAdminedOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdminedOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAdminedOrganization",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organization/{orgname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdminedOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAdminedOrganizationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAdminedOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganizationApplication Deletes the application under this organization.
*/
func (a *Client) DeleteOrganizationApplication(params *DeleteOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationApplicationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganizationApplication",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organization/{orgname}/applications/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationApplicationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganizationApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationInformation Get information on the specified application.
*/
func (a *Client) GetApplicationInformation(params *GetApplicationInformationParams) (*GetApplicationInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationInformationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApplicationInformation",
		Method:             "GET",
		PathPattern:        "/api/v1/app/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApplicationInformationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationInformation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganization Get the details for the specified organization.
*/
func (a *Client) GetOrganization(params *GetOrganizationParams) (*GetOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganization",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationApplication Retrieves the application with the specified client_id under the specified organization.
*/
func (a *Client) GetOrganizationApplication(params *GetOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationApplication",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/applications/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationApplications List the applications for the specified organization.
*/
func (a *Client) GetOrganizationApplications(params *GetOrganizationApplicationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationApplicationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationApplications",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationApplicationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationApplicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationApplications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationCollaborators List outside collaborators of the specified organization.
*/
func (a *Client) GetOrganizationCollaborators(params *GetOrganizationCollaboratorsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationCollaboratorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationCollaboratorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationCollaborators",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/collaborators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationCollaboratorsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationCollaboratorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationCollaborators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationMember Retrieves the details of a member of the organization.
*/
func (a *Client) GetOrganizationMember(params *GetOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationMember",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/members/{membername}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationMembers List the human members of the specified organization.
*/
func (a *Client) GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationMembers",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveOrganizationMember Removes a member from an organization, revoking all its repository priviledges and removing
        it from all teams in the organization.
*/
func (a *Client) RemoveOrganizationMember(params *RemoveOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeOrganizationMember",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organization/{orgname}/members/{membername}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveOrganizationMemberNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeOrganizationMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrganizationApplication Updates an application under this organization.
*/
func (a *Client) UpdateOrganizationApplication(params *UpdateOrganizationApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationApplication",
		Method:             "PUT",
		PathPattern:        "/api/v1/organization/{orgname}/applications/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
