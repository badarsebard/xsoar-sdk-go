# InvestigationPlaybook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dirty** | Pointer to **bool** |  | [optional] 
**ReadyPlaybookInputs** | Pointer to **map[string]map[string]map[string]interface{}** |  | [optional] 
**ReplacedPlaybook** | Pointer to **bool** | Indicate whether this playbook has new history during this session | [optional] 
**ShardID** | Pointer to **int64** |  | [optional] 
**UpdatedOperatorIDs** | Pointer to **bool** |  | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**AutoExtracting** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncidentCreateDate** | Pointer to **time.Time** | Incident create date | [optional] 
**Inputs** | Pointer to [**[]PlaybookInput**](PlaybookInput.md) | PlaybookInputs - array of PlaybookInput | [optional] 
**InvPBDebugInfo** | Pointer to [**InvPlaybookDebugInfo**](InvPlaybookDebugInfo.md) |  | [optional] 
**InvestigationId** | Pointer to **string** |  | [optional] 
**IsTIM** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Outputs** | Pointer to [**[]PlaybookOutput**](PlaybookOutput.md) | PlaybookOutputs - array of PlaybookOutput | [optional] 
**PbHistory** | Pointer to [**[]InvestigationPlaybookData**](InvestigationPlaybookData.md) | in: body | [optional] 
**PendingTasks** | Pointer to **map[string]map[string]interface{}** | Tasks that are marked for running, but cannot yet run due to incomplete parents | [optional] 
**PlaybookId** | Pointer to **string** |  | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**RecoveryAttempts** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ServerId** | Pointer to **string** | Holds the ID of the responsible cluster app server | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**StartTaskId** | Pointer to **string** | FirstTask is the root task of the playbook | [optional] 
**State** | Pointer to **string** | InvestigationPlaybookState indicates the state of the running playbook | [optional] 
**SubPlaybookInputs** | Pointer to [**map[string][]PlaybookInput**](array.md) |  | [optional] 
**SubPlaybookOutputs** | Pointer to [**map[string][]PlaybookOutput**](array.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]InvestigationPlaybookTask**](InvestigationPlaybookTask.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**View** | Pointer to **map[string]interface{}** | PlaybookView represents the view in client of playbook graph | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInvestigationPlaybook

`func NewInvestigationPlaybook() *InvestigationPlaybook`

NewInvestigationPlaybook instantiates a new InvestigationPlaybook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationPlaybookWithDefaults

`func NewInvestigationPlaybookWithDefaults() *InvestigationPlaybook`

NewInvestigationPlaybookWithDefaults instantiates a new InvestigationPlaybook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirty

`func (o *InvestigationPlaybook) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *InvestigationPlaybook) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *InvestigationPlaybook) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *InvestigationPlaybook) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetReadyPlaybookInputs

`func (o *InvestigationPlaybook) GetReadyPlaybookInputs() map[string]map[string]map[string]interface{}`

GetReadyPlaybookInputs returns the ReadyPlaybookInputs field if non-nil, zero value otherwise.

### GetReadyPlaybookInputsOk

`func (o *InvestigationPlaybook) GetReadyPlaybookInputsOk() (*map[string]map[string]map[string]interface{}, bool)`

GetReadyPlaybookInputsOk returns a tuple with the ReadyPlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyPlaybookInputs

`func (o *InvestigationPlaybook) SetReadyPlaybookInputs(v map[string]map[string]map[string]interface{})`

SetReadyPlaybookInputs sets ReadyPlaybookInputs field to given value.

### HasReadyPlaybookInputs

`func (o *InvestigationPlaybook) HasReadyPlaybookInputs() bool`

HasReadyPlaybookInputs returns a boolean if a field has been set.

### GetReplacedPlaybook

`func (o *InvestigationPlaybook) GetReplacedPlaybook() bool`

GetReplacedPlaybook returns the ReplacedPlaybook field if non-nil, zero value otherwise.

### GetReplacedPlaybookOk

`func (o *InvestigationPlaybook) GetReplacedPlaybookOk() (*bool, bool)`

GetReplacedPlaybookOk returns a tuple with the ReplacedPlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedPlaybook

`func (o *InvestigationPlaybook) SetReplacedPlaybook(v bool)`

SetReplacedPlaybook sets ReplacedPlaybook field to given value.

### HasReplacedPlaybook

`func (o *InvestigationPlaybook) HasReplacedPlaybook() bool`

HasReplacedPlaybook returns a boolean if a field has been set.

### GetShardID

`func (o *InvestigationPlaybook) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *InvestigationPlaybook) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *InvestigationPlaybook) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *InvestigationPlaybook) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetUpdatedOperatorIDs

