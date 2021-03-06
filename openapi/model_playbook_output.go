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

// PlaybookOutput PlaybookOutput represents the playbook output
type PlaybookOutput struct {
	ContextPath *string `json:"contextPath,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewPlaybookOutput instantiates a new PlaybookOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybookOutput() *PlaybookOutput {
	this := PlaybookOutput{}
	return &this
}

// NewPlaybookOutputWithDefaults instantiates a new PlaybookOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybookOutputWithDefaults() *PlaybookOutput {
	this := PlaybookOutput{}
	return &this
}

// GetContextPath returns the ContextPath field value if set, zero value otherwise.
func (o *PlaybookOutput) GetContextPath() string {
	if o == nil || o.ContextPath == nil {
		var ret string
		return ret
	}
	return *o.ContextPath
}

// GetContextPathOk returns a tuple with the ContextPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookOutput) GetContextPathOk() (*string, bool) {
	if o == nil || o.ContextPath == nil {
		return nil, false
	}
	return o.ContextPath, true
}

// HasContextPath returns a boolean if a field has been set.
func (o *PlaybookOutput) HasContextPath() bool {
	if o != nil && o.ContextPath != nil {
		return true
	}

	return false
}

// SetContextPath gets a reference to the given string and assigns it to the ContextPath field.
func (o *PlaybookOutput) SetContextPath(v string) {
	o.ContextPath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlaybookOutput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlaybookOutput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlaybookOutput) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlaybookOutput) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookOutput) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlaybookOutput) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlaybookOutput) SetType(v string) {
	o.Type = &v
}

func (o PlaybookOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContextPath != nil {
		toSerialize["contextPath"] = o.ContextPath
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePlaybookOutput struct {
	value *PlaybookOutput
	isSet bool
}

func (v NullablePlaybookOutput) Get() *PlaybookOutput {
	return v.value
}

func (v *NullablePlaybookOutput) Set(val *PlaybookOutput) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybookOutput) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybookOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybookOutput(val *PlaybookOutput) *NullablePlaybookOutput {
	return &NullablePlaybookOutput{value: val, isSet: true}
}

func (v NullablePlaybookOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybookOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
