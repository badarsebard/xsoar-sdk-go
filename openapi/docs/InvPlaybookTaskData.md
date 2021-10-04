# InvPlaybookTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAfter** | Pointer to **bool** |  | [optional] 
**AutomationScript** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Loop** | Pointer to [**TaskLoop**](TaskLoop.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeighborInvPBTaskId** | Pointer to **string** |  | [optional] 
**PlaybookId** | Pointer to **string** |  | [optional] 
**ScriptArguments** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**SeparateContext** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | TaskType is the Task in the playbook context as a node | [optional] 

## Methods

### NewInvPlaybookTaskData

`func NewInvPlaybookTaskData() *InvPlaybookTaskData`

NewInvPlaybookTaskData instantiates a new InvPlaybookTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvPlaybookTaskDataWithDefaults

`func NewInvPlaybookTaskDataWithDefaults() *InvPlaybookTaskData`

NewInvPlaybookTaskDataWithDefaults instantiates a new InvPlaybookTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAfter

`func (o *InvPlaybookTaskData) GetAddAfter() bool`

GetAddAfter returns the AddAfter field if non-nil, zero value otherwise.

### GetAddAfterOk

`func (o *InvPlaybookTaskData) GetAddAfterOk() (*bool, bool)`

GetAddAfterOk returns a tuple with the AddAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAfter

`func (o *InvPlaybookTaskData) SetAddAfter(v bool)`

SetAddAfter sets AddAfter field to given value.

### HasAddAfter

`func (o *InvPlaybookTaskData) HasAddAfter() bool`

HasAddAfter returns a boolean if a field has been set.

### GetAutomationScript

`func (o *InvPlaybookTaskData) GetAutomationScript() string`

GetAutomationScript returns the AutomationScript field if non-nil, zero value otherwise.

### GetAutomationScriptOk

`func (o *InvPlaybookTaskData) GetAutomationScriptOk() (*string, bool)`

GetAutomationScriptOk returns a tuple with the AutomationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationScript

`func (o *InvPlaybookTaskData) SetAutomationScript(v string)`

SetAutomationScript sets AutomationScript field to given value.

### HasAutomationScript

`func (o *InvPlaybookTaskData) HasAutomationScript() bool`

HasAutomationScript returns a boolean if a field has been set.

### GetDescription

`func (o *InvPlaybookTaskData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvPlaybookTaskData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvPlaybookTaskData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvPlaybookTaskData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoop

`func (o *InvPlaybookTaskData) GetLoop() TaskLoop`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *InvPlaybookTaskData) GetLoopOk() (*TaskLoop, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *InvPlaybookTaskData) SetLoop(v TaskLoop)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *InvPlaybookTaskData) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetName

`func (o *InvPlaybookTaskData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvPlaybookTaskData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvPlaybookTaskData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvPlaybookTaskData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeighborInvPBTaskId

`func (o *InvPlaybookTaskData) GetNeighborInvPBTaskId() string`

GetNeighborInvPBTaskId returns the NeighborInvPBTaskId field if non-nil, zero value otherwise.

### GetNeighborInvPBTaskIdOk

`func (o *InvPlaybookTaskData) GetNeighborInvPBTaskIdOk() (*string, bool)`

GetNeighborInvPBTaskIdOk returns a tuple with the NeighborInvPBTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInvPBTaskId

`func (o *InvPlaybookTaskData) SetNeighborInvPBTaskId(v string)`

SetNeighborInvPBTaskId sets NeighborInvPBTaskId field to given value.

### HasNeighborInvPBTaskId

`func (o *InvPlaybookTaskData) HasNeighborInvPBTaskId() bool`

HasNeighborInvPBTaskId returns a boolean if a field has been set.

### GetPlaybookId

`func (o *InvPlaybookTaskData) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *InvPlaybookTaskData) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *InvPlaybookTaskData) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *InvPlaybookTaskData) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetScriptArguments

`func (o *InvPlaybookTaskData) GetScriptArguments() map[string]AdvanceArg`

GetScriptArguments returns the ScriptArguments field if non-nil, zero value otherwise.

### GetScriptArgumentsOk

`func (o *InvPlaybookTaskData) GetScriptArgumentsOk() (*map[string]AdvanceArg, bool)`

GetScriptArgumentsOk returns a tuple with the ScriptArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArguments

`func (o *InvPlaybookTaskData) SetScriptArguments(v map[string]AdvanceArg)`

SetScriptArguments sets ScriptArguments field to given value.

### HasScriptArguments

`func (o *InvPlaybookTaskData) HasScriptArguments() bool`

HasScriptArguments returns a boolean if a field has been set.

### GetSeparateContext

`func (o *InvPlaybookTaskData) GetSeparateContext() bool`

GetSeparateContext returns the SeparateContext field if non-nil, zero value otherwise.

### GetSeparateContextOk

`func (o *InvPlaybookTaskData) GetSeparateContextOk() (*bool, bool)`

GetSeparateContextOk returns a tuple with the SeparateContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateContext

`func (o *InvPlaybookTaskData) SetSeparateContext(v bool)`

SetSeparateContext sets SeparateContext field to given value.

### HasSeparateContext

`func (o *InvPlaybookTaskData) HasSeparateContext() bool`

HasSeparateContext returns a boolean if a field has been set.

### GetTags

`func (o *InvPlaybookTaskData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InvPlaybookTaskData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InvPlaybookTaskData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InvPlaybookTaskData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *InvPlaybookTaskData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvPlaybookTaskData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvPlaybookTaskData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvPlaybookTaskData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


