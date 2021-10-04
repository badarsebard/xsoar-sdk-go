# PlaybookTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]TaskCondition**](TaskCondition.md) | Conditions - optional list of conditions to run when task is conditional. we check conditions by their order (e.i. - considering the first one that satisfied) | [optional] 
**ContinueOnError** | Pointer to **bool** |  | [optional] 
**DefaultAssignee** | Pointer to **string** |  | [optional] 
**DefaultAssigneeComplex** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**DefaultReminder** | Pointer to **int64** |  | [optional] 
**EvidenceData** | Pointer to [**EvidenceData**](EvidenceData.md) |  | [optional] 
**FieldMapping** | Pointer to [**[]FieldMapping**](FieldMapping.md) |  | [optional] 
**Form** | Pointer to [**DataCollectionForm**](DataCollectionForm.md) |  | [optional] 
**FormDisplay** | Pointer to [**FormDisplay**](FormDisplay.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IgnoreWorker** | Pointer to **bool** | Do not run this task in a worker | [optional] 
**IsAutoSwitchedToQuietMode** | Pointer to **bool** |  | [optional] 
**IsOverSize** | Pointer to **bool** |  | [optional] 
**Loop** | Pointer to [**TaskLoop**](TaskLoop.md) |  | [optional] 
**Message** | Pointer to [**NotifiableItem**](NotifiableItem.md) |  | [optional] 
**NextTasks** | Pointer to **map[string][]string** |  | [optional] 
**Note** | Pointer to **bool** |  | [optional] 
**QuietMode** | Pointer to **int64** | QuietMode quiet mode for playbook task | [optional] 
**ReputationCalc** | Pointer to **float64** |  | [optional] 
**RestrictedCompletion** | Pointer to **bool** |  | [optional] 
**ScriptArguments** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**SeparateContext** | Pointer to **bool** |  | [optional] 
**SkipUnavailable** | Pointer to **bool** | SkipUnavailable if true then will check if automation exists, integration of that command is installed and active or sub playbook exists in Demisto | [optional] 
**Sla** | Pointer to [**SLA**](SLA.md) |  | [optional] 
**SlaReminder** | Pointer to [**SLA**](SLA.md) |  | [optional] 
**Task** | Pointer to [**Task**](Task.md) |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TimerTriggers** | Pointer to [**[]TimerTrigger**](TimerTrigger.md) | SLA fields | [optional] 
**Type** | Pointer to **string** | TaskType is the Task in the playbook context as a node | [optional] 
**View** | Pointer to **map[string]interface{}** | TaskView represents the view in client of the tasks graph | [optional] 

## Methods

### NewPlaybookTask

`func NewPlaybookTask() *PlaybookTask`

NewPlaybookTask instantiates a new PlaybookTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookTaskWithDefaults

`func NewPlaybookTaskWithDefaults() *PlaybookTask`

NewPlaybookTaskWithDefaults instantiates a new PlaybookTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *PlaybookTask) GetConditions() []TaskCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PlaybookTask) GetConditionsOk() (*[]TaskCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PlaybookTask) SetConditions(v []TaskCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PlaybookTask) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContinueOnError

