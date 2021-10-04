# IncidentSearchResponseWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IncidentWrapper**](IncidentWrapper.md) | in: body | [optional] 
**NotUpdated** | Pointer to **int32** |  | [optional] 
**SearchAfter** | Pointer to **[]string** |  | [optional] 
**SearchAfterElastic** | Pointer to **[]string** |  | [optional] 
**SearchBefore** | Pointer to **[]string** |  | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewIncidentSearchResponseWrapper

`func NewIncidentSearchResponseWrapper() *IncidentSearchResponseWrapper`

NewIncidentSearchResponseWrapper instantiates a new IncidentSearchResponseWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSearchResponseWrapperWithDefaults

`func NewIncidentSearchResponseWrapperWithDefaults() *IncidentSearchResponseWrapper`

NewIncidentSearchResponseWrapperWithDefaults instantiates a new IncidentSearchResponseWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentSearchResponseWrapper) GetData() []IncidentWrapper`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentSearchResponseWrapper) GetDataOk() (*[]IncidentWrapper, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentSearchResponseWrapper) SetData(v []IncidentWrapper)`

SetData sets Data field to given value.

### HasData

`func (o *IncidentSearchResponseWrapper) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNotUpdated

`func (o *IncidentSearchResponseWrapper) GetNotUpdated() int32`

GetNotUpdated returns the NotUpdated field if non-nil, zero value otherwise.

### GetNotUpdatedOk

`func (o *IncidentSearchResponseWrapper) GetNotUpdatedOk() (*int32, bool)`

GetNotUpdatedOk returns a tuple with the NotUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotUpdated

`func (o *IncidentSearchResponseWrapper) SetNotUpdated(v int32)`

SetNotUpdated sets NotUpdated field to given value.

### HasNotUpdated

`func (o *IncidentSearchResponseWrapper) HasNotUpdated() bool`

HasNotUpdated returns a boolean if a field has been set.

### GetSearchAfter

`func (o *IncidentSearchResponseWrapper) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *IncidentSearchResponseWrapper) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *IncidentSearchResponseWrapper) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *IncidentSearchResponseWrapper) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *IncidentSearchResponseWrapper) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *IncidentSearchResponseWrapper) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *IncidentSearchResponseWrapper) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *IncidentSearchResponseWrapper) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *IncidentSearchResponseWrapper) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *IncidentSearchResponseWrapper) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *IncidentSearchResponseWrapper) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *IncidentSearchResponseWrapper) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *IncidentSearchResponseWrapper) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *IncidentSearchResponseWrapper) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *IncidentSearchResponseWrapper) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *IncidentSearchResponseWrapper) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetTotal

`func (o *IncidentSearchResponseWrapper) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncidentSearchResponseWrapper) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncidentSearchResponseWrapper) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IncidentSearchResponseWrapper) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