`func (o *InvestigationPlaybook) GetUpdatedOperatorIDs() bool`

GetUpdatedOperatorIDs returns the UpdatedOperatorIDs field if non-nil, zero value otherwise.

### GetUpdatedOperatorIDsOk

`func (o *InvestigationPlaybook) GetUpdatedOperatorIDsOk() (*bool, bool)`

GetUpdatedOperatorIDsOk returns a tuple with the UpdatedOperatorIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOperatorIDs

`func (o *InvestigationPlaybook) SetUpdatedOperatorIDs(v bool)`

SetUpdatedOperatorIDs sets UpdatedOperatorIDs field to given value.

### HasUpdatedOperatorIDs

`func (o *InvestigationPlaybook) HasUpdatedOperatorIDs() bool`

HasUpdatedOperatorIDs returns a boolean if a field has been set.

### GetAllRead

`func (o *InvestigationPlaybook) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *InvestigationPlaybook) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *InvestigationPlaybook) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *InvestigationPlaybook) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *InvestigationPlaybook) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *InvestigationPlaybook) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *InvestigationPlaybook) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *InvestigationPlaybook) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetAutoExtracting

`func (o *InvestigationPlaybook) GetAutoExtracting() bool`

GetAutoExtracting returns the AutoExtracting field if non-nil, zero value otherwise.

### GetAutoExtractingOk

`func (o *InvestigationPlaybook) GetAutoExtractingOk() (*bool, bool)`

GetAutoExtractingOk returns a tuple with the AutoExtracting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExtracting

`func (o *InvestigationPlaybook) SetAutoExtracting(v bool)`

SetAutoExtracting sets AutoExtracting field to given value.

### HasAutoExtracting

`func (o *InvestigationPlaybook) HasAutoExtracting() bool`

HasAutoExtracting returns a boolean if a field has been set.

### GetComment

`func (o *InvestigationPlaybook) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InvestigationPlaybook) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InvestigationPlaybook) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InvestigationPlaybook) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *InvestigationPlaybook) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *InvestigationPlaybook) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *InvestigationPlaybook) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *InvestigationPlaybook) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetHasRole

`func (o *InvestigationPlaybook) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *InvestigationPlaybook) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *InvestigationPlaybook) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *InvestigationPlaybook) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *InvestigationPlaybook) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *InvestigationPlaybook) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *InvestigationPlaybook) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *InvestigationPlaybook) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *InvestigationPlaybook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvestigationPlaybook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvestigationPlaybook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvestigationPlaybook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentCreateDate

`func (o *InvestigationPlaybook) GetIncidentCreateDate() time.Time`

GetIncidentCreateDate returns the IncidentCreateDate field if non-nil, zero value otherwise.

### GetIncidentCreateDateOk

`func (o *InvestigationPlaybook) GetIncidentCreateDateOk() (*time.Time, bool)`

GetIncidentCreateDateOk returns a tuple with the IncidentCreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCreateDate

`func (o *InvestigationPlaybook) SetIncidentCreateDate(v time.Time)`

SetIncidentCreateDate sets IncidentCreateDate field to given value.

### HasIncidentCreateDate

`func (o *InvestigationPlaybook) HasIncidentCreateDate() bool`

HasIncidentCreateDate returns a boolean if a field has been set.

### GetInputs

`func (o *InvestigationPlaybook) GetInputs() []PlaybookInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *InvestigationPlaybook) GetInputsOk() (*[]PlaybookInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *InvestigationPlaybook) SetInputs(v []PlaybookInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *InvestigationPlaybook) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetInvPBDebugInfo

`func (o *InvestigationPlaybook) GetInvPBDebugInfo() InvPlaybookDebugInfo`

GetInvPBDebugInfo returns the InvPBDebugInfo field if non-nil, zero value otherwise.

### GetInvPBDebugInfoOk

`func (o *InvestigationPlaybook) GetInvPBDebugInfoOk() (*InvPlaybookDebugInfo, bool)`

GetInvPBDebugInfoOk returns a tuple with the InvPBDebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvPBDebugInfo

`func (o *InvestigationPlaybook) SetInvPBDebugInfo(v InvPlaybookDebugInfo)`

SetInvPBDebugInfo sets InvPBDebugInfo field to given value.

### HasInvPBDebugInfo

`func (o *InvestigationPlaybook) HasInvPBDebugInfo() bool`