`func (o *PlaybookTask) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *PlaybookTask) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *PlaybookTask) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *PlaybookTask) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetDefaultAssignee

`func (o *PlaybookTask) GetDefaultAssignee() string`

GetDefaultAssignee returns the DefaultAssignee field if non-nil, zero value otherwise.

### GetDefaultAssigneeOk

`func (o *PlaybookTask) GetDefaultAssigneeOk() (*string, bool)`

GetDefaultAssigneeOk returns a tuple with the DefaultAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAssignee

`func (o *PlaybookTask) SetDefaultAssignee(v string)`

SetDefaultAssignee sets DefaultAssignee field to given value.

### HasDefaultAssignee

`func (o *PlaybookTask) HasDefaultAssignee() bool`

HasDefaultAssignee returns a boolean if a field has been set.

### GetDefaultAssigneeComplex

`func (o *PlaybookTask) GetDefaultAssigneeComplex() AdvanceArg`

GetDefaultAssigneeComplex returns the DefaultAssigneeComplex field if non-nil, zero value otherwise.

### GetDefaultAssigneeComplexOk

`func (o *PlaybookTask) GetDefaultAssigneeComplexOk() (*AdvanceArg, bool)`

GetDefaultAssigneeComplexOk returns a tuple with the DefaultAssigneeComplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAssigneeComplex

`func (o *PlaybookTask) SetDefaultAssigneeComplex(v AdvanceArg)`

SetDefaultAssigneeComplex sets DefaultAssigneeComplex field to given value.

### HasDefaultAssigneeComplex

`func (o *PlaybookTask) HasDefaultAssigneeComplex() bool`

HasDefaultAssigneeComplex returns a boolean if a field has been set.

### GetDefaultReminder

`func (o *PlaybookTask) GetDefaultReminder() int64`

GetDefaultReminder returns the DefaultReminder field if non-nil, zero value otherwise.

### GetDefaultReminderOk

`func (o *PlaybookTask) GetDefaultReminderOk() (*int64, bool)`

GetDefaultReminderOk returns a tuple with the DefaultReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReminder

`func (o *PlaybookTask) SetDefaultReminder(v int64)`

SetDefaultReminder sets DefaultReminder field to given value.

### HasDefaultReminder

`func (o *PlaybookTask) HasDefaultReminder() bool`

HasDefaultReminder returns a boolean if a field has been set.

### GetEvidenceData

`func (o *PlaybookTask) GetEvidenceData() EvidenceData`

GetEvidenceData returns the EvidenceData field if non-nil, zero value otherwise.

### GetEvidenceDataOk

`func (o *PlaybookTask) GetEvidenceDataOk() (*EvidenceData, bool)`

GetEvidenceDataOk returns a tuple with the EvidenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceData

`func (o *PlaybookTask) SetEvidenceData(v EvidenceData)`

SetEvidenceData sets EvidenceData field to given value.

### HasEvidenceData

`func (o *PlaybookTask) HasEvidenceData() bool`

HasEvidenceData returns a boolean if a field has been set.

### GetFieldMapping

`func (o *PlaybookTask) GetFieldMapping() []FieldMapping`

GetFieldMapping returns the FieldMapping field if non-nil, zero value otherwise.

### GetFieldMappingOk

`func (o *PlaybookTask) GetFieldMappingOk() (*[]FieldMapping, bool)`

GetFieldMappingOk returns a tuple with the FieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMapping

`func (o *PlaybookTask) SetFieldMapping(v []FieldMapping)`

SetFieldMapping sets FieldMapping field to given value.

### HasFieldMapping

`func (o *PlaybookTask) HasFieldMapping() bool`

HasFieldMapping returns a boolean if a field has been set.

### GetForm

`func (o *PlaybookTask) GetForm() DataCollectionForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PlaybookTask) GetFormOk() (*DataCollectionForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PlaybookTask) SetForm(v DataCollectionForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *PlaybookTask) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetFormDisplay

`func (o *PlaybookTask) GetFormDisplay() FormDisplay`

GetFormDisplay returns the FormDisplay field if non-nil, zero value otherwise.

### GetFormDisplayOk

`func (o *PlaybookTask) GetFormDisplayOk() (*FormDisplay, bool)`

GetFormDisplayOk returns a tuple with the FormDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormDisplay

`func (o *PlaybookTask) SetFormDisplay(v FormDisplay)`

SetFormDisplay sets FormDisplay field to given value.

### HasFormDisplay

`func (o *PlaybookTask) HasFormDisplay() bool`

HasFormDisplay returns a boolean if a field has been set.

### GetId

