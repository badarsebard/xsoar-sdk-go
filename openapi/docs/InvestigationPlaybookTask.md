# InvestigationPlaybookTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**AssigneeSet** | Pointer to **bool** |  | [optional] 
**BlockingTasks** | Pointer to **[]string** |  | [optional] 
**CalculatedDescription** | Pointer to **string** |  | [optional] 
**CalculatedTaskName** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **bool** | Whether this task had any comments or not | [optional] 
**CompletedBy** | Pointer to **string** |  | [optional] 
**CompletedCount** | Pointer to **int64** |  | [optional] 
**CompletedDate** | Pointer to **time.Time** |  | [optional] 
**Conditions** | Pointer to [**[]TaskCondition**](TaskCondition.md) | Conditions - optional list of conditions to run when task is conditional. we check conditions by their order (e.i. - considering the first one that satisfied) | [optional] 
**ContinueOnError** | Pointer to **bool** |  | [optional] 
**DefaultAssignee** | Pointer to **string** |  | [optional] 
**DefaultAssigneeComplex** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**DefaultReminder** | Pointer to **int64** |  | [optional] 
**DoNotSaveTaskHistory** | Pointer to **bool** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**DueDateDuration** | Pointer to **int64** |  | [optional] 
**DueDateSet** | Pointer to **bool** |  | [optional] 
**Entries** | Pointer to **[]string** |  | [optional] 
**EvidenceData** | Pointer to [**EvidenceData**](EvidenceData.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 
**FieldMapping** | Pointer to [**[]FieldMapping**](FieldMapping.md) |  | [optional] 
**ForEachIndex** | Pointer to **int64** | Parameters needed for loops | [optional] 
**ForEachInputs** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 
**Form** | Pointer to [**DataCollectionForm**](DataCollectionForm.md) |  | [optional] 
**FormDisplay** | Pointer to [**FormDisplay**](FormDisplay.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IgnoreWorker** | Pointer to **bool** | Do not run this task in a worker | [optional] 
**Indent** | Pointer to **int64** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**IsAutoSwitchedToQuietMode** | Pointer to **bool** |  | [optional] 
**IsOverSize** | Pointer to **bool** |  | [optional] 
**Loop** | Pointer to [**TaskLoop**](TaskLoop.md) |  | [optional] 
**Message** | Pointer to [**NotifiableItem**](NotifiableItem.md) |  | [optional] 
**MissingContentItem** | Pointer to **string** | content item which caused task to skip | [optional] 
**NextPollingTime** | Pointer to **time.Time** | NextPollTime Time of the next scheduled command execution | [optional] 
**NextTasks** | Pointer to **map[string][]string** |  | [optional] 
**Note** | Pointer to **bool** |  | [optional] 
**Outputs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ParentBlockCount** | Pointer to **int64** | the number of tasks that are waiting on blocked in subplaybooks of this task | [optional] 
**ParentPlaybookID** | Pointer to **string** |  | [optional] 
**Patched** | Pointer to **bool** | Indicates whether this task was patched to InvPB and did not originally belong to the playbook | [optional] 
**PlaybookInputs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PollingEntries** | Pointer to **[]string** | PollingEntryIDs IDs of entries that are still polling | [optional] 
**PollingTimeoutTime** | Pointer to **time.Time** |  | [optional] 
**PreviousTasks** | Pointer to **map[string][]string** |  | [optional] 
**QuietMode** | Pointer to **int64** | QuietMode quiet mode for playbook task | [optional] 
**Reminder** | Pointer to **int64** | Duration in minutes, this field is not persisted here | [optional] 
**ReputationCalc** | Pointer to **float64** |  | [optional] 
**RestrictedCompletion** | Pointer to **bool** |  | [optional] 
**ScriptArguments** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**SeparateContext** | Pointer to **bool** |  | [optional] 
**Skip** | Pointer to **bool** | Skip - if true then this task will be skipped and all the tasks which comes after this task and depend on it will skip (WillNotExecute) | [optional] 
**SkipUnavailable** | Pointer to **bool** | SkipUnavailable if true then will check if automation exists, integration of that command is installed and active or sub playbook exists in Demisto | [optional] 
**Sla** | Pointer to [**SLA**](SLA.md) |  | [optional] 
**SlaReminder** | Pointer to [**SLA**](SLA.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**StartedExecutingHarmfulScript** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** | TaskState indicates the state of the task during the incident/investigation execution | [optional] 
**SubPlaybook** | Pointer to [**InvestigationPlaybook**](InvestigationPlaybook.md) |  | [optional] 
**Task** | Pointer to [**Task**](Task.md) |  | [optional] 
**TaskCompleteData** | Pointer to [**[]InvPlaybookTaskCompleteData**](InvPlaybookTaskCompleteData.md) | History complete data | [optional] 
**TaskDebug** | Pointer to [**InvTaskDebug**](InvTaskDebug.md) |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskSize** | Pointer to **int64** |  | [optional] 
**TimerTriggers** | Pointer to [**[]TimerTrigger**](TimerTrigger.md) | SLA fields | [optional] 
**TimesPolled** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** | TaskType is the Task in the playbook context as a node | [optional] 
**View** | Pointer to **map[string]interface{}** | TaskView represents the view in client of the tasks graph | [optional] 
**WillNotExecuteCount** | Pointer to **int64** |  | [optional] 
**WillNotExecuteReason** | Pointer to **string** |  | [optional] 

## Methods

### NewInvestigationPlaybookTask

`func NewInvestigationPlaybookTask() *InvestigationPlaybookTask`

NewInvestigationPlaybookTask instantiates a new InvestigationPlaybookTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationPlaybookTaskWithDefaults

`func NewInvestigationPlaybookTaskWithDefaults() *InvestigationPlaybookTask`

NewInvestigationPlaybookTaskWithDefaults instantiates a new InvestigationPlaybookTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *InvestigationPlaybookTask) GetArguments() map[string]map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *InvestigationPlaybookTask) GetArgumentsOk() (*map[string]map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *InvestigationPlaybookTask) SetArguments(v map[string]map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *InvestigationPlaybookTask) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetAssignee

`func (o *InvestigationPlaybookTask) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InvestigationPlaybookTask) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InvestigationPlaybookTask) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InvestigationPlaybookTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAssigneeSet

`func (o *InvestigationPlaybookTask) GetAssigneeSet() bool`

GetAssigneeSet returns the AssigneeSet field if non-nil, zero value otherwise.

### GetAssigneeSetOk

`func (o *InvestigationPlaybookTask) GetAssigneeSetOk() (*bool, bool)`

GetAssigneeSetOk returns a tuple with the AssigneeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeSet

`func (o *InvestigationPlaybookTask) SetAssigneeSet(v bool)`

SetAssigneeSet sets AssigneeSet field to given value.

### HasAssigneeSet

`func (o *InvestigationPlaybookTask) HasAssigneeSet() bool`

HasAssigneeSet returns a boolean if a field has been set.

### GetBlockingTasks

`func (o *InvestigationPlaybookTask) GetBlockingTasks() []string`

GetBlockingTasks returns the BlockingTasks field if non-nil, zero value otherwise.

### GetBlockingTasksOk

`func (o *InvestigationPlaybookTask) GetBlockingTasksOk() (*[]string, bool)`

GetBlockingTasksOk returns a tuple with the BlockingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingTasks

`func (o *InvestigationPlaybookTask) SetBlockingTasks(v []string)`

SetBlockingTasks sets BlockingTasks field to given value.

### HasBlockingTasks

`func (o *InvestigationPlaybookTask) HasBlockingTasks() bool`

HasBlockingTasks returns a boolean if a field has been set.

### GetCalculatedDescription

`func (o *InvestigationPlaybookTask) GetCalculatedDescription() string`

GetCalculatedDescription returns the CalculatedDescription field if non-nil, zero value otherwise.

### GetCalculatedDescriptionOk

`func (o *InvestigationPlaybookTask) GetCalculatedDescriptionOk() (*string, bool)`

GetCalculatedDescriptionOk returns a tuple with the CalculatedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedDescription

`func (o *InvestigationPlaybookTask) SetCalculatedDescription(v string)`

SetCalculatedDescription sets CalculatedDescription field to given value.

### HasCalculatedDescription

`func (o *InvestigationPlaybookTask) HasCalculatedDescription() bool`

HasCalculatedDescription returns a boolean if a field has been set.

### GetCalculatedTaskName

`func (o *InvestigationPlaybookTask) GetCalculatedTaskName() string`

GetCalculatedTaskName returns the CalculatedTaskName field if non-nil, zero value otherwise.

### GetCalculatedTaskNameOk

`func (o *InvestigationPlaybookTask) GetCalculatedTaskNameOk() (*string, bool)`

GetCalculatedTaskNameOk returns a tuple with the CalculatedTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedTaskName

`func (o *InvestigationPlaybookTask) SetCalculatedTaskName(v string)`

SetCalculatedTaskName sets CalculatedTaskName field to given value.

### HasCalculatedTaskName

`func (o *InvestigationPlaybookTask) HasCalculatedTaskName() bool`

HasCalculatedTaskName returns a boolean if a field has been set.

### GetComments

`func (o *InvestigationPlaybookTask) GetComments() bool`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InvestigationPlaybookTask) GetCommentsOk() (*bool, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InvestigationPlaybookTask) SetComments(v bool)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InvestigationPlaybookTask) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompletedBy

`func (o *InvestigationPlaybookTask) GetCompletedBy() string`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *InvestigationPlaybookTask) GetCompletedByOk() (*string, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *InvestigationPlaybookTask) SetCompletedBy(v string)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *InvestigationPlaybookTask) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCompletedCount

`func (o *InvestigationPlaybookTask) GetCompletedCount() int64`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *InvestigationPlaybookTask) GetCompletedCountOk() (*int64, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *InvestigationPlaybookTask) SetCompletedCount(v int64)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *InvestigationPlaybookTask) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCompletedDate

`func (o *InvestigationPlaybookTask) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *InvestigationPlaybookTask) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *InvestigationPlaybookTask) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *InvestigationPlaybookTask) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetConditions

`func (o *InvestigationPlaybookTask) GetConditions() []TaskCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InvestigationPlaybookTask) GetConditionsOk() (*[]TaskCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InvestigationPlaybookTask) SetConditions(v []TaskCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InvestigationPlaybookTask) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContinueOnError

`func (o *InvestigationPlaybookTask) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *InvestigationPlaybookTask) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *InvestigationPlaybookTask) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *InvestigationPlaybookTask) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetDefaultAssignee

`func (o *InvestigationPlaybookTask) GetDefaultAssignee() string`

GetDefaultAssignee returns the DefaultAssignee field if non-nil, zero value otherwise.

### GetDefaultAssigneeOk

`func (o *InvestigationPlaybookTask) GetDefaultAssigneeOk() (*string, bool)`

GetDefaultAssigneeOk returns a tuple with the DefaultAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAssignee

`func (o *InvestigationPlaybookTask) SetDefaultAssignee(v string)`

SetDefaultAssignee sets DefaultAssignee field to given value.

### HasDefaultAssignee

`func (o *InvestigationPlaybookTask) HasDefaultAssignee() bool`

HasDefaultAssignee returns a boolean if a field has been set.

### GetDefaultAssigneeComplex

`func (o *InvestigationPlaybookTask) GetDefaultAssigneeComplex() AdvanceArg`

