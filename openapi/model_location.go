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

// Location struct for Location
type Location struct {
	ArrayPositions *ArrayPositions `json:"array_positions,omitempty"`
	End            *int32          `json:"end,omitempty"`
	// Pos is the position of the term within the field, starting at 1
	Pos *int32 `json:"pos,omitempty"`
	// Start and End are the byte offsets of the term in the field
	Start *int32 `json:"start,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetArrayPositions returns the ArrayPositions field value if set, zero value otherwise.
func (o *Location) GetArrayPositions() ArrayPositions {
	if o == nil || o.ArrayPositions == nil {
		var ret ArrayPositions
		return ret
	}
	return *o.ArrayPositions
}

// GetArrayPositionsOk returns a tuple with the ArrayPositions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetArrayPositionsOk() (*ArrayPositions, bool) {
	if o == nil || o.ArrayPositions == nil {
		return nil, false
	}
	return o.ArrayPositions, true
}

// HasArrayPositions returns a boolean if a field has been set.
func (o *Location) HasArrayPositions() bool {
	if o != nil && o.ArrayPositions != nil {
		return true
	}

	return false
}

// SetArrayPositions gets a reference to the given ArrayPositions and assigns it to the ArrayPositions field.
func (o *Location) SetArrayPositions(v ArrayPositions) {
	o.ArrayPositions = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Location) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Location) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *Location) SetEnd(v int32) {
	o.End = &v
}

// GetPos returns the Pos field value if set, zero value otherwise.
func (o *Location) GetPos() int32 {
	if o == nil || o.Pos == nil {
		var ret int32
		return ret
	}
	return *o.Pos
}

// GetPosOk returns a tuple with the Pos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPosOk() (*int32, bool) {
	if o == nil || o.Pos == nil {
		return nil, false
	}
	return o.Pos, true
}

// HasPos returns a boolean if a field has been set.
func (o *Location) HasPos() bool {
	if o != nil && o.Pos != nil {
		return true
	}

	return false
}

// SetPos gets a reference to the given int32 and assigns it to the Pos field.
func (o *Location) SetPos(v int32) {
	o.Pos = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Location) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Location) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *Location) SetStart(v int32) {
	o.Start = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArrayPositions != nil {
		toSerialize["array_positions"] = o.ArrayPositions
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Pos != nil {
		toSerialize["pos"] = o.Pos
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
