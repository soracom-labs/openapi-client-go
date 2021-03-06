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

// ListRolesResponse struct for ListRolesResponse
type ListRolesResponse struct {
	CreateDateTime *int64 `json:"createDateTime,omitempty"`
	Description *string `json:"description,omitempty"`
	RoleId *string `json:"roleId,omitempty"`
	UpdateDateTime *int64 `json:"updateDateTime,omitempty"`
}

// NewListRolesResponse instantiates a new ListRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesResponse() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// NewListRolesResponseWithDefaults instantiates a new ListRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesResponseWithDefaults() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *ListRolesResponse) GetCreateDateTime() int64 {
	if o == nil || o.CreateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetCreateDateTimeOk() (*int64, bool) {
	if o == nil || o.CreateDateTime == nil {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *ListRolesResponse) HasCreateDateTime() bool {
	if o != nil && o.CreateDateTime != nil {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given int64 and assigns it to the CreateDateTime field.
func (o *ListRolesResponse) SetCreateDateTime(v int64) {
	o.CreateDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListRolesResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListRolesResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListRolesResponse) SetDescription(v string) {
	o.Description = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ListRolesResponse) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ListRolesResponse) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ListRolesResponse) SetRoleId(v string) {
	o.RoleId = &v
}

// GetUpdateDateTime returns the UpdateDateTime field value if set, zero value otherwise.
func (o *ListRolesResponse) GetUpdateDateTime() int64 {
	if o == nil || o.UpdateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDateTime
}

// GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetUpdateDateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateDateTime == nil {
		return nil, false
	}
	return o.UpdateDateTime, true
}

// HasUpdateDateTime returns a boolean if a field has been set.
func (o *ListRolesResponse) HasUpdateDateTime() bool {
	if o != nil && o.UpdateDateTime != nil {
		return true
	}

	return false
}

// SetUpdateDateTime gets a reference to the given int64 and assigns it to the UpdateDateTime field.
func (o *ListRolesResponse) SetUpdateDateTime(v int64) {
	o.UpdateDateTime = &v
}

func (o ListRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateDateTime != nil {
		toSerialize["createDateTime"] = o.CreateDateTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RoleId != nil {
		toSerialize["roleId"] = o.RoleId
	}
	if o.UpdateDateTime != nil {
		toSerialize["updateDateTime"] = o.UpdateDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableListRolesResponse struct {
	value *ListRolesResponse
	isSet bool
}

func (v NullableListRolesResponse) Get() *ListRolesResponse {
	return v.value
}

func (v *NullableListRolesResponse) Set(val *ListRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesResponse(val *ListRolesResponse) *NullableListRolesResponse {
	return &NullableListRolesResponse{value: val, isSet: true}
}

func (v NullableListRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


