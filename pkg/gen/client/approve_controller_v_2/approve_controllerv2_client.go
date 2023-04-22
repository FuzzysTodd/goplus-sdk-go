// Code generated by go-swagger; DO NOT EDIT.

package approve_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new approve controller v 2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for approve controller v 2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddressNFT1155ApproveListUsingGET1(params *AddressNFT1155ApproveListUsingGET1Params, opts ...ClientOption) (*AddressNFT1155ApproveListUsingGET1OK, error)

	AddressNFT721ApproveListUsingGET1(params *AddressNFT721ApproveListUsingGET1Params, opts ...ClientOption) (*AddressNFT721ApproveListUsingGET1OK, error)

	AddressTokenApproveListUsingGET1(params *AddressTokenApproveListUsingGET1Params, opts ...ClientOption) (*AddressTokenApproveListUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddressNFT1155ApproveListUsingGET1 authorizations authorized address list n f t1155
*/
func (a *Client) AddressNFT1155ApproveListUsingGET1(params *AddressNFT1155ApproveListUsingGET1Params, opts ...ClientOption) (*AddressNFT1155ApproveListUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressNFT1155ApproveListUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressNFT1155ApproveListUsingGET_1",
		Method:             "GET",
		PathPattern:        "/api/v2/nft1155_approval_security/{chainId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddressNFT1155ApproveListUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*AddressNFT1155ApproveListUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressNFT1155ApproveListUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddressNFT721ApproveListUsingGET1 authorizations authorized address list n f t721
*/
func (a *Client) AddressNFT721ApproveListUsingGET1(params *AddressNFT721ApproveListUsingGET1Params, opts ...ClientOption) (*AddressNFT721ApproveListUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressNFT721ApproveListUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressNFT721ApproveListUsingGET_1",
		Method:             "GET",
		PathPattern:        "/api/v2/nft721_approval_security/{chainId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddressNFT721ApproveListUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*AddressNFT721ApproveListUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressNFT721ApproveListUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddressTokenApproveListUsingGET1 authorizations e r c20 authorized address list
*/
func (a *Client) AddressTokenApproveListUsingGET1(params *AddressTokenApproveListUsingGET1Params, opts ...ClientOption) (*AddressTokenApproveListUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressTokenApproveListUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressTokenApproveListUsingGET_1",
		Method:             "GET",
		PathPattern:        "/api/v2/token_approval_security/{chainId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddressTokenApproveListUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*AddressTokenApproveListUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressTokenApproveListUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
