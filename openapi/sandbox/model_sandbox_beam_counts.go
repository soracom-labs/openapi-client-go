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

// SandboxBeamCounts struct for SandboxBeamCounts
type SandboxBeamCounts struct {
	Count *int64 `json:"count,omitempty"`
}

// NewSandboxBeamCounts instantiates a new SandboxBeamCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBeamCounts() *SandboxBeamCounts {
	this := SandboxBeamCounts{}
	return &this
}

// NewSandboxBeamCountsWithDefaults instantiates a new SandboxBeamCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBeamCountsWithDefaults() *SandboxBeamCounts {
	this := SandboxBeamCounts{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SandboxBeamCounts) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBeamCounts) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SandboxBeamCounts) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SandboxBeamCounts) SetCount(v int64) {
	o.Count = &v
}

func (o SandboxBeamCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxBeamCounts struct {
	value *SandboxBeamCounts
	isSet bool
}

func (v NullableSandboxBeamCounts) Get() *SandboxBeamCounts {
	return v.value
}

func (v *NullableSandboxBeamCounts) Set(val *SandboxBeamCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBeamCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBeamCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBeamCounts(val *SandboxBeamCounts) *NullableSandboxBeamCounts {
	return &NullableSandboxBeamCounts{value: val, isSet: true}
}

func (v NullableSandboxBeamCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBeamCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


