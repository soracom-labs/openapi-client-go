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

// FileEntry struct for FileEntry
type FileEntry struct {
	// Content length of the file
	ContentLength *int64 `json:"contentLength,omitempty"`
	// Content type of the file
	ContentType *string `json:"contentType,omitempty"`
	// Created time of the file
	CreatedTime *int64 `json:"createdTime,omitempty"`
	// Parent directory name
	Directory *string `json:"directory,omitempty"`
	// Etag of the file
	Etag *string `json:"etag,omitempty"`
	// Absolute path of the file
	FilePath *string `json:"filePath,omitempty"`
	// File name
	Filename *string `json:"filename,omitempty"`
	// Whether the entry is directory or not
	IsDirectory *bool `json:"isDirectory,omitempty"`
	// Last modified time of the file
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`
}

// NewFileEntry instantiates a new FileEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileEntry() *FileEntry {
	this := FileEntry{}
	return &this
}

// NewFileEntryWithDefaults instantiates a new FileEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileEntryWithDefaults() *FileEntry {
	this := FileEntry{}
	return &this
}

// GetContentLength returns the ContentLength field value if set, zero value otherwise.
func (o *FileEntry) GetContentLength() int64 {
	if o == nil || o.ContentLength == nil {
		var ret int64
		return ret
	}
	return *o.ContentLength
}

// GetContentLengthOk returns a tuple with the ContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetContentLengthOk() (*int64, bool) {
	if o == nil || o.ContentLength == nil {
		return nil, false
	}
	return o.ContentLength, true
}

// HasContentLength returns a boolean if a field has been set.
func (o *FileEntry) HasContentLength() bool {
	if o != nil && o.ContentLength != nil {
		return true
	}

	return false
}

// SetContentLength gets a reference to the given int64 and assigns it to the ContentLength field.
func (o *FileEntry) SetContentLength(v int64) {
	o.ContentLength = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *FileEntry) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *FileEntry) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *FileEntry) SetContentType(v string) {
	o.ContentType = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *FileEntry) GetCreatedTime() int64 {
	if o == nil || o.CreatedTime == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetCreatedTimeOk() (*int64, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *FileEntry) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int64 and assigns it to the CreatedTime field.
func (o *FileEntry) SetCreatedTime(v int64) {
	o.CreatedTime = &v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *FileEntry) GetDirectory() string {
	if o == nil || o.Directory == nil {
		var ret string
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetDirectoryOk() (*string, bool) {
	if o == nil || o.Directory == nil {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *FileEntry) HasDirectory() bool {
	if o != nil && o.Directory != nil {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given string and assigns it to the Directory field.
func (o *FileEntry) SetDirectory(v string) {
	o.Directory = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *FileEntry) GetEtag() string {
	if o == nil || o.Etag == nil {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetEtagOk() (*string, bool) {
	if o == nil || o.Etag == nil {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *FileEntry) HasEtag() bool {
	if o != nil && o.Etag != nil {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *FileEntry) SetEtag(v string) {
	o.Etag = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *FileEntry) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *FileEntry) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *FileEntry) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FileEntry) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FileEntry) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FileEntry) SetFilename(v string) {
	o.Filename = &v
}

// GetIsDirectory returns the IsDirectory field value if set, zero value otherwise.
func (o *FileEntry) GetIsDirectory() bool {
	if o == nil || o.IsDirectory == nil {
		var ret bool
		return ret
	}
	return *o.IsDirectory
}

// GetIsDirectoryOk returns a tuple with the IsDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetIsDirectoryOk() (*bool, bool) {
	if o == nil || o.IsDirectory == nil {
		return nil, false
	}
	return o.IsDirectory, true
}

// HasIsDirectory returns a boolean if a field has been set.
func (o *FileEntry) HasIsDirectory() bool {
	if o != nil && o.IsDirectory != nil {
		return true
	}

	return false
}

// SetIsDirectory gets a reference to the given bool and assigns it to the IsDirectory field.
func (o *FileEntry) SetIsDirectory(v bool) {
	o.IsDirectory = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *FileEntry) GetLastModifiedTime() int64 {
	if o == nil || o.LastModifiedTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetLastModifiedTimeOk() (*int64, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *FileEntry) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given int64 and assigns it to the LastModifiedTime field.
func (o *FileEntry) SetLastModifiedTime(v int64) {
	o.LastModifiedTime = &v
}

func (o FileEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentLength != nil {
		toSerialize["contentLength"] = o.ContentLength
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.Directory != nil {
		toSerialize["directory"] = o.Directory
	}
	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.IsDirectory != nil {
		toSerialize["isDirectory"] = o.IsDirectory
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableFileEntry struct {
	value *FileEntry
	isSet bool
}

func (v NullableFileEntry) Get() *FileEntry {
	return v.value
}

func (v *NullableFileEntry) Set(val *FileEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFileEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFileEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileEntry(val *FileEntry) *NullableFileEntry {
	return &NullableFileEntry{value: val, isSet: true}
}

func (v NullableFileEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


