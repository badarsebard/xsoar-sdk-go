# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | color used to identify the group | [optional] 
**Count** | Pointer to **int64** | The number of participants in the group | [optional] 
**Data** | Pointer to **[]int64** | The data value provided in array of integer values. | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**FloatData** | Pointer to **[]float64** | The data value provided in array of float values. | [optional] 
**Groups** | Pointer to [**Groups**](Groups.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pivot** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 
**Z** | Pointer to **float64** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *Group) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Group) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Group) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Group) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCount

`func (o *Group) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Group) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Group) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Group) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *Group) GetData() []int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Group) GetDataOk() (*[]int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Group) SetData(v []int64)`

SetData sets Data field to given value.

### HasData

`func (o *Group) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDataType

`func (o *Group) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Group) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Group) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Group) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetFloatData

`func (o *Group) GetFloatData() []float64`

GetFloatData returns the FloatData field if non-nil, zero value otherwise.

### GetFloatDataOk

`func (o *Group) GetFloatDataOk() (*[]float64, bool)`

GetFloatDataOk returns a tuple with the FloatData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatData

`func (o *Group) SetFloatData(v []float64)`

SetFloatData sets FloatData field to given value.

### HasFloatData

`func (o *Group) HasFloatData() bool`

HasFloatData returns a boolean if a field has been set.

### GetGroups

`func (o *Group) GetGroups() Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Group) GetGroupsOk() (*Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Group) SetGroups(v Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Group) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPivot

`func (o *Group) GetPivot() string`

GetPivot returns the Pivot field if non-nil, zero value otherwise.

### GetPivotOk

`func (o *Group) GetPivotOk() (*string, bool)`

GetPivotOk returns a tuple with the Pivot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivot

`func (o *Group) SetPivot(v string)`

SetPivot sets Pivot field to given value.

### HasPivot

`func (o *Group) HasPivot() bool`

HasPivot returns a boolean if a field has been set.

### GetQuery

`func (o *Group) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Group) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Group) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Group) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetX

`func (o *Group) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Group) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Group) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *Group) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *Group) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Group) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Group) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *Group) HasY() bool`

HasY returns a boolean if a field has been set.

### GetZ

`func (o *Group) GetZ() float64`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *Group) GetZOk() (*float64, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *Group) SetZ(v float64)`

SetZ sets Z field to given value.

### HasZ

`func (o *Group) HasZ() bool`

HasZ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


