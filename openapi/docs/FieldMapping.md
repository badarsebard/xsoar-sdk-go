# FieldMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **string** |  | [optional] 
**Output** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewFieldMapping

`func NewFieldMapping() *FieldMapping`

NewFieldMapping instantiates a new FieldMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingWithDefaults

`func NewFieldMappingWithDefaults() *FieldMapping`

NewFieldMappingWithDefaults instantiates a new FieldMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *FieldMapping) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *FieldMapping) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *FieldMapping) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *FieldMapping) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetOutput

`func (o *FieldMapping) GetOutput() AdvanceArg`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *FieldMapping) GetOutputOk() (*AdvanceArg, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *FieldMapping) SetOutput(v AdvanceArg)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *FieldMapping) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


