# ComplexArg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]ArgFilter**](ArgFilter.md) |  | [optional] 
**Root** | Pointer to **string** |  | [optional] 
**Transformers** | Pointer to [**[]ArgTransformer**](ArgTransformer.md) |  | [optional] 

## Methods

### NewComplexArg

`func NewComplexArg() *ComplexArg`

NewComplexArg instantiates a new ComplexArg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexArgWithDefaults

`func NewComplexArgWithDefaults() *ComplexArg`

NewComplexArgWithDefaults instantiates a new ComplexArg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessor

`func (o *ComplexArg) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *ComplexArg) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *ComplexArg) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.

### HasAccessor

`func (o *ComplexArg) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.

### GetFilters

`func (o *ComplexArg) GetFilters() []ArgFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ComplexArg) GetFiltersOk() (*[]ArgFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ComplexArg) SetFilters(v []ArgFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ComplexArg) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRoot

`func (o *ComplexArg) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ComplexArg) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ComplexArg) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *ComplexArg) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTransformers

`func (o *ComplexArg) GetTransformers() []ArgTransformer`

GetTransformers returns the Transformers field if non-nil, zero value otherwise.

### GetTransformersOk

`func (o *ComplexArg) GetTransformersOk() (*[]ArgTransformer, bool)`

GetTransformersOk returns a tuple with the Transformers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformers

`func (o *ComplexArg) SetTransformers(v []ArgTransformer)`

SetTransformers sets Transformers field to given value.

### HasTransformers

`func (o *ComplexArg) HasTransformers() bool`

HasTransformers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


