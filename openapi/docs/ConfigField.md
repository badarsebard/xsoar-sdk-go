# ConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**DisplayPassword** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HiddenPassword** | Pointer to **bool** |  | [optional] 
**HiddenUsername** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **float64** | ConfigDataType holds the type of a configuration field or command argument | [optional] 

## Methods

### NewConfigField

`func NewConfigField() *ConfigField`

NewConfigField instantiates a new ConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFieldWithDefaults

`func NewConfigFieldWithDefaults() *ConfigField`

NewConfigFieldWithDefaults instantiates a new ConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *ConfigField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ConfigField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ConfigField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDisplay

`func (o *ConfigField) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigField) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigField) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConfigField) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDisplayPassword

`func (o *ConfigField) GetDisplayPassword() string`

GetDisplayPassword returns the DisplayPassword field if non-nil, zero value otherwise.

### GetDisplayPasswordOk

`func (o *ConfigField) GetDisplayPasswordOk() (*string, bool)`

GetDisplayPasswordOk returns a tuple with the DisplayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPassword

`func (o *ConfigField) SetDisplayPassword(v string)`

SetDisplayPassword sets DisplayPassword field to given value.

### HasDisplayPassword

`func (o *ConfigField) HasDisplayPassword() bool`

HasDisplayPassword returns a boolean if a field has been set.

### GetHidden

`func (o *ConfigField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ConfigField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ConfigField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ConfigField) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHiddenPassword

`func (o *ConfigField) GetHiddenPassword() bool`

GetHiddenPassword returns the HiddenPassword field if non-nil, zero value otherwise.

### GetHiddenPasswordOk

`func (o *ConfigField) GetHiddenPasswordOk() (*bool, bool)`

GetHiddenPasswordOk returns a tuple with the HiddenPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenPassword

`func (o *ConfigField) SetHiddenPassword(v bool)`

SetHiddenPassword sets HiddenPassword field to given value.

### HasHiddenPassword

`func (o *ConfigField) HasHiddenPassword() bool`

HasHiddenPassword returns a boolean if a field has been set.

### GetHiddenUsername

`func (o *ConfigField) GetHiddenUsername() bool`

GetHiddenUsername returns the HiddenUsername field if non-nil, zero value otherwise.

### GetHiddenUsernameOk

`func (o *ConfigField) GetHiddenUsernameOk() (*bool, bool)`

GetHiddenUsernameOk returns a tuple with the HiddenUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenUsername

`func (o *ConfigField) SetHiddenUsername(v bool)`

SetHiddenUsername sets HiddenUsername field to given value.

### HasHiddenUsername

`func (o *ConfigField) HasHiddenUsername() bool`

HasHiddenUsername returns a boolean if a field has been set.

### GetInfo

`func (o *ConfigField) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigField) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigField) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigField) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetName

`func (o *ConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ConfigField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigField) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *ConfigField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConfigField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConfigField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ConfigField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *ConfigField) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigField) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigField) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


