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

// WhitelistedIndicator WhitelistedIndicator Excluded indicator
type WhitelistedIndicator struct {
	Highlight      *map[string][]string `json:"highlight,omitempty"`
	Id             *string              `json:"id,omitempty"`
	Locked         *bool                `json:"locked,omitempty"`
	Modified       *time.Time           `json:"modified,omitempty"`
	NumericId      *int64               `json:"numericId,omitempty"`
	PrimaryTerm    *int64               `json:"primaryTerm,omitempty"`
	Reason         *string              `json:"reason,omitempty"`
	Reputations    *[]string            `json:"reputations,omitempty"`
	SequenceNumber *int64               `json:"sequenceNumber,omitempty"`
	SortValues     *[]string            `json:"sortValues,omitempty"`
	Type           *string              `json:"type,omitempty"`
	User           *string              `json:"user,omitempty"`
	Value          *string              `json:"value,omitempty"`
	Version        *int64               `json:"version,omitempty"`
	WhitelistTime  *time.Time           `json:"whitelistTime,omitempty"`
}

// NewWhitelistedIndicator instantiates a new WhitelistedIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhitelistedIndicator() *WhitelistedIndicator {
	this := WhitelistedIndicator{}
	return &this
}

// NewWhitelistedIndicatorWithDefaults instantiates a new WhitelistedIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhitelistedIndicatorWithDefaults() *WhitelistedIndicator {
	this := WhitelistedIndicator{}
	return &this
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetHighlight() map[string][]string {
	if o == nil || o.Highlight == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetHighlightOk() (*map[string][]string, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given map[string][]string and assigns it to the Highlight field.
func (o *WhitelistedIndicator) SetHighlight(v map[string][]string) {
	o.Highlight = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhitelistedIndicator) SetId(v string) {
	o.Id = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *WhitelistedIndicator) SetLocked(v bool) {
	o.Locked = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *WhitelistedIndicator) SetModified(v time.Time) {
	o.Modified = &v
}

// GetNumericId returns the NumericId field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetNumericId() int64 {
	if o == nil || o.NumericId == nil {
		var ret int64
		return ret
	}
	return *o.NumericId
}

// GetNumericIdOk returns a tuple with the NumericId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetNumericIdOk() (*int64, bool) {
	if o == nil || o.NumericId == nil {
		return nil, false
	}
	return o.NumericId, true
}

// HasNumericId returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasNumericId() bool {
	if o != nil && o.NumericId != nil {
		return true
	}

	return false
}

// SetNumericId gets a reference to the given int64 and assigns it to the NumericId field.
func (o *WhitelistedIndicator) SetNumericId(v int64) {
	o.NumericId = &v
}

// GetPrimaryTerm returns the PrimaryTerm field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetPrimaryTerm() int64 {
	if o == nil || o.PrimaryTerm == nil {
		var ret int64
		return ret
	}
	return *o.PrimaryTerm
}

// GetPrimaryTermOk returns a tuple with the PrimaryTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetPrimaryTermOk() (*int64, bool) {
	if o == nil || o.PrimaryTerm == nil {
		return nil, false
	}
	return o.PrimaryTerm, true
}

// HasPrimaryTerm returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasPrimaryTerm() bool {
	if o != nil && o.PrimaryTerm != nil {
		return true
	}

	return false
}

// SetPrimaryTerm gets a reference to the given int64 and assigns it to the PrimaryTerm field.
func (o *WhitelistedIndicator) SetPrimaryTerm(v int64) {
	o.PrimaryTerm = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *WhitelistedIndicator) SetReason(v string) {
	o.Reason = &v
}

// GetReputations returns the Reputations field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetReputations() []string {
	if o == nil || o.Reputations == nil {
		var ret []string
		return ret
	}
	return *o.Reputations
}

// GetReputationsOk returns a tuple with the Reputations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetReputationsOk() (*[]string, bool) {
	if o == nil || o.Reputations == nil {
		return nil, false
	}
	return o.Reputations, true
}

// HasReputations returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasReputations() bool {
	if o != nil && o.Reputations != nil {
		return true
	}

	return false
}

// SetReputations gets a reference to the given []string and assigns it to the Reputations field.
func (o *WhitelistedIndicator) SetReputations(v []string) {
	o.Reputations = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetSequenceNumber() int64 {
	if o == nil || o.SequenceNumber == nil {
		var ret int64
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetSequenceNumberOk() (*int64, bool) {
	if o == nil || o.SequenceNumber == nil {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber != nil {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int64 and assigns it to the SequenceNumber field.
func (o *WhitelistedIndicator) SetSequenceNumber(v int64) {
	o.SequenceNumber = &v
}

// GetSortValues returns the SortValues field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetSortValues() []string {
	if o == nil || o.SortValues == nil {
		var ret []string
		return ret
	}
	return *o.SortValues
}

// GetSortValuesOk returns a tuple with the SortValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetSortValuesOk() (*[]string, bool) {
	if o == nil || o.SortValues == nil {
		return nil, false
	}
	return o.SortValues, true
}

// HasSortValues returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasSortValues() bool {
	if o != nil && o.SortValues != nil {
		return true
	}

	return false
}

// SetSortValues gets a reference to the given []string and assigns it to the SortValues field.
func (o *WhitelistedIndicator) SetSortValues(v []string) {
	o.SortValues = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhitelistedIndicator) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *WhitelistedIndicator) SetUser(v string) {
	o.User = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WhitelistedIndicator) SetValue(v string) {
	o.Value = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WhitelistedIndicator) SetVersion(v int64) {
	o.Version = &v
}

// GetWhitelistTime returns the WhitelistTime field value if set, zero value otherwise.
func (o *WhitelistedIndicator) GetWhitelistTime() time.Time {
	if o == nil || o.WhitelistTime == nil {
		var ret time.Time
		return ret
	}
	return *o.WhitelistTime
}

// GetWhitelistTimeOk returns a tuple with the WhitelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistedIndicator) GetWhitelistTimeOk() (*time.Time, bool) {
	if o == nil || o.WhitelistTime == nil {
		return nil, false
	}
	return o.WhitelistTime, true
}

// HasWhitelistTime returns a boolean if a field has been set.
func (o *WhitelistedIndicator) HasWhitelistTime() bool {
	if o != nil && o.WhitelistTime != nil {
		return true
	}

	return false
}

// SetWhitelistTime gets a reference to the given time.Time and assigns it to the WhitelistTime field.
func (o *WhitelistedIndicator) SetWhitelistTime(v time.Time) {
	o.WhitelistTime = &v
}

func (o WhitelistedIndicator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.NumericId != nil {
		toSerialize["numericId"] = o.NumericId
	}
	if o.PrimaryTerm != nil {
		toSerialize["primaryTerm"] = o.PrimaryTerm
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Reputations != nil {
		toSerialize["reputations"] = o.Reputations
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
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.WhitelistTime != nil {
		toSerialize["whitelistTime"] = o.WhitelistTime
	}
	return json.Marshal(toSerialize)
}

type NullableWhitelistedIndicator struct {
	value *WhitelistedIndicator
	isSet bool
}

func (v NullableWhitelistedIndicator) Get() *WhitelistedIndicator {
	return v.value
}

func (v *NullableWhitelistedIndicator) Set(val *WhitelistedIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableWhitelistedIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableWhitelistedIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhitelistedIndicator(val *WhitelistedIndicator) *NullableWhitelistedIndicator {
	return &NullableWhitelistedIndicator{value: val, isSet: true}
}

func (v NullableWhitelistedIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhitelistedIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
