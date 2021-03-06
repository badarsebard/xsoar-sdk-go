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

// ContainersInfo Info - holds all containers info
type ContainersInfo struct {
	All      *int64 `json:"all,omitempty"`
	Inactive *int64 `json:"inactive,omitempty"`
	Running  *int64 `json:"running,omitempty"`
}

// NewContainersInfo instantiates a new ContainersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainersInfo() *ContainersInfo {
	this := ContainersInfo{}
	return &this
}

// NewContainersInfoWithDefaults instantiates a new ContainersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainersInfoWithDefaults() *ContainersInfo {
	this := ContainersInfo{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *ContainersInfo) GetAll() int64 {
	if o == nil || o.All == nil {
		var ret int64
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersInfo) GetAllOk() (*int64, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *ContainersInfo) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given int64 and assigns it to the All field.
func (o *ContainersInfo) SetAll(v int64) {
	o.All = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *ContainersInfo) GetInactive() int64 {
	if o == nil || o.Inactive == nil {
		var ret int64
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersInfo) GetInactiveOk() (*int64, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *ContainersInfo) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given int64 and assigns it to the Inactive field.
func (o *ContainersInfo) SetInactive(v int64) {
	o.Inactive = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *ContainersInfo) GetRunning() int64 {
	if o == nil || o.Running == nil {
		var ret int64
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersInfo) GetRunningOk() (*int64, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *ContainersInfo) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int64 and assigns it to the Running field.
func (o *ContainersInfo) SetRunning(v int64) {
	o.Running = &v
}

func (o ContainersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	if o.Running != nil {
		toSerialize["running"] = o.Running
	}
	return json.Marshal(toSerialize)
}

type NullableContainersInfo struct {
	value *ContainersInfo
	isSet bool
}

func (v NullableContainersInfo) Get() *ContainersInfo {
	return v.value
}

func (v *NullableContainersInfo) Set(val *ContainersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableContainersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableContainersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainersInfo(val *ContainersInfo) *NullableContainersInfo {
	return &NullableContainersInfo{value: val, isSet: true}
}

func (v NullableContainersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
