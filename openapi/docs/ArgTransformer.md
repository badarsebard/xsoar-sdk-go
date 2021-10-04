# ArgTransformer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to [**map[string]OperatorArgument**](OperatorArgument.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewArgTransformer

`func NewArgTransformer() *ArgTransformer`

NewArgTransformer instantiates a new ArgTransformer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgTransformerWithDefaults

`func NewArgTransformerWithDefaults() *ArgTransformer`

NewArgTransformerWithDefaults instantiates a new ArgTransformer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ArgTransformer) GetArgs() map[string]OperatorArgument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ArgTransformer) GetArgsOk() (*map[string]OperatorArgument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ArgTransformer) SetArgs(v map[string]OperatorArgument)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ArgTransformer) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetOperator

`func (o *ArgTransformer) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ArgTransformer) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ArgTransformer) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ArgTransformer) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


