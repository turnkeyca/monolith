// Code generated by go-swagger; DO NOT EDIT.

package pet

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

// NewDeletePetParams creates a new DeletePetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePetParams() *DeletePetParams {
	return &DeletePetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePetParamsWithTimeout creates a new DeletePetParams object
// with the ability to set a timeout on a request.
func NewDeletePetParamsWithTimeout(timeout time.Duration) *DeletePetParams {
	return &DeletePetParams{
		timeout: timeout,
	}
}

// NewDeletePetParamsWithContext creates a new DeletePetParams object
// with the ability to set a context for a request.
func NewDeletePetParamsWithContext(ctx context.Context) *DeletePetParams {
	return &DeletePetParams{
		Context: ctx,
	}
}

// NewDeletePetParamsWithHTTPClient creates a new DeletePetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePetParamsWithHTTPClient(client *http.Client) *DeletePetParams {
	return &DeletePetParams{
		HTTPClient: client,
	}
}

/* DeletePetParams contains all the parameters to send to the API endpoint
   for the delete pet operation.

   Typically these are written to a http.Request.
*/
type DeletePetParams struct {

	// Token.
	Token string

	/* ID.

	   The id of the pet for which the operation relates
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete pet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePetParams) WithDefaults() *DeletePetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete pet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete pet params
func (o *DeletePetParams) WithTimeout(timeout time.Duration) *DeletePetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pet params
func (o *DeletePetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pet params
func (o *DeletePetParams) WithContext(ctx context.Context) *DeletePetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pet params
func (o *DeletePetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pet params
func (o *DeletePetParams) WithHTTPClient(client *http.Client) *DeletePetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pet params
func (o *DeletePetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the delete pet params
func (o *DeletePetParams) WithToken(token string) *DeletePetParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete pet params
func (o *DeletePetParams) SetToken(token string) {
	o.Token = token
}

// WithID adds the id to the delete pet params
func (o *DeletePetParams) WithID(id string) *DeletePetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete pet params
func (o *DeletePetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Token
	if err := r.SetHeaderParam("Token", o.Token); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.Id); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
