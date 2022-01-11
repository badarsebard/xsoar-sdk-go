# InvestigationPlaybookData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadyPlaybookInputs** | Pointer to **map[string]map[string]map[string]interface{}** |  | [optional] 
**AutoExtracting** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**PlaybookInputs**](PlaybookInputs.md) |  | [optional] 
**InvestigationId** | Pointer to **string** |  | [optional] 
**IsTIM** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**PlaybookOutputs**](PlaybookOutputs.md) |  | [optional] 
**PendingTasks** | Pointer to **map[string]map[string]interface{}** | Tasks that are marked for running, but cannot yet run due to incomplete parents | [optional] 
**PlaybookId** | Pointer to **string** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**StartTaskId** | Pointer to **string** | FirstTask is the root task of the playbook | [optional] 
**State** | Pointer to **string** | InvestigationPlaybookState indicates the state of the running playbook | [optional] 
**SubPlaybookInputs** | Pointer to [**map[string]PlaybookInputs**](PlaybookInputs.md) |  | [optional] 
**SubPlaybookOutputs** | Pointer to [**map[string]PlaybookOutputs**](PlaybookOutputs.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]InvestigationPlaybookTask**](InvestigationPlaybookTask.md) |  | [optional] 
**View** | Pointer to **map[string]interface{}** | PlaybookView represents the view in client of playbook graph | [optional] 

## Methods

### NewInvestigationPlaybookData

`func NewInvestigationPlaybookData() *InvestigationPlaybookData`

NewInvestigationPlaybookData instantiates a new InvestigationPlaybookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationPlaybookDataWithDefaults

`func NewInvestigationPlaybookDataWithDefaults() *InvestigationPlaybookData`

NewInvestigationPlaybookDataWithDefaults instantiates a new InvestigationPlaybookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadyPlaybookInputs

`func (o *InvestigationPlaybookData) GetReadyPlaybookInputs() map[string]map[string]map[string]interface{}`

GetReadyPlaybookInputs returns the ReadyPlaybookInputs field if non-nil, zero value otherwise.

### GetReadyPlaybookInputsOk

`func (o *InvestigationPlaybookData) GetReadyPlaybookInputsOk() (*map[string]map[string]map[string]interface{}, bool)`

GetReadyPlaybookInputsOk returns a tuple with the ReadyPlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyPlaybookInputs

`func (o *InvestigationPlaybookData) SetReadyPlaybookInputs(v map[string]map[string]map[string]interface{})`

SetReadyPlaybookInputs sets ReadyPlaybookInputs field to given value.

### HasReadyPlaybookInputs

`func (o *InvestigationPlaybookData) HasReadyPlaybookInputs() bool`

HasReadyPlaybookInputs returns a boolean if a field has been set.

### GetAutoExtracting

`func (o *InvestigationPlaybookData) GetAutoExtracting() bool`

GetAutoExtracting returns the AutoExtracting field if non-nil, zero value otherwise.

### GetAutoExtractingOk

`func (o *InvestigationPlaybookData) GetAutoExtractingOk() (*bool, bool)`

GetAutoExtractingOk returns a tuple with the AutoExtracting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExtracting

`func (o *InvestigationPlaybookData) SetAutoExtracting(v bool)`

SetAutoExtracting sets AutoExtracting field to given value.

### HasAutoExtracting

`func (o *InvestigationPlaybookData) HasAutoExtracting() bool`

HasAutoExtracting returns a boolean if a field has been set.

### GetComment

`func (o *InvestigationPlaybookData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InvestigationPlaybookData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InvestigationPlaybookData) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InvestigationPlaybookData) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInputs

`func (o *InvestigationPlaybookData) GetInputs() PlaybookInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *InvestigationPlaybookData) GetInputsOk() (*PlaybookInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *InvestigationPlaybookData) SetInputs(v PlaybookInputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *InvestigationPlaybookData) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetInvestigationId

`func (o *InvestigationPlaybookData) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *InvestigationPlaybookData) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *InvestigationPlaybookData) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *InvestigationPlaybookData) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsTIM

`func (o *InvestigationPlaybookData) GetIsTIM() bool`

GetIsTIM returns the IsTIM field if non-nil, zero value otherwise.

### GetIsTIMOk

`func (o *InvestigationPlaybookData) GetIsTIMOk() (*bool, bool)`

GetIsTIMOk returns a tuple with the IsTIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTIM

`func (o *InvestigationPlaybookData) SetIsTIM(v bool)`

SetIsTIM sets IsTIM field to given value.

### HasIsTIM

`func (o *InvestigationPlaybookData) HasIsTIM() bool`

HasIsTIM returns a boolean if a field has been set.

### GetName

`func (o *InvestigationPlaybookData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationPlaybookData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationPlaybookData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvestigationPlaybookData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *InvestigationPlaybookData) GetOutputs() PlaybookOutputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *InvestigationPlaybookData) GetOutputsOk() (*PlaybookOutputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *InvestigationPlaybookData) SetOutputs(v PlaybookOutputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *InvestigationPlaybookData) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPendingTasks

`func (o *InvestigationPlaybookData) GetPendingTasks() map[string]map[string]interface{}`

GetPendingTasks returns the PendingTasks field if non-nil, zero value otherwise.

### GetPendingTasksOk

`func (o *InvestigationPlaybookData) GetPendingTasksOk() (*map[string]map[string]interface{}, bool)`

GetPendingTasksOk returns a tuple with the PendingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTasks

`func (o *InvestigationPlaybookData) SetPendingTasks(v map[string]map[string]interface{})`

SetPendingTasks sets PendingTasks field to given value.

### HasPendingTasks

`func (o *InvestigationPlaybookData) HasPendingTasks() bool`

HasPendingTasks returns a boolean if a field has been set.

### GetPlaybookId

`func (o *InvestigationPlaybookData) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *InvestigationPlaybookData) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *InvestigationPlaybookData) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *InvestigationPlaybookData) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetQuiet

`func (o *InvestigationPlaybookData) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *InvestigationPlaybookData) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *InvestigationPlaybookData) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *InvestigationPlaybookData) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetStartDate

`func (o *InvestigationPlaybookData) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvestigationPlaybookData) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvestigationPlaybookData) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvestigationPlaybookData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartTaskId

`func (o *InvestigationPlaybookData) GetStartTaskId() string`

GetStartTaskId returns the StartTaskId field if non-nil, zero value otherwise.

### GetStartTaskIdOk

`func (o *InvestigationPlaybookData) GetStartTaskIdOk() (*string, bool)`

GetStartTaskIdOk returns a tuple with the StartTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTaskId

`func (o *InvestigationPlaybookData) SetStartTaskId(v string)`

SetStartTaskId sets StartTaskId field to given value.

### HasStartTaskId

`func (o *InvestigationPlaybookData) HasStartTaskId() bool`

HasStartTaskId returns a boolean if a field has been set.

### GetState

`func (o *InvestigationPlaybookData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvestigationPlaybookData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvestigationPlaybookData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InvestigationPlaybookData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubPlaybookInputs

`func (o *InvestigationPlaybookData) GetSubPlaybookInputs() map[string]PlaybookInputs`

GetSubPlaybookInputs returns the SubPlaybookInputs field if non-nil, zero value otherwise.

### GetSubPlaybookInputsOk

`func (o *InvestigationPlaybookData) GetSubPlaybookInputsOk() (*map[string]PlaybookInputs, bool)`

GetSubPlaybookInputsOk returns a tuple with the SubPlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPlaybookInputs

`func (o *InvestigationPlaybookData) SetSubPlaybookInputs(v map[string]PlaybookInputs)`

SetSubPlaybookInputs sets SubPlaybookInputs field to given value.

### HasSubPlaybookInputs

`func (o *InvestigationPlaybookData) HasSubPlaybookInputs() bool`

HasSubPlaybookInputs returns a boolean if a field has been set.

### GetSubPlaybookOutputs

`func (o *InvestigationPlaybookData) GetSubPlaybookOutputs() map[string]PlaybookOutputs`

GetSubPlaybookOutputs returns the SubPlaybookOutputs field if non-nil, zero value otherwise.

### GetSubPlaybookOutputsOk

`func (o *InvestigationPlaybookData) GetSubPlaybookOutputsOk() (*map[string]PlaybookOutputs, bool)`

GetSubPlaybookOutputsOk returns a tuple with the SubPlaybookOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPlaybookOutputs

`func (o *InvestigationPlaybookData) SetSubPlaybookOutputs(v map[string]PlaybookOutputs)`

SetSubPlaybookOutputs sets SubPlaybookOutputs field to given value.

### HasSubPlaybookOutputs

`func (o *InvestigationPlaybookData) HasSubPlaybookOutputs() bool`

HasSubPlaybookOutputs returns a boolean if a field has been set.

### GetTasks

`func (o *InvestigationPlaybookData) GetTasks() map[string]InvestigationPlaybookTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InvestigationPlaybookData) GetTasksOk() (*map[string]InvestigationPlaybookTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InvestigationPlaybookData) SetTasks(v map[string]InvestigationPlaybookTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InvestigationPlaybookData) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetView

`func (o *InvestigationPlaybookData) GetView() map[string]interface{}`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *InvestigationPlaybookData) GetViewOk() (*map[string]interface{}, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *InvestigationPlaybookData) SetView(v map[string]interface{})`

SetView sets View field to given value.

### HasView

`func (o *InvestigationPlaybookData) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


