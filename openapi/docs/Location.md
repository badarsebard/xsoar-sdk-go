# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayPositions** | Pointer to **[]int32** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Pos** | Pointer to **int32** | Pos is the position of the term within the field, starting at 1 | [optional] 
**Start** | Pointer to **int32** | Start and End are the byte offsets of the term in the field | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayPositions

`func (o *Location) GetArrayPositions() []int32`

GetArrayPositions returns the ArrayPositions field if non-nil, zero value otherwise.

### GetArrayPositionsOk

`func (o *Location) GetArrayPositionsOk() (*[]int32, bool)`

GetArrayPositionsOk returns a tuple with the ArrayPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayPositions

`func (o *Location) SetArrayPositions(v []int32)`

SetArrayPositions sets ArrayPositions field to given value.

### HasArrayPositions

`func (o *Location) HasArrayPositions() bool`

HasArrayPositions returns a boolean if a field has been set.

### GetEnd

`func (o *Location) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Location) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Location) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Location) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPos

`func (o *Location) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *Location) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *Location) SetPos(v int32)`

SetPos sets Pos field to given value.

### HasPos

`func (o *Location) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetStart

`func (o *Location) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Location) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Location) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Location) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


