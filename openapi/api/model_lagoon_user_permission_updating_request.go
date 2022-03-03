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

// LagoonUserPermissionUpdatingRequest struct for LagoonUserPermissionUpdatingRequest
type LagoonUserPermissionUpdatingRequest struct {
	// A role that represents the permission.
	Role *string `json:"role,omitempty"`
}

// NewLagoonUserPermissionUpdatingRequest instantiates a new LagoonUserPermissionUpdatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonUserPermissionUpdatingRequest() *LagoonUserPermissionUpdatingRequest {
	this := LagoonUserPermissionUpdatingRequest{}
	return &this
}

// NewLagoonUserPermissionUpdatingRequestWithDefaults instantiates a new LagoonUserPermissionUpdatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonUserPermissionUpdatingRequestWithDefaults() *LagoonUserPermissionUpdatingRequest {
	this := LagoonUserPermissionUpdatingRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *LagoonUserPermissionUpdatingRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUserPermissionUpdatingRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *LagoonUserPermissionUpdatingRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *LagoonUserPermissionUpdatingRequest) SetRole(v string) {
	o.Role = &v
}

func (o LagoonUserPermissionUpdatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonUserPermissionUpdatingRequest struct {
	value *LagoonUserPermissionUpdatingRequest
	isSet bool
}

func (v NullableLagoonUserPermissionUpdatingRequest) Get() *LagoonUserPermissionUpdatingRequest {
	return v.value
}

func (v *NullableLagoonUserPermissionUpdatingRequest) Set(val *LagoonUserPermissionUpdatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonUserPermissionUpdatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonUserPermissionUpdatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonUserPermissionUpdatingRequest(val *LagoonUserPermissionUpdatingRequest) *NullableLagoonUserPermissionUpdatingRequest {
	return &NullableLagoonUserPermissionUpdatingRequest{value: val, isSet: true}
}

func (v NullableLagoonUserPermissionUpdatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonUserPermissionUpdatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

