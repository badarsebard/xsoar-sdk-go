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

// TerminalOptions TerminalOptions - terminal options to use in case of using pty
type TerminalOptions struct {
	Echo           *int32  `json:"Echo,omitempty"`
	Terminal       *bool   `json:"Terminal,omitempty"`
	TerminalHeight *int64  `json:"TerminalHeight,omitempty"`
	TerminalType   *string `json:"TerminalType,omitempty"`
	TerminalWidth  *int64  `json:"TerminalWidth,omitempty"`
	TyISpeed       *int32  `json:"TyISpeed,omitempty"`
	TyOSpeed       *int32  `json:"TyOSpeed,omitempty"`
}

// NewTerminalOptions instantiates a new TerminalOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalOptions() *TerminalOptions {
	this := TerminalOptions{}
	return &this
}

// NewTerminalOptionsWithDefaults instantiates a new TerminalOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalOptionsWithDefaults() *TerminalOptions {
	this := TerminalOptions{}
	return &this
}

// GetEcho returns the Echo field value if set, zero value otherwise.
func (o *TerminalOptions) GetEcho() int32 {
	if o == nil || o.Echo == nil {
		var ret int32
		return ret
	}
	return *o.Echo
}

// GetEchoOk returns a tuple with the Echo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetEchoOk() (*int32, bool) {
	if o == nil || o.Echo == nil {
		return nil, false
	}
	return o.Echo, true
}

// HasEcho returns a boolean if a field has been set.
func (o *TerminalOptions) HasEcho() bool {
	if o != nil && o.Echo != nil {
		return true
	}

	return false
}

// SetEcho gets a reference to the given int32 and assigns it to the Echo field.
func (o *TerminalOptions) SetEcho(v int32) {
	o.Echo = &v
}

// GetTerminal returns the Terminal field value if set, zero value otherwise.
func (o *TerminalOptions) GetTerminal() bool {
	if o == nil || o.Terminal == nil {
		var ret bool
		return ret
	}
	return *o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTerminalOk() (*bool, bool) {
	if o == nil || o.Terminal == nil {
		return nil, false
	}
	return o.Terminal, true
}

// HasTerminal returns a boolean if a field has been set.
func (o *TerminalOptions) HasTerminal() bool {
	if o != nil && o.Terminal != nil {
		return true
	}

	return false
}

// SetTerminal gets a reference to the given bool and assigns it to the Terminal field.
func (o *TerminalOptions) SetTerminal(v bool) {
	o.Terminal = &v
}

// GetTerminalHeight returns the TerminalHeight field value if set, zero value otherwise.
func (o *TerminalOptions) GetTerminalHeight() int64 {
	if o == nil || o.TerminalHeight == nil {
		var ret int64
		return ret
	}
	return *o.TerminalHeight
}

// GetTerminalHeightOk returns a tuple with the TerminalHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTerminalHeightOk() (*int64, bool) {
	if o == nil || o.TerminalHeight == nil {
		return nil, false
	}
	return o.TerminalHeight, true
}

// HasTerminalHeight returns a boolean if a field has been set.
func (o *TerminalOptions) HasTerminalHeight() bool {
	if o != nil && o.TerminalHeight != nil {
		return true
	}

	return false
}

// SetTerminalHeight gets a reference to the given int64 and assigns it to the TerminalHeight field.
func (o *TerminalOptions) SetTerminalHeight(v int64) {
	o.TerminalHeight = &v
}

// GetTerminalType returns the TerminalType field value if set, zero value otherwise.
func (o *TerminalOptions) GetTerminalType() string {
	if o == nil || o.TerminalType == nil {
		var ret string
		return ret
	}
	return *o.TerminalType
}

// GetTerminalTypeOk returns a tuple with the TerminalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTerminalTypeOk() (*string, bool) {
	if o == nil || o.TerminalType == nil {
		return nil, false
	}
	return o.TerminalType, true
}

// HasTerminalType returns a boolean if a field has been set.
func (o *TerminalOptions) HasTerminalType() bool {
	if o != nil && o.TerminalType != nil {
		return true
	}

	return false
}

