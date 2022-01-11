# FeedIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationSource** | Pointer to [**ExpirationSource**](ExpirationSource.md) |  | [optional] 
**BypassExclusionList** | Pointer to **bool** |  | [optional] 
**ClassifierId** | Pointer to **string** |  | [optional] 
**ClassifierVersion** | Pointer to **int64** |  | [optional] 
**Comments** | Pointer to [**[]FeedIndicatorComment**](FeedIndicatorComment.md) |  | [optional] 
**ExpirationInterval** | Pointer to **int64** |  | [optional] 
**ExpirationPolicy** | Pointer to **string** |  | [optional] 
**FetchTime** | Pointer to **time.Time** |  | [optional] 
**Fields** | Pointer to [**CustomFields**](CustomFields.md) |  | [optional] 
**IsEnrichment** | Pointer to **bool** |  | [optional] 
**MapperId** | Pointer to **string** |  | [optional] 
**MapperVersion** | Pointer to **int64** |  | [optional] 
**ModifiedTime** | Pointer to **time.Time** |  | [optional] 
**ModuleId** | Pointer to **string** |  | [optional] 
**RawJSON** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**RelationshipsAPI**](RelationshipsAPI.md) |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**SourceBrand** | Pointer to **string** |  | [optional] 
**SourceInstance** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewFeedIndicator

`func NewFeedIndicator() *FeedIndicator`

NewFeedIndicator instantiates a new FeedIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedIndicatorWithDefaults

`func NewFeedIndicatorWithDefaults() *FeedIndicator`

NewFeedIndicatorWithDefaults instantiates a new FeedIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationSource

`func (o *FeedIndicator) GetExpirationSource() ExpirationSource`

GetExpirationSource returns the ExpirationSource field if non-nil, zero value otherwise.

### GetExpirationSourceOk

`func (o *FeedIndicator) GetExpirationSourceOk() (*ExpirationSource, bool)`

GetExpirationSourceOk returns a tuple with the ExpirationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSource

`func (o *FeedIndicator) SetExpirationSource(v ExpirationSource)`

SetExpirationSource sets ExpirationSource field to given value.

### HasExpirationSource

`func (o *FeedIndicator) HasExpirationSource() bool`

HasExpirationSource returns a boolean if a field has been set.

### GetBypassExclusionList

`func (o *FeedIndicator) GetBypassExclusionList() bool`

GetBypassExclusionList returns the BypassExclusionList field if non-nil, zero value otherwise.

### GetBypassExclusionListOk

`func (o *FeedIndicator) GetBypassExclusionListOk() (*bool, bool)`

GetBypassExclusionListOk returns a tuple with the BypassExclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassExclusionList

`func (o *FeedIndicator) SetBypassExclusionList(v bool)`

SetBypassExclusionList sets BypassExclusionList field to given value.

### HasBypassExclusionList

`func (o *FeedIndicator) HasBypassExclusionList() bool`

HasBypassExclusionList returns a boolean if a field has been set.

### GetClassifierId

`func (o *FeedIndicator) GetClassifierId() string`

GetClassifierId returns the ClassifierId field if non-nil, zero value otherwise.

### GetClassifierIdOk

`func (o *FeedIndicator) GetClassifierIdOk() (*string, bool)`

GetClassifierIdOk returns a tuple with the ClassifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierId

`func (o *FeedIndicator) SetClassifierId(v string)`

SetClassifierId sets ClassifierId field to given value.

### HasClassifierId

`func (o *FeedIndicator) HasClassifierId() bool`

HasClassifierId returns a boolean if a field has been set.

### GetClassifierVersion

`func (o *FeedIndicator) GetClassifierVersion() int64`

GetClassifierVersion returns the ClassifierVersion field if non-nil, zero value otherwise.

### GetClassifierVersionOk

`func (o *FeedIndicator) GetClassifierVersionOk() (*int64, bool)`

GetClassifierVersionOk returns a tuple with the ClassifierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierVersion

`func (o *FeedIndicator) SetClassifierVersion(v int64)`

SetClassifierVersion sets ClassifierVersion field to given value.

### HasClassifierVersion

`func (o *FeedIndicator) HasClassifierVersion() bool`

HasClassifierVersion returns a boolean if a field has been set.

### GetComments

`func (o *FeedIndicator) GetComments() []FeedIndicatorComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *FeedIndicator) GetCommentsOk() (*[]FeedIndicatorComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *FeedIndicator) SetComments(v []FeedIndicatorComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *FeedIndicator) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetExpirationInterval

`func (o *FeedIndicator) GetExpirationInterval() int64`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *FeedIndicator) GetExpirationIntervalOk() (*int64, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *FeedIndicator) SetExpirationInterval(v int64)`

SetExpirationInterval sets ExpirationInterval field to given value.

### HasExpirationInterval

`func (o *FeedIndicator) HasExpirationInterval() bool`

HasExpirationInterval returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *FeedIndicator) GetExpirationPolicy() string`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *FeedIndicator) GetExpirationPolicyOk() (*string, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *FeedIndicator) SetExpirationPolicy(v string)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *FeedIndicator) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetFetchTime

`func (o *FeedIndicator) GetFetchTime() time.Time`

GetFetchTime returns the FetchTime field if non-nil, zero value otherwise.

### GetFetchTimeOk

`func (o *FeedIndicator) GetFetchTimeOk() (*time.Time, bool)`

GetFetchTimeOk returns a tuple with the FetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTime

`func (o *FeedIndicator) SetFetchTime(v time.Time)`

SetFetchTime sets FetchTime field to given value.

### HasFetchTime

`func (o *FeedIndicator) HasFetchTime() bool`

HasFetchTime returns a boolean if a field has been set.

### GetFields

`func (o *FeedIndicator) GetFields() CustomFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FeedIndicator) GetFieldsOk() (*CustomFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FeedIndicator) SetFields(v CustomFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FeedIndicator) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIsEnrichment

`func (o *FeedIndicator) GetIsEnrichment() bool`

GetIsEnrichment returns the IsEnrichment field if non-nil, zero value otherwise.

### GetIsEnrichmentOk

`func (o *FeedIndicator) GetIsEnrichmentOk() (*bool, bool)`

GetIsEnrichmentOk returns a tuple with the IsEnrichment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnrichment

`func (o *FeedIndicator) SetIsEnrichment(v bool)`

SetIsEnrichment sets IsEnrichment field to given value.

### HasIsEnrichment

`func (o *FeedIndicator) HasIsEnrichment() bool`

HasIsEnrichment returns a boolean if a field has been set.

### GetMapperId

`func (o *FeedIndicator) GetMapperId() string`

GetMapperId returns the MapperId field if non-nil, zero value otherwise.

### GetMapperIdOk

`func (o *FeedIndicator) GetMapperIdOk() (*string, bool)`

GetMapperIdOk returns a tuple with the MapperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperId

`func (o *FeedIndicator) SetMapperId(v string)`

SetMapperId sets MapperId field to given value.

### HasMapperId

`func (o *FeedIndicator) HasMapperId() bool`

HasMapperId returns a boolean if a field has been set.

### GetMapperVersion

`func (o *FeedIndicator) GetMapperVersion() int64`

GetMapperVersion returns the MapperVersion field if non-nil, zero value otherwise.

### GetMapperVersionOk

`func (o *FeedIndicator) GetMapperVersionOk() (*int64, bool)`

GetMapperVersionOk returns a tuple with the MapperVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperVersion

`func (o *FeedIndicator) SetMapperVersion(v int64)`

SetMapperVersion sets MapperVersion field to given value.

### HasMapperVersion

`func (o *FeedIndicator) HasMapperVersion() bool`

HasMapperVersion returns a boolean if a field has been set.

### GetModifiedTime

`func (o *FeedIndicator) GetModifiedTime() time.Time`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *FeedIndicator) GetModifiedTimeOk() (*time.Time, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *FeedIndicator) SetModifiedTime(v time.Time)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *FeedIndicator) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetModuleId

`func (o *FeedIndicator) GetModuleId() string`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *FeedIndicator) GetModuleIdOk() (*string, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *FeedIndicator) SetModuleId(v string)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *FeedIndicator) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetRawJSON

`func (o *FeedIndicator) GetRawJSON() map[string]map[string]interface{}`

GetRawJSON returns the RawJSON field if non-nil, zero value otherwise.

### GetRawJSONOk

`func (o *FeedIndicator) GetRawJSONOk() (*map[string]map[string]interface{}, bool)`

GetRawJSONOk returns a tuple with the RawJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawJSON

`func (o *FeedIndicator) SetRawJSON(v map[string]map[string]interface{})`

SetRawJSON sets RawJSON field to given value.

### HasRawJSON

`func (o *FeedIndicator) HasRawJSON() bool`

HasRawJSON returns a boolean if a field has been set.

### GetRelationships

`func (o *FeedIndicator) GetRelationships() RelationshipsAPI`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FeedIndicator) GetRelationshipsOk() (*RelationshipsAPI, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FeedIndicator) SetRelationships(v RelationshipsAPI)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FeedIndicator) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetReliability

`func (o *FeedIndicator) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *FeedIndicator) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *FeedIndicator) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *FeedIndicator) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetScore

`func (o *FeedIndicator) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FeedIndicator) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FeedIndicator) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *FeedIndicator) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSourceBrand

`func (o *FeedIndicator) GetSourceBrand() string`

GetSourceBrand returns the SourceBrand field if non-nil, zero value otherwise.

### GetSourceBrandOk

`func (o *FeedIndicator) GetSourceBrandOk() (*string, bool)`

GetSourceBrandOk returns a tuple with the SourceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrand

`func (o *FeedIndicator) SetSourceBrand(v string)`

SetSourceBrand sets SourceBrand field to given value.

### HasSourceBrand

`func (o *FeedIndicator) HasSourceBrand() bool`

HasSourceBrand returns a boolean if a field has been set.

### GetSourceInstance

`func (o *FeedIndicator) GetSourceInstance() string`

GetSourceInstance returns the SourceInstance field if non-nil, zero value otherwise.

### GetSourceInstanceOk

`func (o *FeedIndicator) GetSourceInstanceOk() (*string, bool)`

GetSourceInstanceOk returns a tuple with the SourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstance

`func (o *FeedIndicator) SetSourceInstance(v string)`

SetSourceInstance sets SourceInstance field to given value.

### HasSourceInstance

`func (o *FeedIndicator) HasSourceInstance() bool`

HasSourceInstance returns a boolean if a field has been set.

### GetTimestamp

`func (o *FeedIndicator) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FeedIndicator) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FeedIndicator) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FeedIndicator) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *FeedIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedIndicator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeedIndicator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *FeedIndicator) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeedIndicator) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeedIndicator) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeedIndicator) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


