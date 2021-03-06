/*
Cortex XSOAR API

This is the public REST API to integrate with the Cortex XSOAR server. HTTP request can be sent using any HTTP-client.  For an example dedicated client take a look at: https://github.com/demisto/demisto-py.  Requests must include API-key that can be generated in the Cortex XSOAR web client under 'Settings' -> 'Integrations' -> 'API keys'   Optimistic Locking and Versioning\\:  When using Cortex XSOAR REST API, you will need to make sure to work on the latest version of the item (incident, entry, etc.), otherwise, you will get a DB version error (which not allow you to override a newer item). In addition, you can pass 'version\\: -1' to force data override (make sure that other users data might be lost).  Assume that Alice and Bob both read the same data from Cortex XSOAR server, then they both changed the data, and then both tried to write the new versions back to the server. Whose changes should be saved? Alice’s? Bob’s? To solve this, each data item in Cortex XSOAR has a numeric incremental version. If Alice saved an item with version 4 and Bob trying to save the same item with version 3, Cortex XSOAR will rollback Bob request and returns a DB version conflict error. Bob will need to get the latest item and work on it so Alice work will not get lost.  Example request using 'curl'\\:  ``` curl 'https://hostname:443/incidents/search' -H 'content-type: application/json' -H 'accept: application/json' -H 'Authorization: <API Key goes here>' --data-binary '{\"filter\":{\"query\":\"-status:closed -category:job\",\"period\":{\"by\":\"day\",\"fromValue\":7}}}' --compressed ```

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LayoutSection struct for LayoutSection
type LayoutSection struct {
	Description    *string                 `json:"description,omitempty"`
	Fields         *[]LayoutField          `json:"fields,omitempty"`
	Highlight      *map[string][]string    `json:"highlight,omitempty"`
	Id             *string                 `json:"id,omitempty"`
	IsVisible      *bool                   `json:"isVisible,omitempty"`
	Modified       *time.Time              `json:"modified,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	NumericId      *int64                  `json:"numericId,omitempty"`
	PrimaryTerm    *int64                  `json:"primaryTerm,omitempty"`
	Query          *map[string]interface{} `json:"query,omitempty"`
	QueryType      *string                 `json:"queryType,omitempty"`
	ReadOnly       *bool                   `json:"readOnly,omitempty"`
	SequenceNumber *int64                  `json:"sequenceNumber,omitempty"`
	SortValues     *[]string               `json:"sortValues,omitempty"`
	Type           *string                 `json:"type,omitempty"`
	Version        *int64                  `json:"version,omitempty"`
}

// NewLayoutSection instantiates a new LayoutSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutSection() *LayoutSection {
	this := LayoutSection{}
	return &this
}

// NewLayoutSectionWithDefaults instantiates a new LayoutSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutSectionWithDefaults() *LayoutSection {
	this := LayoutSection{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LayoutSection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LayoutSection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LayoutSection) SetDescription(v string) {
	o.Description = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *LayoutSection) GetFields() []LayoutField {
	if o == nil || o.Fields == nil {
		var ret []LayoutField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetFieldsOk() (*[]LayoutField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *LayoutSection) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []LayoutField and assigns it to the Fields field.
func (o *LayoutSection) SetFields(v []LayoutField) {
	o.Fields = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *LayoutSection) GetHighlight() map[string][]string {
	if o == nil || o.Highlight == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetHighlightOk() (*map[string][]string, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *LayoutSection) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given map[string][]string and assigns it to the Highlight field.
func (o *LayoutSection) SetHighlight(v map[string][]string) {
	o.Highlight = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LayoutSection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LayoutSection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LayoutSection) SetId(v string) {
	o.Id = &v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *LayoutSection) GetIsVisible() bool {
	if o == nil || o.IsVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetIsVisibleOk() (*bool, bool) {
	if o == nil || o.IsVisible == nil {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *LayoutSection) HasIsVisible() bool {
	if o != nil && o.IsVisible != nil {
		return true
	}

	return false
}

// SetIsVisible gets a reference to the given bool and assigns it to the IsVisible field.
func (o *LayoutSection) SetIsVisible(v bool) {
	o.IsVisible = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *LayoutSection) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *LayoutSection) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *LayoutSection) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LayoutSection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LayoutSection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LayoutSection) SetName(v string) {
	o.Name = &v
}

// GetNumericId returns the NumericId field value if set, zero value otherwise.
func (o *LayoutSection) GetNumericId() int64 {
	if o == nil || o.NumericId == nil {
		var ret int64
		return ret
	}
	return *o.NumericId
}

// GetNumericIdOk returns a tuple with the NumericId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetNumericIdOk() (*int64, bool) {
	if o == nil || o.NumericId == nil {
		return nil, false
	}
	return o.NumericId, true
}

// HasNumericId returns a boolean if a field has been set.
func (o *LayoutSection) HasNumericId() bool {
	if o != nil && o.NumericId != nil {
		return true
	}

	return false
}

// SetNumericId gets a reference to the given int64 and assigns it to the NumericId field.
func (o *LayoutSection) SetNumericId(v int64) {
	o.NumericId = &v
}

// GetPrimaryTerm returns the PrimaryTerm field value if set, zero value otherwise.
func (o *LayoutSection) GetPrimaryTerm() int64 {
	if o == nil || o.PrimaryTerm == nil {
		var ret int64
		return ret
	}
	return *o.PrimaryTerm
}

// GetPrimaryTermOk returns a tuple with the PrimaryTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetPrimaryTermOk() (*int64, bool) {
	if o == nil || o.PrimaryTerm == nil {
		return nil, false
	}
	return o.PrimaryTerm, true
}

// HasPrimaryTerm returns a boolean if a field has been set.
func (o *LayoutSection) HasPrimaryTerm() bool {
	if o != nil && o.PrimaryTerm != nil {
		return true
	}

	return false
}

// SetPrimaryTerm gets a reference to the given int64 and assigns it to the PrimaryTerm field.
func (o *LayoutSection) SetPrimaryTerm(v int64) {
	o.PrimaryTerm = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LayoutSection) GetQuery() map[string]interface{} {
	if o == nil || o.Query == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetQueryOk() (*map[string]interface{}, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LayoutSection) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]interface{} and assigns it to the Query field.
func (o *LayoutSection) SetQuery(v map[string]interface{}) {
	o.Query = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *LayoutSection) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *LayoutSection) HasQueryType() bool {
	if o != nil && o.QueryType != nil {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *LayoutSection) SetQueryType(v string) {
	o.QueryType = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *LayoutSection) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *LayoutSection) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *LayoutSection) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *LayoutSection) GetSequenceNumber() int64 {
	if o == nil || o.SequenceNumber == nil {
		var ret int64
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetSequenceNumberOk() (*int64, bool) {
	if o == nil || o.SequenceNumber == nil {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *LayoutSection) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber != nil {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int64 and assigns it to the SequenceNumber field.
func (o *LayoutSection) SetSequenceNumber(v int64) {
	o.SequenceNumber = &v
}

// GetSortValues returns the SortValues field value if set, zero value otherwise.
func (o *LayoutSection) GetSortValues() []string {
	if o == nil || o.SortValues == nil {
		var ret []string
		return ret
	}
	return *o.SortValues
}

// GetSortValuesOk returns a tuple with the SortValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetSortValuesOk() (*[]string, bool) {
	if o == nil || o.SortValues == nil {
		return nil, false
	}
	return o.SortValues, true
}

// HasSortValues returns a boolean if a field has been set.
func (o *LayoutSection) HasSortValues() bool {
	if o != nil && o.SortValues != nil {
		return true
	}

	return false
}

// SetSortValues gets a reference to the given []string and assigns it to the SortValues field.
func (o *LayoutSection) SetSortValues(v []string) {
	o.SortValues = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LayoutSection) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LayoutSection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LayoutSection) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LayoutSection) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSection) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LayoutSection) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *LayoutSection) SetVersion(v int64) {
	o.Version = &v
}

func (o LayoutSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsVisible != nil {
		toSerialize["isVisible"] = o.IsVisible
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumericId != nil {
		toSerialize["numericId"] = o.NumericId
	}
	if o.PrimaryTerm != nil {
		toSerialize["primaryTerm"] = o.PrimaryTerm
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryType != nil {
		toSerialize["queryType"] = o.QueryType
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.SequenceNumber != nil {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	if o.SortValues != nil {
		toSerialize["sortValues"] = o.SortValues
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableLayoutSection struct {
	value *LayoutSection
	isSet bool
}

func (v NullableLayoutSection) Get() *LayoutSection {
	return v.value
}

func (v *NullableLayoutSection) Set(val *LayoutSection) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutSection) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutSection(val *LayoutSection) *NullableLayoutSection {
	return &NullableLayoutSection{value: val, isSet: true}
}

func (v NullableLayoutSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
