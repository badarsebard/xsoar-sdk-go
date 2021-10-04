# GridColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**FieldCalcScript** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**SelectValues** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int64** |  | [optional] 

## Methods

### NewGridColumn

`func NewGridColumn() *GridColumn`

NewGridColumn instantiates a new GridColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridColumnWithDefaults

`func NewGridColumnWithDefaults() *GridColumn`

NewGridColumnWithDefaults instantiates a new GridColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GridColumn) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GridColumn) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GridColumn) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GridColumn) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFieldCalcScript

`func (o *GridColumn) GetFieldCalcScript() string`

GetFieldCalcScript returns the FieldCalcScript field if non-nil, zero value otherwise.

### GetFieldCalcScriptOk

`func (o *GridColumn) GetFieldCalcScriptOk() (*string, bool)`

GetFieldCalcScriptOk returns a tuple with the FieldCalcScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCalcScript

`func (o *GridColumn) SetFieldCalcScript(v string)`

SetFieldCalcScript sets FieldCalcScript field to given value.

### HasFieldCalcScript

`func (o *GridColumn) HasFieldCalcScript() bool`

HasFieldCalcScript returns a boolean if a field has been set.

### GetIsDefault

`func (o *GridColumn) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GridColumn) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GridColumn) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GridColumn) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *GridColumn) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *GridColumn) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *GridColumn) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *GridColumn) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetKey

`func (o *GridColumn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GridColumn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GridColumn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GridColumn) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRequired

`func (o *GridColumn) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GridColumn) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GridColumn) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GridColumn) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetScript

`func (o *GridColumn) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *GridColumn) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *GridColumn) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *GridColumn) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSelectValues

`func (o *GridColumn) GetSelectValues() []string`

GetSelectValues returns the SelectValues field if non-nil, zero value otherwise.

### GetSelectValuesOk

`func (o *GridColumn) GetSelectValuesOk() (*[]string, bool)`

GetSelectValuesOk returns a tuple with the SelectValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectValues

`func (o *GridColumn) SetSelectValues(v []string)`

SetSelectValues sets SelectValues field to given value.

### HasSelectValues

`func (o *GridColumn) HasSelectValues() bool`

HasSelectValues returns a boolean if a field has been set.

### GetType

`func (o *GridColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridColumn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GridColumn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *GridColumn) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GridColumn) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GridColumn) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *GridColumn) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


