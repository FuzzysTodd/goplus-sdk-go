// Code generated by go-swagger; DO NOT EDIT.

package dapp_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

// GetDappInfoUsingGETReader is a Reader for the GetDappInfoUsingGET structure.
type GetDappInfoUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDappInfoUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDappInfoUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDappInfoUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDappInfoUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDappInfoUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDappInfoUsingGETOK creates a GetDappInfoUsingGETOK with default headers values
func NewGetDappInfoUsingGETOK() *GetDappInfoUsingGETOK {
	return &GetDappInfoUsingGETOK{}
}

/*
GetDappInfoUsingGETOK handles this case with default header values.

OK
*/
type GetDappInfoUsingGETOK struct {
	Payload *models.ResponseWrapperDappContractSecurityResponse
}

func (o *GetDappInfoUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/dapp_security][%d] getDappInfoUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDappInfoUsingGETOK) GetPayload() *models.ResponseWrapperDappContractSecurityResponse {
	return o.Payload
}

func (o *GetDappInfoUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDappContractSecurityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDappInfoUsingGETUnauthorized creates a GetDappInfoUsingGETUnauthorized with default headers values
func NewGetDappInfoUsingGETUnauthorized() *GetDappInfoUsingGETUnauthorized {
	return &GetDappInfoUsingGETUnauthorized{}
}

/*
GetDappInfoUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDappInfoUsingGETUnauthorized struct {
}

func (o *GetDappInfoUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/dapp_security][%d] getDappInfoUsingGETUnauthorized ", 401)
}

func (o *GetDappInfoUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDappInfoUsingGETForbidden creates a GetDappInfoUsingGETForbidden with default headers values
func NewGetDappInfoUsingGETForbidden() *GetDappInfoUsingGETForbidden {
	return &GetDappInfoUsingGETForbidden{}
}

/*
GetDappInfoUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetDappInfoUsingGETForbidden struct {
}

func (o *GetDappInfoUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/dapp_security][%d] getDappInfoUsingGETForbidden ", 403)
}

func (o *GetDappInfoUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDappInfoUsingGETNotFound creates a GetDappInfoUsingGETNotFound with default headers values
func NewGetDappInfoUsingGETNotFound() *GetDappInfoUsingGETNotFound {
	return &GetDappInfoUsingGETNotFound{}
}

/*
GetDappInfoUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetDappInfoUsingGETNotFound struct {
}

func (o *GetDappInfoUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/dapp_security][%d] getDappInfoUsingGETNotFound ", 404)
}

func (o *GetDappInfoUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}