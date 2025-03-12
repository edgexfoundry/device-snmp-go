// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package informational

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListEnumeratedCapabilitiesReader is a Reader for the ListEnumeratedCapabilities structure.
type ListEnumeratedCapabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEnumeratedCapabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEnumeratedCapabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEnumeratedCapabilitiesOK creates a ListEnumeratedCapabilitiesOK with default headers values
func NewListEnumeratedCapabilitiesOK() *ListEnumeratedCapabilitiesOK {
	return &ListEnumeratedCapabilitiesOK{}
}

/* ListEnumeratedCapabilitiesOK describes a response with status code 200, with default header values.

A typed and enumerated list of capabilities
*/
type ListEnumeratedCapabilitiesOK struct {
	Payload *rest_model.ListEnumeratedCapabilitiesEnvelope
}

func (o *ListEnumeratedCapabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /enumerated-capabilities][%d] listEnumeratedCapabilitiesOK  %+v", 200, o.Payload)
}
func (o *ListEnumeratedCapabilitiesOK) GetPayload() *rest_model.ListEnumeratedCapabilitiesEnvelope {
	return o.Payload
}

func (o *ListEnumeratedCapabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEnumeratedCapabilitiesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
