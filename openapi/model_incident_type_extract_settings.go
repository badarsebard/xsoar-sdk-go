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

// IncidentTypeExtractSettings struct for IncidentTypeExtractSettings
type IncidentTypeExtractSettings struct {
	FieldCliNameToExtractSettings *map[string]FieldExtractSetting `json:"fieldCliNameToExtractSettings,omitempty"`
	Mode                          *string                         `json:"mode,omitempty"`
}

// NewIncidentTypeExtractSettings instantiates a new IncidentTypeExtractSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTypeExtractSettings() *IncidentTypeExtractSettings {
	this := IncidentTypeExtractSettings{}
	return &this
}

// NewIncidentTypeExtractSettingsWithDefaults instantiates a new IncidentTypeExtractSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTypeExtractSettingsWithDefaults() *IncidentTypeExtractSettings {
	this := IncidentTypeExtractSettings{}
	return &this
}

// GetFieldCliNameToExtractSettings returns the FieldCliNameToExtractSettings field value if set, zero value otherwise.
func (o *IncidentTypeExtractSettings) GetFieldCliNameToExtractSettings() map[string]FieldExtractSetting {
	if o == nil || o.FieldCliNameToExtractSettings == nil {
		var ret map[string]FieldExtractSetting
		return ret
	}
	return *o.FieldCliNameToExtractSettings
}

// GetFieldCliNameToExtractSettingsOk returns a tuple with the FieldCliNameToExtractSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeExtractSettings) GetFieldCliNameToExtractSettingsOk() (*map[string]FieldExtractSetting, bool) {
	if o == nil || o.FieldCliNameToExtractSettings == nil {
		return nil, false
	}
	return o.FieldCliNameToExtractSettings, true
}

// HasFieldCliNameToExtractSettings returns a boolean if a field has been set.
func (o *IncidentTypeExtractSettings) HasFieldCliNameToExtractSettings() bool {
	if o != nil && o.FieldCliNameToExtractSettings != nil {
		return true
	}

	return false
}

// SetFieldCliNameToExtractSettings gets a reference to the given map[string]FieldExtractSetting and assigns it to the FieldCliNameToExtractSettings field.
func (o *IncidentTypeExtractSettings) SetFieldCliNameToExtractSettings(v map[string]FieldExtractSetting) {
	o.FieldCliNameToExtractSettings = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IncidentTypeExtractSettings) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeExtractSettings) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IncidentTypeExtractSettings) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *IncidentTypeExtractSettings) SetMode(v string) {
	o.Mode = &v
}

func (o IncidentTypeExtractSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldCliNameToExtractSettings != nil {
		toSerialize["fieldCliNameToExtractSettings"] = o.FieldCliNameToExtractSettings
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTypeExtractSettings struct {
	value *IncidentTypeExtractSettings
	isSet bool
}

func (v NullableIncidentTypeExtractSettings) Get() *IncidentTypeExtractSettings {
	return v.value
}

func (v *NullableIncidentTypeExtractSettings) Set(val *IncidentTypeExtractSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTypeExtractSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTypeExtractSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTypeExtractSettings(val *IncidentTypeExtractSettings) *NullableIncidentTypeExtractSettings {
	return &NullableIncidentTypeExtractSettings{value: val, isSet: true}
}

func (v NullableIncidentTypeExtractSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTypeExtractSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
