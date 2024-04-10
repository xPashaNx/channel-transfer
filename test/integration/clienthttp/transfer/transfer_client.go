// Code generated by go-swagger; DO NOT EDIT.

package transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transfer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transfer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TransferByAdmin(params *TransferByAdminParams, opts ...ClientOption) (*TransferByAdminOK, error)

	TransferByCustomer(params *TransferByCustomerParams, opts ...ClientOption) (*TransferByCustomerOK, error)

	TransferStatus(params *TransferStatusParams, opts ...ClientOption) (*TransferStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TransferByAdmin transfers cross chanel be admin
*/
func (a *Client) TransferByAdmin(params *TransferByAdminParams, opts ...ClientOption) (*TransferByAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferByAdminParams()
	}

	op := &runtime.ClientOperation{
		ID:                 "transferByAdmin",
		Method:             "POST",
		PathPattern:        "/v1/transfer/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransferByAdminReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}

	success, ok := result.(*TransferByAdminOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransferByAdminDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TransferByCustomer transfers cross chanel be customer
*/
func (a *Client) TransferByCustomer(params *TransferByCustomerParams, opts ...ClientOption) (*TransferByCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferByCustomerParams()
	}

	op := &runtime.ClientOperation{
		ID:                 "transferByCustomer",
		Method:             "POST",
		PathPattern:        "/v1/transfer/customer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransferByCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}

	success, ok := result.(*TransferByCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransferByCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TransferStatus gets status transfer by id transfer
*/
func (a *Client) TransferStatus(params *TransferStatusParams, opts ...ClientOption) (*TransferStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferStatusParams()
	}

	op := &runtime.ClientOperation{
		ID:                 "transferStatus",
		Method:             "GET",
		PathPattern:        "/v1/status/{idTransfer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransferStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}

	success, ok := result.(*TransferStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransferStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
