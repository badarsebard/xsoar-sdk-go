# SearchStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidences** | Pointer to [**[]SearchStats**](SearchStats.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchStatsResponse

`func NewSearchStatsResponse() *SearchStatsResponse`

NewSearchStatsResponse instantiates a new SearchStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsResponseWithDefaults

`func NewSearchStatsResponseWithDefaults() *SearchStatsResponse`

NewSearchStatsResponseWithDefaults instantiates a new SearchStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidences

`func (o *SearchStatsResponse) GetEvidences() []SearchStats`

GetEvidences returns the Evidences field if non-nil, zero value otherwise.

### GetEvidencesOk

`func (o *SearchStatsResponse) GetEvidencesOk() (*[]SearchStats, bool)`

GetEvidencesOk returns a tuple with the Evidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidences

`func (o *SearchStatsResponse) SetEvidences(v []SearchStats)`

SetEvidences sets Evidences field to given value.

### HasEvidences

`func (o *SearchStatsResponse) HasEvidences() bool`

HasEvidences returns a boolean if a field has been set.

### GetTotal

`func (o *SearchStatsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchStatsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchStatsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchStatsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


