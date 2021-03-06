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

// DataTrafficStats struct for DataTrafficStats
type DataTrafficStats struct {
	DownloadByteSizeTotal *int64 `json:"downloadByteSizeTotal,omitempty"`
	DownloadPacketSizeTotal *int64 `json:"downloadPacketSizeTotal,omitempty"`
	UploadByteSizeTotal *int64 `json:"uploadByteSizeTotal,omitempty"`
	UploadPacketSizeTotal *int64 `json:"uploadPacketSizeTotal,omitempty"`
}

// NewDataTrafficStats instantiates a new DataTrafficStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTrafficStats() *DataTrafficStats {
	this := DataTrafficStats{}
	return &this
}

// NewDataTrafficStatsWithDefaults instantiates a new DataTrafficStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTrafficStatsWithDefaults() *DataTrafficStats {
	this := DataTrafficStats{}
	return &this
}

// GetDownloadByteSizeTotal returns the DownloadByteSizeTotal field value if set, zero value otherwise.
func (o *DataTrafficStats) GetDownloadByteSizeTotal() int64 {
	if o == nil || o.DownloadByteSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.DownloadByteSizeTotal
}

// GetDownloadByteSizeTotalOk returns a tuple with the DownloadByteSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficStats) GetDownloadByteSizeTotalOk() (*int64, bool) {
	if o == nil || o.DownloadByteSizeTotal == nil {
		return nil, false
	}
	return o.DownloadByteSizeTotal, true
}

// HasDownloadByteSizeTotal returns a boolean if a field has been set.
func (o *DataTrafficStats) HasDownloadByteSizeTotal() bool {
	if o != nil && o.DownloadByteSizeTotal != nil {
		return true
	}

	return false
}

// SetDownloadByteSizeTotal gets a reference to the given int64 and assigns it to the DownloadByteSizeTotal field.
func (o *DataTrafficStats) SetDownloadByteSizeTotal(v int64) {
	o.DownloadByteSizeTotal = &v
}

// GetDownloadPacketSizeTotal returns the DownloadPacketSizeTotal field value if set, zero value otherwise.
func (o *DataTrafficStats) GetDownloadPacketSizeTotal() int64 {
	if o == nil || o.DownloadPacketSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.DownloadPacketSizeTotal
}

// GetDownloadPacketSizeTotalOk returns a tuple with the DownloadPacketSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficStats) GetDownloadPacketSizeTotalOk() (*int64, bool) {
	if o == nil || o.DownloadPacketSizeTotal == nil {
		return nil, false
	}
	return o.DownloadPacketSizeTotal, true
}

// HasDownloadPacketSizeTotal returns a boolean if a field has been set.
func (o *DataTrafficStats) HasDownloadPacketSizeTotal() bool {
	if o != nil && o.DownloadPacketSizeTotal != nil {
		return true
	}

	return false
}

// SetDownloadPacketSizeTotal gets a reference to the given int64 and assigns it to the DownloadPacketSizeTotal field.
func (o *DataTrafficStats) SetDownloadPacketSizeTotal(v int64) {
	o.DownloadPacketSizeTotal = &v
}

// GetUploadByteSizeTotal returns the UploadByteSizeTotal field value if set, zero value otherwise.
func (o *DataTrafficStats) GetUploadByteSizeTotal() int64 {
	if o == nil || o.UploadByteSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.UploadByteSizeTotal
}

// GetUploadByteSizeTotalOk returns a tuple with the UploadByteSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficStats) GetUploadByteSizeTotalOk() (*int64, bool) {
	if o == nil || o.UploadByteSizeTotal == nil {
		return nil, false
	}
	return o.UploadByteSizeTotal, true
}

// HasUploadByteSizeTotal returns a boolean if a field has been set.
func (o *DataTrafficStats) HasUploadByteSizeTotal() bool {
	if o != nil && o.UploadByteSizeTotal != nil {
		return true
	}

	return false
}

// SetUploadByteSizeTotal gets a reference to the given int64 and assigns it to the UploadByteSizeTotal field.
func (o *DataTrafficStats) SetUploadByteSizeTotal(v int64) {
	o.UploadByteSizeTotal = &v
}

// GetUploadPacketSizeTotal returns the UploadPacketSizeTotal field value if set, zero value otherwise.
func (o *DataTrafficStats) GetUploadPacketSizeTotal() int64 {
	if o == nil || o.UploadPacketSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.UploadPacketSizeTotal
}

// GetUploadPacketSizeTotalOk returns a tuple with the UploadPacketSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficStats) GetUploadPacketSizeTotalOk() (*int64, bool) {
	if o == nil || o.UploadPacketSizeTotal == nil {
		return nil, false
	}
	return o.UploadPacketSizeTotal, true
}

// HasUploadPacketSizeTotal returns a boolean if a field has been set.
func (o *DataTrafficStats) HasUploadPacketSizeTotal() bool {
	if o != nil && o.UploadPacketSizeTotal != nil {
		return true
	}

	return false
}

// SetUploadPacketSizeTotal gets a reference to the given int64 and assigns it to the UploadPacketSizeTotal field.
func (o *DataTrafficStats) SetUploadPacketSizeTotal(v int64) {
	o.UploadPacketSizeTotal = &v
}

func (o DataTrafficStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadByteSizeTotal != nil {
		toSerialize["downloadByteSizeTotal"] = o.DownloadByteSizeTotal
	}
	if o.DownloadPacketSizeTotal != nil {
		toSerialize["downloadPacketSizeTotal"] = o.DownloadPacketSizeTotal
	}
	if o.UploadByteSizeTotal != nil {
		toSerialize["uploadByteSizeTotal"] = o.UploadByteSizeTotal
	}
	if o.UploadPacketSizeTotal != nil {
		toSerialize["uploadPacketSizeTotal"] = o.UploadPacketSizeTotal
	}
	return json.Marshal(toSerialize)
}

type NullableDataTrafficStats struct {
	value *DataTrafficStats
	isSet bool
}

func (v NullableDataTrafficStats) Get() *DataTrafficStats {
	return v.value
}

func (v *NullableDataTrafficStats) Set(val *DataTrafficStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTrafficStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTrafficStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTrafficStats(val *DataTrafficStats) *NullableDataTrafficStats {
	return &NullableDataTrafficStats{value: val, isSet: true}
}

func (v NullableDataTrafficStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTrafficStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


