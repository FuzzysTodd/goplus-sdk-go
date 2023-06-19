// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperDappContractSecurityResponse ResponseWrapperDappContractSecurityResponse
//
// swagger:model ResponseWrapperDappContractSecurityResponse
type ResponseWrapperDappContractSecurityResponse struct {

	// Code 1：Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result *DappContractSecurityResponse `json:"result,omitempty"`
}

// Validate validates this response wrapper dapp contract security response
func (m *ResponseWrapperDappContractSecurityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperDappContractSecurityResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperDappContractSecurityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperDappContractSecurityResponse) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperDappContractSecurityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
