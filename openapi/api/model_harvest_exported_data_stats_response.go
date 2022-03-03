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

// HarvestExportedDataStatsResponse struct for HarvestExportedDataStatsResponse
type HarvestExportedDataStatsResponse struct {
	// exportedBytes
	ExportedBytes *int32 `json:"exportedBytes,omitempty"`
	// yearMonth
	YearMonth *string `json:"yearMonth,omitempty"`
}

// NewHarvestExportedDataStatsResponse instantiates a new HarvestExportedDataStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvestExportedDataStatsResponse() *HarvestExportedDataStatsResponse {
	this := HarvestExportedDataStatsResponse{}
	return &this
}

// NewHarvestExportedDataStatsResponseWithDefaults instantiates a new HarvestExportedDataStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvestExportedDataStatsResponseWithDefaults() *HarvestExportedDataStatsResponse {
	this := HarvestExportedDataStatsResponse{}
	return &this
}

// GetExportedBytes returns the ExportedBytes field value if set, zero value otherwise.
func (o *HarvestExportedDataStatsResponse) GetExportedBytes() int32 {
	if o == nil || o.ExportedBytes == nil {
		var ret int32
		return ret
	}
	return *o.ExportedBytes
}

// GetExportedBytesOk returns a tuple with the ExportedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvestExportedDataStatsResponse) GetExportedBytesOk() (*int32, bool) {
	if o == nil || o.ExportedBytes == nil {
		return nil, false
	}
	return o.ExportedBytes, true
}

// HasExportedBytes returns a boolean if a field has been set.
func (o *HarvestExportedDataStatsResponse) HasExportedBytes() bool {
	if o != nil && o.ExportedBytes != nil {
		return true
	}

	return false
}

// SetExportedBytes gets a reference to the given int32 and assigns it to the ExportedBytes field.
func (o *HarvestExportedDataStatsResponse) SetExportedBytes(v int32) {
	o.ExportedBytes = &v
}

// GetYearMonth returns the YearMonth field value if set, zero value otherwise.
func (o *HarvestExportedDataStatsResponse) GetYearMonth() string {
	if o == nil || o.YearMonth == nil {
		var ret string
		return ret
	}
	return *o.YearMonth
}

// GetYearMonthOk returns a tuple with the YearMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvestExportedDataStatsResponse) GetYearMonthOk() (*string, bool) {
	if o == nil || o.YearMonth == nil {
		return nil, false
	}
	return o.YearMonth, true
}

// HasYearMonth returns a boolean if a field has been set.
func (o *HarvestExportedDataStatsResponse) HasYearMonth() bool {
	if o != nil && o.YearMonth != nil {
		return true
	}

	return false
}

// SetYearMonth gets a reference to the given string and assigns it to the YearMonth field.
func (o *HarvestExportedDataStatsResponse) SetYearMonth(v string) {
	o.YearMonth = &v
}

func (o HarvestExportedDataStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExportedBytes != nil {
		toSerialize["exportedBytes"] = o.ExportedBytes
	}
	if o.YearMonth != nil {
		toSerialize["yearMonth"] = o.YearMonth
	}
	return json.Marshal(toSerialize)
}

type NullableHarvestExportedDataStatsResponse struct {
	value *HarvestExportedDataStatsResponse
	isSet bool
}

func (v NullableHarvestExportedDataStatsResponse) Get() *HarvestExportedDataStatsResponse {
	return v.value
}

func (v *NullableHarvestExportedDataStatsResponse) Set(val *HarvestExportedDataStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvestExportedDataStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvestExportedDataStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvestExportedDataStatsResponse(val *HarvestExportedDataStatsResponse) *NullableHarvestExportedDataStatsResponse {
	return &NullableHarvestExportedDataStatsResponse{value: val, isSet: true}
}

func (v NullableHarvestExportedDataStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvestExportedDataStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


