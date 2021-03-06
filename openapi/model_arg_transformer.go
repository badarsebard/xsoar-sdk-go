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

// ArgTransformer struct for ArgTransformer
type ArgTransformer struct {
	Args     *map[string]OperatorArgument `json:"args,omitempty"`
	Operator *string                      `json:"operator,omitempty"`
}

// NewArgTransformer instantiates a new ArgTransformer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgTransformer() *ArgTransformer {
	this := ArgTransformer{}
	return &this
}

// NewArgTransformerWithDefaults instantiates a new ArgTransformer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgTransformerWithDefaults() *ArgTransformer {
	this := ArgTransformer{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ArgTransformer) GetArgs() map[string]OperatorArgument {
	if o == nil || o.Args == nil {
		var ret map[string]OperatorArgument
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgTransformer) GetArgsOk() (*map[string]OperatorArgument, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ArgTransformer) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]OperatorArgument and assigns it to the Args field.
func (o *ArgTransformer) SetArgs(v map[string]OperatorArgument) {
	o.Args = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ArgTransformer) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgTransformer) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ArgTransformer) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ArgTransformer) SetOperator(v string) {
	o.Operator = &v
}

func (o ArgTransformer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	return json.Marshal(toSerialize)
}

type NullableArgTransformer struct {
	value *ArgTransformer
	isSet bool
}

func (v NullableArgTransformer) Get() *ArgTransformer {
	return v.value
}

func (v *NullableArgTransformer) Set(val *ArgTransformer) {
	v.value = val
	v.isSet = true
}

func (v NullableArgTransformer) IsSet() bool {
	return v.isSet
}

func (v *NullableArgTransformer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgTransformer(val *ArgTransformer) *NullableArgTransformer {
	return &NullableArgTransformer{value: val, isSet: true}
}

func (v NullableArgTransformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgTransformer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
