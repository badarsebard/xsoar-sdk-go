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

// ComplexArg ComplexArg - all info that is necessary to compute argument value from context Root - the root slice (or object) from to work against Filters - a slice of filters to apply to the root object, one after another (e.i. with AND condition between each one) Accessor - the key to access to each result after filter Transformers - a slice to transformers to apply on the result of the accessed key from each result e.g. if user want to take from context all File.DisplayName where File.Extension is 'EXE', and the result in uppercase than: Root: is \"File\" Filters: will hold the \"File.DisplayName where File.Extension is 'EXE'\" query Accessor: is \"DisplayName\" Transformers: will hold the uppercase transformation
type ComplexArg struct {
	Accessor     *string              `json:"accessor,omitempty"`
	Filters      *[][]ArgAtomicFilter `json:"filters,omitempty"`
	Root         *string              `json:"root,omitempty"`
	Transformers *[]ArgTransformer    `json:"transformers,omitempty"`
}

// NewComplexArg instantiates a new ComplexArg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplexArg() *ComplexArg {
	this := ComplexArg{}
	return &this
}

// NewComplexArgWithDefaults instantiates a new ComplexArg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplexArgWithDefaults() *ComplexArg {
	this := ComplexArg{}
	return &this
}

// GetAccessor returns the Accessor field value if set, zero value otherwise.
func (o *ComplexArg) GetAccessor() string {
	if o == nil || o.Accessor == nil {
		var ret string
		return ret
	}
	return *o.Accessor
}

// GetAccessorOk returns a tuple with the Accessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexArg) GetAccessorOk() (*string, bool) {
	if o == nil || o.Accessor == nil {
		return nil, false
	}
	return o.Accessor, true
}

// HasAccessor returns a boolean if a field has been set.
func (o *ComplexArg) HasAccessor() bool {
	if o != nil && o.Accessor != nil {
		return true
	}

	return false
}

// SetAccessor gets a reference to the given string and assigns it to the Accessor field.
func (o *ComplexArg) SetAccessor(v string) {
	o.Accessor = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ComplexArg) GetFilters() [][]ArgAtomicFilter {
	if o == nil || o.Filters == nil {
		var ret [][]ArgAtomicFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexArg) GetFiltersOk() (*[][]ArgAtomicFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ComplexArg) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given [][]ArgAtomicFilter and assigns it to the Filters field.
func (o *ComplexArg) SetFilters(v [][]ArgAtomicFilter) {
	o.Filters = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *ComplexArg) GetRoot() string {
	if o == nil || o.Root == nil {
		var ret string
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexArg) GetRootOk() (*string, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *ComplexArg) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given string and assigns it to the Root field.
func (o *ComplexArg) SetRoot(v string) {
	o.Root = &v
}

// GetTransformers returns the Transformers field value if set, zero value otherwise.
func (o *ComplexArg) GetTransformers() []ArgTransformer {
	if o == nil || o.Transformers == nil {
		var ret []ArgTransformer
		return ret
	}
	return *o.Transformers
}

// GetTransformersOk returns a tuple with the Transformers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexArg) GetTransformersOk() (*[]ArgTransformer, bool) {
	if o == nil || o.Transformers == nil {
		return nil, false
	}
	return o.Transformers, true
}

// HasTransformers returns a boolean if a field has been set.
func (o *ComplexArg) HasTransformers() bool {
	if o != nil && o.Transformers != nil {
		return true
	}

	return false
}

// SetTransformers gets a reference to the given []ArgTransformer and assigns it to the Transformers field.
func (o *ComplexArg) SetTransformers(v []ArgTransformer) {
	o.Transformers = &v
}

func (o ComplexArg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessor != nil {
		toSerialize["accessor"] = o.Accessor
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	if o.Transformers != nil {
		toSerialize["transformers"] = o.Transformers
	}
	return json.Marshal(toSerialize)
}

type NullableComplexArg struct {
	value *ComplexArg
	isSet bool
}

func (v NullableComplexArg) Get() *ComplexArg {
	return v.value
}

func (v *NullableComplexArg) Set(val *ComplexArg) {
	v.value = val
	v.isSet = true
}

func (v NullableComplexArg) IsSet() bool {
	return v.isSet
}

func (v *NullableComplexArg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplexArg(val *ComplexArg) *NullableComplexArg {
	return &NullableComplexArg{value: val, isSet: true}
}

func (v NullableComplexArg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplexArg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
