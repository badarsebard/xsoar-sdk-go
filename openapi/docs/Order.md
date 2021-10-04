# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asc** | Pointer to **bool** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**FieldType** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsc

`func (o *Order) GetAsc() bool`

GetAsc returns the Asc field if non-nil, zero value otherwise.

### GetAscOk

`func (o *Order) GetAscOk() (*bool, bool)`

GetAscOk returns a tuple with the Asc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsc

`func (o *Order) SetAsc(v bool)`

SetAsc sets Asc field to given value.

### HasAsc

`func (o *Order) HasAsc() bool`

HasAsc returns a boolean if a field has been set.

### GetField

`func (o *Order) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Order) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Order) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Order) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFieldType

`func (o *Order) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *Order) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *Order) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *Order) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