HasInvPBDebugInfo returns a boolean if a field has been set.

### GetInvestigationId

`func (o *InvestigationPlaybook) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *InvestigationPlaybook) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *InvestigationPlaybook) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *InvestigationPlaybook) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsTIM

`func (o *InvestigationPlaybook) GetIsTIM() bool`

GetIsTIM returns the IsTIM field if non-nil, zero value otherwise.

### GetIsTIMOk

`func (o *InvestigationPlaybook) GetIsTIMOk() (*bool, bool)`

GetIsTIMOk returns a tuple with the IsTIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTIM

`func (o *InvestigationPlaybook) SetIsTIM(v bool)`

SetIsTIM sets IsTIM field to given value.

### HasIsTIM

`func (o *InvestigationPlaybook) HasIsTIM() bool`

HasIsTIM returns a boolean if a field has been set.

### GetModified

`func (o *InvestigationPlaybook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InvestigationPlaybook) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InvestigationPlaybook) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *InvestigationPlaybook) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *InvestigationPlaybook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationPlaybook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationPlaybook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvestigationPlaybook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *InvestigationPlaybook) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *InvestigationPlaybook) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *InvestigationPlaybook) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *InvestigationPlaybook) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOutputs

`func (o *InvestigationPlaybook) GetOutputs() []PlaybookOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *InvestigationPlaybook) GetOutputsOk() (*[]PlaybookOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *InvestigationPlaybook) SetOutputs(v []PlaybookOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *InvestigationPlaybook) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPbHistory

`func (o *InvestigationPlaybook) GetPbHistory() []InvestigationPlaybookData`

GetPbHistory returns the PbHistory field if non-nil, zero value otherwise.

### GetPbHistoryOk

`func (o *InvestigationPlaybook) GetPbHistoryOk() (*[]InvestigationPlaybookData, bool)`

GetPbHistoryOk returns a tuple with the PbHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbHistory

`func (o *InvestigationPlaybook) SetPbHistory(v []InvestigationPlaybookData)`

SetPbHistory sets PbHistory field to given value.

### HasPbHistory

`func (o *InvestigationPlaybook) HasPbHistory() bool`

HasPbHistory returns a boolean if a field has been set.

### GetPendingTasks

`func (o *InvestigationPlaybook) GetPendingTasks() map[string]map[string]interface{}`

GetPendingTasks returns the PendingTasks field if non-nil, zero value otherwise.

### GetPendingTasksOk

`func (o *InvestigationPlaybook) GetPendingTasksOk() (*map[string]map[string]interface{}, bool)`

GetPendingTasksOk returns a tuple with the PendingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTasks

`func (o *InvestigationPlaybook) SetPendingTasks(v map[string]map[string]interface{})`

SetPendingTasks sets PendingTasks field to given value.

### HasPendingTasks

`func (o *InvestigationPlaybook) HasPendingTasks() bool`

HasPendingTasks returns a boolean if a field has been set.

### GetPlaybookId

`func (o *InvestigationPlaybook) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *InvestigationPlaybook) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *InvestigationPlaybook) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *InvestigationPlaybook) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *InvestigationPlaybook) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *InvestigationPlaybook) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *InvestigationPlaybook) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *InvestigationPlaybook) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *InvestigationPlaybook) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *InvestigationPlaybook) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *InvestigationPlaybook) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *InvestigationPlaybook) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *InvestigationPlaybook) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *InvestigationPlaybook) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *InvestigationPlaybook) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *InvestigationPlaybook) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *InvestigationPlaybook) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *InvestigationPlaybook) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *InvestigationPlaybook) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *InvestigationPlaybook) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetQuiet

`func (o *InvestigationPlaybook) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *InvestigationPlaybook) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *InvestigationPlaybook) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *InvestigationPlaybook) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetRecoveryAttempts

`func (o *InvestigationPlaybook) GetRecoveryAttempts() int64`

GetRecoveryAttempts returns the RecoveryAttempts field if non-nil, zero value otherwise.

### GetRecoveryAttemptsOk

`func (o *InvestigationPlaybook) GetRecoveryAttemptsOk() (*int64, bool)`

GetRecoveryAttemptsOk returns a tuple with the RecoveryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAttempts

`func (o *InvestigationPlaybook) SetRecoveryAttempts(v int64)`

SetRecoveryAttempts sets RecoveryAttempts field to given value.

### HasRecoveryAttempts

`func (o *InvestigationPlaybook) HasRecoveryAttempts() bool`

HasRecoveryAttempts returns a boolean if a field has been set.

### GetRoles

`func (o *InvestigationPlaybook) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InvestigationPlaybook) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InvestigationPlaybook) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InvestigationPlaybook) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *InvestigationPlaybook) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *InvestigationPlaybook) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *InvestigationPlaybook) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *InvestigationPlaybook) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetServerId

