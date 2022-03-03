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

// ArcSessionCreateResponse struct for ArcSessionCreateResponse
type ArcSessionCreateResponse struct {
	ArcAllowedIPs []string `json:"arcAllowedIPs,omitempty"`
	ArcClientPeerIpAddress *string `json:"arcClientPeerIpAddress,omitempty"`
	ArcClientPeerPublicKey *string `json:"arcClientPeerPublicKey,omitempty"`
	ArcServerEndpoint *string `json:"arcServerEndpoint,omitempty"`
	ArcServerPeerPublicKey *string `json:"arcServerPeerPublicKey,omitempty"`
}

// NewArcSessionCreateResponse instantiates a new ArcSessionCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcSessionCreateResponse() *ArcSessionCreateResponse {
	this := ArcSessionCreateResponse{}
	return &this
}

// NewArcSessionCreateResponseWithDefaults instantiates a new ArcSessionCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcSessionCreateResponseWithDefaults() *ArcSessionCreateResponse {
	this := ArcSessionCreateResponse{}
	return &this
}

// GetArcAllowedIPs returns the ArcAllowedIPs field value if set, zero value otherwise.
func (o *ArcSessionCreateResponse) GetArcAllowedIPs() []string {
	if o == nil || o.ArcAllowedIPs == nil {
		var ret []string
		return ret
	}
	return o.ArcAllowedIPs
}

// GetArcAllowedIPsOk returns a tuple with the ArcAllowedIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionCreateResponse) GetArcAllowedIPsOk() ([]string, bool) {
	if o == nil || o.ArcAllowedIPs == nil {
		return nil, false
	}
	return o.ArcAllowedIPs, true
}

// HasArcAllowedIPs returns a boolean if a field has been set.
func (o *ArcSessionCreateResponse) HasArcAllowedIPs() bool {
	if o != nil && o.ArcAllowedIPs != nil {
		return true
	}

	return false
}

// SetArcAllowedIPs gets a reference to the given []string and assigns it to the ArcAllowedIPs field.
func (o *ArcSessionCreateResponse) SetArcAllowedIPs(v []string) {
	o.ArcAllowedIPs = v
}

// GetArcClientPeerIpAddress returns the ArcClientPeerIpAddress field value if set, zero value otherwise.
func (o *ArcSessionCreateResponse) GetArcClientPeerIpAddress() string {
	if o == nil || o.ArcClientPeerIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ArcClientPeerIpAddress
}

// GetArcClientPeerIpAddressOk returns a tuple with the ArcClientPeerIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionCreateResponse) GetArcClientPeerIpAddressOk() (*string, bool) {
	if o == nil || o.ArcClientPeerIpAddress == nil {
		return nil, false
	}
	return o.ArcClientPeerIpAddress, true
}

// HasArcClientPeerIpAddress returns a boolean if a field has been set.
func (o *ArcSessionCreateResponse) HasArcClientPeerIpAddress() bool {
	if o != nil && o.ArcClientPeerIpAddress != nil {
		return true
	}

	return false
}

// SetArcClientPeerIpAddress gets a reference to the given string and assigns it to the ArcClientPeerIpAddress field.
func (o *ArcSessionCreateResponse) SetArcClientPeerIpAddress(v string) {
	o.ArcClientPeerIpAddress = &v
}

// GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field value if set, zero value otherwise.
func (o *ArcSessionCreateResponse) GetArcClientPeerPublicKey() string {
	if o == nil || o.ArcClientPeerPublicKey == nil {
		var ret string
		return ret
	}
	return *o.ArcClientPeerPublicKey
}

// GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionCreateResponse) GetArcClientPeerPublicKeyOk() (*string, bool) {
	if o == nil || o.ArcClientPeerPublicKey == nil {
		return nil, false
	}
	return o.ArcClientPeerPublicKey, true
}

// HasArcClientPeerPublicKey returns a boolean if a field has been set.
func (o *ArcSessionCreateResponse) HasArcClientPeerPublicKey() bool {
	if o != nil && o.ArcClientPeerPublicKey != nil {
		return true
	}

	return false
}

