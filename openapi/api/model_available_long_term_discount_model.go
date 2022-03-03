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

// AvailableLongTermDiscountModel struct for AvailableLongTermDiscountModel
type AvailableLongTermDiscountModel struct {
	// Contract months
	ContractTermMonth *int32 `json:"contractTermMonth,omitempty"`
	// Unit price
	UnitPrice *float64 `json:"unitPrice,omitempty"`
	// Payment type
	VolumeDiscountPaymentType *string `json:"volumeDiscountPaymentType,omitempty"`
	// Long term discount type
	VolumeDiscountType *string `json:"volumeDiscountType,omitempty"`
}

// NewAvailableLongTermDiscountModel instantiates a new AvailableLongTermDiscountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableLongTermDiscountModel() *AvailableLongTermDiscountModel {
	this := AvailableLongTermDiscountModel{}
	var contractTermMonth int32 = 12
	this.ContractTermMonth = &contractTermMonth
	return &this
}

// NewAvailableLongTermDiscountModelWithDefaults instantiates a new AvailableLongTermDiscountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableLongTermDiscountModelWithDefaults() *AvailableLongTermDiscountModel {
	this := AvailableLongTermDiscountModel{}
	var contractTermMonth int32 = 12
	this.ContractTermMonth = &contractTermMonth
	return &this
}

// GetContractTermMonth returns the ContractTermMonth field value if set, zero value otherwise.
func (o *AvailableLongTermDiscountModel) GetContractTermMonth() int32 {
	if o == nil || o.ContractTermMonth == nil {
		var ret int32
		return ret
	}
	return *o.ContractTermMonth
}

// GetContractTermMonthOk returns a tuple with the ContractTermMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableLongTermDiscountModel) GetContractTermMonthOk() (*int32, bool) {
	if o == nil || o.ContractTermMonth == nil {
		return nil, false
	}
	return o.ContractTermMonth, true
}

// HasContractTermMonth returns a boolean if a field has been set.
func (o *AvailableLongTermDiscountModel) HasContractTermMonth() bool {
	if o != nil && o.ContractTermMonth != nil {
		return true
	}

	return false
}

// SetContractTermMonth gets a reference to the given int32 and assigns it to the ContractTermMonth field.
func (o *AvailableLongTermDiscountModel) SetContractTermMonth(v int32) {
	o.ContractTermMonth = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *AvailableLongTermDiscountModel) GetUnitPrice() float64 {
	if o == nil || o.UnitPrice == nil {
		var ret float64
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableLongTermDiscountModel) GetUnitPriceOk() (*float64, bool) {
	if o == nil || o.UnitPrice == nil {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *AvailableLongTermDiscountModel) HasUnitPrice() bool {
	if o != nil && o.UnitPrice != nil {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float64 and assigns it to the UnitPrice field.
func (o *AvailableLongTermDiscountModel) SetUnitPrice(v float64) {
	o.UnitPrice = &v
}

// GetVolumeDiscountPaymentType returns the VolumeDiscountPaymentType field value if set, zero value otherwise.
func (o *AvailableLongTermDiscountModel) GetVolumeDiscountPaymentType() string {
	if o == nil || o.VolumeDiscountPaymentType == nil {
		var ret string
		return ret
	}
	return *o.VolumeDiscountPaymentType
}

// GetVolumeDiscountPaymentTypeOk returns a tuple with the VolumeDiscountPaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableLongTermDiscountModel) GetVolumeDiscountPaymentTypeOk() (*string, bool) {
	if o == nil || o.VolumeDiscountPaymentType == nil {
		return nil, false
	}
	return o.VolumeDiscountPaymentType, true
}

// HasVolumeDiscountPaymentType returns a boolean if a field has been set.
func (o *AvailableLongTermDiscountModel) HasVolumeDiscountPaymentType() bool {
	if o != nil && o.VolumeDiscountPaymentType != nil {
		return true
	}

	return false
}

// SetVolumeDiscountPaymentType gets a reference to the given string and assigns it to the VolumeDiscountPaymentType field.
func (o *AvailableLongTermDiscountModel) SetVolumeDiscountPaymentType(v string) {
	o.VolumeDiscountPaymentType = &v
}

// GetVolumeDiscountType returns the VolumeDiscountType field value if set, zero value otherwise.
func (o *AvailableLongTermDiscountModel) GetVolumeDiscountType() string {
	if o == nil || o.VolumeDiscountType == nil {
		var ret string
		return ret
	}
	return *o.VolumeDiscountType
}

// GetVolumeDiscountTypeOk returns a tuple with the VolumeDiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableLongTermDiscountModel) GetVolumeDiscountTypeOk() (*string, bool) {
	if o == nil || o.VolumeDiscountType == nil {
		return nil, false
	}
	return o.VolumeDiscountType, true
}

// HasVolumeDiscountType returns a boolean if a field has been set.
func (o *AvailableLongTermDiscountModel) HasVolumeDiscountType() bool {
	if o != nil && o.VolumeDiscountType != nil {
		return true
	}

	return false
}

// SetVolumeDiscountType gets a reference to the given string and assigns it to the VolumeDiscountType field.
func (o *AvailableLongTermDiscountModel) SetVolumeDiscountType(v string) {
	o.VolumeDiscountType = &v
}

func (o AvailableLongTermDiscountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractTermMonth != nil {
		toSerialize["contractTermMonth"] = o.ContractTermMonth
	}
	if o.UnitPrice != nil {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if o.VolumeDiscountPaymentType != nil {
		toSerialize["volumeDiscountPaymentType"] = o.VolumeDiscountPaymentType
	}
	if o.VolumeDiscountType != nil {
		toSerialize["volumeDiscountType"] = o.VolumeDiscountType
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableLongTermDiscountModel struct {
	value *AvailableLongTermDiscountModel
	isSet bool
}

func (v NullableAvailableLongTermDiscountModel) Get() *AvailableLongTermDiscountModel {
	return v.value
}

func (v *NullableAvailableLongTermDiscountModel) Set(val *AvailableLongTermDiscountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableLongTermDiscountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableLongTermDiscountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableLongTermDiscountModel(val *AvailableLongTermDiscountModel) *NullableAvailableLongTermDiscountModel {
	return &NullableAvailableLongTermDiscountModel{value: val, isSet: true}
}

func (v NullableAvailableLongTermDiscountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableLongTermDiscountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


