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

// FeedIndicator struct for FeedIndicator
type FeedIndicator struct {
	ExpirationSource    *ExpirationSource                  `json:"ExpirationSource,omitempty"`
	BypassExclusionList *bool                              `json:"bypassExclusionList,omitempty"`
	ClassifierId        *string                            `json:"classifierId,omitempty"`
	ClassifierVersion   *int64                             `json:"classifierVersion,omitempty"`
	Comments            *[]FeedIndicatorComment            `json:"comments,omitempty"`
	ExpirationInterval  *int64                             `json:"expirationInterval,omitempty"`
	ExpirationPolicy    *string                            `json:"expirationPolicy,omitempty"`
	FetchTime           *time.Time                         `json:"fetchTime,omitempty"`
	Fields              *CustomFields                      `json:"fields,omitempty"`
	IsEnrichment        *bool                              `json:"isEnrichment,omitempty"`
	MapperId            *string                            `json:"mapperId,omitempty"`
	MapperVersion       *int64                             `json:"mapperVersion,omitempty"`
	ModifiedTime        *time.Time                         `json:"modifiedTime,omitempty"`
	ModuleId            *string                            `json:"moduleId,omitempty"`
	RawJSON             *map[string]map[string]interface{} `json:"rawJSON,omitempty"`
	Relationships       *RelationshipsAPI                  `json:"relationships,omitempty"`
	Reliability         *string                            `json:"reliability,omitempty"`
	Score               *int64                             `json:"score,omitempty"`
	SourceBrand         *string                            `json:"sourceBrand,omitempty"`
	SourceInstance      *string                            `json:"sourceInstance,omitempty"`
	Timestamp           *time.Time                         `json:"timestamp,omitempty"`
	Type                *string                            `json:"type,omitempty"`
	Value               *string                            `json:"value,omitempty"`
}

// NewFeedIndicator instantiates a new FeedIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedIndicator() *FeedIndicator {
	this := FeedIndicator{}
	return &this
}

// NewFeedIndicatorWithDefaults instantiates a new FeedIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedIndicatorWithDefaults() *FeedIndicator {
	this := FeedIndicator{}
	return &this
}

// GetExpirationSource returns the ExpirationSource field value if set, zero value otherwise.
func (o *FeedIndicator) GetExpirationSource() ExpirationSource {
	if o == nil || o.ExpirationSource == nil {
		var ret ExpirationSource
		return ret
	}
	return *o.ExpirationSource
}

// GetExpirationSourceOk returns a tuple with the ExpirationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetExpirationSourceOk() (*ExpirationSource, bool) {
	if o == nil || o.ExpirationSource == nil {
		return nil, false
	}
	return o.ExpirationSource, true
}

// HasExpirationSource returns a boolean if a field has been set.
func (o *FeedIndicator) HasExpirationSource() bool {
	if o != nil && o.ExpirationSource != nil {
		return true
	}

	return false
}

// SetExpirationSource gets a reference to the given ExpirationSource and assigns it to the ExpirationSource field.
func (o *FeedIndicator) SetExpirationSource(v ExpirationSource) {
	o.ExpirationSource = &v
}

// GetBypassExclusionList returns the BypassExclusionList field value if set, zero value otherwise.
func (o *FeedIndicator) GetBypassExclusionList() bool {
	if o == nil || o.BypassExclusionList == nil {
		var ret bool
		return ret
	}
	return *o.BypassExclusionList
}

// GetBypassExclusionListOk returns a tuple with the BypassExclusionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetBypassExclusionListOk() (*bool, bool) {
	if o == nil || o.BypassExclusionList == nil {
		return nil, false
	}
	return o.BypassExclusionList, true
}

// HasBypassExclusionList returns a boolean if a field has been set.
func (o *FeedIndicator) HasBypassExclusionList() bool {
	if o != nil && o.BypassExclusionList != nil {
		return true
	}

	return false
}

// SetBypassExclusionList gets a reference to the given bool and assigns it to the BypassExclusionList field.
func (o *FeedIndicator) SetBypassExclusionList(v bool) {
	o.BypassExclusionList = &v
}

// GetClassifierId returns the ClassifierId field value if set, zero value otherwise.
func (o *FeedIndicator) GetClassifierId() string {
	if o == nil || o.ClassifierId == nil {
		var ret string
		return ret
	}
	return *o.ClassifierId
}

// GetClassifierIdOk returns a tuple with the ClassifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetClassifierIdOk() (*string, bool) {
	if o == nil || o.ClassifierId == nil {
		return nil, false
	}
	return o.ClassifierId, true
}

