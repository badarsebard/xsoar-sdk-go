# Playbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Brands** | Pointer to **[]string** |  | [optional] 
**Commands** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Detached** | Pointer to **bool** |  | [optional] 
**EncTasks** | Pointer to **string** |  | [optional] 
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**PlaybookInputs**](PlaybookInputs.md) |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**MissingScriptsIds** | Pointer to **[]string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameRaw** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Outputs** | Pointer to [**PlaybookOutputs**](PlaybookOutputs.md) |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**ScriptIds** | Pointer to **[]string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**SourcePlaybookID** | Pointer to **string** |  | [optional] 
**StartTaskId** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TaskIds** | Pointer to **[]string** | auto generated field that will contain all task ids in this playbook Needed for searching with bleve | [optional] 
**Tasks** | Pointer to [**map[string]PlaybookTask**](PlaybookTask.md) |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**View** | Pointer to **map[string]interface{}** | PlaybookView represents the view in client of playbook graph | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPlaybook

`func NewPlaybook() *Playbook`

NewPlaybook instantiates a new Playbook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookWithDefaults

`func NewPlaybookWithDefaults() *Playbook`

NewPlaybookWithDefaults instantiates a new Playbook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRead

`func (o *Playbook) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *Playbook) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *Playbook) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *Playbook) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *Playbook) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *Playbook) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *Playbook) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *Playbook) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetBrands

`func (o *Playbook) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *Playbook) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *Playbook) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *Playbook) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCommands

`func (o *Playbook) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *Playbook) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *Playbook) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *Playbook) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetComment

`func (o *Playbook) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Playbook) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Playbook) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Playbook) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommitMessage

`func (o *Playbook) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *Playbook) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *Playbook) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *Playbook) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *Playbook) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *Playbook) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *Playbook) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *Playbook) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDeprecated

`func (o *Playbook) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Playbook) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Playbook) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Playbook) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDetached

`func (o *Playbook) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *Playbook) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *Playbook) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *Playbook) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetEncTasks

`func (o *Playbook) GetEncTasks() string`

GetEncTasks returns the EncTasks field if non-nil, zero value otherwise.

### GetEncTasksOk

`func (o *Playbook) GetEncTasksOk() (*string, bool)`

GetEncTasksOk returns a tuple with the EncTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncTasks

`func (o *Playbook) SetEncTasks(v string)`

SetEncTasks sets EncTasks field to given value.

### HasEncTasks

`func (o *Playbook) HasEncTasks() bool`

HasEncTasks returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *Playbook) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *Playbook) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *Playbook) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *Playbook) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHasRole

`func (o *Playbook) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *Playbook) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *Playbook) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *Playbook) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHidden

`func (o *Playbook) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Playbook) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Playbook) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Playbook) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHighlight

`func (o *Playbook) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Playbook) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Playbook) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Playbook) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Playbook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Playbook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Playbook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Playbook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *Playbook) GetInputs() PlaybookInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *Playbook) GetInputsOk() (*PlaybookInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *Playbook) SetInputs(v PlaybookInputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *Playbook) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetItemVersion

`func (o *Playbook) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *Playbook) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *Playbook) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *Playbook) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLocked

`func (o *Playbook) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Playbook) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Playbook) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Playbook) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMissingScriptsIds

`func (o *Playbook) GetMissingScriptsIds() []string`

GetMissingScriptsIds returns the MissingScriptsIds field if non-nil, zero value otherwise.

### GetMissingScriptsIdsOk

`func (o *Playbook) GetMissingScriptsIdsOk() (*[]string, bool)`

GetMissingScriptsIdsOk returns a tuple with the MissingScriptsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingScriptsIds

`func (o *Playbook) SetMissingScriptsIds(v []string)`

SetMissingScriptsIds sets MissingScriptsIds field to given value.

### HasMissingScriptsIds

`func (o *Playbook) HasMissingScriptsIds() bool`

HasMissingScriptsIds returns a boolean if a field has been set.

### GetModified

`func (o *Playbook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Playbook) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Playbook) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Playbook) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Playbook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Playbook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Playbook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Playbook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameRaw

`func (o *Playbook) GetNameRaw() string`

GetNameRaw returns the NameRaw field if non-nil, zero value otherwise.

### GetNameRawOk

`func (o *Playbook) GetNameRawOk() (*string, bool)`

GetNameRawOk returns a tuple with the NameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRaw

`func (o *Playbook) SetNameRaw(v string)`

SetNameRaw sets NameRaw field to given value.

### HasNameRaw

`func (o *Playbook) HasNameRaw() bool`

HasNameRaw returns a boolean if a field has been set.

### GetNumericId

`func (o *Playbook) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Playbook) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Playbook) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Playbook) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOutputs

