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

// StatsResponseWithReferenceLine struct for StatsResponseWithReferenceLine
type StatsResponseWithReferenceLine struct {
	// Groups is a list of group entities
	Groups         *[]Group `json:"groups,omitempty"`
	ReferenceLineY *float64 `json:"referenceLineY,omitempty"`
}

// NewStatsResponseWithReferenceLine instantiates a new StatsResponseWithReferenceLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsResponseWithReferenceLine() *StatsResponseWithReferenceLine {
	this := StatsResponseWithReferenceLine{}
	return &this
}

// NewStatsResponseWithReferenceLineWithDefaults instantiates a new StatsResponseWithReferenceLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsResponseWithReferenceLineWithDefaults() *StatsResponseWithReferenceLine {
	this := StatsResponseWithReferenceLine{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *StatsResponseWithReferenceLine) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsResponseWithReferenceLine) GetGroupsOk() (*[]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *StatsResponseWithReferenceLine) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *StatsResponseWithReferenceLine) SetGroups(v []Group) {
	o.Groups = &v
}

// GetReferenceLineY returns the ReferenceLineY field value if set, zero value otherwise.
func (o *StatsResponseWithReferenceLine) GetReferenceLineY() float64 {
	if o == nil || o.ReferenceLineY == nil {
		var ret float64
		return ret
	}
	return *o.ReferenceLineY
}

// GetReferenceLineYOk returns a tuple with the ReferenceLineY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsResponseWithReferenceLine) GetReferenceLineYOk() (*float64, bool) {
	if o == nil || o.ReferenceLineY == nil {
		return nil, false
	}
	return o.ReferenceLineY, true
}

// HasReferenceLineY returns a boolean if a field has been set.
func (o *StatsResponseWithReferenceLine) HasReferenceLineY() bool {
	if o != nil && o.ReferenceLineY != nil {
		return true
	}

	return false
}

// SetReferenceLineY gets a reference to the given float64 and assigns it to the ReferenceLineY field.
func (o *StatsResponseWithReferenceLine) SetReferenceLineY(v float64) {
	o.ReferenceLineY = &v
}

func (o StatsResponseWithReferenceLine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.ReferenceLineY != nil {
		toSerialize["referenceLineY"] = o.ReferenceLineY
	}
	return json.Marshal(toSerialize)
}

type NullableStatsResponseWithReferenceLine struct {
	value *StatsResponseWithReferenceLine
	isSet bool
}

func (v NullableStatsResponseWithReferenceLine) Get() *StatsResponseWithReferenceLine {
	return v.value
}

func (v *NullableStatsResponseWithReferenceLine) Set(val *StatsResponseWithReferenceLine) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsResponseWithReferenceLine) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsResponseWithReferenceLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsResponseWithReferenceLine(val *StatsResponseWithReferenceLine) *NullableStatsResponseWithReferenceLine {
	return &NullableStatsResponseWithReferenceLine{value: val, isSet: true}
}

func (v NullableStatsResponseWithReferenceLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsResponseWithReferenceLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
