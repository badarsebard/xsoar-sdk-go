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

// DBotScore DBotScore - Contain the score of a specific brand for a specific insight
type DBotScore struct {
	Content          *string                            `json:"content,omitempty"`
	ContentFormat    *string                            `json:"contentFormat,omitempty"`
	Context          *map[string]map[string]interface{} `json:"context,omitempty"`
	IsTypedIndicator *bool                              `json:"isTypedIndicator,omitempty"`
	Reliability      *string                            `json:"reliability,omitempty"`
	Score            *int64                             `json:"score,omitempty"`
	// We need to track when the score changes to know if we need to re-calculate the overall score
	ScoreChangeTimestamp *time.Time `json:"scoreChangeTimestamp,omitempty"`
	Timestamp            *time.Time `json:"timestamp,omitempty"`
	Type                 *string    `json:"type,omitempty"`
}

// NewDBotScore instantiates a new DBotScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBotScore() *DBotScore {
	this := DBotScore{}
	return &this
}

// NewDBotScoreWithDefaults instantiates a new DBotScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBotScoreWithDefaults() *DBotScore {
	this := DBotScore{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DBotScore) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DBotScore) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DBotScore) SetContent(v string) {
	o.Content = &v
}

// GetContentFormat returns the ContentFormat field value if set, zero value otherwise.
func (o *DBotScore) GetContentFormat() string {
	if o == nil || o.ContentFormat == nil {
		var ret string
		return ret
	}
	return *o.ContentFormat
}

// GetContentFormatOk returns a tuple with the ContentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetContentFormatOk() (*string, bool) {
	if o == nil || o.ContentFormat == nil {
		return nil, false
	}
	return o.ContentFormat, true
}

// HasContentFormat returns a boolean if a field has been set.
func (o *DBotScore) HasContentFormat() bool {
	if o != nil && o.ContentFormat != nil {
		return true
	}

	return false
}

// SetContentFormat gets a reference to the given string and assigns it to the ContentFormat field.
func (o *DBotScore) SetContentFormat(v string) {
	o.ContentFormat = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DBotScore) GetContext() map[string]map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetContextOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DBotScore) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]map[string]interface{} and assigns it to the Context field.
func (o *DBotScore) SetContext(v map[string]map[string]interface{}) {
	o.Context = &v
}

// GetIsTypedIndicator returns the IsTypedIndicator field value if set, zero value otherwise.
func (o *DBotScore) GetIsTypedIndicator() bool {
	if o == nil || o.IsTypedIndicator == nil {
		var ret bool
		return ret
	}
	return *o.IsTypedIndicator
}

// GetIsTypedIndicatorOk returns a tuple with the IsTypedIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetIsTypedIndicatorOk() (*bool, bool) {
	if o == nil || o.IsTypedIndicator == nil {
		return nil, false
	}
	return o.IsTypedIndicator, true
}

// HasIsTypedIndicator returns a boolean if a field has been set.
func (o *DBotScore) HasIsTypedIndicator() bool {
	if o != nil && o.IsTypedIndicator != nil {
		return true
	}

	return false
}

// SetIsTypedIndicator gets a reference to the given bool and assigns it to the IsTypedIndicator field.
func (o *DBotScore) SetIsTypedIndicator(v bool) {
	o.IsTypedIndicator = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *DBotScore) GetReliability() string {
	if o == nil || o.Reliability == nil {
		var ret string
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetReliabilityOk() (*string, bool) {
	if o == nil || o.Reliability == nil {
		return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *DBotScore) HasReliability() bool {
	if o != nil && o.Reliability != nil {
		return true
	}

	return false
}

// SetReliability gets a reference to the given string and assigns it to the Reliability field.
func (o *DBotScore) SetReliability(v string) {
	o.Reliability = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *DBotScore) GetScore() int64 {
	if o == nil || o.Score == nil {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetScoreOk() (*int64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *DBotScore) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *DBotScore) SetScore(v int64) {
	o.Score = &v
}

// GetScoreChangeTimestamp returns the ScoreChangeTimestamp field value if set, zero value otherwise.
func (o *DBotScore) GetScoreChangeTimestamp() time.Time {
	if o == nil || o.ScoreChangeTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.ScoreChangeTimestamp
}

// GetScoreChangeTimestampOk returns a tuple with the ScoreChangeTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetScoreChangeTimestampOk() (*time.Time, bool) {
	if o == nil || o.ScoreChangeTimestamp == nil {
		return nil, false
	}
	return o.ScoreChangeTimestamp, true
}

// HasScoreChangeTimestamp returns a boolean if a field has been set.
func (o *DBotScore) HasScoreChangeTimestamp() bool {
	if o != nil && o.ScoreChangeTimestamp != nil {
		return true
	}

	return false
}

// SetScoreChangeTimestamp gets a reference to the given time.Time and assigns it to the ScoreChangeTimestamp field.
func (o *DBotScore) SetScoreChangeTimestamp(v time.Time) {
	o.ScoreChangeTimestamp = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DBotScore) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DBotScore) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DBotScore) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DBotScore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBotScore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DBotScore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DBotScore) SetType(v string) {
	o.Type = &v
}

func (o DBotScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentFormat != nil {
		toSerialize["contentFormat"] = o.ContentFormat
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.IsTypedIndicator != nil {
		toSerialize["isTypedIndicator"] = o.IsTypedIndicator
	}
	if o.Reliability != nil {
		toSerialize["reliability"] = o.Reliability
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.ScoreChangeTimestamp != nil {
		toSerialize["scoreChangeTimestamp"] = o.ScoreChangeTimestamp
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDBotScore struct {
	value *DBotScore
	isSet bool
}

func (v NullableDBotScore) Get() *DBotScore {
	return v.value
}

func (v *NullableDBotScore) Set(val *DBotScore) {
	v.value = val
	v.isSet = true
}

func (v NullableDBotScore) IsSet() bool {
	return v.isSet
}

func (v *NullableDBotScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBotScore(val *DBotScore) *NullableDBotScore {
	return &NullableDBotScore{value: val, isSet: true}
}

func (v NullableDBotScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBotScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
