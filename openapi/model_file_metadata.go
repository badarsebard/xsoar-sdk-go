/*
Cortex XSOAR API

This is the public REST API to integrate with the Cortex XSOAR server. HTTP request can be sent using any HTTP-client.  For an example dedicated client take a look at: https://github.com/demisto/demisto-py.  Requests must include API-key that can be generated in the Cortex XSOAR web client under 'Settings' -> 'Integrations' -> 'API keys'   Optimistic Locking and Versioning\\:  When using Cortex XSOAR REST API, you will need to make sure to work on the latest version of the item (incident, entry, etc.), otherwise, you will get a DB version error (which not allow you to override a newer item). In addition, you can pass 'version\\: -1' to force data override (make sure that other users data might be lost).  Assume that Alice and Bob both read the same data from Cortex XSOAR server, then they both changed the data, and then both tried to write the new versions back to the server. Whose changes should be saved? Alice’s? Bob’s? To solve this, each data item in Cortex XSOAR has a numeric incremental version. If Alice saved an item with version 4 and Bob trying to save the same item with version 3, Cortex XSOAR will rollback Bob request and returns a DB version conflict error. Bob will need to get the latest item and work on it so Alice work will not get lost.  Example request using 'curl'\\:  ``` curl 'https://hostname:443/incidents/search' -H 'content-type: application/json' -H 'accept: application/json' -H 'Authorization: <API Key goes here>' --data-binary '{\"filter\":{\"query\":\"-status:closed -category:job\",\"period\":{\"by\":\"day\",\"fromValue\":7}}}' --compressed ```

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FileMetadata struct for FileMetadata
type FileMetadata struct {
	Info        *string `json:"info,omitempty"`
	IsMediaFile *bool   `json:"isMediaFile,omitempty"`
	Md5         *string `json:"md5,omitempty"`
	Sha1        *string `json:"sha1,omitempty"`
	Sha256      *string `json:"sha256,omitempty"`
	Sha512      *string `json:"sha512,omitempty"`
	Size        *int64  `json:"size,omitempty"`
	Ssdeep      *string `json:"ssdeep,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewFileMetadata instantiates a new FileMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileMetadata() *FileMetadata {
	this := FileMetadata{}
	return &this
}

// NewFileMetadataWithDefaults instantiates a new FileMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileMetadataWithDefaults() *FileMetadata {
	this := FileMetadata{}
	return &this
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *FileMetadata) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *FileMetadata) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *FileMetadata) SetInfo(v string) {
	o.Info = &v
}

// GetIsMediaFile returns the IsMediaFile field value if set, zero value otherwise.
func (o *FileMetadata) GetIsMediaFile() bool {
	if o == nil || o.IsMediaFile == nil {
		var ret bool
		return ret
	}
	return *o.IsMediaFile
}

// GetIsMediaFileOk returns a tuple with the IsMediaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetIsMediaFileOk() (*bool, bool) {
	if o == nil || o.IsMediaFile == nil {
		return nil, false
	}
	return o.IsMediaFile, true
}

// HasIsMediaFile returns a boolean if a field has been set.
func (o *FileMetadata) HasIsMediaFile() bool {
	if o != nil && o.IsMediaFile != nil {
		return true
	}

	return false
}

// SetIsMediaFile gets a reference to the given bool and assigns it to the IsMediaFile field.
func (o *FileMetadata) SetIsMediaFile(v bool) {
	o.IsMediaFile = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *FileMetadata) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *FileMetadata) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *FileMetadata) SetMd5(v string) {
	o.Md5 = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *FileMetadata) GetSha1() string {
	if o == nil || o.Sha1 == nil {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSha1Ok() (*string, bool) {
	if o == nil || o.Sha1 == nil {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *FileMetadata) HasSha1() bool {
	if o != nil && o.Sha1 != nil {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *FileMetadata) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *FileMetadata) GetSha256() string {
	if o == nil || o.Sha256 == nil {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSha256Ok() (*string, bool) {
	if o == nil || o.Sha256 == nil {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *FileMetadata) HasSha256() bool {
	if o != nil && o.Sha256 != nil {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *FileMetadata) SetSha256(v string) {
	o.Sha256 = &v
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise.
func (o *FileMetadata) GetSha512() string {
	if o == nil || o.Sha512 == nil {
		var ret string
		return ret
	}
	return *o.Sha512
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSha512Ok() (*string, bool) {
	if o == nil || o.Sha512 == nil {
		return nil, false
	}
	return o.Sha512, true
}

// HasSha512 returns a boolean if a field has been set.
func (o *FileMetadata) HasSha512() bool {
	if o != nil && o.Sha512 != nil {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given string and assigns it to the Sha512 field.
func (o *FileMetadata) SetSha512(v string) {
	o.Sha512 = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileMetadata) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileMetadata) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FileMetadata) SetSize(v int64) {
	o.Size = &v
}

// GetSsdeep returns the Ssdeep field value if set, zero value otherwise.
func (o *FileMetadata) GetSsdeep() string {
	if o == nil || o.Ssdeep == nil {
		var ret string
		return ret
	}
	return *o.Ssdeep
}

// GetSsdeepOk returns a tuple with the Ssdeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSsdeepOk() (*string, bool) {
	if o == nil || o.Ssdeep == nil {
		return nil, false
	}
	return o.Ssdeep, true
}

// HasSsdeep returns a boolean if a field has been set.
func (o *FileMetadata) HasSsdeep() bool {
	if o != nil && o.Ssdeep != nil {
		return true
	}

	return false
}

// SetSsdeep gets a reference to the given string and assigns it to the Ssdeep field.
func (o *FileMetadata) SetSsdeep(v string) {
	o.Ssdeep = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileMetadata) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileMetadata) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileMetadata) SetType(v string) {
	o.Type = &v
}

func (o FileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.IsMediaFile != nil {
		toSerialize["isMediaFile"] = o.IsMediaFile
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.Sha1 != nil {
		toSerialize["sha1"] = o.Sha1
	}
	if o.Sha256 != nil {
		toSerialize["sha256"] = o.Sha256
	}
	if o.Sha512 != nil {
		toSerialize["sha512"] = o.Sha512
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Ssdeep != nil {
		toSerialize["ssdeep"] = o.Ssdeep
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFileMetadata struct {
	value *FileMetadata
	isSet bool
}

func (v NullableFileMetadata) Get() *FileMetadata {
	return v.value
}

func (v *NullableFileMetadata) Set(val *FileMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFileMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFileMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileMetadata(val *FileMetadata) *NullableFileMetadata {
	return &NullableFileMetadata{value: val, isSet: true}
}

func (v NullableFileMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
