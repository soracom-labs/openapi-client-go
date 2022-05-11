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

// LocationRegistrationStatusForNetwork struct for LocationRegistrationStatusForNetwork
type LocationRegistrationStatusForNetwork struct {
	// Timestamp of the last successful location registration (Unix time, in milliseconds)
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`
	// PLMN ID of the visited network.
	Vplmn *string `json:"vplmn,omitempty"`
}

// NewLocationRegistrationStatusForNetwork instantiates a new LocationRegistrationStatusForNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRegistrationStatusForNetwork() *LocationRegistrationStatusForNetwork {
	this := LocationRegistrationStatusForNetwork{}
	return &this
}

// NewLocationRegistrationStatusForNetworkWithDefaults instantiates a new LocationRegistrationStatusForNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRegistrationStatusForNetworkWithDefaults() *LocationRegistrationStatusForNetwork {
	this := LocationRegistrationStatusForNetwork{}
	return &this
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *LocationRegistrationStatusForNetwork) GetLastModifiedTime() int64 {
	if o == nil || o.LastModifiedTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegistrationStatusForNetwork) GetLastModifiedTimeOk() (*int64, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *LocationRegistrationStatusForNetwork) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given int64 and assigns it to the LastModifiedTime field.
func (o *LocationRegistrationStatusForNetwork) SetLastModifiedTime(v int64) {
	o.LastModifiedTime = &v
}

// GetVplmn returns the Vplmn field value if set, zero value otherwise.
func (o *LocationRegistrationStatusForNetwork) GetVplmn() string {
	if o == nil || o.Vplmn == nil {
		var ret string
		return ret
	}
	return *o.Vplmn
}

// GetVplmnOk returns a tuple with the Vplmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegistrationStatusForNetwork) GetVplmnOk() (*string, bool) {
	if o == nil || o.Vplmn == nil {
		return nil, false
	}
	return o.Vplmn, true
}

// HasVplmn returns a boolean if a field has been set.
func (o *LocationRegistrationStatusForNetwork) HasVplmn() bool {
	if o != nil && o.Vplmn != nil {
		return true
	}

	return false
}

// SetVplmn gets a reference to the given string and assigns it to the Vplmn field.
func (o *LocationRegistrationStatusForNetwork) SetVplmn(v string) {
	o.Vplmn = &v
}

func (o LocationRegistrationStatusForNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.Vplmn != nil {
		toSerialize["vplmn"] = o.Vplmn
	}
	return json.Marshal(toSerialize)
}

type NullableLocationRegistrationStatusForNetwork struct {
	value *LocationRegistrationStatusForNetwork
	isSet bool
}

func (v NullableLocationRegistrationStatusForNetwork) Get() *LocationRegistrationStatusForNetwork {
	return v.value
}

func (v *NullableLocationRegistrationStatusForNetwork) Set(val *LocationRegistrationStatusForNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRegistrationStatusForNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRegistrationStatusForNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRegistrationStatusForNetwork(val *LocationRegistrationStatusForNetwork) *NullableLocationRegistrationStatusForNetwork {
	return &NullableLocationRegistrationStatusForNetwork{value: val, isSet: true}
}

func (v NullableLocationRegistrationStatusForNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRegistrationStatusForNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


