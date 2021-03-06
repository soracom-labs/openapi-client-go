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

// MFAIssueRevokingTokenRequest struct for MFAIssueRevokingTokenRequest
type MFAIssueRevokingTokenRequest struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewMFAIssueRevokingTokenRequest instantiates a new MFAIssueRevokingTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAIssueRevokingTokenRequest() *MFAIssueRevokingTokenRequest {
	this := MFAIssueRevokingTokenRequest{}
	return &this
}

// NewMFAIssueRevokingTokenRequestWithDefaults instantiates a new MFAIssueRevokingTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAIssueRevokingTokenRequestWithDefaults() *MFAIssueRevokingTokenRequest {
	this := MFAIssueRevokingTokenRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MFAIssueRevokingTokenRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAIssueRevokingTokenRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MFAIssueRevokingTokenRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MFAIssueRevokingTokenRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MFAIssueRevokingTokenRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAIssueRevokingTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MFAIssueRevokingTokenRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MFAIssueRevokingTokenRequest) SetPassword(v string) {
	o.Password = &v
}

func (o MFAIssueRevokingTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableMFAIssueRevokingTokenRequest struct {
	value *MFAIssueRevokingTokenRequest
	isSet bool
}

func (v NullableMFAIssueRevokingTokenRequest) Get() *MFAIssueRevokingTokenRequest {
	return v.value
}

func (v *NullableMFAIssueRevokingTokenRequest) Set(val *MFAIssueRevokingTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAIssueRevokingTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAIssueRevokingTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAIssueRevokingTokenRequest(val *MFAIssueRevokingTokenRequest) *NullableMFAIssueRevokingTokenRequest {
	return &NullableMFAIssueRevokingTokenRequest{value: val, isSet: true}
}

func (v NullableMFAIssueRevokingTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAIssueRevokingTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