`func (o *Playbook) GetOutputs() PlaybookOutputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *Playbook) GetOutputsOk() (*PlaybookOutputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *Playbook) SetOutputs(v PlaybookOutputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *Playbook) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPackID

`func (o *Playbook) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *Playbook) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *Playbook) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *Playbook) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *Playbook) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *Playbook) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *Playbook) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *Playbook) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPrevName

`func (o *Playbook) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *Playbook) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *Playbook) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *Playbook) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *Playbook) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *Playbook) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *Playbook) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *Playbook) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *Playbook) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *Playbook) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *Playbook) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *Playbook) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *Playbook) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *Playbook) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *Playbook) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *Playbook) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Playbook) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Playbook) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Playbook) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Playbook) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPrivate

`func (o *Playbook) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Playbook) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Playbook) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Playbook) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *Playbook) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *Playbook) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *Playbook) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *Playbook) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetQuiet

`func (o *Playbook) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *Playbook) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *Playbook) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *Playbook) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetRoles

`func (o *Playbook) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Playbook) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Playbook) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Playbook) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScriptIds

`func (o *Playbook) GetScriptIds() []string`

GetScriptIds returns the ScriptIds field if non-nil, zero value otherwise.

### GetScriptIdsOk

`func (o *Playbook) GetScriptIdsOk() (*[]string, bool)`

GetScriptIdsOk returns a tuple with the ScriptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptIds

`func (o *Playbook) SetScriptIds(v []string)`

SetScriptIds sets ScriptIds field to given value.

### HasScriptIds

`func (o *Playbook) HasScriptIds() bool`

HasScriptIds returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Playbook) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Playbook) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Playbook) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Playbook) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *Playbook) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *Playbook) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *Playbook) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *Playbook) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *Playbook) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Playbook) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Playbook) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Playbook) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourcePlaybookID

`func (o *Playbook) GetSourcePlaybookID() string`

GetSourcePlaybookID returns the SourcePlaybookID field if non-nil, zero value otherwise.

### GetSourcePlaybookIDOk

`func (o *Playbook) GetSourcePlaybookIDOk() (*string, bool)`

GetSourcePlaybookIDOk returns a tuple with the SourcePlaybookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePlaybookID

`func (o *Playbook) SetSourcePlaybookID(v string)`

SetSourcePlaybookID sets SourcePlaybookID field to given value.

### HasSourcePlaybookID

`func (o *Playbook) HasSourcePlaybookID() bool`

HasSourcePlaybookID returns a boolean if a field has been set.

### GetStartTaskId

`func (o *Playbook) GetStartTaskId() string`

GetStartTaskId returns the StartTaskId field if non-nil, zero value otherwise.

### GetStartTaskIdOk

`func (o *Playbook) GetStartTaskIdOk() (*string, bool)`

GetStartTaskIdOk returns a tuple with the StartTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTaskId

`func (o *Playbook) SetStartTaskId(v string)`

SetStartTaskId sets StartTaskId field to given value.

### HasStartTaskId

`func (o *Playbook) HasStartTaskId() bool`

HasStartTaskId returns a boolean if a field has been set.

### GetSystem

`func (o *Playbook) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Playbook) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Playbook) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Playbook) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *Playbook) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Playbook) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Playbook) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Playbook) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskIds

`func (o *Playbook) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *Playbook) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *Playbook) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *Playbook) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetTasks

`func (o *Playbook) GetTasks() map[string]PlaybookTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Playbook) GetTasksOk() (*map[string]PlaybookTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Playbook) SetTasks(v map[string]PlaybookTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Playbook) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetToServerVersion

`func (o *Playbook) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *Playbook) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *Playbook) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *Playbook) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *Playbook) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *Playbook) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *Playbook) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *Playbook) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *Playbook) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *Playbook) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *Playbook) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *Playbook) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *Playbook) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Playbook) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Playbook) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Playbook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetView

`func (o *Playbook) GetView() map[string]interface{}`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Playbook) GetViewOk() (*map[string]interface{}, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Playbook) SetView(v map[string]interface{})`

SetView sets View field to given value.

### HasView

`func (o *Playbook) HasView() bool`

HasView returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *Playbook) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *Playbook) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *Playbook) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *Playbook) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *Playbook) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *Playbook) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *Playbook) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *Playbook) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *Playbook) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *Playbook) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *Playbook) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *Playbook) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


