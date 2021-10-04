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

// OperatorArgument struct for OperatorArgument
type OperatorArgument struct {
	IsContext *bool       `json:"isContext,omitempty"`
	Value     *AdvanceArg `json:"value,omitempty"`
}

// NewOperatorArgument instantiates a new OperatorArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorArgument() *OperatorArgument {
	this := OperatorArgument{}
	return &this
}

// NewOperatorArgumentWithDefaults instantiates a new OperatorArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorArgumentWithDefaults() *OperatorArgument {
	this := OperatorArgument{}
	return &this
}

// GetIsContext returns the IsContext field value if set, zero value otherwise.
func (o *OperatorArgument) GetIsContext() bool {
	if o == nil || o.IsContext == nil {
		var ret bool
		return ret
	}
	return *o.IsContext
}

// GetIsContextOk returns a tuple with the IsContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorArgument) GetIsContextOk() (*bool, bool) {
	if o == nil || o.IsContext == nil {
		return nil, false
	}
	return o.IsContext, true
}

// HasIsContext returns a boolean if a field has been set.
func (o *OperatorArgument) HasIsContext() bool {
	if o != nil && o.IsContext != nil {
		return true
	}

	return false
}

// SetIsContext gets a reference to the given bool and assigns it to the IsContext field.
func (o *OperatorArgument) SetIsContext(v bool) {
	o.IsContext = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OperatorArgument) GetValue() AdvanceArg {
	if o == nil || o.Value == nil {
		var ret AdvanceArg
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorArgument) GetValueOk() (*AdvanceArg, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OperatorArgument) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AdvanceArg and assigns it to the Value field.
func (o *OperatorArgument) SetValue(v AdvanceArg) {
	o.Value = &v
}

func (o OperatorArgument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsContext != nil {
		toSerialize["isContext"] = o.IsContext
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorArgument struct {
	value *OperatorArgument
	isSet bool
}

func (v NullableOperatorArgument) Get() *OperatorArgument {
	return v.value
}

func (v *NullableOperatorArgument) Set(val *OperatorArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorArgument(val *OperatorArgument) *NullableOperatorArgument {
	return &NullableOperatorArgument{value: val, isSet: true}
}

func (v NullableOperatorArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