GetDefaultAssigneeComplex returns the DefaultAssigneeComplex field if non-nil, zero value otherwise.

### GetDefaultAssigneeComplexOk

`func (o *InvestigationPlaybookTask) GetDefaultAssigneeComplexOk() (*AdvanceArg, bool)`

GetDefaultAssigneeComplexOk returns a tuple with the DefaultAssigneeComplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAssigneeComplex

`func (o *InvestigationPlaybookTask) SetDefaultAssigneeComplex(v AdvanceArg)`

SetDefaultAssigneeComplex sets DefaultAssigneeComplex field to given value.

### HasDefaultAssigneeComplex

`func (o *InvestigationPlaybookTask) HasDefaultAssigneeComplex() bool`

HasDefaultAssigneeComplex returns a boolean if a field has been set.

### GetDefaultReminder

`func (o *InvestigationPlaybookTask) GetDefaultReminder() int64`

GetDefaultReminder returns the DefaultReminder field if non-nil, zero value otherwise.

### GetDefaultReminderOk

`func (o *InvestigationPlaybookTask) GetDefaultReminderOk() (*int64, bool)`

GetDefaultReminderOk returns a tuple with the DefaultReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReminder

`func (o *InvestigationPlaybookTask) SetDefaultReminder(v int64)`

SetDefaultReminder sets DefaultReminder field to given value.

### HasDefaultReminder

`func (o *InvestigationPlaybookTask) HasDefaultReminder() bool`

HasDefaultReminder returns a boolean if a field has been set.

### GetDoNotSaveTaskHistory

`func (o *InvestigationPlaybookTask) GetDoNotSaveTaskHistory() bool`

GetDoNotSaveTaskHistory returns the DoNotSaveTaskHistory field if non-nil, zero value otherwise.

### GetDoNotSaveTaskHistoryOk

`func (o *InvestigationPlaybookTask) GetDoNotSaveTaskHistoryOk() (*bool, bool)`

GetDoNotSaveTaskHistoryOk returns a tuple with the DoNotSaveTaskHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotSaveTaskHistory

`func (o *InvestigationPlaybookTask) SetDoNotSaveTaskHistory(v bool)`

SetDoNotSaveTaskHistory sets DoNotSaveTaskHistory field to given value.

### HasDoNotSaveTaskHistory

`func (o *InvestigationPlaybookTask) HasDoNotSaveTaskHistory() bool`

HasDoNotSaveTaskHistory returns a boolean if a field has been set.

### GetDueDate

`func (o *InvestigationPlaybookTask) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvestigationPlaybookTask) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvestigationPlaybookTask) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvestigationPlaybookTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDueDateDuration

`func (o *InvestigationPlaybookTask) GetDueDateDuration() int64`

GetDueDateDuration returns the DueDateDuration field if non-nil, zero value otherwise.

### GetDueDateDurationOk

`func (o *InvestigationPlaybookTask) GetDueDateDurationOk() (*int64, bool)`

GetDueDateDurationOk returns a tuple with the DueDateDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateDuration

`func (o *InvestigationPlaybookTask) SetDueDateDuration(v int64)`

SetDueDateDuration sets DueDateDuration field to given value.

### HasDueDateDuration

`func (o *InvestigationPlaybookTask) HasDueDateDuration() bool`

HasDueDateDuration returns a boolean if a field has been set.

### GetDueDateSet

`func (o *InvestigationPlaybookTask) GetDueDateSet() bool`

GetDueDateSet returns the DueDateSet field if non-nil, zero value otherwise.

### GetDueDateSetOk

`func (o *InvestigationPlaybookTask) GetDueDateSetOk() (*bool, bool)`

GetDueDateSetOk returns a tuple with the DueDateSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateSet

`func (o *InvestigationPlaybookTask) SetDueDateSet(v bool)`

SetDueDateSet sets DueDateSet field to given value.

### HasDueDateSet

`func (o *InvestigationPlaybookTask) HasDueDateSet() bool`

HasDueDateSet returns a boolean if a field has been set.

### GetEntries

`func (o *InvestigationPlaybookTask) GetEntries() []string`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InvestigationPlaybookTask) GetEntriesOk() (*[]string, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InvestigationPlaybookTask) SetEntries(v []string)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InvestigationPlaybookTask) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetEvidenceData

`func (o *InvestigationPlaybookTask) GetEvidenceData() EvidenceData`

GetEvidenceData returns the EvidenceData field if non-nil, zero value otherwise.

### GetEvidenceDataOk

`func (o *InvestigationPlaybookTask) GetEvidenceDataOk() (*EvidenceData, bool)`

GetEvidenceDataOk returns a tuple with the EvidenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceData

`func (o *InvestigationPlaybookTask) SetEvidenceData(v EvidenceData)`

SetEvidenceData sets EvidenceData field to given value.

### HasEvidenceData

`func (o *InvestigationPlaybookTask) HasEvidenceData() bool`

HasEvidenceData returns a boolean if a field has been set.

### GetExecutionCount

`func (o *InvestigationPlaybookTask) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *InvestigationPlaybookTask) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *InvestigationPlaybookTask) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *InvestigationPlaybookTask) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetFieldMapping

`func (o *InvestigationPlaybookTask) GetFieldMapping() []FieldMapping`

GetFieldMapping returns the FieldMapping field if non-nil, zero value otherwise.

### GetFieldMappingOk

`func (o *InvestigationPlaybookTask) GetFieldMappingOk() (*[]FieldMapping, bool)`

