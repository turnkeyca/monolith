// Code generated by go-swagger; DO NOT EDIT.

package employment

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

// NewDeleteEmploymentParams creates a new DeleteEmploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEmploymentParams() *DeleteEmploymentParams {
	return &DeleteEmploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEmploymentParamsWithTimeout creates a new DeleteEmploymentParams object
// with the ability to set a timeout on a request.
func NewDeleteEmploymentParamsWithTimeout(timeout time.Duration) *DeleteEmploymentParams {
	return &DeleteEmploymentParams{
		timeout: timeout,
	}
}

// NewDeleteEmploymentParamsWithContext creates a new DeleteEmploymentParams object
// with the ability to set a context for a request.
func NewDeleteEmploymentParamsWithContext(ctx context.Context) *DeleteEmploymentParams {
	return &DeleteEmploymentParams{
		Context: ctx,
	}
}

// NewDeleteEmploymentParamsWithHTTPClient creates a new DeleteEmploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEmploymentParamsWithHTTPClient(client *http.Client) *DeleteEmploymentParams {
	return &DeleteEmploymentParams{
		HTTPClient: client,
	}
}

/* DeleteEmploymentParams contains all the parameters to send to the API endpoint
   for the delete employment operation.

   Typically these are written to a http.Request.
*/
type DeleteEmploymentParams struct {

	// Token.
	Token string

	/* ID.

	   The id of the employment for which the operation relates
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete employment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEmploymentParams) WithDefaults() *DeleteEmploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete employment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEmploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete employment params
func (o *DeleteEmploymentParams) WithTimeout(timeout time.Duration) *DeleteEmploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete employment params
func (o *DeleteEmploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete employment params
func (o *DeleteEmploymentParams) WithContext(ctx context.Context) *DeleteEmploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete employment params
func (o *DeleteEmploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete employment params
func (o *DeleteEmploymentParams) WithHTTPClient(client *http.Client) *DeleteEmploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete employment params
func (o *DeleteEmploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the delete employment params
func (o *DeleteEmploymentParams) WithToken(token string) *DeleteEmploymentParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete employment params
func (o *DeleteEmploymentParams) SetToken(token string) {
	o.Token = token
}

// WithID adds the id to the delete employment params
func (o *DeleteEmploymentParams) WithID(id string) *DeleteEmploymentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete employment params
func (o *DeleteEmploymentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEmploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Token
	if err := r.SetHeaderParam("Token", o.Token); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
