/*
SORACOM SANDBOX API

SORACOM SANDBOX API v1

API version: 20160218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sandbox

import (
	"encoding/json"
)

// SandboxCreateSubscriberResponse struct for SandboxCreateSubscriberResponse
type SandboxCreateSubscriberResponse struct {
	Tags *map[string]string `json:"tags,omitempty"`
	Apn *string `json:"apn,omitempty"`
	CreatedAt *int64 `json:"createdAt,omitempty"`
	ExpiryTime *int64 `json:"expiryTime,omitempty"`
	Imsi *string `json:"imsi,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`
	Msisdn *string `json:"msisdn,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	RegistrationSecret *string `json:"registrationSecret,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	SpeedClass *string `json:"speedClass,omitempty"`
	Status *string `json:"status,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
}

// NewSandboxCreateSubscriberResponse instantiates a new SandboxCreateSubscriberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxCreateSubscriberResponse() *SandboxCreateSubscriberResponse {
	this := SandboxCreateSubscriberResponse{}
	return &this
}

// NewSandboxCreateSubscriberResponseWithDefaults instantiates a new SandboxCreateSubscriberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxCreateSubscriberResponseWithDefaults() *SandboxCreateSubscriberResponse {
	this := SandboxCreateSubscriberResponse{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SandboxCreateSubscriberResponse) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetApn() string {
	if o == nil || o.Apn == nil {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetApnOk() (*string, bool) {
	if o == nil || o.Apn == nil {
		return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasApn() bool {
	if o != nil && o.Apn != nil {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *SandboxCreateSubscriberResponse) SetApn(v string) {
	o.Apn = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SandboxCreateSubscriberResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetExpiryTime() int64 {
	if o == nil || o.ExpiryTime == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetExpiryTimeOk() (*int64, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given int64 and assigns it to the ExpiryTime field.
func (o *SandboxCreateSubscriberResponse) SetExpiryTime(v int64) {
	o.ExpiryTime = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *SandboxCreateSubscriberResponse) SetImsi(v string) {
	o.Imsi = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SandboxCreateSubscriberResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLastModifiedAt returns the LastModifiedAt field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetLastModifiedAt() int64 {
	if o == nil || o.LastModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetLastModifiedAtOk() (*int64, bool) {
	if o == nil || o.LastModifiedAt == nil {
		return nil, false
	}
	return o.LastModifiedAt, true
}

// HasLastModifiedAt returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasLastModifiedAt() bool {
	if o != nil && o.LastModifiedAt != nil {
		return true
	}

	return false
}

// SetLastModifiedAt gets a reference to the given int64 and assigns it to the LastModifiedAt field.
func (o *SandboxCreateSubscriberResponse) SetLastModifiedAt(v int64) {
	o.LastModifiedAt = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetMsisdn() string {
	if o == nil || o.Msisdn == nil {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetMsisdnOk() (*string, bool) {
	if o == nil || o.Msisdn == nil {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasMsisdn() bool {
	if o != nil && o.Msisdn != nil {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *SandboxCreateSubscriberResponse) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *SandboxCreateSubscriberResponse) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetRegistrationSecret returns the RegistrationSecret field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetRegistrationSecret() string {
	if o == nil || o.RegistrationSecret == nil {
		var ret string
		return ret
	}
	return *o.RegistrationSecret
}

// GetRegistrationSecretOk returns a tuple with the RegistrationSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetRegistrationSecretOk() (*string, bool) {
	if o == nil || o.RegistrationSecret == nil {
		return nil, false
	}
	return o.RegistrationSecret, true
}

// HasRegistrationSecret returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasRegistrationSecret() bool {
	if o != nil && o.RegistrationSecret != nil {
		return true
	}

	return false
}

// SetRegistrationSecret gets a reference to the given string and assigns it to the RegistrationSecret field.
func (o *SandboxCreateSubscriberResponse) SetRegistrationSecret(v string) {
	o.RegistrationSecret = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *SandboxCreateSubscriberResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSpeedClass returns the SpeedClass field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetSpeedClass() string {
	if o == nil || o.SpeedClass == nil {
		var ret string
		return ret
	}
	return *o.SpeedClass
}

// GetSpeedClassOk returns a tuple with the SpeedClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetSpeedClassOk() (*string, bool) {
	if o == nil || o.SpeedClass == nil {
		return nil, false
	}
	return o.SpeedClass, true
}

// HasSpeedClass returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasSpeedClass() bool {
	if o != nil && o.SpeedClass != nil {
		return true
	}

	return false
}

// SetSpeedClass gets a reference to the given string and assigns it to the SpeedClass field.
func (o *SandboxCreateSubscriberResponse) SetSpeedClass(v string) {
	o.SpeedClass = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SandboxCreateSubscriberResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SandboxCreateSubscriberResponse) GetSubscription() string {
	if o == nil || o.Subscription == nil {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxCreateSubscriberResponse) GetSubscriptionOk() (*string, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SandboxCreateSubscriberResponse) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *SandboxCreateSubscriberResponse) SetSubscription(v string) {
	o.Subscription = &v
}

func (o SandboxCreateSubscriberResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Apn != nil {
		toSerialize["apn"] = o.Apn
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ExpiryTime != nil {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.LastModifiedAt != nil {
		toSerialize["lastModifiedAt"] = o.LastModifiedAt
	}
	if o.Msisdn != nil {
		toSerialize["msisdn"] = o.Msisdn
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.RegistrationSecret != nil {
		toSerialize["registrationSecret"] = o.RegistrationSecret
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.SpeedClass != nil {
		toSerialize["speedClass"] = o.SpeedClass
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxCreateSubscriberResponse struct {
	value *SandboxCreateSubscriberResponse
	isSet bool
}

func (v NullableSandboxCreateSubscriberResponse) Get() *SandboxCreateSubscriberResponse {
	return v.value
}

func (v *NullableSandboxCreateSubscriberResponse) Set(val *SandboxCreateSubscriberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxCreateSubscriberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxCreateSubscriberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxCreateSubscriberResponse(val *SandboxCreateSubscriberResponse) *NullableSandboxCreateSubscriberResponse {
	return &NullableSandboxCreateSubscriberResponse{value: val, isSet: true}
}

func (v NullableSandboxCreateSubscriberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxCreateSubscriberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


