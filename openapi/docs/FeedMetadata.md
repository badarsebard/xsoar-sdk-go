# FeedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BypassExclusionList** | Pointer to **bool** |  | [optional] 
**ClassifierId** | Pointer to **string** |  | [optional] 
**ClassifierVersion** | Pointer to **int64** |  | [optional] 
**ExpirationInterval** | Pointer to **int64** |  | [optional] 
**ExpirationPolicy** | Pointer to **string** |  | [optional] 
**FetchTime** | Pointer to **time.Time** |  | [optional] 
**MapperId** | Pointer to **string** |  | [optional] 
**MapperVersion** | Pointer to **int64** |  | [optional] 
**ModuleId** | Pointer to **string** |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**SourceBrand** | Pointer to **string** |  | [optional] 
**SourceInstance** | Pointer to **string** |  | [optional] 

## Methods

### NewFeedMetadata

`func NewFeedMetadata() *FeedMetadata`

NewFeedMetadata instantiates a new FeedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedMetadataWithDefaults

`func NewFeedMetadataWithDefaults() *FeedMetadata`

NewFeedMetadataWithDefaults instantiates a new FeedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBypassExclusionList

`func (o *FeedMetadata) GetBypassExclusionList() bool`

GetBypassExclusionList returns the BypassExclusionList field if non-nil, zero value otherwise.

### GetBypassExclusionListOk

`func (o *FeedMetadata) GetBypassExclusionListOk() (*bool, bool)`

GetBypassExclusionListOk returns a tuple with the BypassExclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassExclusionList

`func (o *FeedMetadata) SetBypassExclusionList(v bool)`

SetBypassExclusionList sets BypassExclusionList field to given value.

### HasBypassExclusionList

`func (o *FeedMetadata) HasBypassExclusionList() bool`

HasBypassExclusionList returns a boolean if a field has been set.

### GetClassifierId

`func (o *FeedMetadata) GetClassifierId() string`

GetClassifierId returns the ClassifierId field if non-nil, zero value otherwise.

### GetClassifierIdOk

`func (o *FeedMetadata) GetClassifierIdOk() (*string, bool)`

GetClassifierIdOk returns a tuple with the ClassifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierId

`func (o *FeedMetadata) SetClassifierId(v string)`

SetClassifierId sets ClassifierId field to given value.

### HasClassifierId

`func (o *FeedMetadata) HasClassifierId() bool`

HasClassifierId returns a boolean if a field has been set.

### GetClassifierVersion

`func (o *FeedMetadata) GetClassifierVersion() int64`

GetClassifierVersion returns the ClassifierVersion field if non-nil, zero value otherwise.

### GetClassifierVersionOk

`func (o *FeedMetadata) GetClassifierVersionOk() (*int64, bool)`

GetClassifierVersionOk returns a tuple with the ClassifierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierVersion

`func (o *FeedMetadata) SetClassifierVersion(v int64)`

SetClassifierVersion sets ClassifierVersion field to given value.

### HasClassifierVersion

`func (o *FeedMetadata) HasClassifierVersion() bool`

HasClassifierVersion returns a boolean if a field has been set.

### GetExpirationInterval

`func (o *FeedMetadata) GetExpirationInterval() int64`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *FeedMetadata) GetExpirationIntervalOk() (*int64, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *FeedMetadata) SetExpirationInterval(v int64)`

SetExpirationInterval sets ExpirationInterval field to given value.

### HasExpirationInterval

`func (o *FeedMetadata) HasExpirationInterval() bool`

HasExpirationInterval returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *FeedMetadata) GetExpirationPolicy() string`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *FeedMetadata) GetExpirationPolicyOk() (*string, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *FeedMetadata) SetExpirationPolicy(v string)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *FeedMetadata) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetFetchTime

`func (o *FeedMetadata) GetFetchTime() time.Time`

GetFetchTime returns the FetchTime field if non-nil, zero value otherwise.

### GetFetchTimeOk

`func (o *FeedMetadata) GetFetchTimeOk() (*time.Time, bool)`

GetFetchTimeOk returns a tuple with the FetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTime

`func (o *FeedMetadata) SetFetchTime(v time.Time)`

SetFetchTime sets FetchTime field to given value.

### HasFetchTime

`func (o *FeedMetadata) HasFetchTime() bool`

HasFetchTime returns a boolean if a field has been set.

### GetMapperId

`func (o *FeedMetadata) GetMapperId() string`

GetMapperId returns the MapperId field if non-nil, zero value otherwise.

### GetMapperIdOk

`func (o *FeedMetadata) GetMapperIdOk() (*string, bool)`

GetMapperIdOk returns a tuple with the MapperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperId

`func (o *FeedMetadata) SetMapperId(v string)`

SetMapperId sets MapperId field to given value.

### HasMapperId

`func (o *FeedMetadata) HasMapperId() bool`

HasMapperId returns a boolean if a field has been set.

### GetMapperVersion

`func (o *FeedMetadata) GetMapperVersion() int64`

GetMapperVersion returns the MapperVersion field if non-nil, zero value otherwise.

### GetMapperVersionOk

`func (o *FeedMetadata) GetMapperVersionOk() (*int64, bool)`

GetMapperVersionOk returns a tuple with the MapperVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperVersion

`func (o *FeedMetadata) SetMapperVersion(v int64)`

SetMapperVersion sets MapperVersion field to given value.

### HasMapperVersion

`func (o *FeedMetadata) HasMapperVersion() bool`

HasMapperVersion returns a boolean if a field has been set.

### GetModuleId

`func (o *FeedMetadata) GetModuleId() string`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *FeedMetadata) GetModuleIdOk() (*string, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *FeedMetadata) SetModuleId(v string)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *FeedMetadata) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetReliability

`func (o *FeedMetadata) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *FeedMetadata) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *FeedMetadata) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *FeedMetadata) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetScore

`func (o *FeedMetadata) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FeedMetadata) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FeedMetadata) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *FeedMetadata) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSourceBrand

`func (o *FeedMetadata) GetSourceBrand() string`

GetSourceBrand returns the SourceBrand field if non-nil, zero value otherwise.

### GetSourceBrandOk

`func (o *FeedMetadata) GetSourceBrandOk() (*string, bool)`

GetSourceBrandOk returns a tuple with the SourceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrand

`func (o *FeedMetadata) SetSourceBrand(v string)`

SetSourceBrand sets SourceBrand field to given value.

### HasSourceBrand

`func (o *FeedMetadata) HasSourceBrand() bool`

HasSourceBrand returns a boolean if a field has been set.

### GetSourceInstance

`func (o *FeedMetadata) GetSourceInstance() string`

GetSourceInstance returns the SourceInstance field if non-nil, zero value otherwise.

### GetSourceInstanceOk

`func (o *FeedMetadata) GetSourceInstanceOk() (*string, bool)`

GetSourceInstanceOk returns a tuple with the SourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstance

`func (o *FeedMetadata) SetSourceInstance(v string)`

SetSourceInstance sets SourceInstance field to given value.

### HasSourceInstance

`func (o *FeedMetadata) HasSourceInstance() bool`

HasSourceInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


