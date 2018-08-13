// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 organization id recipes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 organization id recipes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRecipeInOrganization creates recipe in organization

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) CreateRecipeInOrganization(params *CreateRecipeInOrganizationParams) (*CreateRecipeInOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecipeInOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRecipeInOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/recipes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecipeInOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRecipeInOrganizationOK), nil

}

/*
DeleteRecipeInOrganization deletes recipe by name in organization

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipeInOrganization(params *DeleteRecipeInOrganizationParams) (*DeleteRecipeInOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipeInOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRecipeInOrganization",
		Method:             "DELETE",
		PathPattern:        "/v3/{organizationId}/recipes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipeInOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRecipeInOrganizationOK), nil

}

/*
GetRecipeInOrganization gets recipe by name in organization

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipeInOrganization(params *GetRecipeInOrganizationParams) (*GetRecipeInOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipeInOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecipeInOrganization",
		Method:             "GET",
		PathPattern:        "/v3/{organizationId}/recipes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipeInOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipeInOrganizationOK), nil

}

/*
ListRecipesByOrganization lists recipes for the given organization

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) ListRecipesByOrganization(params *ListRecipesByOrganizationParams) (*ListRecipesByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRecipesByOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRecipesByOrganization",
		Method:             "GET",
		PathPattern:        "/v3/{organizationId}/recipes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRecipesByOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRecipesByOrganizationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
