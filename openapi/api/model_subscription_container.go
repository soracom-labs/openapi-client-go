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

// SubscriptionContainer struct for SubscriptionContainer
type SubscriptionContainer struct {
	ContainerId *int64 `json:"containerId,omitempty"`
	Downloaded *bool `json:"downloaded,omitempty"`
	Subscriber *map[string]Subscriber `json:"subscriber,omitempty"`
}

// NewSubscriptionContainer instantiates a new SubscriptionContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionContainer() *SubscriptionContainer {
	this := SubscriptionContainer{}
	return &this
}

// NewSubscriptionContainerWithDefaults instantiates a new SubscriptionContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionContainerWithDefaults() *SubscriptionContainer {
	this := SubscriptionContainer{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *SubscriptionContainer) GetContainerId() int64 {
	if o == nil || o.ContainerId == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContainer) GetContainerIdOk() (*int64, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *SubscriptionContainer) HasContainerId() bool {
	if o != nil && o.ContainerId != nil {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *SubscriptionContainer) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetDownloaded returns the Downloaded field value if set, zero value otherwise.
func (o *SubscriptionContainer) GetDownloaded() bool {
	if o == nil || o.Downloaded == nil {
		var ret bool
		return ret
	}
	return *o.Downloaded
}

// GetDownloadedOk returns a tuple with the Downloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContainer) GetDownloadedOk() (*bool, bool) {
	if o == nil || o.Downloaded == nil {
		return nil, false
	}
	return o.Downloaded, true
}

// HasDownloaded returns a boolean if a field has been set.
func (o *SubscriptionContainer) HasDownloaded() bool {
	if o != nil && o.Downloaded != nil {
		return true
	}

	return false
}

// SetDownloaded gets a reference to the given bool and assigns it to the Downloaded field.
func (o *SubscriptionContainer) SetDownloaded(v bool) {
	o.Downloaded = &v
}

// GetSubscriber returns the Subscriber field value if set, zero value otherwise.
func (o *SubscriptionContainer) GetSubscriber() map[string]Subscriber {
	if o == nil || o.Subscriber == nil {
		var ret map[string]Subscriber
		return ret
	}
	return *o.Subscriber
}

// GetSubscriberOk returns a tuple with the Subscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContainer) GetSubscriberOk() (*map[string]Subscriber, bool) {
	if o == nil || o.Subscriber == nil {
		return nil, false
	}
	return o.Subscriber, true
}

// HasSubscriber returns a boolean if a field has been set.
func (o *SubscriptionContainer) HasSubscriber() bool {
	if o != nil && o.Subscriber != nil {
		return true
	}

	return false
}

// SetSubscriber gets a reference to the given map[string]Subscriber and assigns it to the Subscriber field.
func (o *SubscriptionContainer) SetSubscriber(v map[string]Subscriber) {
	o.Subscriber = &v
}

func (o SubscriptionContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerId != nil {
		toSerialize["containerId"] = o.ContainerId
	}
	if o.Downloaded != nil {
		toSerialize["downloaded"] = o.Downloaded
	}
	if o.Subscriber != nil {
		toSerialize["subscriber"] = o.Subscriber
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionContainer struct {
	value *SubscriptionContainer
	isSet bool
}

func (v NullableSubscriptionContainer) Get() *SubscriptionContainer {
	return v.value
}

func (v *NullableSubscriptionContainer) Set(val *SubscriptionContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionContainer(val *SubscriptionContainer) *NullableSubscriptionContainer {
	return &NullableSubscriptionContainer{value: val, isSet: true}
}

func (v NullableSubscriptionContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