GetFieldMappingOk returns a tuple with the FieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMapping

`func (o *InvestigationPlaybookTask) SetFieldMapping(v []FieldMapping)`

SetFieldMapping sets FieldMapping field to given value.

### HasFieldMapping

`func (o *InvestigationPlaybookTask) HasFieldMapping() bool`

HasFieldMapping returns a boolean if a field has been set.

### GetForEachIndex

`func (o *InvestigationPlaybookTask) GetForEachIndex() int64`

GetForEachIndex returns the ForEachIndex field if non-nil, zero value otherwise.

### GetForEachIndexOk

`func (o *InvestigationPlaybookTask) GetForEachIndexOk() (*int64, bool)`

GetForEachIndexOk returns a tuple with the ForEachIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForEachIndex

`func (o *InvestigationPlaybookTask) SetForEachIndex(v int64)`

SetForEachIndex sets ForEachIndex field to given value.

### HasForEachIndex

`func (o *InvestigationPlaybookTask) HasForEachIndex() bool`

HasForEachIndex returns a boolean if a field has been set.

### GetForEachInputs

`func (o *InvestigationPlaybookTask) GetForEachInputs() map[string][]map[string]interface{}`

GetForEachInputs returns the ForEachInputs field if non-nil, zero value otherwise.

### GetForEachInputsOk

`func (o *InvestigationPlaybookTask) GetForEachInputsOk() (*map[string][]map[string]interface{}, bool)`

GetForEachInputsOk returns a tuple with the ForEachInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForEachInputs

`func (o *InvestigationPlaybookTask) SetForEachInputs(v map[string][]map[string]interface{})`

SetForEachInputs sets ForEachInputs field to given value.

### HasForEachInputs

`func (o *InvestigationPlaybookTask) HasForEachInputs() bool`

HasForEachInputs returns a boolean if a field has been set.

### GetForm

