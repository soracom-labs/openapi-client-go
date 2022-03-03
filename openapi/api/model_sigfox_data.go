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

// SigfoxData struct for SigfoxData
type SigfoxData struct {
	Data *string `json:"data,omitempty"`
}

// NewSigfoxData instantiates a new SigfoxData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigfoxData() *SigfoxData {
	this := SigfoxData{}
	return &this
}

// NewSigfoxDataWithDefaults instantiates a new SigfoxData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigfoxDataWithDefaults() *SigfoxData {
	this := SigfoxData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SigfoxData) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxData) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SigfoxData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *SigfoxData) SetData(v string) {
	o.Data = &v
}

func (o SigfoxData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSigfoxData struct {
	value *SigfoxData
	isSet bool
}

func (v NullableSigfoxData) Get() *SigfoxData {
	return v.value
}

func (v *NullableSigfoxData) Set(val *SigfoxData) {
	v.value = val
	v.isSet = true
}

func (v NullableSigfoxData) IsSet() bool {
	return v.isSet
}

func (v *NullableSigfoxData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigfoxData(val *SigfoxData) *NullableSigfoxData {
	return &NullableSigfoxData{value: val, isSet: true}
}

func (v NullableSigfoxData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigfoxData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


