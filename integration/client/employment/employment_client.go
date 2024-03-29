// Code generated by go-swagger; DO NOT EDIT.

package employment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new employment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for employment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEmployment(params *CreateEmploymentParams, opts ...ClientOption) (*CreateEmploymentNoContent, error)

	DeleteEmployment(params *DeleteEmploymentParams, opts ...ClientOption) (*DeleteEmploymentNoContent, error)

	GetEmployment(params *GetEmploymentParams, opts ...ClientOption) (*GetEmploymentOK, error)

	GetEmploymentsByUserID(params *GetEmploymentsByUserIDParams, opts ...ClientOption) (*GetEmploymentsByUserIDOK, error)

	UpdateEmployment(params *UpdateEmploymentParams, opts ...ClientOption) (*UpdateEmploymentNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEmployment create a new employment
*/
func (a *Client) CreateEmployment(params *CreateEmploymentParams, opts ...ClientOption) (*CreateEmploymentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEmploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createEmployment",
		Method:             "POST",
		PathPattern:        "/v1/employment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEmploymentReader{formats: a.formats},
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
	success, ok := result.(*CreateEmploymentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createEmployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEmployment delete a employment
*/
func (a *Client) DeleteEmployment(params *DeleteEmploymentParams, opts ...ClientOption) (*DeleteEmploymentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEmploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteEmployment",
		Method:             "DELETE",
		PathPattern:        "/v1/employment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEmploymentReader{formats: a.formats},
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
	success, ok := result.(*DeleteEmploymentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEmployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEmployment return an employment
*/
func (a *Client) GetEmployment(params *GetEmploymentParams, opts ...ClientOption) (*GetEmploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEmployment",
		Method:             "GET",
		PathPattern:        "/v1/employment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmploymentReader{formats: a.formats},
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
	success, ok := result.(*GetEmploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEmployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEmploymentsByUserID return employments for a user
*/
func (a *Client) GetEmploymentsByUserID(params *GetEmploymentsByUserIDParams, opts ...ClientOption) (*GetEmploymentsByUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmploymentsByUserIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEmploymentsByUserId",
		Method:             "GET",
		PathPattern:        "/v1/employment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmploymentsByUserIDReader{formats: a.formats},
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
	success, ok := result.(*GetEmploymentsByUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEmploymentsByUserId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEmployment update a employment
*/
func (a *Client) UpdateEmployment(params *UpdateEmploymentParams, opts ...ClientOption) (*UpdateEmploymentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEmploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEmployment",
		Method:             "PUT",
		PathPattern:        "/v1/employment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEmploymentReader{formats: a.formats},
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
	success, ok := result.(*UpdateEmploymentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateEmployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
