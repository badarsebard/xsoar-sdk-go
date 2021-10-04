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

// ContentItemFields ContentItemFields holds ContentItem fields
type ContentItemFields struct {
	FromServerVersion     *Version  `json:"fromServerVersion,omitempty"`
	ItemVersion           *Version  `json:"itemVersion,omitempty"`
	PackID                *string   `json:"packID,omitempty"`
	PackPropagationLabels *[]string `json:"packPropagationLabels,omitempty"`
	PropagationLabels     *[]string `json:"propagationLabels,omitempty"`
	ToServerVersion       *Version  `json:"toServerVersion,omitempty"`
}

// NewContentItemFields instantiates a new ContentItemFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentItemFields() *ContentItemFields {
	this := ContentItemFields{}
	return &this
}

// NewContentItemFieldsWithDefaults instantiates a new ContentItemFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentItemFieldsWithDefaults() *ContentItemFields {
	this := ContentItemFields{}
	return &this
}

// GetFromServerVersion returns the FromServerVersion field value if set, zero value otherwise.
func (o *ContentItemFields) GetFromServerVersion() Version {
	if o == nil || o.FromServerVersion == nil {
		var ret Version
		return ret
	}
	return *o.FromServerVersion
}

// GetFromServerVersionOk returns a tuple with the FromServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetFromServerVersionOk() (*Version, bool) {
	if o == nil || o.FromServerVersion == nil {
		return nil, false
	}
	return o.FromServerVersion, true
}

// HasFromServerVersion returns a boolean if a field has been set.
func (o *ContentItemFields) HasFromServerVersion() bool {
	if o != nil && o.FromServerVersion != nil {
		return true
	}

	return false
}

// SetFromServerVersion gets a reference to the given Version and assigns it to the FromServerVersion field.
func (o *ContentItemFields) SetFromServerVersion(v Version) {
	o.FromServerVersion = &v
}

// GetItemVersion returns the ItemVersion field value if set, zero value otherwise.
func (o *ContentItemFields) GetItemVersion() Version {
	if o == nil || o.ItemVersion == nil {
		var ret Version
		return ret
	}
	return *o.ItemVersion
}

// GetItemVersionOk returns a tuple with the ItemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetItemVersionOk() (*Version, bool) {
	if o == nil || o.ItemVersion == nil {
		return nil, false
	}
	return o.ItemVersion, true
}

// HasItemVersion returns a boolean if a field has been set.
func (o *ContentItemFields) HasItemVersion() bool {
	if o != nil && o.ItemVersion != nil {
		return true
	}

	return false
}

// SetItemVersion gets a reference to the given Version and assigns it to the ItemVersion field.
func (o *ContentItemFields) SetItemVersion(v Version) {
	o.ItemVersion = &v
}

// GetPackID returns the PackID field value if set, zero value otherwise.
func (o *ContentItemFields) GetPackID() string {
	if o == nil || o.PackID == nil {
		var ret string
		return ret
	}
	return *o.PackID
}

// GetPackIDOk returns a tuple with the PackID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetPackIDOk() (*string, bool) {
	if o == nil || o.PackID == nil {
		return nil, false
	}
	return o.PackID, true
}

// HasPackID returns a boolean if a field has been set.
func (o *ContentItemFields) HasPackID() bool {
	if o != nil && o.PackID != nil {
		return true
	}

	return false
}

// SetPackID gets a reference to the given string and assigns it to the PackID field.
func (o *ContentItemFields) SetPackID(v string) {
	o.PackID = &v
}

// GetPackPropagationLabels returns the PackPropagationLabels field value if set, zero value otherwise.
func (o *ContentItemFields) GetPackPropagationLabels() []string {
	if o == nil || o.PackPropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PackPropagationLabels
}

// GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetPackPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PackPropagationLabels == nil {
		return nil, false
	}
	return o.PackPropagationLabels, true
}

// HasPackPropagationLabels returns a boolean if a field has been set.
func (o *ContentItemFields) HasPackPropagationLabels() bool {
	if o != nil && o.PackPropagationLabels != nil {
		return true
	}

	return false
}

// SetPackPropagationLabels gets a reference to the given []string and assigns it to the PackPropagationLabels field.
func (o *ContentItemFields) SetPackPropagationLabels(v []string) {
	o.PackPropagationLabels = &v
}

// GetPropagationLabels returns the PropagationLabels field value if set, zero value otherwise.
func (o *ContentItemFields) GetPropagationLabels() []string {
	if o == nil || o.PropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PropagationLabels
}

// GetPropagationLabelsOk returns a tuple with the PropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PropagationLabels == nil {
		return nil, false
	}
	return o.PropagationLabels, true
}

// HasPropagationLabels returns a boolean if a field has been set.
func (o *ContentItemFields) HasPropagationLabels() bool {
	if o != nil && o.PropagationLabels != nil {
		return true
	}

	return false
}

// SetPropagationLabels gets a reference to the given []string and assigns it to the PropagationLabels field.
func (o *ContentItemFields) SetPropagationLabels(v []string) {
	o.PropagationLabels = &v
}

// GetToServerVersion returns the ToServerVersion field value if set, zero value otherwise.
func (o *ContentItemFields) GetToServerVersion() Version {
	if o == nil || o.ToServerVersion == nil {
		var ret Version
		return ret
	}
	return *o.ToServerVersion
}

// GetToServerVersionOk returns a tuple with the ToServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemFields) GetToServerVersionOk() (*Version, bool) {
	if o == nil || o.ToServerVersion == nil {
		return nil, false
	}
	return o.ToServerVersion, true
}

// HasToServerVersion returns a boolean if a field has been set.
func (o *ContentItemFields) HasToServerVersion() bool {
	if o != nil && o.ToServerVersion != nil {
		return true
	}

	return false
}

// SetToServerVersion gets a reference to the given Version and assigns it to the ToServerVersion field.
func (o *ContentItemFields) SetToServerVersion(v Version) {
	o.ToServerVersion = &v
}

func (o ContentItemFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromServerVersion != nil {
		toSerialize["fromServerVersion"] = o.FromServerVersion
	}
	if o.ItemVersion != nil {
		toSerialize["itemVersion"] = o.ItemVersion
	}
	if o.PackID != nil {
		toSerialize["packID"] = o.PackID
	}
	if o.PackPropagationLabels != nil {
		toSerialize["packPropagationLabels"] = o.PackPropagationLabels
	}
	if o.PropagationLabels != nil {
		toSerialize["propagationLabels"] = o.PropagationLabels
	}
	if o.ToServerVersion != nil {
		toSerialize["toServerVersion"] = o.ToServerVersion
	}
	return json.Marshal(toSerialize)
}

type NullableContentItemFields struct {
	value *ContentItemFields
	isSet bool
}

func (v NullableContentItemFields) Get() *ContentItemFields {
	return v.value
}

func (v *NullableContentItemFields) Set(val *ContentItemFields) {
	v.value = val
	v.isSet = true
}

func (v NullableContentItemFields) IsSet() bool {
	return v.isSet
}

func (v *NullableContentItemFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentItemFields(val *ContentItemFields) *NullableContentItemFields {
	return &NullableContentItemFields{value: val, isSet: true}
}

func (v NullableContentItemFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentItemFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
