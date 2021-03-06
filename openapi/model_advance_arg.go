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

// AdvanceArg Simple: plain text such as \"hello\" (constant) or ${File.ID} (DT expression) Complex: struct with all info that is necessary to compute argument value from context (Root, Filters, Accessor & Transformers)
type AdvanceArg struct {
	Complex *ComplexArg `json:"complex,omitempty"`
	Simple  *string     `json:"simple,omitempty"`
}

// NewAdvanceArg instantiates a new AdvanceArg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvanceArg() *AdvanceArg {
	this := AdvanceArg{}
	return &this
}

// NewAdvanceArgWithDefaults instantiates a new AdvanceArg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvanceArgWithDefaults() *AdvanceArg {
	this := AdvanceArg{}
	return &this
}

// GetComplex returns the Complex field value if set, zero value otherwise.
func (o *AdvanceArg) GetComplex() ComplexArg {
	if o == nil || o.Complex == nil {
		var ret ComplexArg
		return ret
	}
	return *o.Complex
}

// GetComplexOk returns a tuple with the Complex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceArg) GetComplexOk() (*ComplexArg, bool) {
	if o == nil || o.Complex == nil {
		return nil, false
	}
	return o.Complex, true
}

// HasComplex returns a boolean if a field has been set.
func (o *AdvanceArg) HasComplex() bool {
	if o != nil && o.Complex != nil {
		return true
	}

	return false
}

// SetComplex gets a reference to the given ComplexArg and assigns it to the Complex field.
func (o *AdvanceArg) SetComplex(v ComplexArg) {
	o.Complex = &v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *AdvanceArg) GetSimple() string {
	if o == nil || o.Simple == nil {
		var ret string
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceArg) GetSimpleOk() (*string, bool) {
	if o == nil || o.Simple == nil {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *AdvanceArg) HasSimple() bool {
	if o != nil && o.Simple != nil {
		return true
	}

	return false
}

// SetSimple gets a reference to the given string and assigns it to the Simple field.
func (o *AdvanceArg) SetSimple(v string) {
	o.Simple = &v
}

func (o AdvanceArg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Complex != nil {
		toSerialize["complex"] = o.Complex
	}
	if o.Simple != nil {
		toSerialize["simple"] = o.Simple
	}
	return json.Marshal(toSerialize)
}

type NullableAdvanceArg struct {
	value *AdvanceArg
	isSet bool
}

func (v NullableAdvanceArg) Get() *AdvanceArg {
	return v.value
}

func (v *NullableAdvanceArg) Set(val *AdvanceArg) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvanceArg) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvanceArg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvanceArg(val *AdvanceArg) *NullableAdvanceArg {
	return &NullableAdvanceArg{value: val, isSet: true}
}

func (v NullableAdvanceArg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvanceArg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
