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

// UpdateDataBatch struct for UpdateDataBatch
type UpdateDataBatch struct {
	CustomFields          *map[string]map[string]interface{} `json:"CustomFields,omitempty"`
	All                   *bool                              `json:"all,omitempty"`
	CloseNotes            *string                            `json:"closeNotes,omitempty"`
	CloseReason           *string                            `json:"closeReason,omitempty"`
	Columns               *[]string                          `json:"columns,omitempty"`
	Data                  *map[string]map[string]interface{} `json:"data,omitempty"`
	Filter                *IncidentFilter                    `json:"filter,omitempty"`
	Force                 *bool                              `json:"force,omitempty"`
	Ids                   *[]string                          `json:"ids,omitempty"`
	Line                  *string                            `json:"line,omitempty"`
	OriginalIncidentId    *string                            `json:"originalIncidentId,omitempty"`
	OverrideInvestigation *bool                              `json:"overrideInvestigation,omitempty"`
}

// NewUpdateDataBatch instantiates a new UpdateDataBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataBatch() *UpdateDataBatch {
	this := UpdateDataBatch{}
	return &this
}

// NewUpdateDataBatchWithDefaults instantiates a new UpdateDataBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataBatchWithDefaults() *UpdateDataBatch {
	this := UpdateDataBatch{}
	return &this
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetCustomFields() map[string]map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]map[string]interface{} and assigns it to the CustomFields field.
func (o *UpdateDataBatch) SetCustomFields(v map[string]map[string]interface{}) {
	o.CustomFields = &v
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *UpdateDataBatch) SetAll(v bool) {
	o.All = &v
}

// GetCloseNotes returns the CloseNotes field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetCloseNotes() string {
	if o == nil || o.CloseNotes == nil {
		var ret string
		return ret
	}
	return *o.CloseNotes
}

// GetCloseNotesOk returns a tuple with the CloseNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetCloseNotesOk() (*string, bool) {
	if o == nil || o.CloseNotes == nil {
		return nil, false
	}
	return o.CloseNotes, true
}

// HasCloseNotes returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasCloseNotes() bool {
	if o != nil && o.CloseNotes != nil {
		return true
	}

	return false
}

// SetCloseNotes gets a reference to the given string and assigns it to the CloseNotes field.
func (o *UpdateDataBatch) SetCloseNotes(v string) {
	o.CloseNotes = &v
}

// GetCloseReason returns the CloseReason field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetCloseReason() string {
	if o == nil || o.CloseReason == nil {
		var ret string
		return ret
	}
	return *o.CloseReason
}

// GetCloseReasonOk returns a tuple with the CloseReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetCloseReasonOk() (*string, bool) {
	if o == nil || o.CloseReason == nil {
		return nil, false
	}
	return o.CloseReason, true
}

// HasCloseReason returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasCloseReason() bool {
	if o != nil && o.CloseReason != nil {
		return true
	}

	return false
}

// SetCloseReason gets a reference to the given string and assigns it to the CloseReason field.
func (o *UpdateDataBatch) SetCloseReason(v string) {
	o.CloseReason = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *UpdateDataBatch) SetColumns(v []string) {
	o.Columns = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetData() map[string]map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetDataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *UpdateDataBatch) SetData(v map[string]map[string]interface{}) {
	o.Data = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetFilter() IncidentFilter {
	if o == nil || o.Filter == nil {
		var ret IncidentFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetFilterOk() (*IncidentFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given IncidentFilter and assigns it to the Filter field.
func (o *UpdateDataBatch) SetFilter(v IncidentFilter) {
	o.Filter = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *UpdateDataBatch) SetForce(v bool) {
	o.Force = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *UpdateDataBatch) SetIds(v []string) {
	o.Ids = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetLine() string {
	if o == nil || o.Line == nil {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetLineOk() (*string, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *UpdateDataBatch) SetLine(v string) {
	o.Line = &v
}

// GetOriginalIncidentId returns the OriginalIncidentId field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetOriginalIncidentId() string {
	if o == nil || o.OriginalIncidentId == nil {
		var ret string
		return ret
	}
	return *o.OriginalIncidentId
}

// GetOriginalIncidentIdOk returns a tuple with the OriginalIncidentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetOriginalIncidentIdOk() (*string, bool) {
	if o == nil || o.OriginalIncidentId == nil {
		return nil, false
	}
	return o.OriginalIncidentId, true
}

// HasOriginalIncidentId returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasOriginalIncidentId() bool {
	if o != nil && o.OriginalIncidentId != nil {
		return true
	}

	return false
}

// SetOriginalIncidentId gets a reference to the given string and assigns it to the OriginalIncidentId field.
func (o *UpdateDataBatch) SetOriginalIncidentId(v string) {
	o.OriginalIncidentId = &v
}

// GetOverrideInvestigation returns the OverrideInvestigation field value if set, zero value otherwise.
func (o *UpdateDataBatch) GetOverrideInvestigation() bool {
	if o == nil || o.OverrideInvestigation == nil {
		var ret bool
		return ret
	}
	return *o.OverrideInvestigation
}

// GetOverrideInvestigationOk returns a tuple with the OverrideInvestigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataBatch) GetOverrideInvestigationOk() (*bool, bool) {
	if o == nil || o.OverrideInvestigation == nil {
		return nil, false
	}
	return o.OverrideInvestigation, true
}

// HasOverrideInvestigation returns a boolean if a field has been set.
func (o *UpdateDataBatch) HasOverrideInvestigation() bool {
	if o != nil && o.OverrideInvestigation != nil {
		return true
	}

	return false
}

// SetOverrideInvestigation gets a reference to the given bool and assigns it to the OverrideInvestigation field.
func (o *UpdateDataBatch) SetOverrideInvestigation(v bool) {
	o.OverrideInvestigation = &v
}

func (o UpdateDataBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomFields != nil {
		toSerialize["CustomFields"] = o.CustomFields
	}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.CloseNotes != nil {
		toSerialize["closeNotes"] = o.CloseNotes
	}
	if o.CloseReason != nil {
		toSerialize["closeReason"] = o.CloseReason
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.OriginalIncidentId != nil {
		toSerialize["originalIncidentId"] = o.OriginalIncidentId
	}
	if o.OverrideInvestigation != nil {
		toSerialize["overrideInvestigation"] = o.OverrideInvestigation
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDataBatch struct {
	value *UpdateDataBatch
	isSet bool
}

func (v NullableUpdateDataBatch) Get() *UpdateDataBatch {
	return v.value
}

func (v *NullableUpdateDataBatch) Set(val *UpdateDataBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataBatch(val *UpdateDataBatch) *NullableUpdateDataBatch {
	return &NullableUpdateDataBatch{value: val, isSet: true}
}

func (v NullableUpdateDataBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
