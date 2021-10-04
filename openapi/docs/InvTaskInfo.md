# InvTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]TaskCondition**](TaskCondition.md) |  | [optional] 
**InTaskID** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**InvId** | Pointer to **string** |  | [optional] 
**LoopArgs** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**LoopCondition** | Pointer to [**[][]ArgAtomicFilter**]([]ArgAtomicFilter.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewInvTaskInfo

`func NewInvTaskInfo() *InvTaskInfo`

NewInvTaskInfo instantiates a new InvTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvTaskInfoWithDefaults

`func NewInvTaskInfoWithDefaults() *InvTaskInfo`

NewInvTaskInfoWithDefaults instantiates a new InvTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *InvTaskInfo) GetArgs() map[string]AdvanceArg`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *InvTaskInfo) GetArgsOk() (*map[string]AdvanceArg, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *InvTaskInfo) SetArgs(v map[string]AdvanceArg)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *InvTaskInfo) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetComment

`func (o *InvTaskInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InvTaskInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InvTaskInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InvTaskInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConditions

`func (o *InvTaskInfo) GetConditions() []TaskCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InvTaskInfo) GetConditionsOk() (*[]TaskCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InvTaskInfo) SetConditions(v []TaskCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InvTaskInfo) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetInTaskID

`func (o *InvTaskInfo) GetInTaskID() string`

GetInTaskID returns the InTaskID field if non-nil, zero value otherwise.

### GetInTaskIDOk

`func (o *InvTaskInfo) GetInTaskIDOk() (*string, bool)`

GetInTaskIDOk returns a tuple with the InTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTaskID

`func (o *InvTaskInfo) SetInTaskID(v string)`

SetInTaskID sets InTaskID field to given value.

### HasInTaskID

`func (o *InvTaskInfo) HasInTaskID() bool`

HasInTaskID returns a boolean if a field has been set.

### GetInput

`func (o *InvTaskInfo) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InvTaskInfo) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InvTaskInfo) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *InvTaskInfo) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetInvId

`func (o *InvTaskInfo) GetInvId() string`

GetInvId returns the InvId field if non-nil, zero value otherwise.

### GetInvIdOk

`func (o *InvTaskInfo) GetInvIdOk() (*string, bool)`

GetInvIdOk returns a tuple with the InvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvId

`func (o *InvTaskInfo) SetInvId(v string)`

SetInvId sets InvId field to given value.

### HasInvId

`func (o *InvTaskInfo) HasInvId() bool`

HasInvId returns a boolean if a field has been set.

### GetLoopArgs

`func (o *InvTaskInfo) GetLoopArgs() map[string]AdvanceArg`

GetLoopArgs returns the LoopArgs field if non-nil, zero value otherwise.

### GetLoopArgsOk

`func (o *InvTaskInfo) GetLoopArgsOk() (*map[string]AdvanceArg, bool)`

GetLoopArgsOk returns a tuple with the LoopArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopArgs

`func (o *InvTaskInfo) SetLoopArgs(v map[string]AdvanceArg)`

SetLoopArgs sets LoopArgs field to given value.

### HasLoopArgs

`func (o *InvTaskInfo) HasLoopArgs() bool`

HasLoopArgs returns a boolean if a field has been set.

### GetLoopCondition

`func (o *InvTaskInfo) GetLoopCondition() [][]ArgAtomicFilter`

GetLoopCondition returns the LoopCondition field if non-nil, zero value otherwise.

### GetLoopConditionOk

`func (o *InvTaskInfo) GetLoopConditionOk() (*[][]ArgAtomicFilter, bool)`

GetLoopConditionOk returns a tuple with the LoopCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopCondition

`func (o *InvTaskInfo) SetLoopCondition(v [][]ArgAtomicFilter)`

SetLoopCondition sets LoopCondition field to given value.

### HasLoopCondition

`func (o *InvTaskInfo) HasLoopCondition() bool`

HasLoopCondition returns a boolean if a field has been set.

### GetVersion

`func (o *InvTaskInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InvTaskInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InvTaskInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InvTaskInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


