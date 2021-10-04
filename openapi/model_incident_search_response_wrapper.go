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

// IncidentSearchResponseWrapper IncidentSearchResponseWrapper is an extension for the IncidentSearchResponse type, which holds list of IncidentWrapper(s)
type IncidentSearchResponseWrapper struct {
	// in: body
	Data                *[]IncidentWrapper `json:"data,omitempty"`
	NotUpdated          *int32             `json:"notUpdated,omitempty"`
	SearchAfter         *[]string          `json:"searchAfter,omitempty"`
	SearchAfterElastic  *[]string          `json:"searchAfterElastic,omitempty"`
	SearchBefore        *[]string          `json:"searchBefore,omitempty"`
	SearchBeforeElastic *[]string          `json:"searchBeforeElastic,omitempty"`
	Total               *int64             `json:"total,omitempty"`
}

// NewIncidentSearchResponseWrapper instantiates a new IncidentSearchResponseWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentSearchResponseWrapper() *IncidentSearchResponseWrapper {
	this := IncidentSearchResponseWrapper{}
	return &this
}

// NewIncidentSearchResponseWrapperWithDefaults instantiates a new IncidentSearchResponseWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentSearchResponseWrapperWithDefaults() *IncidentSearchResponseWrapper {
	this := IncidentSearchResponseWrapper{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetData() []IncidentWrapper {
	if o == nil || o.Data == nil {
		var ret []IncidentWrapper
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetDataOk() (*[]IncidentWrapper, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IncidentWrapper and assigns it to the Data field.
func (o *IncidentSearchResponseWrapper) SetData(v []IncidentWrapper) {
	o.Data = &v
}

// GetNotUpdated returns the NotUpdated field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetNotUpdated() int32 {
	if o == nil || o.NotUpdated == nil {
		var ret int32
		return ret
	}
	return *o.NotUpdated
}

// GetNotUpdatedOk returns a tuple with the NotUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetNotUpdatedOk() (*int32, bool) {
	if o == nil || o.NotUpdated == nil {
		return nil, false
	}
	return o.NotUpdated, true
}

// HasNotUpdated returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasNotUpdated() bool {
	if o != nil && o.NotUpdated != nil {
		return true
	}

	return false
}

// SetNotUpdated gets a reference to the given int32 and assigns it to the NotUpdated field.
func (o *IncidentSearchResponseWrapper) SetNotUpdated(v int32) {
	o.NotUpdated = &v
}

// GetSearchAfter returns the SearchAfter field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetSearchAfter() []string {
	if o == nil || o.SearchAfter == nil {
		var ret []string
		return ret
	}
	return *o.SearchAfter
}

// GetSearchAfterOk returns a tuple with the SearchAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetSearchAfterOk() (*[]string, bool) {
	if o == nil || o.SearchAfter == nil {
		return nil, false
	}
	return o.SearchAfter, true
}

// HasSearchAfter returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasSearchAfter() bool {
	if o != nil && o.SearchAfter != nil {
		return true
	}

	return false
}

// SetSearchAfter gets a reference to the given []string and assigns it to the SearchAfter field.
func (o *IncidentSearchResponseWrapper) SetSearchAfter(v []string) {
	o.SearchAfter = &v
}

// GetSearchAfterElastic returns the SearchAfterElastic field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetSearchAfterElastic() []string {
	if o == nil || o.SearchAfterElastic == nil {
		var ret []string
		return ret
	}
	return *o.SearchAfterElastic
}

// GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetSearchAfterElasticOk() (*[]string, bool) {
	if o == nil || o.SearchAfterElastic == nil {
		return nil, false
	}
	return o.SearchAfterElastic, true
}

// HasSearchAfterElastic returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasSearchAfterElastic() bool {
	if o != nil && o.SearchAfterElastic != nil {
		return true
	}

	return false
}

// SetSearchAfterElastic gets a reference to the given []string and assigns it to the SearchAfterElastic field.
func (o *IncidentSearchResponseWrapper) SetSearchAfterElastic(v []string) {
	o.SearchAfterElastic = &v
}

// GetSearchBefore returns the SearchBefore field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetSearchBefore() []string {
	if o == nil || o.SearchBefore == nil {
		var ret []string
		return ret
	}
	return *o.SearchBefore
}

// GetSearchBeforeOk returns a tuple with the SearchBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetSearchBeforeOk() (*[]string, bool) {
	if o == nil || o.SearchBefore == nil {
		return nil, false
	}
	return o.SearchBefore, true
}

// HasSearchBefore returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasSearchBefore() bool {
	if o != nil && o.SearchBefore != nil {
		return true
	}

	return false
}

// SetSearchBefore gets a reference to the given []string and assigns it to the SearchBefore field.
func (o *IncidentSearchResponseWrapper) SetSearchBefore(v []string) {
	o.SearchBefore = &v
}

// GetSearchBeforeElastic returns the SearchBeforeElastic field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetSearchBeforeElastic() []string {
	if o == nil || o.SearchBeforeElastic == nil {
		var ret []string
		return ret
	}
	return *o.SearchBeforeElastic
}

// GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetSearchBeforeElasticOk() (*[]string, bool) {
	if o == nil || o.SearchBeforeElastic == nil {
		return nil, false
	}
	return o.SearchBeforeElastic, true
}

// HasSearchBeforeElastic returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasSearchBeforeElastic() bool {
	if o != nil && o.SearchBeforeElastic != nil {
		return true
	}

	return false
}

// SetSearchBeforeElastic gets a reference to the given []string and assigns it to the SearchBeforeElastic field.
func (o *IncidentSearchResponseWrapper) SetSearchBeforeElastic(v []string) {
	o.SearchBeforeElastic = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *IncidentSearchResponseWrapper) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseWrapper) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *IncidentSearchResponseWrapper) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *IncidentSearchResponseWrapper) SetTotal(v int64) {
	o.Total = &v
}

func (o IncidentSearchResponseWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.NotUpdated != nil {
		toSerialize["notUpdated"] = o.NotUpdated
	}
	if o.SearchAfter != nil {
		toSerialize["searchAfter"] = o.SearchAfter
	}
	if o.SearchAfterElastic != nil {
		toSerialize["searchAfterElastic"] = o.SearchAfterElastic
	}
	if o.SearchBefore != nil {
		toSerialize["searchBefore"] = o.SearchBefore
	}
	if o.SearchBeforeElastic != nil {
		toSerialize["searchBeforeElastic"] = o.SearchBeforeElastic
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentSearchResponseWrapper struct {
	value *IncidentSearchResponseWrapper
	isSet bool
}

func (v NullableIncidentSearchResponseWrapper) Get() *IncidentSearchResponseWrapper {
	return v.value
}

func (v *NullableIncidentSearchResponseWrapper) Set(val *IncidentSearchResponseWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentSearchResponseWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentSearchResponseWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentSearchResponseWrapper(val *IncidentSearchResponseWrapper) *NullableIncidentSearchResponseWrapper {
	return &NullableIncidentSearchResponseWrapper{value: val, isSet: true}
}

func (v NullableIncidentSearchResponseWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentSearchResponseWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
