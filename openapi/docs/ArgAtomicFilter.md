# ArgAtomicFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Left** | Pointer to [**OperatorArgument**](OperatorArgument.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Right** | Pointer to [**OperatorArgument**](OperatorArgument.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewArgAtomicFilter

`func NewArgAtomicFilter() *ArgAtomicFilter`

NewArgAtomicFilter instantiates a new ArgAtomicFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgAtomicFilterWithDefaults

`func NewArgAtomicFilterWithDefaults() *ArgAtomicFilter`

NewArgAtomicFilterWithDefaults instantiates a new ArgAtomicFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreCase

`func (o *ArgAtomicFilter) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *ArgAtomicFilter) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *ArgAtomicFilter) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *ArgAtomicFilter) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetLeft

`func (o *ArgAtomicFilter) GetLeft() OperatorArgument`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *ArgAtomicFilter) GetLeftOk() (*OperatorArgument, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *ArgAtomicFilter) SetLeft(v OperatorArgument)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *ArgAtomicFilter) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetOperator

`func (o *ArgAtomicFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ArgAtomicFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ArgAtomicFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ArgAtomicFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetRight

`func (o *ArgAtomicFilter) GetRight() OperatorArgument`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ArgAtomicFilter) GetRightOk() (*OperatorArgument, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ArgAtomicFilter) SetRight(v OperatorArgument)`

SetRight sets Right field to given value.

### HasRight

`func (o *ArgAtomicFilter) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetType

`func (o *ArgAtomicFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgAtomicFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgAtomicFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArgAtomicFilter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


