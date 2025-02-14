// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllocateIP allocates an ip in the given network for a project
*/
func (a *Client) AllocateIP(params *AllocateIPParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateIPCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateIP",
		Method:             "POST",
		PathPattern:        "/v1/ip/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateIPCreated), nil

}

/*
AllocateSpecificIP allocates an specific ip in the given network for a project
*/
func (a *Client) AllocateSpecificIP(params *AllocateSpecificIPParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateSpecificIPCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateSpecificIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateSpecificIP",
		Method:             "POST",
		PathPattern:        "/v1/ip/allocate/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateSpecificIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateSpecificIPCreated), nil

}

/*
DeleteIP deletes an ip and returns the deleted entity
*/
func (a *Client) DeleteIP(params *DeleteIPParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteIP",
		Method:             "DELETE",
		PathPattern:        "/v1/ip/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteIPOK), nil

}

/*
FindIP gets ip by id
*/
func (a *Client) FindIP(params *FindIPParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIP",
		Method:             "GET",
		PathPattern:        "/v1/ip/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindIPOK), nil

}

/*
FindIps gets all ips that match given properties
*/
func (a *Client) FindIps(params *FindIpsParams, authInfo runtime.ClientAuthInfoWriter) (*FindIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIpsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPs",
		Method:             "POST",
		PathPattern:        "/v1/ip/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIpsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindIpsOK), nil

}

/*
ListIps gets all ips
*/
func (a *Client) ListIps(params *ListIpsParams, authInfo runtime.ClientAuthInfoWriter) (*ListIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListIpsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listIPs",
		Method:             "GET",
		PathPattern:        "/v1/ip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListIpsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListIpsOK), nil

}

/*
UpdateIP updates an ip if the ip was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateIP(params *UpdateIPParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateIP",
		Method:             "POST",
		PathPattern:        "/v1/ip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateIPOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
