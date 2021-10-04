# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]Argument**](Argument.md) |  | [optional] 
**Cartesian** | Pointer to **bool** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocsHidden** | Pointer to **bool** |  | [optional] 
**Execution** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Important** | Pointer to [**[]Important**](Important.md) |  | [optional] 
**IndicatorAction** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**Permitted** | Pointer to **bool** |  | [optional] 
**Polling** | Pointer to **bool** |  | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *Command) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Command) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Command) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Command) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCartesian

`func (o *Command) GetCartesian() bool`

GetCartesian returns the Cartesian field if non-nil, zero value otherwise.

### GetCartesianOk

`func (o *Command) GetCartesianOk() (*bool, bool)`

GetCartesianOk returns a tuple with the Cartesian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartesian

`func (o *Command) SetCartesian(v bool)`

SetCartesian sets Cartesian field to given value.

### HasCartesian

`func (o *Command) HasCartesian() bool`

HasCartesian returns a boolean if a field has been set.

### GetDeprecated

`func (o *Command) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Command) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Command) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Command) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Command) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Command) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Command) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Command) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocsHidden

`func (o *Command) GetDocsHidden() bool`

GetDocsHidden returns the DocsHidden field if non-nil, zero value otherwise.

### GetDocsHiddenOk

`func (o *Command) GetDocsHiddenOk() (*bool, bool)`

GetDocsHiddenOk returns a tuple with the DocsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsHidden

`func (o *Command) SetDocsHidden(v bool)`

SetDocsHidden sets DocsHidden field to given value.

### HasDocsHidden

`func (o *Command) HasDocsHidden() bool`

HasDocsHidden returns a boolean if a field has been set.

### GetExecution

`func (o *Command) GetExecution() bool`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *Command) GetExecutionOk() (*bool, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *Command) SetExecution(v bool)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *Command) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetHidden

`func (o *Command) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Command) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Command) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Command) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetImportant

`func (o *Command) GetImportant() []Important`

GetImportant returns the Important field if non-nil, zero value otherwise.

### GetImportantOk

`func (o *Command) GetImportantOk() (*[]Important, bool)`

GetImportantOk returns a tuple with the Important field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportant

`func (o *Command) SetImportant(v []Important)`

SetImportant sets Important field to given value.

### HasImportant

`func (o *Command) HasImportant() bool`

HasImportant returns a boolean if a field has been set.

### GetIndicatorAction

`func (o *Command) GetIndicatorAction() bool`

GetIndicatorAction returns the IndicatorAction field if non-nil, zero value otherwise.

### GetIndicatorActionOk

`func (o *Command) GetIndicatorActionOk() (*bool, bool)`

GetIndicatorActionOk returns a tuple with the IndicatorAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorAction

`func (o *Command) SetIndicatorAction(v bool)`

SetIndicatorAction sets IndicatorAction field to given value.

### HasIndicatorAction

`func (o *Command) HasIndicatorAction() bool`

HasIndicatorAction returns a boolean if a field has been set.

### GetName

`func (o *Command) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Command) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Command) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Command) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *Command) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *Command) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *Command) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *Command) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPermitted

`func (o *Command) GetPermitted() bool`

GetPermitted returns the Permitted field if non-nil, zero value otherwise.

### GetPermittedOk

`func (o *Command) GetPermittedOk() (*bool, bool)`

GetPermittedOk returns a tuple with the Permitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitted

`func (o *Command) SetPermitted(v bool)`

SetPermitted sets Permitted field to given value.

### HasPermitted

`func (o *Command) HasPermitted() bool`

HasPermitted returns a boolean if a field has been set.

### GetPolling

`func (o *Command) GetPolling() bool`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *Command) GetPollingOk() (*bool, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *Command) SetPolling(v bool)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *Command) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetSensitive

`func (o *Command) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *Command) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *Command) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *Command) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetTimeout

`func (o *Command) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Command) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Command) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Command) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


