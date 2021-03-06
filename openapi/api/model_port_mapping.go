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

// PortMapping struct for PortMapping
type PortMapping struct {
	Destination *PortMappingDestination `json:"destination,omitempty"`
	// The duration that remote access is enabled, in seconds.
	Duration *float32 `json:"duration,omitempty"`
	// SORACOM Napter endpoint (IP address and port number) for remote access.
	Endpoint *string `json:"endpoint,omitempty"`
	// SORACOM Napter hostname for remote access.
	Hostname *string `json:"hostname,omitempty"`
	// SORACOM Napter IP Address for remote access.
	IpAddress *string `json:"ipAddress,omitempty"`
	// SORACOM Napter port number for remote access.
	Port *float32 `json:"port,omitempty"`
	Source *PortMappingSource `json:"source,omitempty"`
	// Indicates TLS is required.
	TlsRequired *bool `json:"tlsRequired,omitempty"`
}

// NewPortMapping instantiates a new PortMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortMapping() *PortMapping {
	this := PortMapping{}
	return &this
}

// NewPortMappingWithDefaults instantiates a new PortMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortMappingWithDefaults() *PortMapping {
	this := PortMapping{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *PortMapping) GetDestination() PortMappingDestination {
	if o == nil || o.Destination == nil {
		var ret PortMappingDestination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetDestinationOk() (*PortMappingDestination, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *PortMapping) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given PortMappingDestination and assigns it to the Destination field.
func (o *PortMapping) SetDestination(v PortMappingDestination) {
	o.Destination = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PortMapping) GetDuration() float32 {
	if o == nil || o.Duration == nil {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetDurationOk() (*float32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PortMapping) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *PortMapping) SetDuration(v float32) {
	o.Duration = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PortMapping) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PortMapping) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PortMapping) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *PortMapping) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *PortMapping) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *PortMapping) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *PortMapping) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PortMapping) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *PortMapping) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PortMapping) GetPort() float32 {
	if o == nil || o.Port == nil {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetPortOk() (*float32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PortMapping) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *PortMapping) SetPort(v float32) {
	o.Port = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PortMapping) GetSource() PortMappingSource {
	if o == nil || o.Source == nil {
		var ret PortMappingSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetSourceOk() (*PortMappingSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PortMapping) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given PortMappingSource and assigns it to the Source field.
func (o *PortMapping) SetSource(v PortMappingSource) {
	o.Source = &v
}

// GetTlsRequired returns the TlsRequired field value if set, zero value otherwise.
func (o *PortMapping) GetTlsRequired() bool {
	if o == nil || o.TlsRequired == nil {
		var ret bool
		return ret
	}
	return *o.TlsRequired
}

// GetTlsRequiredOk returns a tuple with the TlsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetTlsRequiredOk() (*bool, bool) {
	if o == nil || o.TlsRequired == nil {
		return nil, false
	}
	return o.TlsRequired, true
}

// HasTlsRequired returns a boolean if a field has been set.
func (o *PortMapping) HasTlsRequired() bool {
	if o != nil && o.TlsRequired != nil {
		return true
	}

	return false
}

// SetTlsRequired gets a reference to the given bool and assigns it to the TlsRequired field.
func (o *PortMapping) SetTlsRequired(v bool) {
	o.TlsRequired = &v
}

func (o PortMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.TlsRequired != nil {
		toSerialize["tlsRequired"] = o.TlsRequired
	}
	return json.Marshal(toSerialize)
}

type NullablePortMapping struct {
	value *PortMapping
	isSet bool
}

func (v NullablePortMapping) Get() *PortMapping {
	return v.value
}

func (v *NullablePortMapping) Set(val *PortMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePortMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePortMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortMapping(val *PortMapping) *NullablePortMapping {
	return &NullablePortMapping{value: val, isSet: true}
}

func (v NullablePortMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


