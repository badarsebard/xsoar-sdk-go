# QueryState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchSize** | Pointer to **int64** |  | [optional] 
**CurrentPage** | Pointer to **int64** |  | [optional] 
**SearchAfter** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryState

`func NewQueryState() *QueryState`

NewQueryState instantiates a new QueryState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStateWithDefaults

`func NewQueryStateWithDefaults() *QueryState`

NewQueryStateWithDefaults instantiates a new QueryState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchSize

`func (o *QueryState) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *QueryState) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *QueryState) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *QueryState) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetCurrentPage

`func (o *QueryState) GetCurrentPage() int64`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *QueryState) GetCurrentPageOk() (*int64, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *QueryState) SetCurrentPage(v int64)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *QueryState) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetSearchAfter

`func (o *QueryState) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *QueryState) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *QueryState) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *QueryState) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetTotalResults

`func (o *QueryState) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *QueryState) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *QueryState) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *QueryState) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


