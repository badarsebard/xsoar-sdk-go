# Mapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DontMapEventToLabels** | Pointer to **bool** | DontMapEventToLabels by default we will map all the fields of the event to incident labels | [optional] 
**InternalMapping** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewMapper

`func NewMapper() *Mapper`

NewMapper instantiates a new Mapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapperWithDefaults

`func NewMapperWithDefaults() *Mapper`

NewMapperWithDefaults instantiates a new Mapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDontMapEventToLabels

`func (o *Mapper) GetDontMapEventToLabels() bool`

GetDontMapEventToLabels returns the DontMapEventToLabels field if non-nil, zero value otherwise.

### GetDontMapEventToLabelsOk

`func (o *Mapper) GetDontMapEventToLabelsOk() (*bool, bool)`

GetDontMapEventToLabelsOk returns a tuple with the DontMapEventToLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontMapEventToLabels

`func (o *Mapper) SetDontMapEventToLabels(v bool)`

SetDontMapEventToLabels sets DontMapEventToLabels field to given value.

### HasDontMapEventToLabels

`func (o *Mapper) HasDontMapEventToLabels() bool`

HasDontMapEventToLabels returns a boolean if a field has been set.

### GetInternalMapping

`func (o *Mapper) GetInternalMapping() map[string]AdvanceArg`

GetInternalMapping returns the InternalMapping field if non-nil, zero value otherwise.

### GetInternalMappingOk

`func (o *Mapper) GetInternalMappingOk() (*map[string]AdvanceArg, bool)`

GetInternalMappingOk returns a tuple with the InternalMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMapping

`func (o *Mapper) SetInternalMapping(v map[string]AdvanceArg)`

SetInternalMapping sets InternalMapping field to given value.

### HasInternalMapping

`func (o *Mapper) HasInternalMapping() bool`

HasInternalMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


