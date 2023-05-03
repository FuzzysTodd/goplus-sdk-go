// Code generated by go-swagger; DO NOT EDIT.

package token_controller

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
	"github.com/go-openapi/swag"
)

// NewGetAccessTokenUsingPOSTParams creates a new GetAccessTokenUsingPOSTParams object
// with the default values initialized.
func NewGetAccessTokenUsingPOSTParams() *GetAccessTokenUsingPOSTParams {
	var ()
	return &GetAccessTokenUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessTokenUsingPOSTParamsWithTimeout creates a new GetAccessTokenUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessTokenUsingPOSTParamsWithTimeout(timeout time.Duration) *GetAccessTokenUsingPOSTParams {
	var ()
	return &GetAccessTokenUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetAccessTokenUsingPOSTParamsWithContext creates a new GetAccessTokenUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessTokenUsingPOSTParamsWithContext(ctx context.Context) *GetAccessTokenUsingPOSTParams {
	var ()
	return &GetAccessTokenUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetAccessTokenUsingPOSTParamsWithHTTPClient creates a new GetAccessTokenUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessTokenUsingPOSTParamsWithHTTPClient(client *http.Client) *GetAccessTokenUsingPOSTParams {
	var ()
	return &GetAccessTokenUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
GetAccessTokenUsingPOSTParams contains all the parameters to send to the API endpoint
for the get access token using p o s t operation typically these are written to a http.Request
*/
type GetAccessTokenUsingPOSTParams struct {

	/*AppKey
	  app_key

	*/
	AppKey string
	/*Sign
	  Concatenate app_key, time, app_secret in turn, and do sha1().app_key = mBOMg20QW11BbtyH4Zh0 \n" +
	            "time = 1647847498 \n" +
	            "app_secret = V6aRfxlPJwN3ViJSIFSCdxPvneajuJsh \n" +
	            "sign = sha1(mBOMg20QW11BbtyH4Zh01647847498V6aRfxlPJwN3ViJSIFSCdxPvneajuJsh)\n" +
	            "        = 7293d385b9225b3c3f232b76ba97255d0e21063e

	*/
	Sign string
	/*Time
	  Quest timestamp (Second)

	*/
	Time int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithTimeout(timeout time.Duration) *GetAccessTokenUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithContext(ctx context.Context) *GetAccessTokenUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithHTTPClient(client *http.Client) *GetAccessTokenUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithAppKey(appKey string) *GetAccessTokenUsingPOSTParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithSign adds the sign to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithSign(sign string) *GetAccessTokenUsingPOSTParams {
	o.SetSign(sign)
	return o
}

// SetSign adds the sign to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetSign(sign string) {
	o.Sign = sign
}

// WithTime adds the time to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) WithTime(time int64) *GetAccessTokenUsingPOSTParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get access token using p o s t params
func (o *GetAccessTokenUsingPOSTParams) SetTime(time int64) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessTokenUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param appKey
	qrAppKey := o.AppKey
	qAppKey := qrAppKey
	if qAppKey != "" {
		if err := r.SetQueryParam("appKey", qAppKey); err != nil {
			return err
		}
	}

	// query param sign
	qrSign := o.Sign
	qSign := qrSign
	if qSign != "" {
		if err := r.SetQueryParam("sign", qSign); err != nil {
			return err
		}
	}

	// query param time
	qrTime := o.Time
	qTime := swag.FormatInt64(qrTime)
	if qTime != "" {
		if err := r.SetQueryParam("time", qTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}