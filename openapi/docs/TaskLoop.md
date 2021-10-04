# TaskLoop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**BuiltinCondition** | Pointer to [**[][]ArgAtomicFilter**]([]ArgAtomicFilter.md) |  | [optional] 
**ExitCondition** | Pointer to **string** |  | [optional] 
**ForEach** | Pointer to **bool** |  | [optional] 
**IsCommand** | Pointer to **bool** |  | [optional] 
**Max** | Pointer to **int64** |  | [optional] 
**ScriptArguments** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**ScriptId** | Pointer to **string** |  | [optional] 
**ScriptName** | Pointer to **string** |  | [optional] 
**Wait** | Pointer to **int64** |  | [optional] 

## Methods

### NewTaskLoop

`func NewTaskLoop() *TaskLoop`

NewTaskLoop instantiates a new TaskLoop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLoopWithDefaults

`func NewTaskLoopWithDefaults() *TaskLoop`

NewTaskLoopWithDefaults instantiates a new TaskLoop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *TaskLoop) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *TaskLoop) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *TaskLoop) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *TaskLoop) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBuiltinCondition

`func (o *TaskLoop) GetBuiltinCondition() [][]ArgAtomicFilter`

GetBuiltinCondition returns the BuiltinCondition field if non-nil, zero value otherwise.

### GetBuiltinConditionOk

`func (o *TaskLoop) GetBuiltinConditionOk() (*[][]ArgAtomicFilter, bool)`

GetBuiltinConditionOk returns a tuple with the BuiltinCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinCondition

`func (o *TaskLoop) SetBuiltinCondition(v [][]ArgAtomicFilter)`

SetBuiltinCondition sets BuiltinCondition field to given value.

### HasBuiltinCondition

`func (o *TaskLoop) HasBuiltinCondition() bool`

HasBuiltinCondition returns a boolean if a field has been set.

### GetExitCondition

`func (o *TaskLoop) GetExitCondition() string`

GetExitCondition returns the ExitCondition field if non-nil, zero value otherwise.

### GetExitConditionOk

`func (o *TaskLoop) GetExitConditionOk() (*string, bool)`

GetExitConditionOk returns a tuple with the ExitCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCondition

`func (o *TaskLoop) SetExitCondition(v string)`

SetExitCondition sets ExitCondition field to given value.

### HasExitCondition

`func (o *TaskLoop) HasExitCondition() bool`

HasExitCondition returns a boolean if a field has been set.

### GetForEach

`func (o *TaskLoop) GetForEach() bool`

GetForEach returns the ForEach field if non-nil, zero value otherwise.

### GetForEachOk

`func (o *TaskLoop) GetForEachOk() (*bool, bool)`

GetForEachOk returns a tuple with the ForEach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForEach

`func (o *TaskLoop) SetForEach(v bool)`

SetForEach sets ForEach field to given value.

### HasForEach

`func (o *TaskLoop) HasForEach() bool`

HasForEach returns a boolean if a field has been set.

### GetIsCommand

`func (o *TaskLoop) GetIsCommand() bool`

GetIsCommand returns the IsCommand field if non-nil, zero value otherwise.

### GetIsCommandOk

`func (o *TaskLoop) GetIsCommandOk() (*bool, bool)`

GetIsCommandOk returns a tuple with the IsCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommand

`func (o *TaskLoop) SetIsCommand(v bool)`

SetIsCommand sets IsCommand field to given value.

### HasIsCommand

`func (o *TaskLoop) HasIsCommand() bool`

HasIsCommand returns a boolean if a field has been set.

### GetMax

`func (o *TaskLoop) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *TaskLoop) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *TaskLoop) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *TaskLoop) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetScriptArguments

`func (o *TaskLoop) GetScriptArguments() map[string]AdvanceArg`

GetScriptArguments returns the ScriptArguments field if non-nil, zero value otherwise.

### GetScriptArgumentsOk

`func (o *TaskLoop) GetScriptArgumentsOk() (*map[string]AdvanceArg, bool)`

GetScriptArgumentsOk returns a tuple with the ScriptArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArguments

`func (o *TaskLoop) SetScriptArguments(v map[string]AdvanceArg)`

SetScriptArguments sets ScriptArguments field to given value.

### HasScriptArguments

`func (o *TaskLoop) HasScriptArguments() bool`

HasScriptArguments returns a boolean if a field has been set.

### GetScriptId

`func (o *TaskLoop) GetScriptId() string`

GetScriptId returns the ScriptId field if non-nil, zero value otherwise.

### GetScriptIdOk

`func (o *TaskLoop) GetScriptIdOk() (*string, bool)`

GetScriptIdOk returns a tuple with the ScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptId

`func (o *TaskLoop) SetScriptId(v string)`

SetScriptId sets ScriptId field to given value.

### HasScriptId

`func (o *TaskLoop) HasScriptId() bool`

HasScriptId returns a boolean if a field has been set.

### GetScriptName

`func (o *TaskLoop) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *TaskLoop) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *TaskLoop) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *TaskLoop) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetWait

`func (o *TaskLoop) GetWait() int64`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *TaskLoop) GetWaitOk() (*int64, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *TaskLoop) SetWait(v int64)`

SetWait sets Wait field to given value.

### HasWait

`func (o *TaskLoop) HasWait() bool`

HasWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


