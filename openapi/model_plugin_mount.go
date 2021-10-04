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

// PluginMount PluginMount plugin mount
type PluginMount struct {
	// description
	Description string `json:"Description"`
	// destination
	Destination string `json:"Destination"`
	// name
	Name string `json:"Name"`
	// options
	Options []string `json:"Options"`
	// settable
	Settable []string `json:"Settable"`
	// source
	Source string `json:"Source"`
	// type
	Type string `json:"Type"`
}

// NewPluginMount instantiates a new PluginMount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginMount(description string, destination string, name string, options []string, settable []string, source string, type_ string) *PluginMount {
	this := PluginMount{}
	this.Description = description
	this.Destination = destination
	this.Name = name
	this.Options = options
	this.Settable = settable
	this.Source = source
	this.Type = type_
	return &this
}

// NewPluginMountWithDefaults instantiates a new PluginMount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginMountWithDefaults() *PluginMount {
	this := PluginMount{}
	return &this
}

// GetDescription returns the Description field value
func (o *PluginMount) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginMount) SetDescription(v string) {
	o.Description = v
}

// GetDestination returns the Destination field value
func (o *PluginMount) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *PluginMount) SetDestination(v string) {
	o.Destination = v
}

// GetName returns the Name field value
func (o *PluginMount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PluginMount) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *PluginMount) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetOptionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *PluginMount) SetOptions(v []string) {
	o.Options = v
}

// GetSettable returns the Settable field value
func (o *PluginMount) GetSettable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetSettableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settable, true
}

// SetSettable sets field value
func (o *PluginMount) SetSettable(v []string) {
	o.Settable = v
}

// GetSource returns the Source field value
func (o *PluginMount) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PluginMount) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *PluginMount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PluginMount) SetType(v string) {
	o.Type = v
}

func (o PluginMount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Destination"] = o.Destination
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Options"] = o.Options
	}
	if true {
		toSerialize["Settable"] = o.Settable
	}
	if true {
		toSerialize["Source"] = o.Source
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePluginMount struct {
	value *PluginMount
	isSet bool
}

func (v NullablePluginMount) Get() *PluginMount {
	return v.value
}

func (v *NullablePluginMount) Set(val *PluginMount) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginMount) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginMount(val *PluginMount) *NullablePluginMount {
	return &NullablePluginMount{value: val, isSet: true}
}

func (v NullablePluginMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
