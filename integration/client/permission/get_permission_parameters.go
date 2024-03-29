// Code generated by go-swagger; DO NOT EDIT.

package permission

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

// NewGetPermissionParams creates a new GetPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPermissionParams() *GetPermissionParams {
	return &GetPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionParamsWithTimeout creates a new GetPermissionParams object
// with the ability to set a timeout on a request.
func NewGetPermissionParamsWithTimeout(timeout time.Duration) *GetPermissionParams {
	return &GetPermissionParams{
		timeout: timeout,
	}
}

// NewGetPermissionParamsWithContext creates a new GetPermissionParams object
// with the ability to set a context for a request.
func NewGetPermissionParamsWithContext(ctx context.Context) *GetPermissionParams {
	return &GetPermissionParams{
		Context: ctx,
	}
}

// NewGetPermissionParamsWithHTTPClient creates a new GetPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPermissionParamsWithHTTPClient(client *http.Client) *GetPermissionParams {
	return &GetPermissionParams{
		HTTPClient: client,
	}
}

/* GetPermissionParams contains all the parameters to send to the API endpoint
   for the get permission operation.

   Typically these are written to a http.Request.
*/
type GetPermissionParams struct {

	// Token.
	Token string

	/* ID.

	   The id of the permission for which the operation relates
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionParams) WithDefaults() *GetPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get permission params
func (o *GetPermissionParams) WithTimeout(timeout time.Duration) *GetPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get permission params
func (o *GetPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get permission params
func (o *GetPermissionParams) WithContext(ctx context.Context) *GetPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get permission params
func (o *GetPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get permission params
func (o *GetPermissionParams) WithHTTPClient(client *http.Client) *GetPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get permission params
func (o *GetPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the get permission params
func (o *GetPermissionParams) WithToken(token string) *GetPermissionParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get permission params
func (o *GetPermissionParams) SetToken(token string) {
	o.Token = token
}

// WithID adds the id to the get permission params
func (o *GetPermissionParams) WithID(id string) *GetPermissionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get permission params
func (o *GetPermissionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
