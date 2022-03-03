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

// LagoonUser struct for LagoonUser
type LagoonUser struct {
	// This value used on login.
	Email *string `json:"email,omitempty"`
	Id *float32 `json:"id,omitempty"`
	// The last login datetime.
	LastSeenAt *string `json:"lastSeenAt,omitempty"`
	// The last login datetime as age.
	LastSeenAtAge *string `json:"lastSeenAtAge,omitempty"`
	// A role that represents the permission.
	Role *string `json:"role,omitempty"`
}

// NewLagoonUser instantiates a new LagoonUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonUser() *LagoonUser {
	this := LagoonUser{}
	return &this
}

// NewLagoonUserWithDefaults instantiates a new LagoonUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonUserWithDefaults() *LagoonUser {
	this := LagoonUser{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LagoonUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LagoonUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LagoonUser) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LagoonUser) GetId() float32 {
	if o == nil || o.Id == nil {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUser) GetIdOk() (*float32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LagoonUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *LagoonUser) SetId(v float32) {
	o.Id = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *LagoonUser) GetLastSeenAt() string {
	if o == nil || o.LastSeenAt == nil {
		var ret string
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUser) GetLastSeenAtOk() (*string, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *LagoonUser) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given string and assigns it to the LastSeenAt field.
func (o *LagoonUser) SetLastSeenAt(v string) {
	o.LastSeenAt = &v
}

// GetLastSeenAtAge returns the LastSeenAtAge field value if set, zero value otherwise.
func (o *LagoonUser) GetLastSeenAtAge() string {
	if o == nil || o.LastSeenAtAge == nil {
		var ret string
		return ret
	}
	return *o.LastSeenAtAge
}

// GetLastSeenAtAgeOk returns a tuple with the LastSeenAtAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUser) GetLastSeenAtAgeOk() (*string, bool) {
	if o == nil || o.LastSeenAtAge == nil {
		return nil, false
	}
	return o.LastSeenAtAge, true
}

// HasLastSeenAtAge returns a boolean if a field has been set.
func (o *LagoonUser) HasLastSeenAtAge() bool {
	if o != nil && o.LastSeenAtAge != nil {
		return true
	}

	return false
}

// SetLastSeenAtAge gets a reference to the given string and assigns it to the LastSeenAtAge field.
func (o *LagoonUser) SetLastSeenAtAge(v string) {
	o.LastSeenAtAge = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *LagoonUser) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUser) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *LagoonUser) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *LagoonUser) SetRole(v string) {
	o.Role = &v
}

func (o LagoonUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastSeenAt != nil {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	if o.LastSeenAtAge != nil {
		toSerialize["lastSeenAtAge"] = o.LastSeenAtAge
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonUser struct {
	value *LagoonUser
	isSet bool
}

func (v NullableLagoonUser) Get() *LagoonUser {
	return v.value
}

func (v *NullableLagoonUser) Set(val *LagoonUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonUser(val *LagoonUser) *NullableLagoonUser {
	return &NullableLagoonUser{value: val, isSet: true}
}

func (v NullableLagoonUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