// SetTerminalType gets a reference to the given string and assigns it to the TerminalType field.
func (o *TerminalOptions) SetTerminalType(v string) {
	o.TerminalType = &v
}

// GetTerminalWidth returns the TerminalWidth field value if set, zero value otherwise.
func (o *TerminalOptions) GetTerminalWidth() int64 {
	if o == nil || o.TerminalWidth == nil {
		var ret int64
		return ret
	}
	return *o.TerminalWidth
}

// GetTerminalWidthOk returns a tuple with the TerminalWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTerminalWidthOk() (*int64, bool) {
	if o == nil || o.TerminalWidth == nil {
		return nil, false
	}
	return o.TerminalWidth, true
}

// HasTerminalWidth returns a boolean if a field has been set.
func (o *TerminalOptions) HasTerminalWidth() bool {
	if o != nil && o.TerminalWidth != nil {
		return true
	}

	return false
}

// SetTerminalWidth gets a reference to the given int64 and assigns it to the TerminalWidth field.
func (o *TerminalOptions) SetTerminalWidth(v int64) {
	o.TerminalWidth = &v
}

// GetTyISpeed returns the TyISpeed field value if set, zero value otherwise.
func (o *TerminalOptions) GetTyISpeed() int32 {
	if o == nil || o.TyISpeed == nil {
		var ret int32
		return ret
	}
	return *o.TyISpeed
}

// GetTyISpeedOk returns a tuple with the TyISpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTyISpeedOk() (*int32, bool) {
	if o == nil || o.TyISpeed == nil {
		return nil, false
	}
	return o.TyISpeed, true
}

// HasTyISpeed returns a boolean if a field has been set.
func (o *TerminalOptions) HasTyISpeed() bool {
	if o != nil && o.TyISpeed != nil {
		return true
	}

	return false
}

// SetTyISpeed gets a reference to the given int32 and assigns it to the TyISpeed field.
func (o *TerminalOptions) SetTyISpeed(v int32) {
	o.TyISpeed = &v
}

// GetTyOSpeed returns the TyOSpeed field value if set, zero value otherwise.
func (o *TerminalOptions) GetTyOSpeed() int32 {
	if o == nil || o.TyOSpeed == nil {
		var ret int32
		return ret
	}
	return *o.TyOSpeed
}

// GetTyOSpeedOk returns a tuple with the TyOSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOptions) GetTyOSpeedOk() (*int32, bool) {
	if o == nil || o.TyOSpeed == nil {
		return nil, false
	}
	return o.TyOSpeed, true
}

// HasTyOSpeed returns a boolean if a field has been set.
func (o *TerminalOptions) HasTyOSpeed() bool {
	if o != nil && o.TyOSpeed != nil {
		return true
	}

	return false
}

// SetTyOSpeed gets a reference to the given int32 and assigns it to the TyOSpeed field.
func (o *TerminalOptions) SetTyOSpeed(v int32) {
	o.TyOSpeed = &v
}

func (o TerminalOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Echo != nil {
		toSerialize["Echo"] = o.Echo
	}
	if o.Terminal != nil {
		toSerialize["Terminal"] = o.Terminal
	}
	if o.TerminalHeight != nil {
		toSerialize["TerminalHeight"] = o.TerminalHeight
	}
	if o.TerminalType != nil {
		toSerialize["TerminalType"] = o.TerminalType
	}
	if o.TerminalWidth != nil {
		toSerialize["TerminalWidth"] = o.TerminalWidth
	}
	if o.TyISpeed != nil {
		toSerialize["TyISpeed"] = o.TyISpeed
	}
	if o.TyOSpeed != nil {
		toSerialize["TyOSpeed"] = o.TyOSpeed
	}
	return json.Marshal(toSerialize)
}

type NullableTerminalOptions struct {
	value *TerminalOptions
	isSet bool
}

func (v NullableTerminalOptions) Get() *TerminalOptions {
	return v.value
}

func (v *NullableTerminalOptions) Set(val *TerminalOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalOptions(val *TerminalOptions) *NullableTerminalOptions {
	return &NullableTerminalOptions{value: val, isSet: true}
}

func (v NullableTerminalOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