`func (o *InvestigationPlaybook) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InvestigationPlaybook) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InvestigationPlaybook) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *InvestigationPlaybook) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetSortValues

`func (o *InvestigationPlaybook) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *InvestigationPlaybook) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *InvestigationPlaybook) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *InvestigationPlaybook) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetStartDate

`func (o *InvestigationPlaybook) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvestigationPlaybook) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvestigationPlaybook) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvestigationPlaybook) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartTaskId

`func (o *InvestigationPlaybook) GetStartTaskId() string`

GetStartTaskId returns the StartTaskId field if non-nil, zero value otherwise.

### GetStartTaskIdOk

`func (o *InvestigationPlaybook) GetStartTaskIdOk() (*string, bool)`

GetStartTaskIdOk returns a tuple with the StartTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTaskId

`func (o *InvestigationPlaybook) SetStartTaskId(v string)`

SetStartTaskId sets StartTaskId field to given value.

### HasStartTaskId

`func (o *InvestigationPlaybook) HasStartTaskId() bool`

HasStartTaskId returns a boolean if a field has been set.

### GetState

`func (o *InvestigationPlaybook) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvestigationPlaybook) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvestigationPlaybook) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InvestigationPlaybook) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubPlaybookInputs

`func (o *InvestigationPlaybook) GetSubPlaybookInputs() map[string][]PlaybookInput`

GetSubPlaybookInputs returns the SubPlaybookInputs field if non-nil, zero value otherwise.

### GetSubPlaybookInputsOk

`func (o *InvestigationPlaybook) GetSubPlaybookInputsOk() (*map[string][]PlaybookInput, bool)`

GetSubPlaybookInputsOk returns a tuple with the SubPlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPlaybookInputs

`func (o *InvestigationPlaybook) SetSubPlaybookInputs(v map[string][]PlaybookInput)`

SetSubPlaybookInputs sets SubPlaybookInputs field to given value.

### HasSubPlaybookInputs

`func (o *InvestigationPlaybook) HasSubPlaybookInputs() bool`

HasSubPlaybookInputs returns a boolean if a field has been set.

### GetSubPlaybookOutputs

`func (o *InvestigationPlaybook) GetSubPlaybookOutputs() map[string][]PlaybookOutput`

GetSubPlaybookOutputs returns the SubPlaybookOutputs field if non-nil, zero value otherwise.

### GetSubPlaybookOutputsOk

`func (o *InvestigationPlaybook) GetSubPlaybookOutputsOk() (*map[string][]PlaybookOutput, bool)`

GetSubPlaybookOutputsOk returns a tuple with the SubPlaybookOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPlaybookOutputs

`func (o *InvestigationPlaybook) SetSubPlaybookOutputs(v map[string][]PlaybookOutput)`

SetSubPlaybookOutputs sets SubPlaybookOutputs field to given value.

### HasSubPlaybookOutputs

`func (o *InvestigationPlaybook) HasSubPlaybookOutputs() bool`

HasSubPlaybookOutputs returns a boolean if a field has been set.

### GetTasks

`func (o *InvestigationPlaybook) GetTasks() map[string]InvestigationPlaybookTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InvestigationPlaybook) GetTasksOk() (*map[string]InvestigationPlaybookTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InvestigationPlaybook) SetTasks(v map[string]InvestigationPlaybookTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InvestigationPlaybook) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetVersion

`func (o *InvestigationPlaybook) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InvestigationPlaybook) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InvestigationPlaybook) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InvestigationPlaybook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetView

`func (o *InvestigationPlaybook) GetView() map[string]interface{}`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *InvestigationPlaybook) GetViewOk() (*map[string]interface{}, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *InvestigationPlaybook) SetView(v map[string]interface{})`

SetView sets View field to given value.

### HasView

`func (o *InvestigationPlaybook) HasView() bool`

HasView returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *InvestigationPlaybook) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *InvestigationPlaybook) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *InvestigationPlaybook) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *InvestigationPlaybook) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *InvestigationPlaybook) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *InvestigationPlaybook) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *InvestigationPlaybook) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *InvestigationPlaybook) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *InvestigationPlaybook) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *InvestigationPlaybook) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *InvestigationPlaybook) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *InvestigationPlaybook) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


