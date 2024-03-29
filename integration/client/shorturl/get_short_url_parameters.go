// Code generated by go-swagger; DO NOT EDIT.

package shorturl

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

// NewGetShortURLParams creates a new GetShortURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetShortURLParams() *GetShortURLParams {
	return &GetShortURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetShortURLParamsWithTimeout creates a new GetShortURLParams object
// with the ability to set a timeout on a request.
func NewGetShortURLParamsWithTimeout(timeout time.Duration) *GetShortURLParams {
	return &GetShortURLParams{
		timeout: timeout,
	}
}

// NewGetShortURLParamsWithContext creates a new GetShortURLParams object
// with the ability to set a context for a request.
func NewGetShortURLParamsWithContext(ctx context.Context) *GetShortURLParams {
	return &GetShortURLParams{
		Context: ctx,
	}
}

// NewGetShortURLParamsWithHTTPClient creates a new GetShortURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetShortURLParamsWithHTTPClient(client *http.Client) *GetShortURLParams {
	return &GetShortURLParams{
		HTTPClient: client,
	}
}

/* GetShortURLParams contains all the parameters to send to the API endpoint
   for the get short Url operation.

   Typically these are written to a http.Request.
*/
type GetShortURLParams struct {

	// Token.
	Token string

	/* URL.

	   The url to be converted
	*/
	URL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get short Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShortURLParams) WithDefaults() *GetShortURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get short Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShortURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get short Url params
func (o *GetShortURLParams) WithTimeout(timeout time.Duration) *GetShortURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get short Url params
func (o *GetShortURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get short Url params
func (o *GetShortURLParams) WithContext(ctx context.Context) *GetShortURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get short Url params
func (o *GetShortURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get short Url params
func (o *GetShortURLParams) WithHTTPClient(client *http.Client) *GetShortURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get short Url params
func (o *GetShortURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the get short Url params
func (o *GetShortURLParams) WithToken(token string) *GetShortURLParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get short Url params
func (o *GetShortURLParams) SetToken(token string) {
	o.Token = token
}

// WithURL adds the url to the get short Url params
func (o *GetShortURLParams) WithURL(url string) *GetShortURLParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the get short Url params
func (o *GetShortURLParams) SetURL(url string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *GetShortURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Token
	if err := r.SetHeaderParam("Token", o.Token); err != nil {
		return err
	}

	// query param url
	qrURL := o.URL
	qURL := qrURL
	if qURL != "" {

		if err := r.SetQueryParam("url", qURL); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
