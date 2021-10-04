# StatsResponseWithReferenceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]Group**](Group.md) | Groups is a list of group entities | [optional] 
**ReferenceLineY** | Pointer to **float64** |  | [optional] 

## Methods

### NewStatsResponseWithReferenceLine

`func NewStatsResponseWithReferenceLine() *StatsResponseWithReferenceLine`

NewStatsResponseWithReferenceLine instantiates a new StatsResponseWithReferenceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsResponseWithReferenceLineWithDefaults

`func NewStatsResponseWithReferenceLineWithDefaults() *StatsResponseWithReferenceLine`

NewStatsResponseWithReferenceLineWithDefaults instantiates a new StatsResponseWithReferenceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *StatsResponseWithReferenceLine) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *StatsResponseWithReferenceLine) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *StatsResponseWithReferenceLine) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *StatsResponseWithReferenceLine) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetReferenceLineY

`func (o *StatsResponseWithReferenceLine) GetReferenceLineY() float64`

GetReferenceLineY returns the ReferenceLineY field if non-nil, zero value otherwise.

### GetReferenceLineYOk

`func (o *StatsResponseWithReferenceLine) GetReferenceLineYOk() (*float64, bool)`

GetReferenceLineYOk returns a tuple with the ReferenceLineY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceLineY

`func (o *StatsResponseWithReferenceLine) SetReferenceLineY(v float64)`

SetReferenceLineY sets ReferenceLineY field to given value.

### HasReferenceLineY

`func (o *StatsResponseWithReferenceLine) HasReferenceLineY() bool`

HasReferenceLineY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


