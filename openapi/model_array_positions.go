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

// ArrayPositions struct for ArrayPositions
type ArrayPositions struct {
	Items []int32
}

// NewArrayPositions instantiates a new ArrayPositions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayPositions() *ArrayPositions {
	this := ArrayPositions{}
	return &this
}

// NewArrayPositionsWithDefaults instantiates a new ArrayPositions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayPositionsWithDefaults() *ArrayPositions {
	this := ArrayPositions{}
	return &this
}

func (o ArrayPositions) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *ArrayPositions) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableArrayPositions struct {
	value *ArrayPositions
	isSet bool
}

func (v NullableArrayPositions) Get() *ArrayPositions {
	return v.value
}

func (v *NullableArrayPositions) Set(val *ArrayPositions) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayPositions) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayPositions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayPositions(val *ArrayPositions) *NullableArrayPositions {
	return &NullableArrayPositions{value: val, isSet: true}
}

func (v NullableArrayPositions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayPositions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
