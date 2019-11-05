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

package dcim

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

// NewDcimCablesReadParams creates a new DcimCablesReadParams object
// with the default values initialized.
func NewDcimCablesReadParams() *DcimCablesReadParams {
	var ()
	return &DcimCablesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesReadParamsWithTimeout creates a new DcimCablesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimCablesReadParamsWithTimeout(timeout time.Duration) *DcimCablesReadParams {
	var ()
	return &DcimCablesReadParams{

		timeout: timeout,
	}
}

// NewDcimCablesReadParamsWithContext creates a new DcimCablesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimCablesReadParamsWithContext(ctx context.Context) *DcimCablesReadParams {
	var ()
	return &DcimCablesReadParams{

		Context: ctx,
	}
}

// NewDcimCablesReadParamsWithHTTPClient creates a new DcimCablesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimCablesReadParamsWithHTTPClient(client *http.Client) *DcimCablesReadParams {
	var ()
	return &DcimCablesReadParams{
		HTTPClient: client,
	}
}

/*DcimCablesReadParams contains all the parameters to send to the API endpoint
for the dcim cables read operation typically these are written to a http.Request
*/
type DcimCablesReadParams struct {

	/*ID
	  A unique integer value identifying this cable.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim cables read params
func (o *DcimCablesReadParams) WithTimeout(timeout time.Duration) *DcimCablesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables read params
func (o *DcimCablesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables read params
func (o *DcimCablesReadParams) WithContext(ctx context.Context) *DcimCablesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables read params
func (o *DcimCablesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables read params
func (o *DcimCablesReadParams) WithHTTPClient(client *http.Client) *DcimCablesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables read params
func (o *DcimCablesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim cables read params
func (o *DcimCablesReadParams) WithID(id int64) *DcimCablesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables read params
func (o *DcimCablesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
