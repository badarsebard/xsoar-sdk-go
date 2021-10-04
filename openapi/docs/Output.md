# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentPath** | Pointer to **string** |  | [optional] 
**ContextPath** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **map[string]interface{}** | Description is either a string or a map from string to interface | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOutput

`func NewOutput() *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentPath

`func (o *Output) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *Output) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *Output) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *Output) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetContextPath

`func (o *Output) GetContextPath() string`

GetContextPath returns the ContextPath field if non-nil, zero value otherwise.

### GetContextPathOk

`func (o *Output) GetContextPathOk() (*string, bool)`

GetContextPathOk returns a tuple with the ContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextPath

`func (o *Output) SetContextPath(v string)`

SetContextPath sets ContextPath field to given value.

### HasContextPath

`func (o *Output) HasContextPath() bool`

HasContextPath returns a boolean if a field has been set.

### GetDescription

`func (o *Output) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Output) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Output) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Output) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Output) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Output) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


