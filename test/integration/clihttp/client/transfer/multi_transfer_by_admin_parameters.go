// Code generated by go-swagger; DO NOT EDIT.

package transfer

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

	"github.com/anoideaopen/channel-transfer/test/integration/clihttp/models"
)

// NewMultiTransferByAdminParams creates a new MultiTransferByAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiTransferByAdminParams() *MultiTransferByAdminParams {
	return &MultiTransferByAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiTransferByAdminParamsWithTimeout creates a new MultiTransferByAdminParams object
// with the ability to set a timeout on a request.
func NewMultiTransferByAdminParamsWithTimeout(timeout time.Duration) *MultiTransferByAdminParams {
	return &MultiTransferByAdminParams{
		timeout: timeout,
	}
}

// NewMultiTransferByAdminParamsWithContext creates a new MultiTransferByAdminParams object
// with the ability to set a context for a request.
func NewMultiTransferByAdminParamsWithContext(ctx context.Context) *MultiTransferByAdminParams {
	return &MultiTransferByAdminParams{
		Context: ctx,
	}
}

// NewMultiTransferByAdminParamsWithHTTPClient creates a new MultiTransferByAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiTransferByAdminParamsWithHTTPClient(client *http.Client) *MultiTransferByAdminParams {
	return &MultiTransferByAdminParams{
		HTTPClient: client,
	}
}

/*
MultiTransferByAdminParams contains all the parameters to send to the API endpoint

	for the multi transfer by admin operation.

	Typically these are written to a http.Request.
*/
type MultiTransferByAdminParams struct {

	// Body.
	Body *models.ChannelTransferMultiTransferBeginAdminRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi transfer by admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiTransferByAdminParams) WithDefaults() *MultiTransferByAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi transfer by admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiTransferByAdminParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi transfer by admin params
func (o *MultiTransferByAdminParams) WithTimeout(timeout time.Duration) *MultiTransferByAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi transfer by admin params
func (o *MultiTransferByAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi transfer by admin params
func (o *MultiTransferByAdminParams) WithContext(ctx context.Context) *MultiTransferByAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi transfer by admin params
func (o *MultiTransferByAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi transfer by admin params
func (o *MultiTransferByAdminParams) WithHTTPClient(client *http.Client) *MultiTransferByAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi transfer by admin params
func (o *MultiTransferByAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the multi transfer by admin params
func (o *MultiTransferByAdminParams) WithBody(body *models.ChannelTransferMultiTransferBeginAdminRequest) *MultiTransferByAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the multi transfer by admin params
func (o *MultiTransferByAdminParams) SetBody(body *models.ChannelTransferMultiTransferBeginAdminRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MultiTransferByAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
