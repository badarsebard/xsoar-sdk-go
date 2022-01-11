# FeedIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BypassExclusionList** | Pointer to **bool** |  | [optional] 
**ClassifierId** | Pointer to **string** |  | [optional] 
**Indicators** | Pointer to [**[]RawFeedIndicator**](RawFeedIndicator.md) |  | [optional] 
**MapperId** | Pointer to **string** |  | [optional] 

## Methods

### NewFeedIndicatorsRequest

`func NewFeedIndicatorsRequest() *FeedIndicatorsRequest`

NewFeedIndicatorsRequest instantiates a new FeedIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedIndicatorsRequestWithDefaults

`func NewFeedIndicatorsRequestWithDefaults() *FeedIndicatorsRequest`

NewFeedIndicatorsRequestWithDefaults instantiates a new FeedIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBypassExclusionList

`func (o *FeedIndicatorsRequest) GetBypassExclusionList() bool`

GetBypassExclusionList returns the BypassExclusionList field if non-nil, zero value otherwise.

### GetBypassExclusionListOk

`func (o *FeedIndicatorsRequest) GetBypassExclusionListOk() (*bool, bool)`

GetBypassExclusionListOk returns a tuple with the BypassExclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassExclusionList

`func (o *FeedIndicatorsRequest) SetBypassExclusionList(v bool)`

SetBypassExclusionList sets BypassExclusionList field to given value.

### HasBypassExclusionList

`func (o *FeedIndicatorsRequest) HasBypassExclusionList() bool`

HasBypassExclusionList returns a boolean if a field has been set.

### GetClassifierId

`func (o *FeedIndicatorsRequest) GetClassifierId() string`

GetClassifierId returns the ClassifierId field if non-nil, zero value otherwise.

### GetClassifierIdOk

`func (o *FeedIndicatorsRequest) GetClassifierIdOk() (*string, bool)`

GetClassifierIdOk returns a tuple with the ClassifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierId

`func (o *FeedIndicatorsRequest) SetClassifierId(v string)`

SetClassifierId sets ClassifierId field to given value.

### HasClassifierId

`func (o *FeedIndicatorsRequest) HasClassifierId() bool`

HasClassifierId returns a boolean if a field has been set.

### GetIndicators

`func (o *FeedIndicatorsRequest) GetIndicators() []RawFeedIndicator`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *FeedIndicatorsRequest) GetIndicatorsOk() (*[]RawFeedIndicator, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *FeedIndicatorsRequest) SetIndicators(v []RawFeedIndicator)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *FeedIndicatorsRequest) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.

### GetMapperId

`func (o *FeedIndicatorsRequest) GetMapperId() string`

GetMapperId returns the MapperId field if non-nil, zero value otherwise.

### GetMapperIdOk

`func (o *FeedIndicatorsRequest) GetMapperIdOk() (*string, bool)`

GetMapperIdOk returns a tuple with the MapperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperId

`func (o *FeedIndicatorsRequest) SetMapperId(v string)`

SetMapperId sets MapperId field to given value.

### HasMapperId

`func (o *FeedIndicatorsRequest) HasMapperId() bool`

HasMapperId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


