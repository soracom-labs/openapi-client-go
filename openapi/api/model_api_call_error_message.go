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

// APICallErrorMessage struct for APICallErrorMessage
type APICallErrorMessage struct {
	// エラーコード
	Code string `json:"code"`
	// エラーメッセージ。リクエスト時にX-Soracom-Langヘッダーに言語(en,ja)を設定するとその言語のメッセージがセットされます。
	Message string `json:"message"`
}

// NewAPICallErrorMessage instantiates a new APICallErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPICallErrorMessage(code string, message string) *APICallErrorMessage {
	this := APICallErrorMessage{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAPICallErrorMessageWithDefaults instantiates a new APICallErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPICallErrorMessageWithDefaults() *APICallErrorMessage {
	this := APICallErrorMessage{}
	return &this
}

// GetCode returns the Code field value
func (o *APICallErrorMessage) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *APICallErrorMessage) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *APICallErrorMessage) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *APICallErrorMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *APICallErrorMessage) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *APICallErrorMessage) SetMessage(v string) {
	o.Message = v
}

func (o APICallErrorMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAPICallErrorMessage struct {
	value *APICallErrorMessage
	isSet bool
}

func (v NullableAPICallErrorMessage) Get() *APICallErrorMessage {
	return v.value
}

func (v *NullableAPICallErrorMessage) Set(val *APICallErrorMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAPICallErrorMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAPICallErrorMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPICallErrorMessage(val *APICallErrorMessage) *NullableAPICallErrorMessage {
	return &NullableAPICallErrorMessage{value: val, isSet: true}
}

func (v NullableAPICallErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPICallErrorMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


