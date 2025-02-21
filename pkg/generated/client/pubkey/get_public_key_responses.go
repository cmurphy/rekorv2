// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2025 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pubkey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekorv2/pkg/generated/models"
)

// GetPublicKeyReader is a Reader for the GetPublicKey structure.
type GetPublicKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPublicKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPublicKeyOK creates a GetPublicKeyOK with default headers values
func NewGetPublicKeyOK() *GetPublicKeyOK {
	return &GetPublicKeyOK{}
}

/*
GetPublicKeyOK describes a response with status code 200, with default header values.

The public key
*/
type GetPublicKeyOK struct {
	Payload string
}

// IsSuccess returns true when this get public key o k response has a 2xx status code
func (o *GetPublicKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get public key o k response has a 3xx status code
func (o *GetPublicKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get public key o k response has a 4xx status code
func (o *GetPublicKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get public key o k response has a 5xx status code
func (o *GetPublicKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get public key o k response a status code equal to that given
func (o *GetPublicKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get public key o k response
func (o *GetPublicKeyOK) Code() int {
	return 200
}

func (o *GetPublicKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/publicKey][%d] getPublicKeyOK %s", 200, payload)
}

func (o *GetPublicKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/publicKey][%d] getPublicKeyOK %s", 200, payload)
}

func (o *GetPublicKeyOK) GetPayload() string {
	return o.Payload
}

func (o *GetPublicKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicKeyDefault creates a GetPublicKeyDefault with default headers values
func NewGetPublicKeyDefault(code int) *GetPublicKeyDefault {
	return &GetPublicKeyDefault{
		_statusCode: code,
	}
}

/*
GetPublicKeyDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetPublicKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get public key default response has a 2xx status code
func (o *GetPublicKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get public key default response has a 3xx status code
func (o *GetPublicKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get public key default response has a 4xx status code
func (o *GetPublicKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get public key default response has a 5xx status code
func (o *GetPublicKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get public key default response a status code equal to that given
func (o *GetPublicKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get public key default response
func (o *GetPublicKeyDefault) Code() int {
	return o._statusCode
}

func (o *GetPublicKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/publicKey][%d] getPublicKey default %s", o._statusCode, payload)
}

func (o *GetPublicKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/publicKey][%d] getPublicKey default %s", o._statusCode, payload)
}

func (o *GetPublicKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPublicKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
