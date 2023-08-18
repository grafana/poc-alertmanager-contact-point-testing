// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package testable_receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostTestReceiversConfigOKCode is the HTTP code returned for type PostTestReceiversConfigOK
const PostTestReceiversConfigOKCode int = 200

/*
PostTestReceiversConfigOK Successfully tested all receivers from provided configuration file

swagger:response postTestReceiversConfigOK
*/
type PostTestReceiversConfigOK struct {

	/*
	  In: Body
	*/
	Payload []*PostTestReceiversConfigOKBodyItems0 `json:"body,omitempty"`
}

// NewPostTestReceiversConfigOK creates PostTestReceiversConfigOK with default headers values
func NewPostTestReceiversConfigOK() *PostTestReceiversConfigOK {

	return &PostTestReceiversConfigOK{}
}

// WithPayload adds the payload to the post test receivers config o k response
func (o *PostTestReceiversConfigOK) WithPayload(payload []*PostTestReceiversConfigOKBodyItems0) *PostTestReceiversConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test receivers config o k response
func (o *PostTestReceiversConfigOK) SetPayload(payload []*PostTestReceiversConfigOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestReceiversConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*PostTestReceiversConfigOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostTestReceiversConfigBadRequestCode is the HTTP code returned for type PostTestReceiversConfigBadRequest
const PostTestReceiversConfigBadRequestCode int = 400

/*
PostTestReceiversConfigBadRequest Bad request

swagger:response postTestReceiversConfigBadRequest
*/
type PostTestReceiversConfigBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostTestReceiversConfigBadRequest creates PostTestReceiversConfigBadRequest with default headers values
func NewPostTestReceiversConfigBadRequest() *PostTestReceiversConfigBadRequest {

	return &PostTestReceiversConfigBadRequest{}
}

// WithPayload adds the payload to the post test receivers config bad request response
func (o *PostTestReceiversConfigBadRequest) WithPayload(payload string) *PostTestReceiversConfigBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test receivers config bad request response
func (o *PostTestReceiversConfigBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestReceiversConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostTestReceiversConfigRequestTimeoutCode is the HTTP code returned for type PostTestReceiversConfigRequestTimeout
const PostTestReceiversConfigRequestTimeoutCode int = 408

/*
PostTestReceiversConfigRequestTimeout Request time out

swagger:response postTestReceiversConfigRequestTimeout
*/
type PostTestReceiversConfigRequestTimeout struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostTestReceiversConfigRequestTimeout creates PostTestReceiversConfigRequestTimeout with default headers values
func NewPostTestReceiversConfigRequestTimeout() *PostTestReceiversConfigRequestTimeout {

	return &PostTestReceiversConfigRequestTimeout{}
}

// WithPayload adds the payload to the post test receivers config request timeout response
func (o *PostTestReceiversConfigRequestTimeout) WithPayload(payload string) *PostTestReceiversConfigRequestTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test receivers config request timeout response
func (o *PostTestReceiversConfigRequestTimeout) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestReceiversConfigRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(408)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostTestReceiversConfigInternalServerErrorCode is the HTTP code returned for type PostTestReceiversConfigInternalServerError
const PostTestReceiversConfigInternalServerErrorCode int = 500

/*
PostTestReceiversConfigInternalServerError Internal server error

swagger:response postTestReceiversConfigInternalServerError
*/
type PostTestReceiversConfigInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostTestReceiversConfigInternalServerError creates PostTestReceiversConfigInternalServerError with default headers values
func NewPostTestReceiversConfigInternalServerError() *PostTestReceiversConfigInternalServerError {

	return &PostTestReceiversConfigInternalServerError{}
}

// WithPayload adds the payload to the post test receivers config internal server error response
func (o *PostTestReceiversConfigInternalServerError) WithPayload(payload string) *PostTestReceiversConfigInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test receivers config internal server error response
func (o *PostTestReceiversConfigInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestReceiversConfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