// SetArcClientPeerPublicKey gets a reference to the given string and assigns it to the ArcClientPeerPublicKey field.
func (o *ArcSessionCreateResponse) SetArcClientPeerPublicKey(v string) {
	o.ArcClientPeerPublicKey = &v
}

// GetArcServerEndpoint returns the ArcServerEndpoint field value if set, zero value otherwise.
func (o *ArcSessionCreateResponse) GetArcServerEndpoint() string {
	if o == nil || o.ArcServerEndpoint == nil {
		var ret string
		return ret
	}
	return *o.ArcServerEndpoint
}

// GetArcServerEndpointOk returns a tuple with the ArcServerEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionCreateResponse) GetArcServerEndpointOk() (*string, bool) {
	if o == nil || o.ArcServerEndpoint == nil {
		return nil, false
	}
	return o.ArcServerEndpoint, true
}

// HasArcServerEndpoint returns a boolean if a field has been set.
func (o *ArcSessionCreateResponse) HasArcServerEndpoint() bool {
	if o != nil && o.ArcServerEndpoint != nil {
		return true
	}

	return false
}

// SetArcServerEndpoint gets a reference to the given string and assigns it to the ArcServerEndpoint field.
func (o *ArcSessionCreateResponse) SetArcServerEndpoint(v string) {
	o.ArcServerEndpoint = &v
}

// GetArcServerPeerPublicKey returns the ArcServerPeerPublicKey field value if set, zero value otherwise.
func (o *ArcSessionCreateResponse) GetArcServerPeerPublicKey() string {
	if o == nil || o.ArcServerPeerPublicKey == nil {
		var ret string
		return ret
	}
	return *o.ArcServerPeerPublicKey
}

// GetArcServerPeerPublicKeyOk returns a tuple with the ArcServerPeerPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionCreateResponse) GetArcServerPeerPublicKeyOk() (*string, bool) {
	if o == nil || o.ArcServerPeerPublicKey == nil {
		return nil, false
	}
	return o.ArcServerPeerPublicKey, true
}

// HasArcServerPeerPublicKey returns a boolean if a field has been set.
func (o *ArcSessionCreateResponse) HasArcServerPeerPublicKey() bool {
	if o != nil && o.ArcServerPeerPublicKey != nil {
		return true
	}

	return false
}

// SetArcServerPeerPublicKey gets a reference to the given string and assigns it to the ArcServerPeerPublicKey field.
func (o *ArcSessionCreateResponse) SetArcServerPeerPublicKey(v string) {
	o.ArcServerPeerPublicKey = &v
}

func (o ArcSessionCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArcAllowedIPs != nil {
		toSerialize["arcAllowedIPs"] = o.ArcAllowedIPs
	}
	if o.ArcClientPeerIpAddress != nil {
		toSerialize["arcClientPeerIpAddress"] = o.ArcClientPeerIpAddress
	}
	if o.ArcClientPeerPublicKey != nil {
		toSerialize["arcClientPeerPublicKey"] = o.ArcClientPeerPublicKey
	}
	if o.ArcServerEndpoint != nil {
		toSerialize["arcServerEndpoint"] = o.ArcServerEndpoint
	}
	if o.ArcServerPeerPublicKey != nil {
		toSerialize["arcServerPeerPublicKey"] = o.ArcServerPeerPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableArcSessionCreateResponse struct {
	value *ArcSessionCreateResponse
	isSet bool
}

func (v NullableArcSessionCreateResponse) Get() *ArcSessionCreateResponse {
	return v.value
}

func (v *NullableArcSessionCreateResponse) Set(val *ArcSessionCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArcSessionCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArcSessionCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArcSessionCreateResponse(val *ArcSessionCreateResponse) *NullableArcSessionCreateResponse {
	return &NullableArcSessionCreateResponse{value: val, isSet: true}
}

func (v NullableArcSessionCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArcSessionCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


