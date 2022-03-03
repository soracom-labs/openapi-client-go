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

// EstimatedVolumeDiscountModel struct for EstimatedVolumeDiscountModel
type EstimatedVolumeDiscountModel struct {
	// Currency
	Currency *string `json:"currency,omitempty"`
	// Order ID
	OrderId *string `json:"orderId,omitempty"`
	// Tax amount
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Total amount
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	VolumeDiscount *VolumeDiscountModel `json:"volumeDiscount,omitempty"`
}

// NewEstimatedVolumeDiscountModel instantiates a new EstimatedVolumeDiscountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedVolumeDiscountModel() *EstimatedVolumeDiscountModel {
	this := EstimatedVolumeDiscountModel{}
	return &this
}

// NewEstimatedVolumeDiscountModelWithDefaults instantiates a new EstimatedVolumeDiscountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedVolumeDiscountModelWithDefaults() *EstimatedVolumeDiscountModel {
	this := EstimatedVolumeDiscountModel{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *EstimatedVolumeDiscountModel) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedVolumeDiscountModel) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *EstimatedVolumeDiscountModel) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *EstimatedVolumeDiscountModel) SetCurrency(v string) {
	o.Currency = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EstimatedVolumeDiscountModel) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedVolumeDiscountModel) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EstimatedVolumeDiscountModel) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EstimatedVolumeDiscountModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *EstimatedVolumeDiscountModel) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedVolumeDiscountModel) GetTaxAmountOk() (*float64, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *EstimatedVolumeDiscountModel) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *EstimatedVolumeDiscountModel) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *EstimatedVolumeDiscountModel) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedVolumeDiscountModel) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *EstimatedVolumeDiscountModel) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *EstimatedVolumeDiscountModel) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetVolumeDiscount returns the VolumeDiscount field value if set, zero value otherwise.
func (o *EstimatedVolumeDiscountModel) GetVolumeDiscount() VolumeDiscountModel {
	if o == nil || o.VolumeDiscount == nil {
		var ret VolumeDiscountModel
		return ret
	}
	return *o.VolumeDiscount
}

// GetVolumeDiscountOk returns a tuple with the VolumeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedVolumeDiscountModel) GetVolumeDiscountOk() (*VolumeDiscountModel, bool) {
	if o == nil || o.VolumeDiscount == nil {
		return nil, false
	}
	return o.VolumeDiscount, true
}

// HasVolumeDiscount returns a boolean if a field has been set.
func (o *EstimatedVolumeDiscountModel) HasVolumeDiscount() bool {
	if o != nil && o.VolumeDiscount != nil {
		return true
	}

	return false
}

// SetVolumeDiscount gets a reference to the given VolumeDiscountModel and assigns it to the VolumeDiscount field.
func (o *EstimatedVolumeDiscountModel) SetVolumeDiscount(v VolumeDiscountModel) {
	o.VolumeDiscount = &v
}

func (o EstimatedVolumeDiscountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.TaxAmount != nil {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.VolumeDiscount != nil {
		toSerialize["volumeDiscount"] = o.VolumeDiscount
	}
	return json.Marshal(toSerialize)
}

type NullableEstimatedVolumeDiscountModel struct {
	value *EstimatedVolumeDiscountModel
	isSet bool
}

func (v NullableEstimatedVolumeDiscountModel) Get() *EstimatedVolumeDiscountModel {
	return v.value
}

func (v *NullableEstimatedVolumeDiscountModel) Set(val *EstimatedVolumeDiscountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedVolumeDiscountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedVolumeDiscountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedVolumeDiscountModel(val *EstimatedVolumeDiscountModel) *NullableEstimatedVolumeDiscountModel {
	return &NullableEstimatedVolumeDiscountModel{value: val, isSet: true}
}

func (v NullableEstimatedVolumeDiscountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedVolumeDiscountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


