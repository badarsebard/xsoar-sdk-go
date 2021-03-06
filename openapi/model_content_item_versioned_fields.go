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

// ContentItemVersionedFields struct for ContentItemVersionedFields
type ContentItemVersionedFields struct {
	CommitMessage                     *string   `json:"commitMessage,omitempty"`
	FromServerVersion                 *string   `json:"fromServerVersion,omitempty"`
	ItemVersion                       *string   `json:"itemVersion,omitempty"`
	PackID                            *string   `json:"packID,omitempty"`
	PackPropagationLabels             *[]string `json:"packPropagationLabels,omitempty"`
	PropagationLabels                 *[]string `json:"propagationLabels,omitempty"`
	ShouldCommit                      *bool     `json:"shouldCommit,omitempty"`
	ToServerVersion                   *string   `json:"toServerVersion,omitempty"`
	VcShouldIgnore                    *bool     `json:"vcShouldIgnore,omitempty"`
	VcShouldKeepItemLegacyProdMachine *bool     `json:"vcShouldKeepItemLegacyProdMachine,omitempty"`
}

// NewContentItemVersionedFields instantiates a new ContentItemVersionedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentItemVersionedFields() *ContentItemVersionedFields {
	this := ContentItemVersionedFields{}
	return &this
}

// NewContentItemVersionedFieldsWithDefaults instantiates a new ContentItemVersionedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentItemVersionedFieldsWithDefaults() *ContentItemVersionedFields {
	this := ContentItemVersionedFields{}
	return &this
}

// GetCommitMessage returns the CommitMessage field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetCommitMessage() string {
	if o == nil || o.CommitMessage == nil {
		var ret string
		return ret
	}
	return *o.CommitMessage
}

// GetCommitMessageOk returns a tuple with the CommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetCommitMessageOk() (*string, bool) {
	if o == nil || o.CommitMessage == nil {
		return nil, false
	}
	return o.CommitMessage, true
}

// HasCommitMessage returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasCommitMessage() bool {
	if o != nil && o.CommitMessage != nil {
		return true
	}

	return false
}

// SetCommitMessage gets a reference to the given string and assigns it to the CommitMessage field.
func (o *ContentItemVersionedFields) SetCommitMessage(v string) {
	o.CommitMessage = &v
}

// GetFromServerVersion returns the FromServerVersion field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetFromServerVersion() string {
	if o == nil || o.FromServerVersion == nil {
		var ret string
		return ret
	}
	return *o.FromServerVersion
}

// GetFromServerVersionOk returns a tuple with the FromServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetFromServerVersionOk() (*string, bool) {
	if o == nil || o.FromServerVersion == nil {
		return nil, false
	}
	return o.FromServerVersion, true
}

// HasFromServerVersion returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasFromServerVersion() bool {
	if o != nil && o.FromServerVersion != nil {
		return true
	}

	return false
}

// SetFromServerVersion gets a reference to the given string and assigns it to the FromServerVersion field.
func (o *ContentItemVersionedFields) SetFromServerVersion(v string) {
	o.FromServerVersion = &v
}

// GetItemVersion returns the ItemVersion field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetItemVersion() string {
	if o == nil || o.ItemVersion == nil {
		var ret string
		return ret
	}
	return *o.ItemVersion
}

// GetItemVersionOk returns a tuple with the ItemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetItemVersionOk() (*string, bool) {
	if o == nil || o.ItemVersion == nil {
		return nil, false
	}
	return o.ItemVersion, true
}

// HasItemVersion returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasItemVersion() bool {
	if o != nil && o.ItemVersion != nil {
		return true
	}

	return false
}

// SetItemVersion gets a reference to the given string and assigns it to the ItemVersion field.
func (o *ContentItemVersionedFields) SetItemVersion(v string) {
	o.ItemVersion = &v
}

// GetPackID returns the PackID field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetPackID() string {
	if o == nil || o.PackID == nil {
		var ret string
		return ret
	}
	return *o.PackID
}

// GetPackIDOk returns a tuple with the PackID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetPackIDOk() (*string, bool) {
	if o == nil || o.PackID == nil {
		return nil, false
	}
	return o.PackID, true
}

// HasPackID returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasPackID() bool {
	if o != nil && o.PackID != nil {
		return true
	}

	return false
}

// SetPackID gets a reference to the given string and assigns it to the PackID field.
func (o *ContentItemVersionedFields) SetPackID(v string) {
	o.PackID = &v
}

// GetPackPropagationLabels returns the PackPropagationLabels field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetPackPropagationLabels() []string {
	if o == nil || o.PackPropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PackPropagationLabels
}

// GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetPackPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PackPropagationLabels == nil {
		return nil, false
	}
	return o.PackPropagationLabels, true
}

// HasPackPropagationLabels returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasPackPropagationLabels() bool {
	if o != nil && o.PackPropagationLabels != nil {
		return true
	}

	return false
}

// SetPackPropagationLabels gets a reference to the given []string and assigns it to the PackPropagationLabels field.
func (o *ContentItemVersionedFields) SetPackPropagationLabels(v []string) {
	o.PackPropagationLabels = &v
}

// GetPropagationLabels returns the PropagationLabels field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetPropagationLabels() []string {
	if o == nil || o.PropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PropagationLabels
}

