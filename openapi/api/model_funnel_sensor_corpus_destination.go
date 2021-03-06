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

// FunnelSensorCorpusDestination struct for FunnelSensorCorpusDestination
type FunnelSensorCorpusDestination struct {
	AdditionalData *string `json:"additionalData,omitempty"`
	Provider *string `json:"provider,omitempty"`
	ResourceUrl *string `json:"resourceUrl,omitempty"`
	Service *string `json:"service,omitempty"`
}

// NewFunnelSensorCorpusDestination instantiates a new FunnelSensorCorpusDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunnelSensorCorpusDestination() *FunnelSensorCorpusDestination {
	this := FunnelSensorCorpusDestination{}
	return &this
}

// NewFunnelSensorCorpusDestinationWithDefaults instantiates a new FunnelSensorCorpusDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunnelSensorCorpusDestinationWithDefaults() *FunnelSensorCorpusDestination {
	this := FunnelSensorCorpusDestination{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *FunnelSensorCorpusDestination) GetAdditionalData() string {
	if o == nil || o.AdditionalData == nil {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSensorCorpusDestination) GetAdditionalDataOk() (*string, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *FunnelSensorCorpusDestination) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *FunnelSensorCorpusDestination) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *FunnelSensorCorpusDestination) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSensorCorpusDestination) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *FunnelSensorCorpusDestination) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *FunnelSensorCorpusDestination) SetProvider(v string) {
	o.Provider = &v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *FunnelSensorCorpusDestination) GetResourceUrl() string {
	if o == nil || o.ResourceUrl == nil {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSensorCorpusDestination) GetResourceUrlOk() (*string, bool) {
	if o == nil || o.ResourceUrl == nil {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *FunnelSensorCorpusDestination) HasResourceUrl() bool {
	if o != nil && o.ResourceUrl != nil {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *FunnelSensorCorpusDestination) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FunnelSensorCorpusDestination) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSensorCorpusDestination) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FunnelSensorCorpusDestination) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FunnelSensorCorpusDestination) SetService(v string) {
	o.Service = &v
}

func (o FunnelSensorCorpusDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.ResourceUrl != nil {
		toSerialize["resourceUrl"] = o.ResourceUrl
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableFunnelSensorCorpusDestination struct {
	value *FunnelSensorCorpusDestination
	isSet bool
}

func (v NullableFunnelSensorCorpusDestination) Get() *FunnelSensorCorpusDestination {
	return v.value
}

func (v *NullableFunnelSensorCorpusDestination) Set(val *FunnelSensorCorpusDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelSensorCorpusDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelSensorCorpusDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelSensorCorpusDestination(val *FunnelSensorCorpusDestination) *NullableFunnelSensorCorpusDestination {
	return &NullableFunnelSensorCorpusDestination{value: val, isSet: true}
}

func (v NullableFunnelSensorCorpusDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelSensorCorpusDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


