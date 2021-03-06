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

// AttachRoleRequest struct for AttachRoleRequest
type AttachRoleRequest struct {
	RoleId *string `json:"roleId,omitempty"`
}

// NewAttachRoleRequest instantiates a new AttachRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachRoleRequest() *AttachRoleRequest {
	this := AttachRoleRequest{}
	return &this
}

// NewAttachRoleRequestWithDefaults instantiates a new AttachRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachRoleRequestWithDefaults() *AttachRoleRequest {
	this := AttachRoleRequest{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *AttachRoleRequest) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachRoleRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *AttachRoleRequest) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *AttachRoleRequest) SetRoleId(v string) {
	o.RoleId = &v
}

func (o AttachRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleId != nil {
		toSerialize["roleId"] = o.RoleId
	}
	return json.Marshal(toSerialize)
}

type NullableAttachRoleRequest struct {
	value *AttachRoleRequest
	isSet bool
}

func (v NullableAttachRoleRequest) Get() *AttachRoleRequest {
	return v.value
}

func (v *NullableAttachRoleRequest) Set(val *AttachRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachRoleRequest(val *AttachRoleRequest) *NullableAttachRoleRequest {
	return &NullableAttachRoleRequest{value: val, isSet: true}
}

func (v NullableAttachRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