`func (o *InvestigationPlaybookTask) GetForm() DataCollectionForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *InvestigationPlaybookTask) GetFormOk() (*DataCollectionForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *InvestigationPlaybookTask) SetForm(v DataCollectionForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *InvestigationPlaybookTask) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetFormDisplay

`func (o *InvestigationPlaybookTask) GetFormDisplay() FormDisplay`

GetFormDisplay returns the FormDisplay field if non-nil, zero value otherwise.

### GetFormDisplayOk

`func (o *InvestigationPlaybookTask) GetFormDisplayOk() (*FormDisplay, bool)`

GetFormDisplayOk returns a tuple with the FormDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormDisplay

`func (o *InvestigationPlaybookTask) SetFormDisplay(v FormDisplay)`

SetFormDisplay sets FormDisplay field to given value.

### HasFormDisplay

`func (o *InvestigationPlaybookTask) HasFormDisplay() bool`

HasFormDisplay returns a boolean if a field has been set.

### GetId

`func (o *InvestigationPlaybookTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvestigationPlaybookTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvestigationPlaybookTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvestigationPlaybookTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreWorker

`func (o *InvestigationPlaybookTask) GetIgnoreWorker() bool`

GetIgnoreWorker returns the IgnoreWorker field if non-nil, zero value otherwise.

### GetIgnoreWorkerOk

`func (o *InvestigationPlaybookTask) GetIgnoreWorkerOk() (*bool, bool)`

GetIgnoreWorkerOk returns a tuple with the IgnoreWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorker

`func (o *InvestigationPlaybookTask) SetIgnoreWorker(v bool)`

SetIgnoreWorker sets IgnoreWorker field to given value.

### HasIgnoreWorker

`func (o *InvestigationPlaybookTask) HasIgnoreWorker() bool`

HasIgnoreWorker returns a boolean if a field has been set.

### GetIndent

`func (o *InvestigationPlaybookTask) GetIndent() int64`

GetIndent returns the Indent field if non-nil, zero value otherwise.

### GetIndentOk

`func (o *InvestigationPlaybookTask) GetIndentOk() (*int64, bool)`

GetIndentOk returns a tuple with the Indent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndent

`func (o *InvestigationPlaybookTask) SetIndent(v int64)`

SetIndent sets Indent field to given value.

### HasIndent

`func (o *InvestigationPlaybookTask) HasIndent() bool`

HasIndent returns a boolean if a field has been set.

### GetInput

`func (o *InvestigationPlaybookTask) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InvestigationPlaybookTask) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InvestigationPlaybookTask) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *InvestigationPlaybookTask) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetIsAutoSwitchedToQuietMode

`func (o *InvestigationPlaybookTask) GetIsAutoSwitchedToQuietMode() bool`

GetIsAutoSwitchedToQuietMode returns the IsAutoSwitchedToQuietMode field if non-nil, zero value otherwise.

### GetIsAutoSwitchedToQuietModeOk

`func (o *InvestigationPlaybookTask) GetIsAutoSwitchedToQuietModeOk() (*bool, bool)`

GetIsAutoSwitchedToQuietModeOk returns a tuple with the IsAutoSwitchedToQuietMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSwitchedToQuietMode

`func (o *InvestigationPlaybookTask) SetIsAutoSwitchedToQuietMode(v bool)`

SetIsAutoSwitchedToQuietMode sets IsAutoSwitchedToQuietMode field to given value.

### HasIsAutoSwitchedToQuietMode

`func (o *InvestigationPlaybookTask) HasIsAutoSwitchedToQuietMode() bool`

HasIsAutoSwitchedToQuietMode returns a boolean if a field has been set.

### GetIsOverSize

`func (o *InvestigationPlaybookTask) GetIsOverSize() bool`

GetIsOverSize returns the IsOverSize field if non-nil, zero value otherwise.

### GetIsOverSizeOk

`func (o *InvestigationPlaybookTask) GetIsOverSizeOk() (*bool, bool)`

GetIsOverSizeOk returns a tuple with the IsOverSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverSize

`func (o *InvestigationPlaybookTask) SetIsOverSize(v bool)`

SetIsOverSize sets IsOverSize field to given value.

### HasIsOverSize

`func (o *InvestigationPlaybookTask) HasIsOverSize() bool`

HasIsOverSize returns a boolean if a field has been set.

### GetLoop

`func (o *InvestigationPlaybookTask) GetLoop() TaskLoop`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *InvestigationPlaybookTask) GetLoopOk() (*TaskLoop, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *InvestigationPlaybookTask) SetLoop(v TaskLoop)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *InvestigationPlaybookTask) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetMessage

`func (o *InvestigationPlaybookTask) GetMessage() NotifiableItem`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvestigationPlaybookTask) GetMessageOk() (*NotifiableItem, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvestigationPlaybookTask) SetMessage(v NotifiableItem)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvestigationPlaybookTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMissingContentItem

`func (o *InvestigationPlaybookTask) GetMissingContentItem() string`

GetMissingContentItem returns the MissingContentItem field if non-nil, zero value otherwise.

### GetMissingContentItemOk

`func (o *InvestigationPlaybookTask) GetMissingContentItemOk() (*string, bool)`

GetMissingContentItemOk returns a tuple with the MissingContentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingContentItem

`func (o *InvestigationPlaybookTask) SetMissingContentItem(v string)`

SetMissingContentItem sets MissingContentItem field to given value.

### HasMissingContentItem

`func (o *InvestigationPlaybookTask) HasMissingContentItem() bool`

HasMissingContentItem returns a boolean if a field has been set.

### GetNextPollingTime

`func (o *InvestigationPlaybookTask) GetNextPollingTime() time.Time`

GetNextPollingTime returns the NextPollingTime field if non-nil, zero value otherwise.

### GetNextPollingTimeOk

`func (o *InvestigationPlaybookTask) GetNextPollingTimeOk() (*time.Time, bool)`

GetNextPollingTimeOk returns a tuple with the NextPollingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPollingTime

`func (o *InvestigationPlaybookTask) SetNextPollingTime(v time.Time)`

SetNextPollingTime sets NextPollingTime field to given value.

### HasNextPollingTime

`func (o *InvestigationPlaybookTask) HasNextPollingTime() bool`

HasNextPollingTime returns a boolean if a field has been set.

### GetNextTasks

`func (o *InvestigationPlaybookTask) GetNextTasks() map[string][]string`

GetNextTasks returns the NextTasks field if non-nil, zero value otherwise.

### GetNextTasksOk

`func (o *InvestigationPlaybookTask) GetNextTasksOk() (*map[string][]string, bool)`

GetNextTasksOk returns a tuple with the NextTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTasks

`func (o *InvestigationPlaybookTask) SetNextTasks(v map[string][]string)`

SetNextTasks sets NextTasks field to given value.

### HasNextTasks

`func (o *InvestigationPlaybookTask) HasNextTasks() bool`

HasNextTasks returns a boolean if a field has been set.

### GetNote

`func (o *InvestigationPlaybookTask) GetNote() bool`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *InvestigationPlaybookTask) GetNoteOk() (*bool, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *InvestigationPlaybookTask) SetNote(v bool)`

SetNote sets Note field to given value.

### HasNote

`func (o *InvestigationPlaybookTask) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutputs

`func (o *InvestigationPlaybookTask) GetOutputs() map[string]map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *InvestigationPlaybookTask) GetOutputsOk() (*map[string]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *InvestigationPlaybookTask) SetOutputs(v map[string]map[string]interface{})`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *InvestigationPlaybookTask) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetParentBlockCount

`func (o *InvestigationPlaybookTask) GetParentBlockCount() int64`

GetParentBlockCount returns the ParentBlockCount field if non-nil, zero value otherwise.

### GetParentBlockCountOk

`func (o *InvestigationPlaybookTask) GetParentBlockCountOk() (*int64, bool)`

GetParentBlockCountOk returns a tuple with the ParentBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockCount

`func (o *InvestigationPlaybookTask) SetParentBlockCount(v int64)`

SetParentBlockCount sets ParentBlockCount field to given value.

### HasParentBlockCount

`func (o *InvestigationPlaybookTask) HasParentBlockCount() bool`

HasParentBlockCount returns a boolean if a field has been set.

### GetParentPlaybookID

`func (o *InvestigationPlaybookTask) GetParentPlaybookID() string`

GetParentPlaybookID returns the ParentPlaybookID field if non-nil, zero value otherwise.

### GetParentPlaybookIDOk

`func (o *InvestigationPlaybookTask) GetParentPlaybookIDOk() (*string, bool)`

GetParentPlaybookIDOk returns a tuple with the ParentPlaybookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPlaybookID

`func (o *InvestigationPlaybookTask) SetParentPlaybookID(v string)`

SetParentPlaybookID sets ParentPlaybookID field to given value.

### HasParentPlaybookID

`func (o *InvestigationPlaybookTask) HasParentPlaybookID() bool`

HasParentPlaybookID returns a boolean if a field has been set.

### GetPatched

`func (o *InvestigationPlaybookTask) GetPatched() bool`

GetPatched returns the Patched field if non-nil, zero value otherwise.

### GetPatchedOk

`func (o *InvestigationPlaybookTask) GetPatchedOk() (*bool, bool)`

GetPatchedOk returns a tuple with the Patched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatched

`func (o *InvestigationPlaybookTask) SetPatched(v bool)`

SetPatched sets Patched field to given value.

### HasPatched

`func (o *InvestigationPlaybookTask) HasPatched() bool`

HasPatched returns a boolean if a field has been set.

### GetPlaybookInputs

`func (o *InvestigationPlaybookTask) GetPlaybookInputs() map[string]map[string]interface{}`

GetPlaybookInputs returns the PlaybookInputs field if non-nil, zero value otherwise.

### GetPlaybookInputsOk

`func (o *InvestigationPlaybookTask) GetPlaybookInputsOk() (*map[string]map[string]interface{}, bool)`

GetPlaybookInputsOk returns a tuple with the PlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookInputs

`func (o *InvestigationPlaybookTask) SetPlaybookInputs(v map[string]map[string]interface{})`

SetPlaybookInputs sets PlaybookInputs field to given value.

### HasPlaybookInputs

`func (o *InvestigationPlaybookTask) HasPlaybookInputs() bool`

HasPlaybookInputs returns a boolean if a field has been set.

### GetPollingEntries

`func (o *InvestigationPlaybookTask) GetPollingEntries() []string`

GetPollingEntries returns the PollingEntries field if non-nil, zero value otherwise.

### GetPollingEntriesOk

`func (o *InvestigationPlaybookTask) GetPollingEntriesOk() (*[]string, bool)`

GetPollingEntriesOk returns a tuple with the PollingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingEntries

`func (o *InvestigationPlaybookTask) SetPollingEntries(v []string)`

SetPollingEntries sets PollingEntries field to given value.

### HasPollingEntries

`func (o *InvestigationPlaybookTask) HasPollingEntries() bool`

HasPollingEntries returns a boolean if a field has been set.

### GetPollingTimeoutTime

`func (o *InvestigationPlaybookTask) GetPollingTimeoutTime() time.Time`

GetPollingTimeoutTime returns the PollingTimeoutTime field if non-nil, zero value otherwise.

### GetPollingTimeoutTimeOk

`func (o *InvestigationPlaybookTask) GetPollingTimeoutTimeOk() (*time.Time, bool)`

GetPollingTimeoutTimeOk returns a tuple with the PollingTimeoutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingTimeoutTime

`func (o *InvestigationPlaybookTask) SetPollingTimeoutTime(v time.Time)`

SetPollingTimeoutTime sets PollingTimeoutTime field to given value.

### HasPollingTimeoutTime

`func (o *InvestigationPlaybookTask) HasPollingTimeoutTime() bool`

HasPollingTimeoutTime returns a boolean if a field has been set.

### GetPreviousTasks

`func (o *InvestigationPlaybookTask) GetPreviousTasks() map[string][]string`

GetPreviousTasks returns the PreviousTasks field if non-nil, zero value otherwise.

### GetPreviousTasksOk

`func (o *InvestigationPlaybookTask) GetPreviousTasksOk() (*map[string][]string, bool)`

GetPreviousTasksOk returns a tuple with the PreviousTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTasks

`func (o *InvestigationPlaybookTask) SetPreviousTasks(v map[string][]string)`

SetPreviousTasks sets PreviousTasks field to given value.

### HasPreviousTasks

`func (o *InvestigationPlaybookTask) HasPreviousTasks() bool`

HasPreviousTasks returns a boolean if a field has been set.

### GetQuietMode

`func (o *InvestigationPlaybookTask) GetQuietMode() int64`

GetQuietMode returns the QuietMode field if non-nil, zero value otherwise.

### GetQuietModeOk

`func (o *InvestigationPlaybookTask) GetQuietModeOk() (*int64, bool)`

GetQuietModeOk returns a tuple with the QuietMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietMode

`func (o *InvestigationPlaybookTask) SetQuietMode(v int64)`

SetQuietMode sets QuietMode field to given value.

### HasQuietMode

`func (o *InvestigationPlaybookTask) HasQuietMode() bool`

HasQuietMode returns a boolean if a field has been set.

### GetReminder

`func (o *InvestigationPlaybookTask) GetReminder() int64`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *InvestigationPlaybookTask) GetReminderOk() (*int64, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *InvestigationPlaybookTask) SetReminder(v int64)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *InvestigationPlaybookTask) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetReputationCalc

`func (o *InvestigationPlaybookTask) GetReputationCalc() float64`

GetReputationCalc returns the ReputationCalc field if non-nil, zero value otherwise.

### GetReputationCalcOk

`func (o *InvestigationPlaybookTask) GetReputationCalcOk() (*float64, bool)`

GetReputationCalcOk returns a tuple with the ReputationCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationCalc

`func (o *InvestigationPlaybookTask) SetReputationCalc(v float64)`

SetReputationCalc sets ReputationCalc field to given value.

### HasReputationCalc

`func (o *InvestigationPlaybookTask) HasReputationCalc() bool`

HasReputationCalc returns a boolean if a field has been set.

### GetRestrictedCompletion

`func (o *InvestigationPlaybookTask) GetRestrictedCompletion() bool`

GetRestrictedCompletion returns the RestrictedCompletion field if non-nil, zero value otherwise.

### GetRestrictedCompletionOk

`func (o *InvestigationPlaybookTask) GetRestrictedCompletionOk() (*bool, bool)`

GetRestrictedCompletionOk returns a tuple with the RestrictedCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCompletion

`func (o *InvestigationPlaybookTask) SetRestrictedCompletion(v bool)`

SetRestrictedCompletion sets RestrictedCompletion field to given value.

### HasRestrictedCompletion

`func (o *InvestigationPlaybookTask) HasRestrictedCompletion() bool`

HasRestrictedCompletion returns a boolean if a field has been set.

### GetScriptArguments

`func (o *InvestigationPlaybookTask) GetScriptArguments() map[string]AdvanceArg`

GetScriptArguments returns the ScriptArguments field if non-nil, zero value otherwise.

### GetScriptArgumentsOk

`func (o *InvestigationPlaybookTask) GetScriptArgumentsOk() (*map[string]AdvanceArg, bool)`

GetScriptArgumentsOk returns a tuple with the ScriptArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArguments

`func (o *InvestigationPlaybookTask) SetScriptArguments(v map[string]AdvanceArg)`

SetScriptArguments sets ScriptArguments field to given value.

### HasScriptArguments

`func (o *InvestigationPlaybookTask) HasScriptArguments() bool`

HasScriptArguments returns a boolean if a field has been set.

### GetSeparateContext

`func (o *InvestigationPlaybookTask) GetSeparateContext() bool`

GetSeparateContext returns the SeparateContext field if non-nil, zero value otherwise.

### GetSeparateContextOk

`func (o *InvestigationPlaybookTask) GetSeparateContextOk() (*bool, bool)`

GetSeparateContextOk returns a tuple with the SeparateContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateContext

`func (o *InvestigationPlaybookTask) SetSeparateContext(v bool)`

SetSeparateContext sets SeparateContext field to given value.

### HasSeparateContext

`func (o *InvestigationPlaybookTask) HasSeparateContext() bool`

HasSeparateContext returns a boolean if a field has been set.

### GetSkip

`func (o *InvestigationPlaybookTask) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *InvestigationPlaybookTask) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *InvestigationPlaybookTask) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *InvestigationPlaybookTask) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSkipUnavailable

