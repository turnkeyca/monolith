// Code generated by go-swagger; DO NOT EDIT.

package reference

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reference API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reference API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateReference(params *CreateReferenceParams, opts ...ClientOption) (*CreateReferenceNoContent, error)

	DeleteReference(params *DeleteReferenceParams, opts ...ClientOption) (*DeleteReferenceNoContent, error)

	GetReference(params *GetReferenceParams, opts ...ClientOption) (*GetReferenceOK, error)

	GetReferencesByUserID(params *GetReferencesByUserIDParams, opts ...ClientOption) (*GetReferencesByUserIDOK, error)

	UpdateReference(params *UpdateReferenceParams, opts ...ClientOption) (*UpdateReferenceNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateReference create a new reference
*/
func (a *Client) CreateReference(params *CreateReferenceParams, opts ...ClientOption) (*CreateReferenceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReferenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createReference",
		Method:             "POST",
		PathPattern:        "/v1/reference",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReferenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateReferenceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createReference: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteReference delete a reference
*/
func (a *Client) DeleteReference(params *DeleteReferenceParams, opts ...ClientOption) (*DeleteReferenceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReferenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteReference",
		Method:             "DELETE",
		PathPattern:        "/v1/reference/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReferenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteReferenceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteReference: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReference return a reference
*/
func (a *Client) GetReference(params *GetReferenceParams, opts ...ClientOption) (*GetReferenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReferenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReference",
		Method:             "GET",
		PathPattern:        "/v1/reference/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReferenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReferenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReference: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReferencesByUserID return all references for a user
*/
func (a *Client) GetReferencesByUserID(params *GetReferencesByUserIDParams, opts ...ClientOption) (*GetReferencesByUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReferencesByUserIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReferencesByUserId",
		Method:             "GET",
		PathPattern:        "/v1/reference",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReferencesByUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReferencesByUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReferencesByUserId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateReference update a reference
*/
func (a *Client) UpdateReference(params *UpdateReferenceParams, opts ...ClientOption) (*UpdateReferenceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReferenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateReference",
		Method:             "PUT",
		PathPattern:        "/v1/reference/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateReferenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateReferenceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateReference: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
