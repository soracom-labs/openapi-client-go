/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Device struct for Device
type Device struct {
	DeviceId *string `json:"device_id,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	LastRegistrationUpdate *time.Time `json:"lastRegistrationUpdate,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	ModelNumber *string `json:"modelNumber,omitempty"`
	Objects map[string]interface{} `json:"objects,omitempty"`
	Online *bool `json:"online,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	RegistrationId *string `json:"registrationId,omitempty"`
	RegistrationLifeTime *int64 `json:"registrationLifeTime,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	var online bool = false
	this.Online = &online
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	var online bool = false
	this.Online = &online
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Device) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Device) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *Device) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *Device) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *Device) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *Device) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *Device) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *Device) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *Device) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Device) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Device) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Device) SetGroupId(v string) {
	o.GroupId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Device) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Device) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *Device) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *Device) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *Device) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *Device) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetLastRegistrationUpdate returns the LastRegistrationUpdate field value if set, zero value otherwise.
func (o *Device) GetLastRegistrationUpdate() time.Time {
	if o == nil || o.LastRegistrationUpdate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRegistrationUpdate
}

// GetLastRegistrationUpdateOk returns a tuple with the LastRegistrationUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastRegistrationUpdateOk() (*time.Time, bool) {
	if o == nil || o.LastRegistrationUpdate == nil {
		return nil, false
	}
	return o.LastRegistrationUpdate, true
}

// HasLastRegistrationUpdate returns a boolean if a field has been set.
func (o *Device) HasLastRegistrationUpdate() bool {
	if o != nil && o.LastRegistrationUpdate != nil {
		return true
	}

	return false
}

// SetLastRegistrationUpdate gets a reference to the given time.Time and assigns it to the LastRegistrationUpdate field.
func (o *Device) SetLastRegistrationUpdate(v time.Time) {
	o.LastRegistrationUpdate = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *Device) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *Device) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *Device) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *Device) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *Device) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *Device) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *Device) GetObjects() map[string]interface{} {
	if o == nil || o.Objects == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetObjectsOk() (map[string]interface{}, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *Device) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given map[string]interface{} and assigns it to the Objects field.
func (o *Device) SetObjects(v map[string]interface{}) {
	o.Objects = v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *Device) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *Device) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *Device) SetOnline(v bool) {
	o.Online = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *Device) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *Device) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *Device) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetRegistrationId returns the RegistrationId field value if set, zero value otherwise.
func (o *Device) GetRegistrationId() string {
	if o == nil || o.RegistrationId == nil {
		var ret string
		return ret
	}
	return *o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetRegistrationIdOk() (*string, bool) {
	if o == nil || o.RegistrationId == nil {
		return nil, false
	}
	return o.RegistrationId, true
}

// HasRegistrationId returns a boolean if a field has been set.
func (o *Device) HasRegistrationId() bool {
	if o != nil && o.RegistrationId != nil {
		return true
	}

	return false
}

// SetRegistrationId gets a reference to the given string and assigns it to the RegistrationId field.
func (o *Device) SetRegistrationId(v string) {
	o.RegistrationId = &v
}

// GetRegistrationLifeTime returns the RegistrationLifeTime field value if set, zero value otherwise.
func (o *Device) GetRegistrationLifeTime() int64 {
	if o == nil || o.RegistrationLifeTime == nil {
		var ret int64
		return ret
	}
	return *o.RegistrationLifeTime
}

// GetRegistrationLifeTimeOk returns a tuple with the RegistrationLifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetRegistrationLifeTimeOk() (*int64, bool) {
	if o == nil || o.RegistrationLifeTime == nil {
		return nil, false
	}
	return o.RegistrationLifeTime, true
}

// HasRegistrationLifeTime returns a boolean if a field has been set.
func (o *Device) HasRegistrationLifeTime() bool {
	if o != nil && o.RegistrationLifeTime != nil {
		return true
	}

	return false
}

// SetRegistrationLifeTime gets a reference to the given int64 and assigns it to the RegistrationLifeTime field.
func (o *Device) SetRegistrationLifeTime(v int64) {
	o.RegistrationLifeTime = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Device) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Device) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Device) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Device) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Device) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.FirmwareVersion != nil {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.LastRegistrationUpdate != nil {
		toSerialize["lastRegistrationUpdate"] = o.LastRegistrationUpdate
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.ModelNumber != nil {
		toSerialize["modelNumber"] = o.ModelNumber
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.RegistrationId != nil {
		toSerialize["registrationId"] = o.RegistrationId
	}
	if o.RegistrationLifeTime != nil {
		toSerialize["registrationLifeTime"] = o.RegistrationLifeTime
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

