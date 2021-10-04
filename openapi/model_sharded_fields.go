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

// ShardedFields struct for ShardedFields
type ShardedFields struct {
	ShardID *int64 `json:"ShardID,omitempty"`
}

// NewShardedFields instantiates a new ShardedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardedFields() *ShardedFields {
	this := ShardedFields{}
	return &this
}

// NewShardedFieldsWithDefaults instantiates a new ShardedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardedFieldsWithDefaults() *ShardedFields {
	this := ShardedFields{}
	return &this
}

// GetShardID returns the ShardID field value if set, zero value otherwise.
func (o *ShardedFields) GetShardID() int64 {
	if o == nil || o.ShardID == nil {
		var ret int64
		return ret
	}
	return *o.ShardID
}

// GetShardIDOk returns a tuple with the ShardID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardedFields) GetShardIDOk() (*int64, bool) {
	if o == nil || o.ShardID == nil {
		return nil, false
	}
	return o.ShardID, true
}

// HasShardID returns a boolean if a field has been set.
func (o *ShardedFields) HasShardID() bool {
	if o != nil && o.ShardID != nil {
		return true
	}

	return false
}

// SetShardID gets a reference to the given int64 and assigns it to the ShardID field.
func (o *ShardedFields) SetShardID(v int64) {
	o.ShardID = &v
}

func (o ShardedFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShardID != nil {
		toSerialize["ShardID"] = o.ShardID
	}
	return json.Marshal(toSerialize)
}

type NullableShardedFields struct {
	value *ShardedFields
	isSet bool
}

func (v NullableShardedFields) Get() *ShardedFields {
	return v.value
}

func (v *NullableShardedFields) Set(val *ShardedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableShardedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableShardedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShardedFields(val *ShardedFields) *NullableShardedFields {
	return &NullableShardedFields{value: val, isSet: true}
}

func (v NullableShardedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShardedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
