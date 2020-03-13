// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasGraphsListParams creates a new ExtrasGraphsListParams object
// with the default values initialized.
func NewExtrasGraphsListParams() *ExtrasGraphsListParams {
	var ()
	return &ExtrasGraphsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsListParamsWithTimeout creates a new ExtrasGraphsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsListParamsWithTimeout(timeout time.Duration) *ExtrasGraphsListParams {
	var ()
	return &ExtrasGraphsListParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsListParamsWithContext creates a new ExtrasGraphsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsListParamsWithContext(ctx context.Context) *ExtrasGraphsListParams {
	var ()
	return &ExtrasGraphsListParams{

		Context: ctx,
	}
}

// NewExtrasGraphsListParamsWithHTTPClient creates a new ExtrasGraphsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsListParamsWithHTTPClient(client *http.Client) *ExtrasGraphsListParams {
	var ()
	return &ExtrasGraphsListParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsListParams contains all the parameters to send to the API endpoint
for the extras graphs list operation typically these are written to a http.Request
*/
type ExtrasGraphsListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*TemplateLanguage*/
	TemplateLanguage *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs list params
func (o *ExtrasGraphsListParams) WithTimeout(timeout time.Duration) *ExtrasGraphsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs list params
func (o *ExtrasGraphsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs list params
func (o *ExtrasGraphsListParams) WithContext(ctx context.Context) *ExtrasGraphsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs list params
func (o *ExtrasGraphsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs list params
func (o *ExtrasGraphsListParams) WithHTTPClient(client *http.Client) *ExtrasGraphsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs list params
func (o *ExtrasGraphsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the extras graphs list params
func (o *ExtrasGraphsListParams) WithLimit(limit *int64) *ExtrasGraphsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras graphs list params
func (o *ExtrasGraphsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras graphs list params
func (o *ExtrasGraphsListParams) WithName(name *string) *ExtrasGraphsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras graphs list params
func (o *ExtrasGraphsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the extras graphs list params
func (o *ExtrasGraphsListParams) WithOffset(offset *int64) *ExtrasGraphsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras graphs list params
func (o *ExtrasGraphsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithTemplateLanguage adds the templateLanguage to the extras graphs list params
func (o *ExtrasGraphsListParams) WithTemplateLanguage(templateLanguage *string) *ExtrasGraphsListParams {
	o.SetTemplateLanguage(templateLanguage)
	return o
}

// SetTemplateLanguage adds the templateLanguage to the extras graphs list params
func (o *ExtrasGraphsListParams) SetTemplateLanguage(templateLanguage *string) {
	o.TemplateLanguage = templateLanguage
}

// WithType adds the typeVar to the extras graphs list params
func (o *ExtrasGraphsListParams) WithType(typeVar *string) *ExtrasGraphsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the extras graphs list params
func (o *ExtrasGraphsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.TemplateLanguage != nil {

		// query param template_language
		var qrTemplateLanguage string
		if o.TemplateLanguage != nil {
			qrTemplateLanguage = *o.TemplateLanguage
		}
		qTemplateLanguage := qrTemplateLanguage
		if qTemplateLanguage != "" {
			if err := r.SetQueryParam("template_language", qTemplateLanguage); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
