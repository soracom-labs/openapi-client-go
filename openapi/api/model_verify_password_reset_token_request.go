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

// VerifyPasswordResetTokenRequest struct for VerifyPasswordResetTokenRequest
type VerifyPasswordResetTokenRequest struct {
	Password string `json:"password"`
	Token string `json:"token"`
}

// NewVerifyPasswordResetTokenRequest instantiates a new VerifyPasswordResetTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPasswordResetTokenRequest(password string, token string) *VerifyPasswordResetTokenRequest {
	this := VerifyPasswordResetTokenRequest{}
	this.Password = password
	this.Token = token
	return &this
}

// NewVerifyPasswordResetTokenRequestWithDefaults instantiates a new VerifyPasswordResetTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPasswordResetTokenRequestWithDefaults() *VerifyPasswordResetTokenRequest {
	this := VerifyPasswordResetTokenRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *VerifyPasswordResetTokenRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VerifyPasswordResetTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VerifyPasswordResetTokenRequest) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *VerifyPasswordResetTokenRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *VerifyPasswordResetTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *VerifyPasswordResetTokenRequest) SetToken(v string) {
	o.Token = v
}

func (o VerifyPasswordResetTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyPasswordResetTokenRequest struct {
	value *VerifyPasswordResetTokenRequest
	isSet bool
}

func (v NullableVerifyPasswordResetTokenRequest) Get() *VerifyPasswordResetTokenRequest {
	return v.value
}

func (v *NullableVerifyPasswordResetTokenRequest) Set(val *VerifyPasswordResetTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPasswordResetTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPasswordResetTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPasswordResetTokenRequest(val *VerifyPasswordResetTokenRequest) *NullableVerifyPasswordResetTokenRequest {
	return &NullableVerifyPasswordResetTokenRequest{value: val, isSet: true}
}

func (v NullableVerifyPasswordResetTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPasswordResetTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


