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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smutel/go-netbox/netbox/models"
)

// ExtrasImageAttachmentsUpdateReader is a Reader for the ExtrasImageAttachmentsUpdate structure.
type ExtrasImageAttachmentsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasImageAttachmentsUpdateOK creates a ExtrasImageAttachmentsUpdateOK with default headers values
func NewExtrasImageAttachmentsUpdateOK() *ExtrasImageAttachmentsUpdateOK {
	return &ExtrasImageAttachmentsUpdateOK{}
}

/*ExtrasImageAttachmentsUpdateOK handles this case with default header values.

ExtrasImageAttachmentsUpdateOK extras image attachments update o k
*/
type ExtrasImageAttachmentsUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extrasImageAttachmentsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
