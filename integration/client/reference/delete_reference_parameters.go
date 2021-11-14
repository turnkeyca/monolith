// Code generated by go-swagger; DO NOT EDIT.

package reference

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

// NewDeleteReferenceParams creates a new DeleteReferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteReferenceParams() *DeleteReferenceParams {
	return &DeleteReferenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReferenceParamsWithTimeout creates a new DeleteReferenceParams object
// with the ability to set a timeout on a request.
func NewDeleteReferenceParamsWithTimeout(timeout time.Duration) *DeleteReferenceParams {
	return &DeleteReferenceParams{
		timeout: timeout,
	}
}

// NewDeleteReferenceParamsWithContext creates a new DeleteReferenceParams object
// with the ability to set a context for a request.
func NewDeleteReferenceParamsWithContext(ctx context.Context) *DeleteReferenceParams {
	return &DeleteReferenceParams{
		Context: ctx,
	}
}

// NewDeleteReferenceParamsWithHTTPClient creates a new DeleteReferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteReferenceParamsWithHTTPClient(client *http.Client) *DeleteReferenceParams {
	return &DeleteReferenceParams{
		HTTPClient: client,
	}
}

/* DeleteReferenceParams contains all the parameters to send to the API endpoint
   for the delete reference operation.

   Typically these are written to a http.Request.
*/
type DeleteReferenceParams struct {

	// Token.
	Token string

	/* ID.

	   The id of the reference for which the operation relates
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete reference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReferenceParams) WithDefaults() *DeleteReferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete reference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete reference params
func (o *DeleteReferenceParams) WithTimeout(timeout time.Duration) *DeleteReferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete reference params
func (o *DeleteReferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete reference params
func (o *DeleteReferenceParams) WithContext(ctx context.Context) *DeleteReferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete reference params
func (o *DeleteReferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete reference params
func (o *DeleteReferenceParams) WithHTTPClient(client *http.Client) *DeleteReferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete reference params
func (o *DeleteReferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the delete reference params
func (o *DeleteReferenceParams) WithToken(token string) *DeleteReferenceParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete reference params
func (o *DeleteReferenceParams) SetToken(token string) {
	o.Token = token
}

// WithID adds the id to the delete reference params
func (o *DeleteReferenceParams) WithID(id string) *DeleteReferenceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete reference params
func (o *DeleteReferenceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
