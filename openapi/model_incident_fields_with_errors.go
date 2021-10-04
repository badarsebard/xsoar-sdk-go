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

// IncidentFieldsWithErrors struct for IncidentFieldsWithErrors
type IncidentFieldsWithErrors struct {
	Error          *string          `json:"error,omitempty"`
	IncidentFields *[]IncidentField `json:"incidentFields,omitempty"`
}

// NewIncidentFieldsWithErrors instantiates a new IncidentFieldsWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentFieldsWithErrors() *IncidentFieldsWithErrors {
	this := IncidentFieldsWithErrors{}
	return &this
}

// NewIncidentFieldsWithErrorsWithDefaults instantiates a new IncidentFieldsWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentFieldsWithErrorsWithDefaults() *IncidentFieldsWithErrors {
	this := IncidentFieldsWithErrors{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IncidentFieldsWithErrors) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldsWithErrors) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IncidentFieldsWithErrors) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *IncidentFieldsWithErrors) SetError(v string) {
	o.Error = &v
}

// GetIncidentFields returns the IncidentFields field value if set, zero value otherwise.
func (o *IncidentFieldsWithErrors) GetIncidentFields() []IncidentField {
	if o == nil || o.IncidentFields == nil {
		var ret []IncidentField
		return ret
	}
	return *o.IncidentFields
}

// GetIncidentFieldsOk returns a tuple with the IncidentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldsWithErrors) GetIncidentFieldsOk() (*[]IncidentField, bool) {
	if o == nil || o.IncidentFields == nil {
		return nil, false
	}
	return o.IncidentFields, true
}

// HasIncidentFields returns a boolean if a field has been set.
func (o *IncidentFieldsWithErrors) HasIncidentFields() bool {
	if o != nil && o.IncidentFields != nil {
		return true
	}

	return false
}

// SetIncidentFields gets a reference to the given []IncidentField and assigns it to the IncidentFields field.
func (o *IncidentFieldsWithErrors) SetIncidentFields(v []IncidentField) {
	o.IncidentFields = &v
}

func (o IncidentFieldsWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.IncidentFields != nil {
		toSerialize["incidentFields"] = o.IncidentFields
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentFieldsWithErrors struct {
	value *IncidentFieldsWithErrors
	isSet bool
}

func (v NullableIncidentFieldsWithErrors) Get() *IncidentFieldsWithErrors {
	return v.value
}

func (v *NullableIncidentFieldsWithErrors) Set(val *IncidentFieldsWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFieldsWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFieldsWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFieldsWithErrors(val *IncidentFieldsWithErrors) *NullableIncidentFieldsWithErrors {
	return &NullableIncidentFieldsWithErrors{value: val, isSet: true}
}

func (v NullableIncidentFieldsWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFieldsWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