// HasClassifierId returns a boolean if a field has been set.
func (o *FeedIndicator) HasClassifierId() bool {
	if o != nil && o.ClassifierId != nil {
		return true
	}

	return false
}

// SetClassifierId gets a reference to the given string and assigns it to the ClassifierId field.
func (o *FeedIndicator) SetClassifierId(v string) {
	o.ClassifierId = &v
}

// GetClassifierVersion returns the ClassifierVersion field value if set, zero value otherwise.
func (o *FeedIndicator) GetClassifierVersion() int64 {
	if o == nil || o.ClassifierVersion == nil {
		var ret int64
		return ret
	}
	return *o.ClassifierVersion
}

// GetClassifierVersionOk returns a tuple with the ClassifierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetClassifierVersionOk() (*int64, bool) {
	if o == nil || o.ClassifierVersion == nil {
		return nil, false
	}
	return o.ClassifierVersion, true
}

// HasClassifierVersion returns a boolean if a field has been set.
func (o *FeedIndicator) HasClassifierVersion() bool {
	if o != nil && o.ClassifierVersion != nil {
		return true
	}

	return false
}

// SetClassifierVersion gets a reference to the given int64 and assigns it to the ClassifierVersion field.
func (o *FeedIndicator) SetClassifierVersion(v int64) {
	o.ClassifierVersion = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *FeedIndicator) GetComments() []FeedIndicatorComment {
	if o == nil || o.Comments == nil {
		var ret []FeedIndicatorComment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetCommentsOk() (*[]FeedIndicatorComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *FeedIndicator) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []FeedIndicatorComment and assigns it to the Comments field.
func (o *FeedIndicator) SetComments(v []FeedIndicatorComment) {
	o.Comments = &v
}

// GetExpirationInterval returns the ExpirationInterval field value if set, zero value otherwise.
func (o *FeedIndicator) GetExpirationInterval() int64 {
	if o == nil || o.ExpirationInterval == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationInterval
}

// GetExpirationIntervalOk returns a tuple with the ExpirationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetExpirationIntervalOk() (*int64, bool) {
	if o == nil || o.ExpirationInterval == nil {
		return nil, false
	}
	return o.ExpirationInterval, true
}

// HasExpirationInterval returns a boolean if a field has been set.
func (o *FeedIndicator) HasExpirationInterval() bool {
	if o != nil && o.ExpirationInterval != nil {
		return true
	}

	return false
}

// SetExpirationInterval gets a reference to the given int64 and assigns it to the ExpirationInterval field.
func (o *FeedIndicator) SetExpirationInterval(v int64) {
	o.ExpirationInterval = &v
}

// GetExpirationPolicy returns the ExpirationPolicy field value if set, zero value otherwise.
func (o *FeedIndicator) GetExpirationPolicy() string {
	if o == nil || o.ExpirationPolicy == nil {
		var ret string
		return ret
	}
	return *o.ExpirationPolicy
}

// GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetExpirationPolicyOk() (*string, bool) {
	if o == nil || o.ExpirationPolicy == nil {
		return nil, false
	}
	return o.ExpirationPolicy, true
}

// HasExpirationPolicy returns a boolean if a field has been set.
func (o *FeedIndicator) HasExpirationPolicy() bool {
	if o != nil && o.ExpirationPolicy != nil {
		return true
	}

	return false
}

// SetExpirationPolicy gets a reference to the given string and assigns it to the ExpirationPolicy field.
func (o *FeedIndicator) SetExpirationPolicy(v string) {
	o.ExpirationPolicy = &v
}

// GetFetchTime returns the FetchTime field value if set, zero value otherwise.
func (o *FeedIndicator) GetFetchTime() time.Time {
	if o == nil || o.FetchTime == nil {
		var ret time.Time
		return ret
	}
	return *o.FetchTime
}

// GetFetchTimeOk returns a tuple with the FetchTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetFetchTimeOk() (*time.Time, bool) {
	if o == nil || o.FetchTime == nil {
		return nil, false
	}
	return o.FetchTime, true
}

// HasFetchTime returns a boolean if a field has been set.
func (o *FeedIndicator) HasFetchTime() bool {
	if o != nil && o.FetchTime != nil {
		return true
	}

	return false
}