`func (o *InvestigationPlaybookTask) GetSkipUnavailable() bool`

GetSkipUnavailable returns the SkipUnavailable field if non-nil, zero value otherwise.

### GetSkipUnavailableOk

`func (o *InvestigationPlaybookTask) GetSkipUnavailableOk() (*bool, bool)`

GetSkipUnavailableOk returns a tuple with the SkipUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUnavailable

`func (o *InvestigationPlaybookTask) SetSkipUnavailable(v bool)`

SetSkipUnavailable sets SkipUnavailable field to given value.

### HasSkipUnavailable

`func (o *InvestigationPlaybookTask) HasSkipUnavailable() bool`

HasSkipUnavailable returns a boolean if a field has been set.

### GetSla

`func (o *InvestigationPlaybookTask) GetSla() SLA`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *InvestigationPlaybookTask) GetSlaOk() (*SLA, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *InvestigationPlaybookTask) SetSla(v SLA)`

SetSla sets Sla field to given value.

### HasSla

`func (o *InvestigationPlaybookTask) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaReminder

`func (o *InvestigationPlaybookTask) GetSlaReminder() SLA`

GetSlaReminder returns the SlaReminder field if non-nil, zero value otherwise.

### GetSlaReminderOk

`func (o *InvestigationPlaybookTask) GetSlaReminderOk() (*SLA, bool)`

