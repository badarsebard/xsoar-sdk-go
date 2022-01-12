# AutomationScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Arguments** | Pointer to [**[]Argument**](Argument.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**ContextKeys** | Pointer to **[]string** |  | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**DependsOn** | Pointer to **map[string][]string** | This fields indicates which commands this script depends on | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Detached** | Pointer to **bool** |  | [optional] 
**DockerImage** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Engine** | Pointer to **string** | Engine that will run the script | [optional] 
**EngineGroup** | Pointer to **string** | EngineGroup that will run the script | [optional] 
**FromServerVersion** | Pointer to **string** |  | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Important** | Pointer to [**[]Important**](Important.md) |  | [optional] 
**ItemVersion** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Outputs** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**Polling** | Pointer to **bool** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Pswd** | Pointer to **string** |  | [optional] 
**RawTags** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**RunAs** | Pointer to **string** |  | [optional] 
**RunOnce** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**ScriptTarget** | Pointer to **int64** | ScriptTarget represents the module where this script should run | [optional] 
**SearchableName** | Pointer to **string** |  | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**SourceScripID** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** | ScriptSubType holds the script type version | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Timeout** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**ToServerVersion** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | ScriptType holds the type of a script | [optional] 
**User** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**VisualScript** | Pointer to **string** |  | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutomationScript

`func NewAutomationScript() *AutomationScript`

NewAutomationScript instantiates a new AutomationScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationScriptWithDefaults

`func NewAutomationScriptWithDefaults() *AutomationScript`

NewAutomationScriptWithDefaults instantiates a new AutomationScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRead

`func (o *AutomationScript) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *AutomationScript) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *AutomationScript) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *AutomationScript) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *AutomationScript) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *AutomationScript) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *AutomationScript) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *AutomationScript) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetArguments

`func (o *AutomationScript) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *AutomationScript) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *AutomationScript) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *AutomationScript) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetComment

`func (o *AutomationScript) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AutomationScript) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AutomationScript) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AutomationScript) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommitMessage

`func (o *AutomationScript) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *AutomationScript) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *AutomationScript) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *AutomationScript) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetContextKeys

`func (o *AutomationScript) GetContextKeys() []string`

GetContextKeys returns the ContextKeys field if non-nil, zero value otherwise.

### GetContextKeysOk

`func (o *AutomationScript) GetContextKeysOk() (*[]string, bool)`

GetContextKeysOk returns a tuple with the ContextKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKeys

`func (o *AutomationScript) SetContextKeys(v []string)`

SetContextKeys sets ContextKeys field to given value.

### HasContextKeys

`func (o *AutomationScript) HasContextKeys() bool`

HasContextKeys returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *AutomationScript) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *AutomationScript) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *AutomationScript) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *AutomationScript) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDependsOn

`func (o *AutomationScript) GetDependsOn() map[string][]string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *AutomationScript) GetDependsOnOk() (*map[string][]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *AutomationScript) SetDependsOn(v map[string][]string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *AutomationScript) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetDeprecated

`func (o *AutomationScript) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AutomationScript) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AutomationScript) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AutomationScript) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDetached

`func (o *AutomationScript) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *AutomationScript) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *AutomationScript) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *AutomationScript) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetDockerImage

`func (o *AutomationScript) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *AutomationScript) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *AutomationScript) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *AutomationScript) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetEnabled

`func (o *AutomationScript) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutomationScript) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutomationScript) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutomationScript) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngine

`func (o *AutomationScript) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *AutomationScript) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *AutomationScript) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *AutomationScript) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetEngineGroup

`func (o *AutomationScript) GetEngineGroup() string`

GetEngineGroup returns the EngineGroup field if non-nil, zero value otherwise.

### GetEngineGroupOk

`func (o *AutomationScript) GetEngineGroupOk() (*string, bool)`

GetEngineGroupOk returns a tuple with the EngineGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineGroup

`func (o *AutomationScript) SetEngineGroup(v string)`

SetEngineGroup sets EngineGroup field to given value.

### HasEngineGroup

`func (o *AutomationScript) HasEngineGroup() bool`

HasEngineGroup returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *AutomationScript) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *AutomationScript) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *AutomationScript) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *AutomationScript) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHasRole

