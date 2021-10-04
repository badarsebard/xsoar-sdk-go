# UpdateIndicatorBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Filter** | Pointer to [**IndicatorFilter**](IndicatorFilter.md) |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateIndicatorBatch

`func NewUpdateIndicatorBatch() *UpdateIndicatorBatch`

NewUpdateIndicatorBatch instantiates a new UpdateIndicatorBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIndicatorBatchWithDefaults

`func NewUpdateIndicatorBatchWithDefaults() *UpdateIndicatorBatch`

NewUpdateIndicatorBatchWithDefaults instantiates a new UpdateIndicatorBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateIndicatorBatch) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateIndicatorBatch) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateIndicatorBatch) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateIndicatorBatch) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetData

`func (o *UpdateIndicatorBatch) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateIndicatorBatch) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateIndicatorBatch) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UpdateIndicatorBatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateIndicatorBatch) GetFilter() IndicatorFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateIndicatorBatch) GetFilterOk() (*IndicatorFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateIndicatorBatch) SetFilter(v IndicatorFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateIndicatorBatch) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIds

`func (o *UpdateIndicatorBatch) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateIndicatorBatch) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateIndicatorBatch) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UpdateIndicatorBatch) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


