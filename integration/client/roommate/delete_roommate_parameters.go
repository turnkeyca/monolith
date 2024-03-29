// Code generated by go-swagger; DO NOT EDIT.

package roommate

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

// NewDeleteRoommateParams creates a new DeleteRoommateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoommateParams() *DeleteRoommateParams {
	return &DeleteRoommateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoommateParamsWithTimeout creates a new DeleteRoommateParams object
// with the ability to set a timeout on a request.
func NewDeleteRoommateParamsWithTimeout(timeout time.Duration) *DeleteRoommateParams {
	return &DeleteRoommateParams{
		timeout: timeout,
	}
}

// NewDeleteRoommateParamsWithContext creates a new DeleteRoommateParams object
// with the ability to set a context for a request.
func NewDeleteRoommateParamsWithContext(ctx context.Context) *DeleteRoommateParams {
	return &DeleteRoommateParams{
		Context: ctx,
	}
}

// NewDeleteRoommateParamsWithHTTPClient creates a new DeleteRoommateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoommateParamsWithHTTPClient(client *http.Client) *DeleteRoommateParams {
	return &DeleteRoommateParams{
		HTTPClient: client,
	}
}

/* DeleteRoommateParams contains all the parameters to send to the API endpoint
   for the delete roommate operation.

   Typically these are written to a http.Request.
*/
type DeleteRoommateParams struct {

	// Token.
	Token string

	/* ID.

	   The id of the roommate for which the operation relates
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete roommate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoommateParams) WithDefaults() *DeleteRoommateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete roommate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoommateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete roommate params
func (o *DeleteRoommateParams) WithTimeout(timeout time.Duration) *DeleteRoommateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete roommate params
func (o *DeleteRoommateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete roommate params
func (o *DeleteRoommateParams) WithContext(ctx context.Context) *DeleteRoommateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete roommate params
func (o *DeleteRoommateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete roommate params
func (o *DeleteRoommateParams) WithHTTPClient(client *http.Client) *DeleteRoommateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete roommate params
func (o *DeleteRoommateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the delete roommate params
func (o *DeleteRoommateParams) WithToken(token string) *DeleteRoommateParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete roommate params
func (o *DeleteRoommateParams) SetToken(token string) {
	o.Token = token
}

// WithID adds the id to the delete roommate params
func (o *DeleteRoommateParams) WithID(id string) *DeleteRoommateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete roommate params
func (o *DeleteRoommateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoommateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