`func (o *AutomationScript) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *AutomationScript) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *AutomationScript) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *AutomationScript) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHidden

`func (o *AutomationScript) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AutomationScript) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AutomationScript) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *AutomationScript) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHighlight

`func (o *AutomationScript) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *AutomationScript) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *AutomationScript) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *AutomationScript) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *AutomationScript) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationScript) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationScript) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationScript) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImportant

`func (o *AutomationScript) GetImportant() []Important`

GetImportant returns the Important field if non-nil, zero value otherwise.

### GetImportantOk

`func (o *AutomationScript) GetImportantOk() (*[]Important, bool)`

GetImportantOk returns a tuple with the Important field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportant

`func (o *AutomationScript) SetImportant(v []Important)`

SetImportant sets Important field to given value.

### HasImportant

`func (o *AutomationScript) HasImportant() bool`

HasImportant returns a boolean if a field has been set.

### GetItemVersion

`func (o *AutomationScript) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *AutomationScript) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *AutomationScript) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *AutomationScript) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLocked

`func (o *AutomationScript) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AutomationScript) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AutomationScript) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AutomationScript) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *AutomationScript) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AutomationScript) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AutomationScript) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AutomationScript) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *AutomationScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationScript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationScript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *AutomationScript) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *AutomationScript) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *AutomationScript) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *AutomationScript) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOutputs

`func (o *AutomationScript) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AutomationScript) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AutomationScript) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AutomationScript) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPackID

`func (o *AutomationScript) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *AutomationScript) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *AutomationScript) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *AutomationScript) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *AutomationScript) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *AutomationScript) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *AutomationScript) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *AutomationScript) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPolling

`func (o *AutomationScript) GetPolling() bool`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *AutomationScript) GetPollingOk() (*bool, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *AutomationScript) SetPolling(v bool)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *AutomationScript) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetPrevName

`func (o *AutomationScript) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *AutomationScript) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *AutomationScript) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *AutomationScript) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *AutomationScript) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *AutomationScript) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *AutomationScript) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *AutomationScript) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *AutomationScript) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *AutomationScript) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *AutomationScript) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *AutomationScript) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *AutomationScript) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *AutomationScript) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *AutomationScript) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *AutomationScript) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *AutomationScript) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *AutomationScript) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *AutomationScript) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *AutomationScript) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPrivate

`func (o *AutomationScript) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AutomationScript) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AutomationScript) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AutomationScript) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *AutomationScript) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *AutomationScript) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *AutomationScript) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *AutomationScript) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetPswd

`func (o *AutomationScript) GetPswd() string`

GetPswd returns the Pswd field if non-nil, zero value otherwise.

### GetPswdOk

`func (o *AutomationScript) GetPswdOk() (*string, bool)`

GetPswdOk returns a tuple with the Pswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPswd

`func (o *AutomationScript) SetPswd(v string)`

SetPswd sets Pswd field to given value.

### HasPswd

`func (o *AutomationScript) HasPswd() bool`

HasPswd returns a boolean if a field has been set.

### GetRawTags

`func (o *AutomationScript) GetRawTags() []string`

GetRawTags returns the RawTags field if non-nil, zero value otherwise.

### GetRawTagsOk

`func (o *AutomationScript) GetRawTagsOk() (*[]string, bool)`

GetRawTagsOk returns a tuple with the RawTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTags

`func (o *AutomationScript) SetRawTags(v []string)`

SetRawTags sets RawTags field to given value.

### HasRawTags

`func (o *AutomationScript) HasRawTags() bool`

HasRawTags returns a boolean if a field has been set.

### GetRoles

`func (o *AutomationScript) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AutomationScript) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AutomationScript) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AutomationScript) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunAs

`func (o *AutomationScript) GetRunAs() string`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *AutomationScript) GetRunAsOk() (*string, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *AutomationScript) SetRunAs(v string)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *AutomationScript) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetRunOnce

`func (o *AutomationScript) GetRunOnce() bool`

GetRunOnce returns the RunOnce field if non-nil, zero value otherwise.

### GetRunOnceOk

`func (o *AutomationScript) GetRunOnceOk() (*bool, bool)`

GetRunOnceOk returns a tuple with the RunOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnce

`func (o *AutomationScript) SetRunOnce(v bool)`