// GetPropagationLabelsOk returns a tuple with the PropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PropagationLabels == nil {
		return nil, false
	}
	return o.PropagationLabels, true
}

// HasPropagationLabels returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasPropagationLabels() bool {
	if o != nil && o.PropagationLabels != nil {
		return true
	}

	return false
}

// SetPropagationLabels gets a reference to the given []string and assigns it to the PropagationLabels field.
func (o *ContentItemVersionedFields) SetPropagationLabels(v []string) {
	o.PropagationLabels = &v
}

// GetShouldCommit returns the ShouldCommit field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetShouldCommit() bool {
	if o == nil || o.ShouldCommit == nil {
		var ret bool
		return ret
	}
	return *o.ShouldCommit
}

// GetShouldCommitOk returns a tuple with the ShouldCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetShouldCommitOk() (*bool, bool) {
	if o == nil || o.ShouldCommit == nil {
		return nil, false
	}
	return o.ShouldCommit, true
}

// HasShouldCommit returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasShouldCommit() bool {
	if o != nil && o.ShouldCommit != nil {
		return true
	}

	return false
}

// SetShouldCommit gets a reference to the given bool and assigns it to the ShouldCommit field.
func (o *ContentItemVersionedFields) SetShouldCommit(v bool) {
	o.ShouldCommit = &v
}

// GetToServerVersion returns the ToServerVersion field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetToServerVersion() string {
	if o == nil || o.ToServerVersion == nil {
		var ret string
		return ret
	}
	return *o.ToServerVersion
}

// GetToServerVersionOk returns a tuple with the ToServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetToServerVersionOk() (*string, bool) {
	if o == nil || o.ToServerVersion == nil {
		return nil, false
	}
	return o.ToServerVersion, true
}

// HasToServerVersion returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasToServerVersion() bool {
	if o != nil && o.ToServerVersion != nil {
		return true
	}

	return false
}

// SetToServerVersion gets a reference to the given string and assigns it to the ToServerVersion field.
func (o *ContentItemVersionedFields) SetToServerVersion(v string) {
	o.ToServerVersion = &v
}

// GetVcShouldIgnore returns the VcShouldIgnore field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetVcShouldIgnore() bool {
	if o == nil || o.VcShouldIgnore == nil {
		var ret bool
		return ret
	}
	return *o.VcShouldIgnore
}

// GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetVcShouldIgnoreOk() (*bool, bool) {
	if o == nil || o.VcShouldIgnore == nil {
		return nil, false
	}
	return o.VcShouldIgnore, true
}

// HasVcShouldIgnore returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasVcShouldIgnore() bool {
	if o != nil && o.VcShouldIgnore != nil {
		return true
	}

	return false
}

// SetVcShouldIgnore gets a reference to the given bool and assigns it to the VcShouldIgnore field.
func (o *ContentItemVersionedFields) SetVcShouldIgnore(v bool) {
	o.VcShouldIgnore = &v
}

// GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field value if set, zero value otherwise.
func (o *ContentItemVersionedFields) GetVcShouldKeepItemLegacyProdMachine() bool {
	if o == nil || o.VcShouldKeepItemLegacyProdMachine == nil {
		var ret bool
		return ret
	}
	return *o.VcShouldKeepItemLegacyProdMachine
}

// GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentItemVersionedFields) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool) {
	if o == nil || o.VcShouldKeepItemLegacyProdMachine == nil {
		return nil, false
	}
	return o.VcShouldKeepItemLegacyProdMachine, true
}

// HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.
func (o *ContentItemVersionedFields) HasVcShouldKeepItemLegacyProdMachine() bool {
	if o != nil && o.VcShouldKeepItemLegacyProdMachine != nil {
		return true
	}

	return false
}

// SetVcShouldKeepItemLegacyProdMachine gets a reference to the given bool and assigns it to the VcShouldKeepItemLegacyProdMachine field.
func (o *ContentItemVersionedFields) SetVcShouldKeepItemLegacyProdMachine(v bool) {
	o.VcShouldKeepItemLegacyProdMachine = &v
}

func (o ContentItemVersionedFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommitMessage != nil {
		toSerialize["commitMessage"] = o.CommitMessage
	}
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
	if o.ShouldCommit != nil {
		toSerialize["shouldCommit"] = o.ShouldCommit
	}
	if o.ToServerVersion != nil {
		toSerialize["toServerVersion"] = o.ToServerVersion
	}
	if o.VcShouldIgnore != nil {
		toSerialize["vcShouldIgnore"] = o.VcShouldIgnore
	}
	if o.VcShouldKeepItemLegacyProdMachine != nil {
		toSerialize["vcShouldKeepItemLegacyProdMachine"] = o.VcShouldKeepItemLegacyProdMachine
	}
	return json.Marshal(toSerialize)
}

type NullableContentItemVersionedFields struct {
	value *ContentItemVersionedFields
	isSet bool
}

func (v NullableContentItemVersionedFields) Get() *ContentItemVersionedFields {
	return v.value
}

func (v *NullableContentItemVersionedFields) Set(val *ContentItemVersionedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableContentItemVersionedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableContentItemVersionedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentItemVersionedFields(val *ContentItemVersionedFields) *NullableContentItemVersionedFields {
	return &NullableContentItemVersionedFields{value: val, isSet: true}
}

func (v NullableContentItemVersionedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentItemVersionedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
