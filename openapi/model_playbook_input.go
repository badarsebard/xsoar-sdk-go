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

// PlaybookInput PlaybookInput represents the playbook input
type PlaybookInput struct {
	Description        *string             `json:"description,omitempty"`
	Key                *string             `json:"key,omitempty"`
	PlaybookInputQuery *PlaybookInputQuery `json:"playbookInputQuery,omitempty"`
	Required           *bool               `json:"required,omitempty"`
	Value              *AdvanceArg         `json:"value,omitempty"`
}

// NewPlaybookInput instantiates a new PlaybookInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybookInput() *PlaybookInput {
	this := PlaybookInput{}
	return &this
}

// NewPlaybookInputWithDefaults instantiates a new PlaybookInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybookInputWithDefaults() *PlaybookInput {
	this := PlaybookInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlaybookInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlaybookInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlaybookInput) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PlaybookInput) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookInput) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PlaybookInput) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PlaybookInput) SetKey(v string) {
	o.Key = &v
}

// GetPlaybookInputQuery returns the PlaybookInputQuery field value if set, zero value otherwise.
func (o *PlaybookInput) GetPlaybookInputQuery() PlaybookInputQuery {
	if o == nil || o.PlaybookInputQuery == nil {
		var ret PlaybookInputQuery
		return ret
	}
	return *o.PlaybookInputQuery
}

// GetPlaybookInputQueryOk returns a tuple with the PlaybookInputQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookInput) GetPlaybookInputQueryOk() (*PlaybookInputQuery, bool) {
	if o == nil || o.PlaybookInputQuery == nil {
		return nil, false
	}
	return o.PlaybookInputQuery, true
}

// HasPlaybookInputQuery returns a boolean if a field has been set.
func (o *PlaybookInput) HasPlaybookInputQuery() bool {
	if o != nil && o.PlaybookInputQuery != nil {
		return true
	}

	return false
}

// SetPlaybookInputQuery gets a reference to the given PlaybookInputQuery and assigns it to the PlaybookInputQuery field.
func (o *PlaybookInput) SetPlaybookInputQuery(v PlaybookInputQuery) {
	o.PlaybookInputQuery = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PlaybookInput) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookInput) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PlaybookInput) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PlaybookInput) SetRequired(v bool) {
	o.Required = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PlaybookInput) GetValue() AdvanceArg {
	if o == nil || o.Value == nil {
		var ret AdvanceArg
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybookInput) GetValueOk() (*AdvanceArg, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlaybookInput) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AdvanceArg and assigns it to the Value field.
func (o *PlaybookInput) SetValue(v AdvanceArg) {
	o.Value = &v
}

func (o PlaybookInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.PlaybookInputQuery != nil {
		toSerialize["playbookInputQuery"] = o.PlaybookInputQuery
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePlaybookInput struct {
	value *PlaybookInput
	isSet bool
}

func (v NullablePlaybookInput) Get() *PlaybookInput {
	return v.value
}

func (v *NullablePlaybookInput) Set(val *PlaybookInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybookInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybookInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybookInput(val *PlaybookInput) *NullablePlaybookInput {
	return &NullablePlaybookInput{value: val, isSet: true}
}

func (v NullablePlaybookInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybookInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