`func (o *PlaybookTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaybookTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaybookTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaybookTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreWorker

`func (o *PlaybookTask) GetIgnoreWorker() bool`

GetIgnoreWorker returns the IgnoreWorker field if non-nil, zero value otherwise.

### GetIgnoreWorkerOk

`func (o *PlaybookTask) GetIgnoreWorkerOk() (*bool, bool)`

GetIgnoreWorkerOk returns a tuple with the IgnoreWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorker

`func (o *PlaybookTask) SetIgnoreWorker(v bool)`

SetIgnoreWorker sets IgnoreWorker field to given value.

### HasIgnoreWorker

`func (o *PlaybookTask) HasIgnoreWorker() bool`

HasIgnoreWorker returns a boolean if a field has been set.

### GetIsAutoSwitchedToQuietMode

`func (o *PlaybookTask) GetIsAutoSwitchedToQuietMode() bool`

GetIsAutoSwitchedToQuietMode returns the IsAutoSwitchedToQuietMode field if non-nil, zero value otherwise.

### GetIsAutoSwitchedToQuietModeOk

`func (o *PlaybookTask) GetIsAutoSwitchedToQuietModeOk() (*bool, bool)`

GetIsAutoSwitchedToQuietModeOk returns a tuple with the IsAutoSwitchedToQuietMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSwitchedToQuietMode

`func (o *PlaybookTask) SetIsAutoSwitchedToQuietMode(v bool)`

SetIsAutoSwitchedToQuietMode sets IsAutoSwitchedToQuietMode field to given value.

### HasIsAutoSwitchedToQuietMode

`func (o *PlaybookTask) HasIsAutoSwitchedToQuietMode() bool`

HasIsAutoSwitchedToQuietMode returns a boolean if a field has been set.

### GetIsOverSize

`func (o *PlaybookTask) GetIsOverSize() bool`

GetIsOverSize returns the IsOverSize field if non-nil, zero value otherwise.

### GetIsOverSizeOk

`func (o *PlaybookTask) GetIsOverSizeOk() (*bool, bool)`

GetIsOverSizeOk returns a tuple with the IsOverSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverSize

`func (o *PlaybookTask) SetIsOverSize(v bool)`

SetIsOverSize sets IsOverSize field to given value.

### HasIsOverSize

`func (o *PlaybookTask) HasIsOverSize() bool`

HasIsOverSize returns a boolean if a field has been set.

### GetLoop

`func (o *PlaybookTask) GetLoop() TaskLoop`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *PlaybookTask) GetLoopOk() (*TaskLoop, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *PlaybookTask) SetLoop(v TaskLoop)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *PlaybookTask) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetMessage

`func (o *PlaybookTask) GetMessage() NotifiableItem`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlaybookTask) GetMessageOk() (*NotifiableItem, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlaybookTask) SetMessage(v NotifiableItem)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlaybookTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNextTasks

`func (o *PlaybookTask) GetNextTasks() map[string][]string`

GetNextTasks returns the NextTasks field if non-nil, zero value otherwise.

### GetNextTasksOk

`func (o *PlaybookTask) GetNextTasksOk() (*map[string][]string, bool)`

GetNextTasksOk returns a tuple with the NextTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTasks

`func (o *PlaybookTask) SetNextTasks(v map[string][]string)`

SetNextTasks sets NextTasks field to given value.

### HasNextTasks

`func (o *PlaybookTask) HasNextTasks() bool`

HasNextTasks returns a boolean if a field has been set.

### GetNote

`func (o *PlaybookTask) GetNote() bool`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PlaybookTask) GetNoteOk() (*bool, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PlaybookTask) SetNote(v bool)`

SetNote sets Note field to given value.

### HasNote

