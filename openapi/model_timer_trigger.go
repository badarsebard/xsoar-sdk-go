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

// TimerTrigger struct for TimerTrigger
type TimerTrigger struct {
	Action    *string `json:"action,omitempty"`
	FieldName *string `json:"fieldName,omitempty"`
}

// NewTimerTrigger instantiates a new TimerTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerTrigger() *TimerTrigger {
	this := TimerTrigger{}
	return &this
}

// NewTimerTriggerWithDefaults instantiates a new TimerTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerTriggerWithDefaults() *TimerTrigger {
	this := TimerTrigger{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TimerTrigger) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTrigger) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TimerTrigger) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TimerTrigger) SetAction(v string) {
	o.Action = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *TimerTrigger) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTrigger) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *TimerTrigger) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *TimerTrigger) SetFieldName(v string) {
	o.FieldName = &v
}

func (o TimerTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}
	return json.Marshal(toSerialize)
}

type NullableTimerTrigger struct {
	value *TimerTrigger
	isSet bool
}

func (v NullableTimerTrigger) Get() *TimerTrigger {
	return v.value
}

func (v *NullableTimerTrigger) Set(val *TimerTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerTrigger(val *TimerTrigger) *NullableTimerTrigger {
	return &NullableTimerTrigger{value: val, isSet: true}
}

func (v NullableTimerTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
