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

// StatsTextResponse struct for StatsTextResponse
type StatsTextResponse struct {
	Groups *Groups `json:"groups,omitempty"`
	// Describe the complete text for the text widget, after placeholders injection.
	Text *string `json:"text,omitempty"`
}

// NewStatsTextResponse instantiates a new StatsTextResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsTextResponse() *StatsTextResponse {
	this := StatsTextResponse{}
	return &this
}

// NewStatsTextResponseWithDefaults instantiates a new StatsTextResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsTextResponseWithDefaults() *StatsTextResponse {
	this := StatsTextResponse{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *StatsTextResponse) GetGroups() Groups {
	if o == nil || o.Groups == nil {
		var ret Groups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsTextResponse) GetGroupsOk() (*Groups, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *StatsTextResponse) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given Groups and assigns it to the Groups field.
func (o *StatsTextResponse) SetGroups(v Groups) {
	o.Groups = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *StatsTextResponse) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsTextResponse) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *StatsTextResponse) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *StatsTextResponse) SetText(v string) {
	o.Text = &v
}

func (o StatsTextResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableStatsTextResponse struct {
	value *StatsTextResponse
	isSet bool
}

func (v NullableStatsTextResponse) Get() *StatsTextResponse {
	return v.value
}

func (v *NullableStatsTextResponse) Set(val *StatsTextResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsTextResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsTextResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsTextResponse(val *StatsTextResponse) *NullableStatsTextResponse {
	return &NullableStatsTextResponse{value: val, isSet: true}
}

func (v NullableStatsTextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsTextResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
