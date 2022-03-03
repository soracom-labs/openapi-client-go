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

// SandboxInsertBeamStatsRequest struct for SandboxInsertBeamStatsRequest
type SandboxInsertBeamStatsRequest struct {
	BeamStatsMap *SandboxInsertBeamStatsRequestBeamStatsMap `json:"beamStatsMap,omitempty"`
	// UNIX time (in milliseconds)
	Unixtime *int64 `json:"unixtime,omitempty"`
}

// NewSandboxInsertBeamStatsRequest instantiates a new SandboxInsertBeamStatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxInsertBeamStatsRequest() *SandboxInsertBeamStatsRequest {
	this := SandboxInsertBeamStatsRequest{}
	return &this
}

// NewSandboxInsertBeamStatsRequestWithDefaults instantiates a new SandboxInsertBeamStatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxInsertBeamStatsRequestWithDefaults() *SandboxInsertBeamStatsRequest {
	this := SandboxInsertBeamStatsRequest{}
	return &this
}

// GetBeamStatsMap returns the BeamStatsMap field value if set, zero value otherwise.
func (o *SandboxInsertBeamStatsRequest) GetBeamStatsMap() SandboxInsertBeamStatsRequestBeamStatsMap {
	if o == nil || o.BeamStatsMap == nil {
		var ret SandboxInsertBeamStatsRequestBeamStatsMap
		return ret
	}
	return *o.BeamStatsMap
}

// GetBeamStatsMapOk returns a tuple with the BeamStatsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxInsertBeamStatsRequest) GetBeamStatsMapOk() (*SandboxInsertBeamStatsRequestBeamStatsMap, bool) {
	if o == nil || o.BeamStatsMap == nil {
		return nil, false
	}
	return o.BeamStatsMap, true
}

// HasBeamStatsMap returns a boolean if a field has been set.
func (o *SandboxInsertBeamStatsRequest) HasBeamStatsMap() bool {
	if o != nil && o.BeamStatsMap != nil {
		return true
	}

	return false
}

// SetBeamStatsMap gets a reference to the given SandboxInsertBeamStatsRequestBeamStatsMap and assigns it to the BeamStatsMap field.
func (o *SandboxInsertBeamStatsRequest) SetBeamStatsMap(v SandboxInsertBeamStatsRequestBeamStatsMap) {
	o.BeamStatsMap = &v
}

// GetUnixtime returns the Unixtime field value if set, zero value otherwise.
func (o *SandboxInsertBeamStatsRequest) GetUnixtime() int64 {
	if o == nil || o.Unixtime == nil {
		var ret int64
		return ret
	}
	return *o.Unixtime
}

// GetUnixtimeOk returns a tuple with the Unixtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxInsertBeamStatsRequest) GetUnixtimeOk() (*int64, bool) {
	if o == nil || o.Unixtime == nil {
		return nil, false
	}
	return o.Unixtime, true
}

// HasUnixtime returns a boolean if a field has been set.
func (o *SandboxInsertBeamStatsRequest) HasUnixtime() bool {
	if o != nil && o.Unixtime != nil {
		return true
	}

	return false
}

// SetUnixtime gets a reference to the given int64 and assigns it to the Unixtime field.
func (o *SandboxInsertBeamStatsRequest) SetUnixtime(v int64) {
	o.Unixtime = &v
}

func (o SandboxInsertBeamStatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BeamStatsMap != nil {
		toSerialize["beamStatsMap"] = o.BeamStatsMap
	}
	if o.Unixtime != nil {
		toSerialize["unixtime"] = o.Unixtime
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxInsertBeamStatsRequest struct {
	value *SandboxInsertBeamStatsRequest
	isSet bool
}

func (v NullableSandboxInsertBeamStatsRequest) Get() *SandboxInsertBeamStatsRequest {
	return v.value
}

func (v *NullableSandboxInsertBeamStatsRequest) Set(val *SandboxInsertBeamStatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxInsertBeamStatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxInsertBeamStatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxInsertBeamStatsRequest(val *SandboxInsertBeamStatsRequest) *NullableSandboxInsertBeamStatsRequest {
	return &NullableSandboxInsertBeamStatsRequest{value: val, isSet: true}
}

func (v NullableSandboxInsertBeamStatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxInsertBeamStatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


