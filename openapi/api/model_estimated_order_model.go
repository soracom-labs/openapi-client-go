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

// EstimatedOrderModel struct for EstimatedOrderModel
type EstimatedOrderModel struct {
	// Email address
	Email *string `json:"email,omitempty"`
	// Order ID
	OrderId *string `json:"orderId,omitempty"`
	// Order item list
	OrderItemList []EstimatedOrderItemModel `json:"orderItemList,omitempty"`
	ShippingAddress *ShippingAddressModel `json:"shippingAddress,omitempty"`
	// Shipping address ID
	ShippingAddressId *string `json:"shippingAddressId,omitempty"`
	// Shipping cost
	ShippingCost *float64 `json:"shippingCost,omitempty"`
	// Tax amount
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Total amount
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

// NewEstimatedOrderModel instantiates a new EstimatedOrderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedOrderModel() *EstimatedOrderModel {
	this := EstimatedOrderModel{}
	return &this
}

// NewEstimatedOrderModelWithDefaults instantiates a new EstimatedOrderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedOrderModelWithDefaults() *EstimatedOrderModel {
	this := EstimatedOrderModel{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EstimatedOrderModel) SetEmail(v string) {
	o.Email = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EstimatedOrderModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderItemList returns the OrderItemList field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetOrderItemList() []EstimatedOrderItemModel {
	if o == nil || o.OrderItemList == nil {
		var ret []EstimatedOrderItemModel
		return ret
	}
	return o.OrderItemList
}

// GetOrderItemListOk returns a tuple with the OrderItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetOrderItemListOk() ([]EstimatedOrderItemModel, bool) {
	if o == nil || o.OrderItemList == nil {
		return nil, false
	}
	return o.OrderItemList, true
}

// HasOrderItemList returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasOrderItemList() bool {
	if o != nil && o.OrderItemList != nil {
		return true
	}

	return false
}

// SetOrderItemList gets a reference to the given []EstimatedOrderItemModel and assigns it to the OrderItemList field.
func (o *EstimatedOrderModel) SetOrderItemList(v []EstimatedOrderItemModel) {
	o.OrderItemList = v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingAddress() ShippingAddressModel {
	if o == nil || o.ShippingAddress == nil {
		var ret ShippingAddressModel
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingAddressOk() (*ShippingAddressModel, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddressModel and assigns it to the ShippingAddress field.
func (o *EstimatedOrderModel) SetShippingAddress(v ShippingAddressModel) {
	o.ShippingAddress = &v
}

// GetShippingAddressId returns the ShippingAddressId field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingAddressId() string {
	if o == nil || o.ShippingAddressId == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingAddressIdOk() (*string, bool) {
	if o == nil || o.ShippingAddressId == nil {
		return nil, false
	}
	return o.ShippingAddressId, true
}

// HasShippingAddressId returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingAddressId() bool {
	if o != nil && o.ShippingAddressId != nil {
		return true
	}

	return false
}

// SetShippingAddressId gets a reference to the given string and assigns it to the ShippingAddressId field.
func (o *EstimatedOrderModel) SetShippingAddressId(v string) {
	o.ShippingAddressId = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingCost() float64 {
	if o == nil || o.ShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingCostOk() (*float64, bool) {
	if o == nil || o.ShippingCost == nil {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingCost() bool {
	if o != nil && o.ShippingCost != nil {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float64 and assigns it to the ShippingCost field.
func (o *EstimatedOrderModel) SetShippingCost(v float64) {
	o.ShippingCost = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTaxAmountOk() (*float64, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *EstimatedOrderModel) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *EstimatedOrderModel) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

func (o EstimatedOrderModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.OrderItemList != nil {
		toSerialize["orderItemList"] = o.OrderItemList
	}
	if o.ShippingAddress != nil {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if o.ShippingAddressId != nil {
		toSerialize["shippingAddressId"] = o.ShippingAddressId
	}
	if o.ShippingCost != nil {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if o.TaxAmount != nil {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableEstimatedOrderModel struct {
	value *EstimatedOrderModel
	isSet bool
}

func (v NullableEstimatedOrderModel) Get() *EstimatedOrderModel {
	return v.value
}

func (v *NullableEstimatedOrderModel) Set(val *EstimatedOrderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedOrderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedOrderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedOrderModel(val *EstimatedOrderModel) *NullableEstimatedOrderModel {
	return &NullableEstimatedOrderModel{value: val, isSet: true}
}

func (v NullableEstimatedOrderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedOrderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


