// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package ipam

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

// NewIpamPrefixesReadParams creates a new IpamPrefixesReadParams object
// with the default values initialized.
func NewIpamPrefixesReadParams() *IpamPrefixesReadParams {
	var ()
	return &IpamPrefixesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesReadParamsWithTimeout creates a new IpamPrefixesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamPrefixesReadParamsWithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	var ()
	return &IpamPrefixesReadParams{

		timeout: timeout,
	}
}

// NewIpamPrefixesReadParamsWithContext creates a new IpamPrefixesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamPrefixesReadParamsWithContext(ctx context.Context) *IpamPrefixesReadParams {
	var ()
	return &IpamPrefixesReadParams{

		Context: ctx,
	}
}

// NewIpamPrefixesReadParamsWithHTTPClient creates a new IpamPrefixesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamPrefixesReadParamsWithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	var ()
	return &IpamPrefixesReadParams{
		HTTPClient: client,
	}
}

/*IpamPrefixesReadParams contains all the parameters to send to the API endpoint
for the ipam prefixes read operation typically these are written to a http.Request
*/
type IpamPrefixesReadParams struct {

	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithContext(ctx context.Context) *IpamPrefixesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithID(id int64) *IpamPrefixesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