GetSlaReminderOk returns a tuple with the SlaReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaReminder

`func (o *InvestigationPlaybookTask) SetSlaReminder(v SLA)`

SetSlaReminder sets SlaReminder field to given value.

### HasSlaReminder

`func (o *InvestigationPlaybookTask) HasSlaReminder() bool`

HasSlaReminder returns a boolean if a field has been set.

### GetStartDate

`func (o *InvestigationPlaybookTask) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvestigationPlaybookTask) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvestigationPlaybookTask) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvestigationPlaybookTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartedExecutingHarmfulScript

`func (o *InvestigationPlaybookTask) GetStartedExecutingHarmfulScript() bool`

GetStartedExecutingHarmfulScript returns the StartedExecutingHarmfulScript field if non-nil, zero value otherwise.

### GetStartedExecutingHarmfulScriptOk

`func (o *InvestigationPlaybookTask) GetStartedExecutingHarmfulScriptOk() (*bool, bool)`

GetStartedExecutingHarmfulScriptOk returns a tuple with the StartedExecutingHarmfulScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedExecutingHarmfulScript

`func (o *InvestigationPlaybookTask) SetStartedExecutingHarmfulScript(v bool)`

SetStartedExecutingHarmfulScript sets StartedExecutingHarmfulScript field to given value.

### HasStartedExecutingHarmfulScript

`func (o *InvestigationPlaybookTask) HasStartedExecutingHarmfulScript() bool`

HasStartedExecutingHarmfulScript returns a boolean if a field has been set.

### GetState

`func (o *InvestigationPlaybookTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvestigationPlaybookTask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvestigationPlaybookTask) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InvestigationPlaybookTask) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubPlaybook

`func (o *InvestigationPlaybookTask) GetSubPlaybook() InvestigationPlaybook`

GetSubPlaybook returns the SubPlaybook field if non-nil, zero value otherwise.

### GetSubPlaybookOk

`func (o *InvestigationPlaybookTask) GetSubPlaybookOk() (*InvestigationPlaybook, bool)`

GetSubPlaybookOk returns a tuple with the SubPlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPlaybook

`func (o *InvestigationPlaybookTask) SetSubPlaybook(v InvestigationPlaybook)`

SetSubPlaybook sets SubPlaybook field to given value.

### HasSubPlaybook

`func (o *InvestigationPlaybookTask) HasSubPlaybook() bool`

HasSubPlaybook returns a boolean if a field has been set.

### GetTask

