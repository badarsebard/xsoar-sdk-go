# OperatorArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsContext** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewOperatorArgument

`func NewOperatorArgument() *OperatorArgument`

NewOperatorArgument instantiates a new OperatorArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorArgumentWithDefaults

`func NewOperatorArgumentWithDefaults() *OperatorArgument`

NewOperatorArgumentWithDefaults instantiates a new OperatorArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsContext

`func (o *OperatorArgument) GetIsContext() bool`

GetIsContext returns the IsContext field if non-nil, zero value otherwise.

### GetIsContextOk

`func (o *OperatorArgument) GetIsContextOk() (*bool, bool)`

GetIsContextOk returns a tuple with the IsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContext

`func (o *OperatorArgument) SetIsContext(v bool)`

SetIsContext sets IsContext field to given value.

### HasIsContext

`func (o *OperatorArgument) HasIsContext() bool`

HasIsContext returns a boolean if a field has been set.

### GetValue

`func (o *OperatorArgument) GetValue() AdvanceArg`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OperatorArgument) GetValueOk() (*AdvanceArg, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OperatorArgument) SetValue(v AdvanceArg)`

SetValue sets Value field to given value.

### HasValue

`func (o *OperatorArgument) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


