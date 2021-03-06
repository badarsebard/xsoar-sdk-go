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

// ContainerChangeResponseItem ContainerChangeResponseItem change item in response to ContainerChanges operation
type ContainerChangeResponseItem struct {
	// Kind of change
	Kind int32 `json:"Kind"`
	// Path to file that has changed
	Path string `json:"Path"`
}

// NewContainerChangeResponseItem instantiates a new ContainerChangeResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerChangeResponseItem(kind int32, path string) *ContainerChangeResponseItem {
	this := ContainerChangeResponseItem{}
	this.Kind = kind
	this.Path = path
	return &this
}

// NewContainerChangeResponseItemWithDefaults instantiates a new ContainerChangeResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerChangeResponseItemWithDefaults() *ContainerChangeResponseItem {
	this := ContainerChangeResponseItem{}
	return &this
}

// GetKind returns the Kind field value
func (o *ContainerChangeResponseItem) GetKind() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ContainerChangeResponseItem) GetKindOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ContainerChangeResponseItem) SetKind(v int32) {
	o.Kind = v
}

// GetPath returns the Path field value
func (o *ContainerChangeResponseItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ContainerChangeResponseItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ContainerChangeResponseItem) SetPath(v string) {
	o.Path = v
}

func (o ContainerChangeResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Kind"] = o.Kind
	}
	if true {
		toSerialize["Path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableContainerChangeResponseItem struct {
	value *ContainerChangeResponseItem
	isSet bool
}

func (v NullableContainerChangeResponseItem) Get() *ContainerChangeResponseItem {
	return v.value
}

func (v *NullableContainerChangeResponseItem) Set(val *ContainerChangeResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerChangeResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerChangeResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerChangeResponseItem(val *ContainerChangeResponseItem) *NullableContainerChangeResponseItem {
	return &NullableContainerChangeResponseItem{value: val, isSet: true}
}

func (v NullableContainerChangeResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerChangeResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
