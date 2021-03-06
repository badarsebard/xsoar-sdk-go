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

// StatsQueryResponseWithError struct for StatsQueryResponseWithError
type StatsQueryResponseWithError struct {
	StatsQueryError    *map[string]interface{} `json:"StatsQueryError,omitempty"`
	StatsQueryResponse *StatsQueryResponse     `json:"statsQueryResponse,omitempty"`
}

// NewStatsQueryResponseWithError instantiates a new StatsQueryResponseWithError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsQueryResponseWithError() *StatsQueryResponseWithError {
	this := StatsQueryResponseWithError{}
	return &this
}

// NewStatsQueryResponseWithErrorWithDefaults instantiates a new StatsQueryResponseWithError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsQueryResponseWithErrorWithDefaults() *StatsQueryResponseWithError {
	this := StatsQueryResponseWithError{}
	return &this
}

// GetStatsQueryError returns the StatsQueryError field value if set, zero value otherwise.
func (o *StatsQueryResponseWithError) GetStatsQueryError() map[string]interface{} {
	if o == nil || o.StatsQueryError == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StatsQueryError
}

// GetStatsQueryErrorOk returns a tuple with the StatsQueryError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsQueryResponseWithError) GetStatsQueryErrorOk() (*map[string]interface{}, bool) {
	if o == nil || o.StatsQueryError == nil {
		return nil, false
	}
	return o.StatsQueryError, true
}

// HasStatsQueryError returns a boolean if a field has been set.
func (o *StatsQueryResponseWithError) HasStatsQueryError() bool {
	if o != nil && o.StatsQueryError != nil {
		return true
	}

	return false
}

// SetStatsQueryError gets a reference to the given map[string]interface{} and assigns it to the StatsQueryError field.
func (o *StatsQueryResponseWithError) SetStatsQueryError(v map[string]interface{}) {
	o.StatsQueryError = &v
}

// GetStatsQueryResponse returns the StatsQueryResponse field value if set, zero value otherwise.
func (o *StatsQueryResponseWithError) GetStatsQueryResponse() StatsQueryResponse {
	if o == nil || o.StatsQueryResponse == nil {
		var ret StatsQueryResponse
		return ret
	}
	return *o.StatsQueryResponse
}

// GetStatsQueryResponseOk returns a tuple with the StatsQueryResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsQueryResponseWithError) GetStatsQueryResponseOk() (*StatsQueryResponse, bool) {
	if o == nil || o.StatsQueryResponse == nil {
		return nil, false
	}
	return o.StatsQueryResponse, true
}

// HasStatsQueryResponse returns a boolean if a field has been set.
func (o *StatsQueryResponseWithError) HasStatsQueryResponse() bool {
	if o != nil && o.StatsQueryResponse != nil {
		return true
	}

	return false
}

// SetStatsQueryResponse gets a reference to the given StatsQueryResponse and assigns it to the StatsQueryResponse field.
func (o *StatsQueryResponseWithError) SetStatsQueryResponse(v StatsQueryResponse) {
	o.StatsQueryResponse = &v
}

func (o StatsQueryResponseWithError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatsQueryError != nil {
		toSerialize["StatsQueryError"] = o.StatsQueryError
	}
	if o.StatsQueryResponse != nil {
		toSerialize["statsQueryResponse"] = o.StatsQueryResponse
	}
	return json.Marshal(toSerialize)
}

type NullableStatsQueryResponseWithError struct {
	value *StatsQueryResponseWithError
	isSet bool
}

func (v NullableStatsQueryResponseWithError) Get() *StatsQueryResponseWithError {
	return v.value
}

func (v *NullableStatsQueryResponseWithError) Set(val *StatsQueryResponseWithError) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsQueryResponseWithError) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsQueryResponseWithError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsQueryResponseWithError(val *StatsQueryResponseWithError) *NullableStatsQueryResponseWithError {
	return &NullableStatsQueryResponseWithError{value: val, isSet: true}
}

func (v NullableStatsQueryResponseWithError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsQueryResponseWithError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
