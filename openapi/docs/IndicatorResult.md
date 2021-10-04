# IndicatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IocObjects** | Pointer to [**[]IocObject**](IocObject.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewIndicatorResult

`func NewIndicatorResult() *IndicatorResult`

NewIndicatorResult instantiates a new IndicatorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorResultWithDefaults

`func NewIndicatorResultWithDefaults() *IndicatorResult`

NewIndicatorResultWithDefaults instantiates a new IndicatorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIocObjects

`func (o *IndicatorResult) GetIocObjects() []IocObject`

GetIocObjects returns the IocObjects field if non-nil, zero value otherwise.

### GetIocObjectsOk

`func (o *IndicatorResult) GetIocObjectsOk() (*[]IocObject, bool)`

GetIocObjectsOk returns a tuple with the IocObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocObjects

`func (o *IndicatorResult) SetIocObjects(v []IocObject)`

SetIocObjects sets IocObjects field to given value.

### HasIocObjects

`func (o *IndicatorResult) HasIocObjects() bool`

HasIocObjects returns a boolean if a field has been set.

### GetTotal

`func (o *IndicatorResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IndicatorResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IndicatorResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IndicatorResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


