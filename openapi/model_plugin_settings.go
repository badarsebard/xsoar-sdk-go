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

// PluginSettings struct for PluginSettings
type PluginSettings struct {
	// args
	Args []string `json:"Args"`
	// devices
	Devices []PluginDevice `json:"Devices"`
	// env
	Env []string `json:"Env"`
	// mounts
	Mounts []PluginMount `json:"Mounts"`
}

// NewPluginSettings instantiates a new PluginSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginSettings(args []string, devices []PluginDevice, env []string, mounts []PluginMount) *PluginSettings {
	this := PluginSettings{}
	this.Args = args
	this.Devices = devices
	this.Env = env
	this.Mounts = mounts
	return &this
}

// NewPluginSettingsWithDefaults instantiates a new PluginSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginSettingsWithDefaults() *PluginSettings {
	this := PluginSettings{}
	return &this
}

// GetArgs returns the Args field value
func (o *PluginSettings) GetArgs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetArgsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *PluginSettings) SetArgs(v []string) {
	o.Args = v
}

// GetDevices returns the Devices field value
func (o *PluginSettings) GetDevices() []PluginDevice {
	if o == nil {
		var ret []PluginDevice
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetDevicesOk() (*[]PluginDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *PluginSettings) SetDevices(v []PluginDevice) {
	o.Devices = v
}

// GetEnv returns the Env field value
func (o *PluginSettings) GetEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetEnvOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value
func (o *PluginSettings) SetEnv(v []string) {
	o.Env = v
}

// GetMounts returns the Mounts field value
func (o *PluginSettings) GetMounts() []PluginMount {
	if o == nil {
		var ret []PluginMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetMountsOk() (*[]PluginMount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mounts, true
}

// SetMounts sets field value
func (o *PluginSettings) SetMounts(v []PluginMount) {
	o.Mounts = v
}

func (o PluginSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Args"] = o.Args
	}
	if true {
		toSerialize["Devices"] = o.Devices
	}
	if true {
		toSerialize["Env"] = o.Env
	}
	if true {
		toSerialize["Mounts"] = o.Mounts
	}
	return json.Marshal(toSerialize)
}

type NullablePluginSettings struct {
	value *PluginSettings
	isSet bool
}

func (v NullablePluginSettings) Get() *PluginSettings {
	return v.value
}

func (v *NullablePluginSettings) Set(val *PluginSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginSettings(val *PluginSettings) *NullablePluginSettings {
	return &NullablePluginSettings{value: val, isSet: true}
}

func (v NullablePluginSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