// SetFetchTime gets a reference to the given time.Time and assigns it to the FetchTime field.
func (o *FeedIndicator) SetFetchTime(v time.Time) {
	o.FetchTime = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *FeedIndicator) GetFields() CustomFields {
	if o == nil || o.Fields == nil {
		var ret CustomFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetFieldsOk() (*CustomFields, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *FeedIndicator) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given CustomFields and assigns it to the Fields field.
func (o *FeedIndicator) SetFields(v CustomFields) {
	o.Fields = &v
}

// GetIsEnrichment returns the IsEnrichment field value if set, zero value otherwise.
func (o *FeedIndicator) GetIsEnrichment() bool {
	if o == nil || o.IsEnrichment == nil {
		var ret bool
		return ret
	}
	return *o.IsEnrichment
}

// GetIsEnrichmentOk returns a tuple with the IsEnrichment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetIsEnrichmentOk() (*bool, bool) {
	if o == nil || o.IsEnrichment == nil {
		return nil, false
	}
	return o.IsEnrichment, true
}

// HasIsEnrichment returns a boolean if a field has been set.
func (o *FeedIndicator) HasIsEnrichment() bool {
	if o != nil && o.IsEnrichment != nil {
		return true
	}

	return false
}

// SetIsEnrichment gets a reference to the given bool and assigns it to the IsEnrichment field.
func (o *FeedIndicator) SetIsEnrichment(v bool) {
	o.IsEnrichment = &v
}

// GetMapperId returns the MapperId field value if set, zero value otherwise.
func (o *FeedIndicator) GetMapperId() string {
	if o == nil || o.MapperId == nil {
		var ret string
		return ret
	}
	return *o.MapperId
}

// GetMapperIdOk returns a tuple with the MapperId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetMapperIdOk() (*string, bool) {
	if o == nil || o.MapperId == nil {
		return nil, false
	}
	return o.MapperId, true
}

// HasMapperId returns a boolean if a field has been set.
func (o *FeedIndicator) HasMapperId() bool {
	if o != nil && o.MapperId != nil {
		return true
	}

	return false
}

// SetMapperId gets a reference to the given string and assigns it to the MapperId field.
func (o *FeedIndicator) SetMapperId(v string) {
	o.MapperId = &v
}

// GetMapperVersion returns the MapperVersion field value if set, zero value otherwise.
func (o *FeedIndicator) GetMapperVersion() int64 {
	if o == nil || o.MapperVersion == nil {
		var ret int64
		return ret
	}
	return *o.MapperVersion
}

// GetMapperVersionOk returns a tuple with the MapperVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetMapperVersionOk() (*int64, bool) {
	if o == nil || o.MapperVersion == nil {
		return nil, false
	}
	return o.MapperVersion, true
}

// HasMapperVersion returns a boolean if a field has been set.
func (o *FeedIndicator) HasMapperVersion() bool {
	if o != nil && o.MapperVersion != nil {
		return true
	}

	return false
}

