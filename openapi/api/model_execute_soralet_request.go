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

// ExecuteSoraletRequest struct for ExecuteSoraletRequest
type ExecuteSoraletRequest struct {
	ContentType string `json:"contentType"`
	Direction string `json:"direction"`
	EncodingType *string `json:"encodingType,omitempty"`
	Payload string `json:"payload"`
	Source map[string]SoraletDataSource `json:"source"`
	Userdata *string `json:"userdata,omitempty"`
	Version string `json:"version"`
}

// NewExecuteSoraletRequest instantiates a new ExecuteSoraletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteSoraletRequest(contentType string, direction string, payload string, source map[string]SoraletDataSource, version string) *ExecuteSoraletRequest {
	this := ExecuteSoraletRequest{}
	this.ContentType = contentType
	this.Direction = direction
	this.Payload = payload
	this.Source = source
	this.Version = version
	return &this
}

// NewExecuteSoraletRequestWithDefaults instantiates a new ExecuteSoraletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteSoraletRequestWithDefaults() *ExecuteSoraletRequest {
	this := ExecuteSoraletRequest{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *ExecuteSoraletRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ExecuteSoraletRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetDirection returns the Direction field value
func (o *ExecuteSoraletRequest) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetDirectionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *ExecuteSoraletRequest) SetDirection(v string) {
	o.Direction = v
}

// GetEncodingType returns the EncodingType field value if set, zero value otherwise.
func (o *ExecuteSoraletRequest) GetEncodingType() string {
	if o == nil || o.EncodingType == nil {
		var ret string
		return ret
	}
	return *o.EncodingType
}

// GetEncodingTypeOk returns a tuple with the EncodingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetEncodingTypeOk() (*string, bool) {
	if o == nil || o.EncodingType == nil {
		return nil, false
	}
	return o.EncodingType, true
}

// HasEncodingType returns a boolean if a field has been set.
func (o *ExecuteSoraletRequest) HasEncodingType() bool {
	if o != nil && o.EncodingType != nil {
		return true
	}

	return false
}

// SetEncodingType gets a reference to the given string and assigns it to the EncodingType field.
func (o *ExecuteSoraletRequest) SetEncodingType(v string) {
	o.EncodingType = &v
}

// GetPayload returns the Payload field value
func (o *ExecuteSoraletRequest) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetPayloadOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *ExecuteSoraletRequest) SetPayload(v string) {
	o.Payload = v
}

// GetSource returns the Source field value
func (o *ExecuteSoraletRequest) GetSource() map[string]SoraletDataSource {
	if o == nil {
		var ret map[string]SoraletDataSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetSourceOk() (*map[string]SoraletDataSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ExecuteSoraletRequest) SetSource(v map[string]SoraletDataSource) {
	o.Source = v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *ExecuteSoraletRequest) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *ExecuteSoraletRequest) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *ExecuteSoraletRequest) SetUserdata(v string) {
	o.Userdata = &v
}

// GetVersion returns the Version field value
func (o *ExecuteSoraletRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ExecuteSoraletRequest) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ExecuteSoraletRequest) SetVersion(v string) {
	o.Version = v
}

func (o ExecuteSoraletRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentType"] = o.ContentType
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if o.EncodingType != nil {
		toSerialize["encodingType"] = o.EncodingType
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteSoraletRequest struct {
	value *ExecuteSoraletRequest
	isSet bool
}

func (v NullableExecuteSoraletRequest) Get() *ExecuteSoraletRequest {
	return v.value
}

func (v *NullableExecuteSoraletRequest) Set(val *ExecuteSoraletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteSoraletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteSoraletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteSoraletRequest(val *ExecuteSoraletRequest) *NullableExecuteSoraletRequest {
	return &NullableExecuteSoraletRequest{value: val, isSet: true}
}

func (v NullableExecuteSoraletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteSoraletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