SetRunOnce sets RunOnce field to given value.

### HasRunOnce

`func (o *AutomationScript) HasRunOnce() bool`

HasRunOnce returns a boolean if a field has been set.

### GetScript

`func (o *AutomationScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *AutomationScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *AutomationScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *AutomationScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetScriptTarget

`func (o *AutomationScript) GetScriptTarget() int64`

GetScriptTarget returns the ScriptTarget field if non-nil, zero value otherwise.

### GetScriptTargetOk

`func (o *AutomationScript) GetScriptTargetOk() (*int64, bool)`

GetScriptTargetOk returns a tuple with the ScriptTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptTarget

`func (o *AutomationScript) SetScriptTarget(v int64)`

SetScriptTarget sets ScriptTarget field to given value.

### HasScriptTarget

`func (o *AutomationScript) HasScriptTarget() bool`

HasScriptTarget returns a boolean if a field has been set.

### GetSearchableName

`func (o *AutomationScript) GetSearchableName() string`

GetSearchableName returns the SearchableName field if non-nil, zero value otherwise.

### GetSearchableNameOk

`func (o *AutomationScript) GetSearchableNameOk() (*string, bool)`

GetSearchableNameOk returns a tuple with the SearchableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableName

`func (o *AutomationScript) SetSearchableName(v string)`

SetSearchableName sets SearchableName field to given value.

### HasSearchableName

`func (o *AutomationScript) HasSearchableName() bool`

HasSearchableName returns a boolean if a field has been set.

### GetSensitive

`func (o *AutomationScript) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *AutomationScript) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *AutomationScript) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *AutomationScript) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *AutomationScript) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *AutomationScript) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *AutomationScript) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *AutomationScript) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *AutomationScript) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *AutomationScript) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *AutomationScript) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *AutomationScript) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *AutomationScript) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *AutomationScript) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *AutomationScript) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *AutomationScript) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceScripID

`func (o *AutomationScript) GetSourceScripID() string`

GetSourceScripID returns the SourceScripID field if non-nil, zero value otherwise.

### GetSourceScripIDOk

`func (o *AutomationScript) GetSourceScripIDOk() (*string, bool)`

GetSourceScripIDOk returns a tuple with the SourceScripID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceScripID

`func (o *AutomationScript) SetSourceScripID(v string)`

SetSourceScripID sets SourceScripID field to given value.

### HasSourceScripID

`func (o *AutomationScript) HasSourceScripID() bool`

HasSourceScripID returns a boolean if a field has been set.

### GetSubtype

`func (o *AutomationScript) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AutomationScript) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AutomationScript) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *AutomationScript) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSystem

`func (o *AutomationScript) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AutomationScript) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AutomationScript) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AutomationScript) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *AutomationScript) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutomationScript) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutomationScript) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutomationScript) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *AutomationScript) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AutomationScript) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AutomationScript) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AutomationScript) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToServerVersion

`func (o *AutomationScript) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *AutomationScript) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *AutomationScript) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *AutomationScript) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetType

`func (o *AutomationScript) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomationScript) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomationScript) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutomationScript) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *AutomationScript) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AutomationScript) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AutomationScript) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AutomationScript) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *AutomationScript) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *AutomationScript) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *AutomationScript) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *AutomationScript) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *AutomationScript) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *AutomationScript) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *AutomationScript) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *AutomationScript) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *AutomationScript) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutomationScript) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutomationScript) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutomationScript) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVisualScript

`func (o *AutomationScript) GetVisualScript() string`

GetVisualScript returns the VisualScript field if non-nil, zero value otherwise.

### GetVisualScriptOk

`func (o *AutomationScript) GetVisualScriptOk() (*string, bool)`

GetVisualScriptOk returns a tuple with the VisualScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualScript

`func (o *AutomationScript) SetVisualScript(v string)`

SetVisualScript sets VisualScript field to given value.

### HasVisualScript

`func (o *AutomationScript) HasVisualScript() bool`

HasVisualScript returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *AutomationScript) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *AutomationScript) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *AutomationScript) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *AutomationScript) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *AutomationScript) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *AutomationScript) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *AutomationScript) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *AutomationScript) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *AutomationScript) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *AutomationScript) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *AutomationScript) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *AutomationScript) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


