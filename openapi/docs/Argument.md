# Argument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsArray** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Predefined** | Pointer to **[]string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Secret** | Pointer to **bool** |  | [optional] 

## Methods

### NewArgument

`func NewArgument() *Argument`

NewArgument instantiates a new Argument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentWithDefaults

`func NewArgumentWithDefaults() *Argument`

NewArgumentWithDefaults instantiates a new Argument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *Argument) GetAuto() string`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *Argument) GetAutoOk() (*string, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *Argument) SetAuto(v string)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *Argument) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetDefault

`func (o *Argument) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Argument) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Argument) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Argument) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultValue

`func (o *Argument) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Argument) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Argument) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *Argument) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDeprecated

`func (o *Argument) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Argument) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Argument) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Argument) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Argument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Argument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Argument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Argument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsArray

`func (o *Argument) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *Argument) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *Argument) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *Argument) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetName

`func (o *Argument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Argument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Argument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Argument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPredefined

`func (o *Argument) GetPredefined() []string`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *Argument) GetPredefinedOk() (*[]string, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *Argument) SetPredefined(v []string)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *Argument) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetRequired

`func (o *Argument) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Argument) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Argument) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Argument) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSecret

`func (o *Argument) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Argument) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Argument) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Argument) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


