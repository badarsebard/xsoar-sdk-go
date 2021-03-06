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

// EngineInfo struct for EngineInfo
type EngineInfo struct {
	// Engine that will run the script
	Engine *string `json:"engine,omitempty"`
	// EngineGroup that will run the script
	EngineGroup *string `json:"engineGroup,omitempty"`
}

// NewEngineInfo instantiates a new EngineInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineInfo() *EngineInfo {
	this := EngineInfo{}
	return &this
}

// NewEngineInfoWithDefaults instantiates a new EngineInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineInfoWithDefaults() *EngineInfo {
	this := EngineInfo{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *EngineInfo) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineInfo) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *EngineInfo) HasEngine() bool {
	if o != nil && o.Engine != nil {
		return true
	}

	return false
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *EngineInfo) SetEngine(v string) {
	o.Engine = &v
}

// GetEngineGroup returns the EngineGroup field value if set, zero value otherwise.
func (o *EngineInfo) GetEngineGroup() string {
	if o == nil || o.EngineGroup == nil {
		var ret string
		return ret
	}
	return *o.EngineGroup
}

// GetEngineGroupOk returns a tuple with the EngineGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineInfo) GetEngineGroupOk() (*string, bool) {
	if o == nil || o.EngineGroup == nil {
		return nil, false
	}
	return o.EngineGroup, true
}

// HasEngineGroup returns a boolean if a field has been set.
func (o *EngineInfo) HasEngineGroup() bool {
	if o != nil && o.EngineGroup != nil {
		return true
	}

	return false
}

// SetEngineGroup gets a reference to the given string and assigns it to the EngineGroup field.
func (o *EngineInfo) SetEngineGroup(v string) {
	o.EngineGroup = &v
}

func (o EngineInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.EngineGroup != nil {
		toSerialize["engineGroup"] = o.EngineGroup
	}
	return json.Marshal(toSerialize)
}

type NullableEngineInfo struct {
	value *EngineInfo
	isSet bool
}

func (v NullableEngineInfo) Get() *EngineInfo {
	return v.value
}

func (v *NullableEngineInfo) Set(val *EngineInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineInfo(val *EngineInfo) *NullableEngineInfo {
	return &NullableEngineInfo{value: val, isSet: true}
}

func (v NullableEngineInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