// SetMapperVersion gets a reference to the given int64 and assigns it to the MapperVersion field.
func (o *FeedIndicator) SetMapperVersion(v int64) {
	o.MapperVersion = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *FeedIndicator) GetModifiedTime() time.Time {
	if o == nil || o.ModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifiedTime == nil {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *FeedIndicator) HasModifiedTime() bool {
	if o != nil && o.ModifiedTime != nil {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given time.Time and assigns it to the ModifiedTime field.
func (o *FeedIndicator) SetModifiedTime(v time.Time) {
	o.ModifiedTime = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *FeedIndicator) GetModuleId() string {
	if o == nil || o.ModuleId == nil {
		var ret string
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetModuleIdOk() (*string, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *FeedIndicator) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given string and assigns it to the ModuleId field.
func (o *FeedIndicator) SetModuleId(v string) {
	o.ModuleId = &v
}

// GetRawJSON returns the RawJSON field value if set, zero value otherwise.
func (o *FeedIndicator) GetRawJSON() map[string]map[string]interface{} {
	if o == nil || o.RawJSON == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.RawJSON
}

// GetRawJSONOk returns a tuple with the RawJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetRawJSONOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.RawJSON == nil {
		return nil, false
	}
	return o.RawJSON, true
}

// HasRawJSON returns a boolean if a field has been set.
func (o *FeedIndicator) HasRawJSON() bool {
	if o != nil && o.RawJSON != nil {
		return true
	}

	return false
}

// SetRawJSON gets a reference to the given map[string]map[string]interface{} and assigns it to the RawJSON field.
func (o *FeedIndicator) SetRawJSON(v map[string]map[string]interface{}) {
	o.RawJSON = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FeedIndicator) GetRelationships() RelationshipsAPI {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsAPI
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetRelationshipsOk() (*RelationshipsAPI, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FeedIndicator) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsAPI and assigns it to the Relationships field.
func (o *FeedIndicator) SetRelationships(v RelationshipsAPI) {
	o.Relationships = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *FeedIndicator) GetReliability() string {
	if o == nil || o.Reliability == nil {
		var ret string
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetReliabilityOk() (*string, bool) {
	if o == nil || o.Reliability == nil {
		return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *FeedIndicator) HasReliability() bool {
	if o != nil && o.Reliability != nil {
		return true
	}

	return false
}

// SetReliability gets a reference to the given string and assigns it to the Reliability field.
func (o *FeedIndicator) SetReliability(v string) {
	o.Reliability = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *FeedIndicator) GetScore() int64 {
	if o == nil || o.Score == nil {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetScoreOk() (*int64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *FeedIndicator) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *FeedIndicator) SetScore(v int64) {
	o.Score = &v
}

// GetSourceBrand returns the SourceBrand field value if set, zero value otherwise.
func (o *FeedIndicator) GetSourceBrand() string {
	if o == nil || o.SourceBrand == nil {
		var ret string
		return ret
	}
	return *o.SourceBrand
}

// GetSourceBrandOk returns a tuple with the SourceBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetSourceBrandOk() (*string, bool) {
	if o == nil || o.SourceBrand == nil {
		return nil, false
	}
	return o.SourceBrand, true
}

// HasSourceBrand returns a boolean if a field has been set.
func (o *FeedIndicator) HasSourceBrand() bool {
	if o != nil && o.SourceBrand != nil {
		return true
	}

	return false
}

// SetSourceBrand gets a reference to the given string and assigns it to the SourceBrand field.
func (o *FeedIndicator) SetSourceBrand(v string) {
	o.SourceBrand = &v
}

// GetSourceInstance returns the SourceInstance field value if set, zero value otherwise.
func (o *FeedIndicator) GetSourceInstance() string {
	if o == nil || o.SourceInstance == nil {
		var ret string
		return ret
	}
	return *o.SourceInstance
}

// GetSourceInstanceOk returns a tuple with the SourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetSourceInstanceOk() (*string, bool) {
	if o == nil || o.SourceInstance == nil {
		return nil, false
	}
	return o.SourceInstance, true
}

// HasSourceInstance returns a boolean if a field has been set.
func (o *FeedIndicator) HasSourceInstance() bool {
	if o != nil && o.SourceInstance != nil {
		return true
	}

	return false
}

// SetSourceInstance gets a reference to the given string and assigns it to the SourceInstance field.
func (o *FeedIndicator) SetSourceInstance(v string) {
	o.SourceInstance = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *FeedIndicator) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FeedIndicator) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *FeedIndicator) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeedIndicator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeedIndicator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeedIndicator) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FeedIndicator) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndicator) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeedIndicator) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FeedIndicator) SetValue(v string) {
	o.Value = &v
}

func (o FeedIndicator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationSource != nil {
		toSerialize["ExpirationSource"] = o.ExpirationSource
	}
	if o.BypassExclusionList != nil {
		toSerialize["bypassExclusionList"] = o.BypassExclusionList
	}
	if o.ClassifierId != nil {
		toSerialize["classifierId"] = o.ClassifierId
	}
	if o.ClassifierVersion != nil {
		toSerialize["classifierVersion"] = o.ClassifierVersion
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.ExpirationInterval != nil {
		toSerialize["expirationInterval"] = o.ExpirationInterval
	}
	if o.ExpirationPolicy != nil {
		toSerialize["expirationPolicy"] = o.ExpirationPolicy
	}
	if o.FetchTime != nil {
		toSerialize["fetchTime"] = o.FetchTime
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IsEnrichment != nil {
		toSerialize["isEnrichment"] = o.IsEnrichment
	}
	if o.MapperId != nil {
		toSerialize["mapperId"] = o.MapperId
	}
	if o.MapperVersion != nil {
		toSerialize["mapperVersion"] = o.MapperVersion
	}
	if o.ModifiedTime != nil {
		toSerialize["modifiedTime"] = o.ModifiedTime
	}
	if o.ModuleId != nil {
		toSerialize["moduleId"] = o.ModuleId
	}
	if o.RawJSON != nil {
		toSerialize["rawJSON"] = o.RawJSON
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Reliability != nil {
		toSerialize["reliability"] = o.Reliability
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.SourceBrand != nil {
		toSerialize["sourceBrand"] = o.SourceBrand
	}
	if o.SourceInstance != nil {
		toSerialize["sourceInstance"] = o.SourceInstance
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFeedIndicator struct {
	value *FeedIndicator
	isSet bool
}

func (v NullableFeedIndicator) Get() *FeedIndicator {
	return v.value
}

func (v *NullableFeedIndicator) Set(val *FeedIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedIndicator(val *FeedIndicator) *NullableFeedIndicator {
	return &NullableFeedIndicator{value: val, isSet: true}
}

func (v NullableFeedIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
