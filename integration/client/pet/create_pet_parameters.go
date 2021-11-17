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

	"github.com/turnkeyca/monolith/integration/models"
)

// NewCreatePetParams creates a new CreatePetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePetParams() *CreatePetParams {
	return &CreatePetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePetParamsWithTimeout creates a new CreatePetParams object
// with the ability to set a timeout on a request.
func NewCreatePetParamsWithTimeout(timeout time.Duration) *CreatePetParams {
	return &CreatePetParams{
		timeout: timeout,
	}
}

// NewCreatePetParamsWithContext creates a new CreatePetParams object
// with the ability to set a context for a request.
func NewCreatePetParamsWithContext(ctx context.Context) *CreatePetParams {
	return &CreatePetParams{
		Context: ctx,
	}
}

// NewCreatePetParamsWithHTTPClient creates a new CreatePetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePetParamsWithHTTPClient(client *http.Client) *CreatePetParams {
	return &CreatePetParams{
		HTTPClient: client,
	}
}

/* CreatePetParams contains all the parameters to send to the API endpoint
   for the create pet operation.

   Typically these are written to a http.Request.
*/
type CreatePetParams struct {

	// Body.
	Body *models.PetDto

	// Token.
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create pet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePetParams) WithDefaults() *CreatePetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pet params
func (o *CreatePetParams) WithTimeout(timeout time.Duration) *CreatePetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pet params
func (o *CreatePetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pet params
func (o *CreatePetParams) WithContext(ctx context.Context) *CreatePetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pet params
func (o *CreatePetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pet params
func (o *CreatePetParams) WithHTTPClient(client *http.Client) *CreatePetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pet params
func (o *CreatePetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pet params
func (o *CreatePetParams) WithBody(body *models.PetDto) *CreatePetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pet params
func (o *CreatePetParams) SetBody(body *models.PetDto) {
	o.Body = body
}

// WithToken adds the token to the create pet params
func (o *CreatePetParams) WithToken(token string) *CreatePetParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the create pet params
func (o *CreatePetParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
