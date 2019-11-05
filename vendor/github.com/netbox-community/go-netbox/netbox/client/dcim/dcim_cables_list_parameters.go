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

// NewDcimCablesListParams creates a new DcimCablesListParams object
// with the default values initialized.
func NewDcimCablesListParams() *DcimCablesListParams {
	var ()
	return &DcimCablesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesListParamsWithTimeout creates a new DcimCablesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimCablesListParamsWithTimeout(timeout time.Duration) *DcimCablesListParams {
	var ()
	return &DcimCablesListParams{

		timeout: timeout,
	}
}

// NewDcimCablesListParamsWithContext creates a new DcimCablesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimCablesListParamsWithContext(ctx context.Context) *DcimCablesListParams {
	var ()
	return &DcimCablesListParams{

		Context: ctx,
	}
}

// NewDcimCablesListParamsWithHTTPClient creates a new DcimCablesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimCablesListParamsWithHTTPClient(client *http.Client) *DcimCablesListParams {
	var ()
	return &DcimCablesListParams{
		HTTPClient: client,
	}
}

/*DcimCablesListParams contains all the parameters to send to the API endpoint
for the dcim cables list operation typically these are written to a http.Request
*/
type DcimCablesListParams struct {

	/*Color*/
	Color *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*ID*/
	ID *string
	/*Label*/
	Label *string
	/*Length*/
	Length *string
	/*LengthUnit*/
	LengthUnit *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Rack*/
	Rack *string
	/*RackID*/
	RackID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Status*/
	Status *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim cables list params
func (o *DcimCablesListParams) WithTimeout(timeout time.Duration) *DcimCablesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables list params
func (o *DcimCablesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables list params
func (o *DcimCablesListParams) WithContext(ctx context.Context) *DcimCablesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables list params
func (o *DcimCablesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables list params
func (o *DcimCablesListParams) WithHTTPClient(client *http.Client) *DcimCablesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables list params
func (o *DcimCablesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the dcim cables list params
func (o *DcimCablesListParams) WithColor(color *string) *DcimCablesListParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the dcim cables list params
func (o *DcimCablesListParams) SetColor(color *string) {
	o.Color = color
}

// WithDevice adds the device to the dcim cables list params
func (o *DcimCablesListParams) WithDevice(device *string) *DcimCablesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim cables list params
func (o *DcimCablesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim cables list params
func (o *DcimCablesListParams) WithDeviceID(deviceID *string) *DcimCablesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim cables list params
func (o *DcimCablesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim cables list params
func (o *DcimCablesListParams) WithID(id *string) *DcimCablesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables list params
func (o *DcimCablesListParams) SetID(id *string) {
	o.ID = id
}

// WithLabel adds the label to the dcim cables list params
func (o *DcimCablesListParams) WithLabel(label *string) *DcimCablesListParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the dcim cables list params
func (o *DcimCablesListParams) SetLabel(label *string) {
	o.Label = label
}

// WithLength adds the length to the dcim cables list params
func (o *DcimCablesListParams) WithLength(length *string) *DcimCablesListParams {
	o.SetLength(length)
	return o
}

// SetLength adds the length to the dcim cables list params
func (o *DcimCablesListParams) SetLength(length *string) {
	o.Length = length
}

// WithLengthUnit adds the lengthUnit to the dcim cables list params
func (o *DcimCablesListParams) WithLengthUnit(lengthUnit *string) *DcimCablesListParams {
	o.SetLengthUnit(lengthUnit)
	return o
}

// SetLengthUnit adds the lengthUnit to the dcim cables list params
func (o *DcimCablesListParams) SetLengthUnit(lengthUnit *string) {
	o.LengthUnit = lengthUnit
}

// WithLimit adds the limit to the dcim cables list params
func (o *DcimCablesListParams) WithLimit(limit *int64) *DcimCablesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim cables list params
func (o *DcimCablesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim cables list params
func (o *DcimCablesListParams) WithOffset(offset *int64) *DcimCablesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim cables list params
func (o *DcimCablesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim cables list params
func (o *DcimCablesListParams) WithQ(q *string) *DcimCablesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim cables list params
func (o *DcimCablesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRack adds the rack to the dcim cables list params
func (o *DcimCablesListParams) WithRack(rack *string) *DcimCablesListParams {
	o.SetRack(rack)
	return o
}

// SetRack adds the rack to the dcim cables list params
func (o *DcimCablesListParams) SetRack(rack *string) {
	o.Rack = rack
}

// WithRackID adds the rackID to the dcim cables list params
func (o *DcimCablesListParams) WithRackID(rackID *string) *DcimCablesListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim cables list params
func (o *DcimCablesListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithSite adds the site to the dcim cables list params
func (o *DcimCablesListParams) WithSite(site *string) *DcimCablesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim cables list params
func (o *DcimCablesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim cables list params
func (o *DcimCablesListParams) WithSiteID(siteID *string) *DcimCablesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim cables list params
func (o *DcimCablesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the dcim cables list params
func (o *DcimCablesListParams) WithStatus(status *string) *DcimCablesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim cables list params
func (o *DcimCablesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTenant adds the tenant to the dcim cables list params
func (o *DcimCablesListParams) WithTenant(tenant *string) *DcimCablesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim cables list params
func (o *DcimCablesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the dcim cables list params
func (o *DcimCablesListParams) WithTenantID(tenantID *string) *DcimCablesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim cables list params
func (o *DcimCablesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithType adds the typeVar to the dcim cables list params
func (o *DcimCablesListParams) WithType(typeVar *string) *DcimCablesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim cables list params
func (o *DcimCablesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string
		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {
			if err := r.SetQueryParam("color", qColor); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Label != nil {

		// query param label
		var qrLabel string
		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {
			if err := r.SetQueryParam("label", qLabel); err != nil {
				return err
			}
		}

	}

	if o.Length != nil {

		// query param length
		var qrLength string
		if o.Length != nil {
			qrLength = *o.Length
		}
		qLength := qrLength
		if qLength != "" {
			if err := r.SetQueryParam("length", qLength); err != nil {
				return err
			}
		}

	}

	if o.LengthUnit != nil {

		// query param length_unit
		var qrLengthUnit string
		if o.LengthUnit != nil {
			qrLengthUnit = *o.LengthUnit
		}
		qLengthUnit := qrLengthUnit
		if qLengthUnit != "" {
			if err := r.SetQueryParam("length_unit", qLengthUnit); err != nil {
				return err
			}
		}

	}

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

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Rack != nil {

		// query param rack
		var qrRack string
		if o.Rack != nil {
			qrRack = *o.Rack
		}
		qRack := qrRack
		if qRack != "" {
			if err := r.SetQueryParam("rack", qRack); err != nil {
				return err
			}
		}

	}

	if o.RackID != nil {

		// query param rack_id
		var qrRackID string
		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := qrRackID
		if qRackID != "" {
			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
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
