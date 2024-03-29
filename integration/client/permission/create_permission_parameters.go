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

	"github.com/turnkeyca/monolith/integration/models"
)

// NewCreatePermissionParams creates a new CreatePermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePermissionParams() *CreatePermissionParams {
	return &CreatePermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePermissionParamsWithTimeout creates a new CreatePermissionParams object
// with the ability to set a timeout on a request.
func NewCreatePermissionParamsWithTimeout(timeout time.Duration) *CreatePermissionParams {
	return &CreatePermissionParams{
		timeout: timeout,
	}
}

// NewCreatePermissionParamsWithContext creates a new CreatePermissionParams object
// with the ability to set a context for a request.
func NewCreatePermissionParamsWithContext(ctx context.Context) *CreatePermissionParams {
	return &CreatePermissionParams{
		Context: ctx,
	}
}

// NewCreatePermissionParamsWithHTTPClient creates a new CreatePermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePermissionParamsWithHTTPClient(client *http.Client) *CreatePermissionParams {
	return &CreatePermissionParams{
		HTTPClient: client,
	}
}

/* CreatePermissionParams contains all the parameters to send to the API endpoint
   for the create permission operation.

   Typically these are written to a http.Request.
*/
type CreatePermissionParams struct {

	// Body.
	Body *models.PermissionDto

	// Token.
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePermissionParams) WithDefaults() *CreatePermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create permission params
func (o *CreatePermissionParams) WithTimeout(timeout time.Duration) *CreatePermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create permission params
func (o *CreatePermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create permission params
func (o *CreatePermissionParams) WithContext(ctx context.Context) *CreatePermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create permission params
func (o *CreatePermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create permission params
func (o *CreatePermissionParams) WithHTTPClient(client *http.Client) *CreatePermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create permission params
func (o *CreatePermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create permission params
func (o *CreatePermissionParams) WithBody(body *models.PermissionDto) *CreatePermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create permission params
func (o *CreatePermissionParams) SetBody(body *models.PermissionDto) {
	o.Body = body
}

// WithToken adds the token to the create permission params
func (o *CreatePermissionParams) WithToken(token string) *CreatePermissionParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the create permission params
func (o *CreatePermissionParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Token
	if err := r.SetHeaderParam("Token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
