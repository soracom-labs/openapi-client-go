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

// GadgetRegistrationRequest struct for GadgetRegistrationRequest
type GadgetRegistrationRequest struct {
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewGadgetRegistrationRequest instantiates a new GadgetRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGadgetRegistrationRequest() *GadgetRegistrationRequest {
	this := GadgetRegistrationRequest{}
	return &this
}

// NewGadgetRegistrationRequestWithDefaults instantiates a new GadgetRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGadgetRegistrationRequestWithDefaults() *GadgetRegistrationRequest {
	this := GadgetRegistrationRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GadgetRegistrationRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GadgetRegistrationRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GadgetRegistrationRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *GadgetRegistrationRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o GadgetRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableGadgetRegistrationRequest struct {
	value *GadgetRegistrationRequest
	isSet bool
}

func (v NullableGadgetRegistrationRequest) Get() *GadgetRegistrationRequest {
	return v.value
}

func (v *NullableGadgetRegistrationRequest) Set(val *GadgetRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGadgetRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGadgetRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGadgetRegistrationRequest(val *GadgetRegistrationRequest) *NullableGadgetRegistrationRequest {
	return &NullableGadgetRegistrationRequest{value: val, isSet: true}
}

func (v NullableGadgetRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGadgetRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


