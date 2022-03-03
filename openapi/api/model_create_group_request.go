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

// CreateGroupRequest struct for CreateGroupRequest
type CreateGroupRequest struct {
	// An object which always contains at least one property \"name\" with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \"name\" property. If the subscriber/SIM does not have a name, an empty string \"\" is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateGroupRequest instantiates a new CreateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupRequest() *CreateGroupRequest {
	this := CreateGroupRequest{}
	return &this
}

// NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupRequestWithDefaults() *CreateGroupRequest {
	this := CreateGroupRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateGroupRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateGroupRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateGroupRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupRequest struct {
	value *CreateGroupRequest
	isSet bool
}

func (v NullableCreateGroupRequest) Get() *CreateGroupRequest {
	return v.value
}

func (v *NullableCreateGroupRequest) Set(val *CreateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupRequest(val *CreateGroupRequest) *NullableCreateGroupRequest {
	return &NullableCreateGroupRequest{value: val, isSet: true}
}

func (v NullableCreateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