`func (o *PlaybookTask) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetQuietMode

`func (o *PlaybookTask) GetQuietMode() int64`

GetQuietMode returns the QuietMode field if non-nil, zero value otherwise.

### GetQuietModeOk

`func (o *PlaybookTask) GetQuietModeOk() (*int64, bool)`

GetQuietModeOk returns a tuple with the QuietMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietMode

`func (o *PlaybookTask) SetQuietMode(v int64)`

SetQuietMode sets QuietMode field to given value.

### HasQuietMode

`func (o *PlaybookTask) HasQuietMode() bool`

HasQuietMode returns a boolean if a field has been set.

### GetReputationCalc

`func (o *PlaybookTask) GetReputationCalc() float64`

GetReputationCalc returns the ReputationCalc field if non-nil, zero value otherwise.

### GetReputationCalcOk

`func (o *PlaybookTask) GetReputationCalcOk() (*float64, bool)`

GetReputationCalcOk returns a tuple with the ReputationCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationCalc

`func (o *PlaybookTask) SetReputationCalc(v float64)`

SetReputationCalc sets ReputationCalc field to given value.

### HasReputationCalc

`func (o *PlaybookTask) HasReputationCalc() bool`

HasReputationCalc returns a boolean if a field has been set.

### GetRestrictedCompletion

`func (o *PlaybookTask) GetRestrictedCompletion() bool`

GetRestrictedCompletion returns the RestrictedCompletion field if non-nil, zero value otherwise.

### GetRestrictedCompletionOk

`func (o *PlaybookTask) GetRestrictedCompletionOk() (*bool, bool)`

GetRestrictedCompletionOk returns a tuple with the RestrictedCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCompletion

`func (o *PlaybookTask) SetRestrictedCompletion(v bool)`

SetRestrictedCompletion sets RestrictedCompletion field to given value.

### HasRestrictedCompletion

`func (o *PlaybookTask) HasRestrictedCompletion() bool`

HasRestrictedCompletion returns a boolean if a field has been set.

### GetScriptArguments

`func (o *PlaybookTask) GetScriptArguments() map[string]AdvanceArg`

GetScriptArguments returns the ScriptArguments field if non-nil, zero value otherwise.

### GetScriptArgumentsOk

`func (o *PlaybookTask) GetScriptArgumentsOk() (*map[string]AdvanceArg, bool)`

GetScriptArgumentsOk returns a tuple with the ScriptArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArguments

`func (o *PlaybookTask) SetScriptArguments(v map[string]AdvanceArg)`

SetScriptArguments sets ScriptArguments field to given value.

### HasScriptArguments

`func (o *PlaybookTask) HasScriptArguments() bool`

HasScriptArguments returns a boolean if a field has been set.

### GetSeparateContext

`func (o *PlaybookTask) GetSeparateContext() bool`

GetSeparateContext returns the SeparateContext field if non-nil, zero value otherwise.

### GetSeparateContextOk

`func (o *PlaybookTask) GetSeparateContextOk() (*bool, bool)`

GetSeparateContextOk returns a tuple with the SeparateContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateContext

`func (o *PlaybookTask) SetSeparateContext(v bool)`

SetSeparateContext sets SeparateContext field to given value.

### HasSeparateContext

`func (o *PlaybookTask) HasSeparateContext() bool`

HasSeparateContext returns a boolean if a field has been set.

### GetSkipUnavailable

`func (o *PlaybookTask) GetSkipUnavailable() bool`

GetSkipUnavailable returns the SkipUnavailable field if non-nil, zero value otherwise.

### GetSkipUnavailableOk

`func (o *PlaybookTask) GetSkipUnavailableOk() (*bool, bool)`

GetSkipUnavailableOk returns a tuple with the SkipUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUnavailable

`func (o *PlaybookTask) SetSkipUnavailable(v bool)`

SetSkipUnavailable sets SkipUnavailable field to given value.

### HasSkipUnavailable

`func (o *PlaybookTask) HasSkipUnavailable() bool`

HasSkipUnavailable returns a boolean if a field has been set.

### GetSla

`func (o *PlaybookTask) GetSla() SLA`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *PlaybookTask) GetSlaOk() (*SLA, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *PlaybookTask) SetSla(v SLA)`

SetSla sets Sla field to given value.

### HasSla

`func (o *PlaybookTask) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaReminder

`func (o *PlaybookTask) GetSlaReminder() SLA`

GetSlaReminder returns the SlaReminder field if non-nil, zero value otherwise.

### GetSlaReminderOk

`func (o *PlaybookTask) GetSlaReminderOk() (*SLA, bool)`

GetSlaReminderOk returns a tuple with the SlaReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaReminder

`func (o *PlaybookTask) SetSlaReminder(v SLA)`

SetSlaReminder sets SlaReminder field to given value.

### HasSlaReminder

`func (o *PlaybookTask) HasSlaReminder() bool`

HasSlaReminder returns a boolean if a field has been set.

### GetTask

`func (o *PlaybookTask) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PlaybookTask) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PlaybookTask) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *PlaybookTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskId

`func (o *PlaybookTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *PlaybookTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *PlaybookTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *PlaybookTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTimerTriggers

`func (o *PlaybookTask) GetTimerTriggers() []TimerTrigger`

GetTimerTriggers returns the TimerTriggers field if non-nil, zero value otherwise.

### GetTimerTriggersOk

`func (o *PlaybookTask) GetTimerTriggersOk() (*[]TimerTrigger, bool)`

GetTimerTriggersOk returns a tuple with the TimerTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerTriggers

`func (o *PlaybookTask) SetTimerTriggers(v []TimerTrigger)`

SetTimerTriggers sets TimerTriggers field to given value.

### HasTimerTriggers

`func (o *PlaybookTask) HasTimerTriggers() bool`

HasTimerTriggers returns a boolean if a field has been set.

### GetType

`func (o *PlaybookTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaybookTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaybookTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaybookTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *PlaybookTask) GetView() map[string]interface{}`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PlaybookTask) GetViewOk() (*map[string]interface{}, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PlaybookTask) SetView(v map[string]interface{})`

SetView sets View field to given value.

### HasView

`func (o *PlaybookTask) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


