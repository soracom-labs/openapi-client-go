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

// UserDetailResponse struct for UserDetailResponse
type UserDetailResponse struct {
	AuthKeyList []AuthKeyResponse `json:"authKeyList,omitempty"`
	CreateDateTime *int64 `json:"createDateTime,omitempty"`
	Description *string `json:"description,omitempty"`
	HasPassword *bool `json:"hasPassword,omitempty"`
	Permission *string `json:"permission,omitempty"`
	RoleList []ListRolesResponse `json:"roleList,omitempty"`
	UpdateDateTime *int64 `json:"updateDateTime,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// NewUserDetailResponse instantiates a new UserDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDetailResponse() *UserDetailResponse {
	this := UserDetailResponse{}
	return &this
}

// NewUserDetailResponseWithDefaults instantiates a new UserDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDetailResponseWithDefaults() *UserDetailResponse {
	this := UserDetailResponse{}
	return &this
}

// GetAuthKeyList returns the AuthKeyList field value if set, zero value otherwise.
func (o *UserDetailResponse) GetAuthKeyList() []AuthKeyResponse {
	if o == nil || o.AuthKeyList == nil {
		var ret []AuthKeyResponse
		return ret
	}
	return o.AuthKeyList
}

// GetAuthKeyListOk returns a tuple with the AuthKeyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetAuthKeyListOk() ([]AuthKeyResponse, bool) {
	if o == nil || o.AuthKeyList == nil {
		return nil, false
	}
	return o.AuthKeyList, true
}

// HasAuthKeyList returns a boolean if a field has been set.
func (o *UserDetailResponse) HasAuthKeyList() bool {
	if o != nil && o.AuthKeyList != nil {
		return true
	}

	return false
}

// SetAuthKeyList gets a reference to the given []AuthKeyResponse and assigns it to the AuthKeyList field.
func (o *UserDetailResponse) SetAuthKeyList(v []AuthKeyResponse) {
	o.AuthKeyList = v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *UserDetailResponse) GetCreateDateTime() int64 {
	if o == nil || o.CreateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetCreateDateTimeOk() (*int64, bool) {
	if o == nil || o.CreateDateTime == nil {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *UserDetailResponse) HasCreateDateTime() bool {
	if o != nil && o.CreateDateTime != nil {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given int64 and assigns it to the CreateDateTime field.
func (o *UserDetailResponse) SetCreateDateTime(v int64) {
	o.CreateDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserDetailResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserDetailResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserDetailResponse) SetDescription(v string) {
	o.Description = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *UserDetailResponse) GetHasPassword() bool {
	if o == nil || o.HasPassword == nil {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetHasPasswordOk() (*bool, bool) {
	if o == nil || o.HasPassword == nil {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *UserDetailResponse) HasHasPassword() bool {
	if o != nil && o.HasPassword != nil {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *UserDetailResponse) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *UserDetailResponse) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *UserDetailResponse) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *UserDetailResponse) SetPermission(v string) {
	o.Permission = &v
}

// GetRoleList returns the RoleList field value if set, zero value otherwise.
func (o *UserDetailResponse) GetRoleList() []ListRolesResponse {
	if o == nil || o.RoleList == nil {
		var ret []ListRolesResponse
		return ret
	}
	return o.RoleList
}

// GetRoleListOk returns a tuple with the RoleList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetRoleListOk() ([]ListRolesResponse, bool) {
	if o == nil || o.RoleList == nil {
		return nil, false
	}
	return o.RoleList, true
}

// HasRoleList returns a boolean if a field has been set.
func (o *UserDetailResponse) HasRoleList() bool {
	if o != nil && o.RoleList != nil {
		return true
	}

	return false
}

// SetRoleList gets a reference to the given []ListRolesResponse and assigns it to the RoleList field.
func (o *UserDetailResponse) SetRoleList(v []ListRolesResponse) {
	o.RoleList = v
}

// GetUpdateDateTime returns the UpdateDateTime field value if set, zero value otherwise.
func (o *UserDetailResponse) GetUpdateDateTime() int64 {
	if o == nil || o.UpdateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDateTime
}

// GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetUpdateDateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateDateTime == nil {
		return nil, false
	}
	return o.UpdateDateTime, true
}

// HasUpdateDateTime returns a boolean if a field has been set.
func (o *UserDetailResponse) HasUpdateDateTime() bool {
	if o != nil && o.UpdateDateTime != nil {
		return true
	}

	return false
}

// SetUpdateDateTime gets a reference to the given int64 and assigns it to the UpdateDateTime field.
func (o *UserDetailResponse) SetUpdateDateTime(v int64) {
	o.UpdateDateTime = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserDetailResponse) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetailResponse) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserDetailResponse) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserDetailResponse) SetUserName(v string) {
	o.UserName = &v
}

func (o UserDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthKeyList != nil {
		toSerialize["authKeyList"] = o.AuthKeyList
	}
	if o.CreateDateTime != nil {
		toSerialize["createDateTime"] = o.CreateDateTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.HasPassword != nil {
		toSerialize["hasPassword"] = o.HasPassword
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.RoleList != nil {
		toSerialize["roleList"] = o.RoleList
	}
	if o.UpdateDateTime != nil {
		toSerialize["updateDateTime"] = o.UpdateDateTime
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableUserDetailResponse struct {
	value *UserDetailResponse
	isSet bool
}

func (v NullableUserDetailResponse) Get() *UserDetailResponse {
	return v.value
}

func (v *NullableUserDetailResponse) Set(val *UserDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDetailResponse(val *UserDetailResponse) *NullableUserDetailResponse {
	return &NullableUserDetailResponse{value: val, isSet: true}
}

func (v NullableUserDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


