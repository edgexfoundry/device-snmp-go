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

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureResponseService posture response service
//
// swagger:model postureResponseService
type PostureResponseService struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// posture query type
	// Required: true
	PostureQueryType *string `json:"postureQueryType"`

	// timeout
	// Required: true
	Timeout *int64 `json:"timeout"`

	// timeout remaining
	// Required: true
	TimeoutRemaining *int64 `json:"timeoutRemaining"`
}

// Validate validates this posture response service
func (m *PostureResponseService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostureQueryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutRemaining(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureResponseService) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostureResponseService) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostureResponseService) validatePostureQueryType(formats strfmt.Registry) error {

	if err := validate.Required("postureQueryType", "body", m.PostureQueryType); err != nil {
		return err
	}

	return nil
}

func (m *PostureResponseService) validateTimeout(formats strfmt.Registry) error {

	if err := validate.Required("timeout", "body", m.Timeout); err != nil {
		return err
	}

	return nil
}

func (m *PostureResponseService) validateTimeoutRemaining(formats strfmt.Registry) error {

	if err := validate.Required("timeoutRemaining", "body", m.TimeoutRemaining); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this posture response service based on context it is used
func (m *PostureResponseService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostureResponseService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureResponseService) UnmarshalBinary(b []byte) error {
	var res PostureResponseService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
