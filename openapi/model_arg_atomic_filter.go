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

// ArgAtomicFilter ArgAtomicFilter - operator with two sides that return true/false
type ArgAtomicFilter struct {
	IgnoreCase *bool             `json:"ignoreCase,omitempty"`
	Left       *OperatorArgument `json:"left,omitempty"`
	Operator   *string           `json:"operator,omitempty"`
	Right      *OperatorArgument `json:"right,omitempty"`
	Type       *string           `json:"type,omitempty"`
}

// NewArgAtomicFilter instantiates a new ArgAtomicFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgAtomicFilter() *ArgAtomicFilter {
	this := ArgAtomicFilter{}
	return &this
}

// NewArgAtomicFilterWithDefaults instantiates a new ArgAtomicFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgAtomicFilterWithDefaults() *ArgAtomicFilter {
	this := ArgAtomicFilter{}
	return &this
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *ArgAtomicFilter) GetIgnoreCase() bool {
	if o == nil || o.IgnoreCase == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgAtomicFilter) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || o.IgnoreCase == nil {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *ArgAtomicFilter) HasIgnoreCase() bool {
	if o != nil && o.IgnoreCase != nil {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *ArgAtomicFilter) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *ArgAtomicFilter) GetLeft() OperatorArgument {
	if o == nil || o.Left == nil {
		var ret OperatorArgument
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgAtomicFilter) GetLeftOk() (*OperatorArgument, bool) {
	if o == nil || o.Left == nil {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *ArgAtomicFilter) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given OperatorArgument and assigns it to the Left field.
func (o *ArgAtomicFilter) SetLeft(v OperatorArgument) {
	o.Left = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ArgAtomicFilter) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgAtomicFilter) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ArgAtomicFilter) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ArgAtomicFilter) SetOperator(v string) {
	o.Operator = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *ArgAtomicFilter) GetRight() OperatorArgument {
	if o == nil || o.Right == nil {
		var ret OperatorArgument
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgAtomicFilter) GetRightOk() (*OperatorArgument, bool) {
	if o == nil || o.Right == nil {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *ArgAtomicFilter) HasRight() bool {
	if o != nil && o.Right != nil {
		return true
	}

	return false
}

// SetRight gets a reference to the given OperatorArgument and assigns it to the Right field.
func (o *ArgAtomicFilter) SetRight(v OperatorArgument) {
	o.Right = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ArgAtomicFilter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgAtomicFilter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ArgAtomicFilter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ArgAtomicFilter) SetType(v string) {
	o.Type = &v
}

func (o ArgAtomicFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IgnoreCase != nil {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	if o.Left != nil {
		toSerialize["left"] = o.Left
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Right != nil {
		toSerialize["right"] = o.Right
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableArgAtomicFilter struct {
	value *ArgAtomicFilter
	isSet bool
}

func (v NullableArgAtomicFilter) Get() *ArgAtomicFilter {
	return v.value
}

func (v *NullableArgAtomicFilter) Set(val *ArgAtomicFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableArgAtomicFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableArgAtomicFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgAtomicFilter(val *ArgAtomicFilter) *NullableArgAtomicFilter {
	return &NullableArgAtomicFilter{value: val, isSet: true}
}

func (v NullableArgAtomicFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgAtomicFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