`func (o *InvestigationPlaybookTask) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *InvestigationPlaybookTask) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *InvestigationPlaybookTask) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *InvestigationPlaybookTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskCompleteData

`func (o *InvestigationPlaybookTask) GetTaskCompleteData() []InvPlaybookTaskCompleteData`

GetTaskCompleteData returns the TaskCompleteData field if non-nil, zero value otherwise.

### GetTaskCompleteDataOk

`func (o *InvestigationPlaybookTask) GetTaskCompleteDataOk() (*[]InvPlaybookTaskCompleteData, bool)`

GetTaskCompleteDataOk returns a tuple with the TaskCompleteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompleteData

`func (o *InvestigationPlaybookTask) SetTaskCompleteData(v []InvPlaybookTaskCompleteData)`

SetTaskCompleteData sets TaskCompleteData field to given value.

### HasTaskCompleteData

`func (o *InvestigationPlaybookTask) HasTaskCompleteData() bool`

HasTaskCompleteData returns a boolean if a field has been set.

### GetTaskDebug

`func (o *InvestigationPlaybookTask) GetTaskDebug() InvTaskDebug`

GetTaskDebug returns the TaskDebug field if non-nil, zero value otherwise.

### GetTaskDebugOk

`func (o *InvestigationPlaybookTask) GetTaskDebugOk() (*InvTaskDebug, bool)`

GetTaskDebugOk returns a tuple with the TaskDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDebug

`func (o *InvestigationPlaybookTask) SetTaskDebug(v InvTaskDebug)`

SetTaskDebug sets TaskDebug field to given value.

### HasTaskDebug

`func (o *InvestigationPlaybookTask) HasTaskDebug() bool`

HasTaskDebug returns a boolean if a field has been set.

### GetTaskId

`func (o *InvestigationPlaybookTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *InvestigationPlaybookTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *InvestigationPlaybookTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *InvestigationPlaybookTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskSize

`func (o *InvestigationPlaybookTask) GetTaskSize() int64`

GetTaskSize returns the TaskSize field if non-nil, zero value otherwise.

### GetTaskSizeOk

`func (o *InvestigationPlaybookTask) GetTaskSizeOk() (*int64, bool)`

GetTaskSizeOk returns a tuple with the TaskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSize

`func (o *InvestigationPlaybookTask) SetTaskSize(v int64)`

SetTaskSize sets TaskSize field to given value.

### HasTaskSize

`func (o *InvestigationPlaybookTask) HasTaskSize() bool`

HasTaskSize returns a boolean if a field has been set.

### GetTimerTriggers

`func (o *InvestigationPlaybookTask) GetTimerTriggers() []TimerTrigger`

GetTimerTriggers returns the TimerTriggers field if non-nil, zero value otherwise.

### GetTimerTriggersOk

`func (o *InvestigationPlaybookTask) GetTimerTriggersOk() (*[]TimerTrigger, bool)`

GetTimerTriggersOk returns a tuple with the TimerTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerTriggers

`func (o *InvestigationPlaybookTask) SetTimerTriggers(v []TimerTrigger)`

SetTimerTriggers sets TimerTriggers field to given value.

### HasTimerTriggers

`func (o *InvestigationPlaybookTask) HasTimerTriggers() bool`

HasTimerTriggers returns a boolean if a field has been set.

### GetTimesPolled

`func (o *InvestigationPlaybookTask) GetTimesPolled() int64`

GetTimesPolled returns the TimesPolled field if non-nil, zero value otherwise.

### GetTimesPolledOk

`func (o *InvestigationPlaybookTask) GetTimesPolledOk() (*int64, bool)`

GetTimesPolledOk returns a tuple with the TimesPolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesPolled

`func (o *InvestigationPlaybookTask) SetTimesPolled(v int64)`

SetTimesPolled sets TimesPolled field to given value.

### HasTimesPolled

`func (o *InvestigationPlaybookTask) HasTimesPolled() bool`

HasTimesPolled returns a boolean if a field has been set.

### GetType

`func (o *InvestigationPlaybookTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvestigationPlaybookTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvestigationPlaybookTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvestigationPlaybookTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *InvestigationPlaybookTask) GetView() map[string]interface{}`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *InvestigationPlaybookTask) GetViewOk() (*map[string]interface{}, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *InvestigationPlaybookTask) SetView(v map[string]interface{})`

SetView sets View field to given value.

### HasView

`func (o *InvestigationPlaybookTask) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWillNotExecuteCount

`func (o *InvestigationPlaybookTask) GetWillNotExecuteCount() int64`

GetWillNotExecuteCount returns the WillNotExecuteCount field if non-nil, zero value otherwise.

### GetWillNotExecuteCountOk

`func (o *InvestigationPlaybookTask) GetWillNotExecuteCountOk() (*int64, bool)`

GetWillNotExecuteCountOk returns a tuple with the WillNotExecuteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillNotExecuteCount

`func (o *InvestigationPlaybookTask) SetWillNotExecuteCount(v int64)`

SetWillNotExecuteCount sets WillNotExecuteCount field to given value.

### HasWillNotExecuteCount

`func (o *InvestigationPlaybookTask) HasWillNotExecuteCount() bool`

HasWillNotExecuteCount returns a boolean if a field has been set.

### GetWillNotExecuteReason

`func (o *InvestigationPlaybookTask) GetWillNotExecuteReason() string`

GetWillNotExecuteReason returns the WillNotExecuteReason field if non-nil, zero value otherwise.

### GetWillNotExecuteReasonOk

`func (o *InvestigationPlaybookTask) GetWillNotExecuteReasonOk() (*string, bool)`

GetWillNotExecuteReasonOk returns a tuple with the WillNotExecuteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillNotExecuteReason

`func (o *InvestigationPlaybookTask) SetWillNotExecuteReason(v string)`

SetWillNotExecuteReason sets WillNotExecuteReason field to given value.

### HasWillNotExecuteReason

`func (o *InvestigationPlaybookTask) HasWillNotExecuteReason() bool`

HasWillNotExecuteReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


