// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetAccessTokenRequest GetAccessTokenRequest
//
// swagger:model GetAccessTokenRequest
type GetAccessTokenRequest struct {

	// app_key
	// Required: true
	AppKey *string `json:"app_key"`

	// Concatenate app_key, time, app_secret in turn, and do sha1().app_key = mBOMg20QW11BbtyH4Zh0 \n" +
	//             "time = 1647847498 \n" +
	//             "app_secret = V6aRfxlPJwN3ViJSIFSCdxPvneajuJsh \n" +
	//             "sign = sha1(mBOMg20QW11BbtyH4Zh01647847498V6aRfxlPJwN3ViJSIFSCdxPvneajuJsh)\n" +
	//             "        = 7293d385b9225b3c3f232b76ba97255d0e21063e
	// Required: true
	Sign *string `json:"sign"`

	// Quest timestamp (Second)
	// Required: true
	Time *int64 `json:"time"`
}

// Validate validates this get access token request
func (m *GetAccessTokenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAccessTokenRequest) validateAppKey(formats strfmt.Registry) error {

	if err := validate.Required("app_key", "body", m.AppKey); err != nil {
		return err
	}

	return nil
}

func (m *GetAccessTokenRequest) validateSign(formats strfmt.Registry) error {

	if err := validate.Required("sign", "body", m.Sign); err != nil {
		return err
	}

	return nil
}

func (m *GetAccessTokenRequest) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAccessTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAccessTokenRequest) UnmarshalBinary(b []byte) error {
	var res GetAccessTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}