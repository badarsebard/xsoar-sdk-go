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

// EvidenceData EvidenceData - all evidence properties to evaluate in task process
type EvidenceData struct {
	// This field must have empty json key
	CustomFields *map[string]AdvanceArg `json:"customFields,omitempty"`
	Description  *AdvanceArg            `json:"description,omitempty"`
	Occurred     *AdvanceArg            `json:"occurred,omitempty"`
	Tags         *AdvanceArg            `json:"tags,omitempty"`
}

// NewEvidenceData instantiates a new EvidenceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvidenceData() *EvidenceData {
	this := EvidenceData{}
	return &this
}

// NewEvidenceDataWithDefaults instantiates a new EvidenceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvidenceDataWithDefaults() *EvidenceData {
	this := EvidenceData{}
	return &this
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EvidenceData) GetCustomFields() map[string]AdvanceArg {
	if o == nil || o.CustomFields == nil {
		var ret map[string]AdvanceArg
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvidenceData) GetCustomFieldsOk() (*map[string]AdvanceArg, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EvidenceData) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]AdvanceArg and assigns it to the CustomFields field.
func (o *EvidenceData) SetCustomFields(v map[string]AdvanceArg) {
	o.CustomFields = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EvidenceData) GetDescription() AdvanceArg {
	if o == nil || o.Description == nil {
		var ret AdvanceArg
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvidenceData) GetDescriptionOk() (*AdvanceArg, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EvidenceData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AdvanceArg and assigns it to the Description field.
func (o *EvidenceData) SetDescription(v AdvanceArg) {
	o.Description = &v
}

// GetOccurred returns the Occurred field value if set, zero value otherwise.
func (o *EvidenceData) GetOccurred() AdvanceArg {
	if o == nil || o.Occurred == nil {
		var ret AdvanceArg
		return ret
	}
	return *o.Occurred
}

// GetOccurredOk returns a tuple with the Occurred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvidenceData) GetOccurredOk() (*AdvanceArg, bool) {
	if o == nil || o.Occurred == nil {
		return nil, false
	}
	return o.Occurred, true
}

// HasOccurred returns a boolean if a field has been set.
func (o *EvidenceData) HasOccurred() bool {
	if o != nil && o.Occurred != nil {
		return true
	}

	return false
}

// SetOccurred gets a reference to the given AdvanceArg and assigns it to the Occurred field.
func (o *EvidenceData) SetOccurred(v AdvanceArg) {
	o.Occurred = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EvidenceData) GetTags() AdvanceArg {
	if o == nil || o.Tags == nil {
		var ret AdvanceArg
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvidenceData) GetTagsOk() (*AdvanceArg, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EvidenceData) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given AdvanceArg and assigns it to the Tags field.
func (o *EvidenceData) SetTags(v AdvanceArg) {
	o.Tags = &v
}

func (o EvidenceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomFields != nil {
		toSerialize["customFields"] = o.CustomFields
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Occurred != nil {
		toSerialize["occurred"] = o.Occurred
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableEvidenceData struct {
	value *EvidenceData
	isSet bool
}

func (v NullableEvidenceData) Get() *EvidenceData {
	return v.value
}

func (v *NullableEvidenceData) Set(val *EvidenceData) {
	v.value = val
	v.isSet = true
}

func (v NullableEvidenceData) IsSet() bool {
	return v.isSet
}

func (v *NullableEvidenceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvidenceData(val *EvidenceData) *NullableEvidenceData {
	return &NullableEvidenceData{value: val, isSet: true}
}

func (v NullableEvidenceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvidenceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
