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

// NewGetEmploymentsByUserIDParams creates a new GetEmploymentsByUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEmploymentsByUserIDParams() *GetEmploymentsByUserIDParams {
	return &GetEmploymentsByUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmploymentsByUserIDParamsWithTimeout creates a new GetEmploymentsByUserIDParams object
// with the ability to set a timeout on a request.
func NewGetEmploymentsByUserIDParamsWithTimeout(timeout time.Duration) *GetEmploymentsByUserIDParams {
	return &GetEmploymentsByUserIDParams{
		timeout: timeout,
	}
}

// NewGetEmploymentsByUserIDParamsWithContext creates a new GetEmploymentsByUserIDParams object
// with the ability to set a context for a request.
func NewGetEmploymentsByUserIDParamsWithContext(ctx context.Context) *GetEmploymentsByUserIDParams {
	return &GetEmploymentsByUserIDParams{
		Context: ctx,
	}
}

// NewGetEmploymentsByUserIDParamsWithHTTPClient creates a new GetEmploymentsByUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEmploymentsByUserIDParamsWithHTTPClient(client *http.Client) *GetEmploymentsByUserIDParams {
	return &GetEmploymentsByUserIDParams{
		HTTPClient: client,
	}
}

/* GetEmploymentsByUserIDParams contains all the parameters to send to the API endpoint
   for the get employments by user Id operation.

   Typically these are written to a http.Request.
*/
type GetEmploymentsByUserIDParams struct {

	// Token.
	Token string

	/* UserID.

	   The user id
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get employments by user Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmploymentsByUserIDParams) WithDefaults() *GetEmploymentsByUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get employments by user Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmploymentsByUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) WithTimeout(timeout time.Duration) *GetEmploymentsByUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) WithContext(ctx context.Context) *GetEmploymentsByUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) WithHTTPClient(client *http.Client) *GetEmploymentsByUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) WithToken(token string) *GetEmploymentsByUserIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) SetToken(token string) {
	o.Token = token
}

// WithUserID adds the userID to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) WithUserID(userID string) *GetEmploymentsByUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get employments by user Id params
func (o *GetEmploymentsByUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmploymentsByUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Token
	if err := r.SetHeaderParam("Token", o.Token); err != nil {
		return err
	}

	// query param userId
	qrUserID := o.UserId
	qUserID := qrUserID
	if qUserID != "" {

		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
