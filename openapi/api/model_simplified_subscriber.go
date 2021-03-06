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

// SimplifiedSubscriber struct for SimplifiedSubscriber
type SimplifiedSubscriber struct {
	Bundles []string `json:"bundles,omitempty"`
	Capabilities *Capabilities `json:"capabilities,omitempty"`
	Imsi *string `json:"imsi,omitempty"`
	Msisdn *string `json:"msisdn,omitempty"`
	Status *string `json:"status,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
}

// NewSimplifiedSubscriber instantiates a new SimplifiedSubscriber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedSubscriber() *SimplifiedSubscriber {
	this := SimplifiedSubscriber{}
	return &this
}

// NewSimplifiedSubscriberWithDefaults instantiates a new SimplifiedSubscriber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedSubscriberWithDefaults() *SimplifiedSubscriber {
	this := SimplifiedSubscriber{}
	return &this
}

// GetBundles returns the Bundles field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetBundles() []string {
	if o == nil || o.Bundles == nil {
		var ret []string
		return ret
	}
	return o.Bundles
}

// GetBundlesOk returns a tuple with the Bundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetBundlesOk() ([]string, bool) {
	if o == nil || o.Bundles == nil {
		return nil, false
	}
	return o.Bundles, true
}

// HasBundles returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasBundles() bool {
	if o != nil && o.Bundles != nil {
		return true
	}

	return false
}

// SetBundles gets a reference to the given []string and assigns it to the Bundles field.
func (o *SimplifiedSubscriber) SetBundles(v []string) {
	o.Bundles = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetCapabilities() Capabilities {
	if o == nil || o.Capabilities == nil {
		var ret Capabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetCapabilitiesOk() (*Capabilities, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given Capabilities and assigns it to the Capabilities field.
func (o *SimplifiedSubscriber) SetCapabilities(v Capabilities) {
	o.Capabilities = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *SimplifiedSubscriber) SetImsi(v string) {
	o.Imsi = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetMsisdn() string {
	if o == nil || o.Msisdn == nil {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetMsisdnOk() (*string, bool) {
	if o == nil || o.Msisdn == nil {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasMsisdn() bool {
	if o != nil && o.Msisdn != nil {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *SimplifiedSubscriber) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SimplifiedSubscriber) SetStatus(v string) {
	o.Status = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SimplifiedSubscriber) GetSubscription() string {
	if o == nil || o.Subscription == nil {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedSubscriber) GetSubscriptionOk() (*string, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SimplifiedSubscriber) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *SimplifiedSubscriber) SetSubscription(v string) {
	o.Subscription = &v
}

func (o SimplifiedSubscriber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bundles != nil {
		toSerialize["bundles"] = o.Bundles
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.Msisdn != nil {
		toSerialize["msisdn"] = o.Msisdn
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableSimplifiedSubscriber struct {
	value *SimplifiedSubscriber
	isSet bool
}

func (v NullableSimplifiedSubscriber) Get() *SimplifiedSubscriber {
	return v.value
}

func (v *NullableSimplifiedSubscriber) Set(val *SimplifiedSubscriber) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedSubscriber) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedSubscriber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedSubscriber(val *SimplifiedSubscriber) *NullableSimplifiedSubscriber {
	return &NullableSimplifiedSubscriber{value: val, isSet: true}
}

func (v NullableSimplifiedSubscriber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedSubscriber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


