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

// MFAAuthenticationRequest struct for MFAAuthenticationRequest
type MFAAuthenticationRequest struct {
	MfaOTPCode *string `json:"mfaOTPCode,omitempty"`
}

// NewMFAAuthenticationRequest instantiates a new MFAAuthenticationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAAuthenticationRequest() *MFAAuthenticationRequest {
	this := MFAAuthenticationRequest{}
	return &this
}

// NewMFAAuthenticationRequestWithDefaults instantiates a new MFAAuthenticationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAAuthenticationRequestWithDefaults() *MFAAuthenticationRequest {
	this := MFAAuthenticationRequest{}
	return &this
}

// GetMfaOTPCode returns the MfaOTPCode field value if set, zero value otherwise.
func (o *MFAAuthenticationRequest) GetMfaOTPCode() string {
	if o == nil || o.MfaOTPCode == nil {
		var ret string
		return ret
	}
	return *o.MfaOTPCode
}

// GetMfaOTPCodeOk returns a tuple with the MfaOTPCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAAuthenticationRequest) GetMfaOTPCodeOk() (*string, bool) {
	if o == nil || o.MfaOTPCode == nil {
		return nil, false
	}
	return o.MfaOTPCode, true
}

// HasMfaOTPCode returns a boolean if a field has been set.
func (o *MFAAuthenticationRequest) HasMfaOTPCode() bool {
	if o != nil && o.MfaOTPCode != nil {
		return true
	}

	return false
}

// SetMfaOTPCode gets a reference to the given string and assigns it to the MfaOTPCode field.
func (o *MFAAuthenticationRequest) SetMfaOTPCode(v string) {
	o.MfaOTPCode = &v
}

func (o MFAAuthenticationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MfaOTPCode != nil {
		toSerialize["mfaOTPCode"] = o.MfaOTPCode
	}
	return json.Marshal(toSerialize)
}

type NullableMFAAuthenticationRequest struct {
	value *MFAAuthenticationRequest
	isSet bool
}

func (v NullableMFAAuthenticationRequest) Get() *MFAAuthenticationRequest {
	return v.value
}

func (v *NullableMFAAuthenticationRequest) Set(val *MFAAuthenticationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAAuthenticationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAAuthenticationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAAuthenticationRequest(val *MFAAuthenticationRequest) *NullableMFAAuthenticationRequest {
	return &NullableMFAAuthenticationRequest{value: val, isSet: true}
}

func (v NullableMFAAuthenticationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAAuthenticationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


