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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewIpamRirsPartialUpdateParams creates a new IpamRirsPartialUpdateParams object
// with the default values initialized.
func NewIpamRirsPartialUpdateParams() *IpamRirsPartialUpdateParams {
	var ()
	return &IpamRirsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsPartialUpdateParamsWithTimeout creates a new IpamRirsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamRirsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRirsPartialUpdateParams {
	var ()
	return &IpamRirsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIpamRirsPartialUpdateParamsWithContext creates a new IpamRirsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamRirsPartialUpdateParamsWithContext(ctx context.Context) *IpamRirsPartialUpdateParams {
	var ()
	return &IpamRirsPartialUpdateParams{

		Context: ctx,
	}
}

// NewIpamRirsPartialUpdateParamsWithHTTPClient creates a new IpamRirsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamRirsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRirsPartialUpdateParams {
	var ()
	return &IpamRirsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IpamRirsPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam rirs partial update operation typically these are written to a http.Request
*/
type IpamRirsPartialUpdateParams struct {

	/*Data*/
	Data *models.RIR
	/*ID
	  A unique integer value identifying this RIR.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRirsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) WithContext(ctx context.Context) *IpamRirsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRirsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) WithData(data *models.RIR) *IpamRirsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) SetData(data *models.RIR) {
	o.Data = data
}

// WithID adds the id to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) WithID(id int64) *IpamRirsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs partial update params
func (o *IpamRirsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
