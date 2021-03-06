/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// APICallError struct for APICallError
type APICallError struct {
	ErrorMessage *APICallErrorMessage `json:"errorMessage,omitempty"`
	HttpStatus *int32 `json:"httpStatus,omitempty"`
}

// NewAPICallError instantiates a new APICallError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPICallError() *APICallError {
	this := APICallError{}
	return &this
}

// NewAPICallErrorWithDefaults instantiates a new APICallError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPICallErrorWithDefaults() *APICallError {
	this := APICallError{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *APICallError) GetErrorMessage() APICallErrorMessage {
	if o == nil || o.ErrorMessage == nil {
		var ret APICallErrorMessage
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APICallError) GetErrorMessageOk() (*APICallErrorMessage, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *APICallError) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given APICallErrorMessage and assigns it to the ErrorMessage field.
func (o *APICallError) SetErrorMessage(v APICallErrorMessage) {
	o.ErrorMessage = &v
}

// GetHttpStatus returns the HttpStatus field value if set, zero value otherwise.
func (o *APICallError) GetHttpStatus() int32 {
	if o == nil || o.HttpStatus == nil {
		var ret int32
		return ret
	}
	return *o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APICallError) GetHttpStatusOk() (*int32, bool) {
	if o == nil || o.HttpStatus == nil {
		return nil, false
	}
	return o.HttpStatus, true
}

// HasHttpStatus returns a boolean if a field has been set.
func (o *APICallError) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// SetHttpStatus gets a reference to the given int32 and assigns it to the HttpStatus field.
func (o *APICallError) SetHttpStatus(v int32) {
	o.HttpStatus = &v
}

func (o APICallError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}
	return json.Marshal(toSerialize)
}

type NullableAPICallError struct {
	value *APICallError
	isSet bool
}

func (v NullableAPICallError) Get() *APICallError {
	return v.value
}

func (v *NullableAPICallError) Set(val *APICallError) {
	v.value = val
	v.isSet = true
}

func (v NullableAPICallError) IsSet() bool {
	return v.isSet
}

func (v *NullableAPICallError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPICallError(val *APICallError) *NullableAPICallError {
	return &NullableAPICallError{value: val, isSet: true}
}

func (v NullableAPICallError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPICallError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


