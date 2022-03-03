/*
SORACOM SANDBOX API

SORACOM SANDBOX API v1

API version: 20160218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sandbox

import (
	"encoding/json"
)

// SandboxAuthResponse struct for SandboxAuthResponse
type SandboxAuthResponse struct {
	ApiKey *string `json:"apiKey,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	Token *string `json:"token,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// NewSandboxAuthResponse instantiates a new SandboxAuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxAuthResponse() *SandboxAuthResponse {
	this := SandboxAuthResponse{}
	return &this
}

// NewSandboxAuthResponseWithDefaults instantiates a new SandboxAuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxAuthResponseWithDefaults() *SandboxAuthResponse {
	this := SandboxAuthResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SandboxAuthResponse) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxAuthResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SandboxAuthResponse) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SandboxAuthResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *SandboxAuthResponse) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxAuthResponse) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *SandboxAuthResponse) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *SandboxAuthResponse) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SandboxAuthResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxAuthResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SandboxAuthResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SandboxAuthResponse) SetToken(v string) {
	o.Token = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SandboxAuthResponse) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxAuthResponse) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SandboxAuthResponse) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SandboxAuthResponse) SetUserName(v string) {
	o.UserName = &v
}

func (o SandboxAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxAuthResponse struct {
	value *SandboxAuthResponse
	isSet bool
}

func (v NullableSandboxAuthResponse) Get() *SandboxAuthResponse {
	return v.value
}

func (v *NullableSandboxAuthResponse) Set(val *SandboxAuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxAuthResponse(val *SandboxAuthResponse) *NullableSandboxAuthResponse {
	return &NullableSandboxAuthResponse{value: val, isSet: true}
}

func (v NullableSandboxAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

