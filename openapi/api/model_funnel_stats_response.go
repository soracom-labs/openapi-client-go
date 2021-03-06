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

// FunnelStatsResponse struct for FunnelStatsResponse
type FunnelStatsResponse struct {
	FunnelStatsMap *map[string]SoracomFunnelStats `json:"funnelStatsMap,omitempty"`
	Unixtime *int64 `json:"unixtime,omitempty"`
}

// NewFunnelStatsResponse instantiates a new FunnelStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunnelStatsResponse() *FunnelStatsResponse {
	this := FunnelStatsResponse{}
	return &this
}

// NewFunnelStatsResponseWithDefaults instantiates a new FunnelStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunnelStatsResponseWithDefaults() *FunnelStatsResponse {
	this := FunnelStatsResponse{}
	return &this
}

// GetFunnelStatsMap returns the FunnelStatsMap field value if set, zero value otherwise.
func (o *FunnelStatsResponse) GetFunnelStatsMap() map[string]SoracomFunnelStats {
	if o == nil || o.FunnelStatsMap == nil {
		var ret map[string]SoracomFunnelStats
		return ret
	}
	return *o.FunnelStatsMap
}

// GetFunnelStatsMapOk returns a tuple with the FunnelStatsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelStatsResponse) GetFunnelStatsMapOk() (*map[string]SoracomFunnelStats, bool) {
	if o == nil || o.FunnelStatsMap == nil {
		return nil, false
	}
	return o.FunnelStatsMap, true
}

// HasFunnelStatsMap returns a boolean if a field has been set.
func (o *FunnelStatsResponse) HasFunnelStatsMap() bool {
	if o != nil && o.FunnelStatsMap != nil {
		return true
	}

	return false
}

// SetFunnelStatsMap gets a reference to the given map[string]SoracomFunnelStats and assigns it to the FunnelStatsMap field.
func (o *FunnelStatsResponse) SetFunnelStatsMap(v map[string]SoracomFunnelStats) {
	o.FunnelStatsMap = &v
}

// GetUnixtime returns the Unixtime field value if set, zero value otherwise.
func (o *FunnelStatsResponse) GetUnixtime() int64 {
	if o == nil || o.Unixtime == nil {
		var ret int64
		return ret
	}
	return *o.Unixtime
}

// GetUnixtimeOk returns a tuple with the Unixtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelStatsResponse) GetUnixtimeOk() (*int64, bool) {
	if o == nil || o.Unixtime == nil {
		return nil, false
	}
	return o.Unixtime, true
}

// HasUnixtime returns a boolean if a field has been set.
func (o *FunnelStatsResponse) HasUnixtime() bool {
	if o != nil && o.Unixtime != nil {
		return true
	}

	return false
}

// SetUnixtime gets a reference to the given int64 and assigns it to the Unixtime field.
func (o *FunnelStatsResponse) SetUnixtime(v int64) {
	o.Unixtime = &v
}

func (o FunnelStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FunnelStatsMap != nil {
		toSerialize["funnelStatsMap"] = o.FunnelStatsMap
	}
	if o.Unixtime != nil {
		toSerialize["unixtime"] = o.Unixtime
	}
	return json.Marshal(toSerialize)
}

type NullableFunnelStatsResponse struct {
	value *FunnelStatsResponse
	isSet bool
}

func (v NullableFunnelStatsResponse) Get() *FunnelStatsResponse {
	return v.value
}

func (v *NullableFunnelStatsResponse) Set(val *FunnelStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelStatsResponse(val *FunnelStatsResponse) *NullableFunnelStatsResponse {
	return &NullableFunnelStatsResponse{value: val, isSet: true}
}

func (v NullableFunnelStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


