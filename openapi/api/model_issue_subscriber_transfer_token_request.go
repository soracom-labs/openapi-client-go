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

// IssueSubscriberTransferTokenRequest struct for IssueSubscriberTransferTokenRequest
type IssueSubscriberTransferTokenRequest struct {
	// 移管先オペレーターEmail
	TransferDestinationOperatorEmail string `json:"transferDestinationOperatorEmail"`
	// 移管先オペレーターID
	TransferDestinationOperatorId string `json:"transferDestinationOperatorId"`
	// 移管する IMSI リスト
	TransferImsi []string `json:"transferImsi"`
}

// NewIssueSubscriberTransferTokenRequest instantiates a new IssueSubscriberTransferTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueSubscriberTransferTokenRequest(transferDestinationOperatorEmail string, transferDestinationOperatorId string, transferImsi []string) *IssueSubscriberTransferTokenRequest {
	this := IssueSubscriberTransferTokenRequest{}
	this.TransferDestinationOperatorEmail = transferDestinationOperatorEmail
	this.TransferDestinationOperatorId = transferDestinationOperatorId
	this.TransferImsi = transferImsi
	return &this
}

// NewIssueSubscriberTransferTokenRequestWithDefaults instantiates a new IssueSubscriberTransferTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueSubscriberTransferTokenRequestWithDefaults() *IssueSubscriberTransferTokenRequest {
	this := IssueSubscriberTransferTokenRequest{}
	return &this
}

// GetTransferDestinationOperatorEmail returns the TransferDestinationOperatorEmail field value
func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferDestinationOperatorEmail
}

// GetTransferDestinationOperatorEmailOk returns a tuple with the TransferDestinationOperatorEmail field value
// and a boolean to check if the value has been set.
func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferDestinationOperatorEmail, true
}

// SetTransferDestinationOperatorEmail sets field value
func (o *IssueSubscriberTransferTokenRequest) SetTransferDestinationOperatorEmail(v string) {
	o.TransferDestinationOperatorEmail = v
}

// GetTransferDestinationOperatorId returns the TransferDestinationOperatorId field value
func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferDestinationOperatorId
}

// GetTransferDestinationOperatorIdOk returns a tuple with the TransferDestinationOperatorId field value
// and a boolean to check if the value has been set.
func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferDestinationOperatorId, true
}

// SetTransferDestinationOperatorId sets field value
func (o *IssueSubscriberTransferTokenRequest) SetTransferDestinationOperatorId(v string) {
	o.TransferDestinationOperatorId = v
}

// GetTransferImsi returns the TransferImsi field value
func (o *IssueSubscriberTransferTokenRequest) GetTransferImsi() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TransferImsi
}

// GetTransferImsiOk returns a tuple with the TransferImsi field value
// and a boolean to check if the value has been set.
func (o *IssueSubscriberTransferTokenRequest) GetTransferImsiOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferImsi, true
}

// SetTransferImsi sets field value
func (o *IssueSubscriberTransferTokenRequest) SetTransferImsi(v []string) {
	o.TransferImsi = v
}

func (o IssueSubscriberTransferTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transferDestinationOperatorEmail"] = o.TransferDestinationOperatorEmail
	}
	if true {
		toSerialize["transferDestinationOperatorId"] = o.TransferDestinationOperatorId
	}
	if true {
		toSerialize["transferImsi"] = o.TransferImsi
	}
	return json.Marshal(toSerialize)
}

type NullableIssueSubscriberTransferTokenRequest struct {
	value *IssueSubscriberTransferTokenRequest
	isSet bool
}

func (v NullableIssueSubscriberTransferTokenRequest) Get() *IssueSubscriberTransferTokenRequest {
	return v.value
}

func (v *NullableIssueSubscriberTransferTokenRequest) Set(val *IssueSubscriberTransferTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueSubscriberTransferTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueSubscriberTransferTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueSubscriberTransferTokenRequest(val *IssueSubscriberTransferTokenRequest) *NullableIssueSubscriberTransferTokenRequest {
	return &NullableIssueSubscriberTransferTokenRequest{value: val, isSet: true}
}

func (v NullableIssueSubscriberTransferTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueSubscriberTransferTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


