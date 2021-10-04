# IncidentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedToAll** | Pointer to **bool** |  | [optional] 
**AssociatedTypes** | Pointer to **[]string** | AssociatedTypes - list of incident (case) types IDs related to specific incident field | [optional] 
**BreachScript** | Pointer to **string** |  | [optional] 
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**CliName** | Pointer to **string** |  | [optional] 
**CloseForm** | Pointer to **bool** |  | [optional] 
**Columns** | Pointer to [**[]GridColumn**](GridColumn.md) |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **bool** |  | [optional] 
**DefaultRows** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EditForm** | Pointer to **bool** |  | [optional] 
**FieldCalcScript** | Pointer to **string** |  | [optional] 
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Group** | Pointer to **float64** | FieldGroup is the field group | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**MergeStrategy** | Pointer to **string** | FieldMergeStrategy strategy for merging of indicator fields | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeverSetAsRequired** | Pointer to **bool** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**OwnerOnly** | Pointer to **bool** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**RunScriptAfterUpdate** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**SelectValues** | Pointer to **[]string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**Sla** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**SystemAssociatedTypes** | Pointer to **[]string** |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unmapped** | Pointer to **bool** |  | [optional] 
**Unsearchable** | Pointer to **bool** |  | [optional] 
**UseAsKpi** | Pointer to **bool** |  | [optional] 
**ValidatedError** | Pointer to **string** |  | [optional] 
**ValidationRegex** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewIncidentField

`func NewIncidentField() *IncidentField`

NewIncidentField instantiates a new IncidentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFieldWithDefaults

`func NewIncidentFieldWithDefaults() *IncidentField`

NewIncidentFieldWithDefaults instantiates a new IncidentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedToAll

`func (o *IncidentField) GetAssociatedToAll() bool`

GetAssociatedToAll returns the AssociatedToAll field if non-nil, zero value otherwise.

### GetAssociatedToAllOk

`func (o *IncidentField) GetAssociatedToAllOk() (*bool, bool)`

GetAssociatedToAllOk returns a tuple with the AssociatedToAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedToAll

`func (o *IncidentField) SetAssociatedToAll(v bool)`

SetAssociatedToAll sets AssociatedToAll field to given value.

### HasAssociatedToAll

`func (o *IncidentField) HasAssociatedToAll() bool`

HasAssociatedToAll returns a boolean if a field has been set.

### GetAssociatedTypes

`func (o *IncidentField) GetAssociatedTypes() []string`

GetAssociatedTypes returns the AssociatedTypes field if non-nil, zero value otherwise.

### GetAssociatedTypesOk

`func (o *IncidentField) GetAssociatedTypesOk() (*[]string, bool)`

GetAssociatedTypesOk returns a tuple with the AssociatedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedTypes

`func (o *IncidentField) SetAssociatedTypes(v []string)`

SetAssociatedTypes sets AssociatedTypes field to given value.

### HasAssociatedTypes

`func (o *IncidentField) HasAssociatedTypes() bool`

HasAssociatedTypes returns a boolean if a field has been set.

### GetBreachScript

`func (o *IncidentField) GetBreachScript() string`

GetBreachScript returns the BreachScript field if non-nil, zero value otherwise.

### GetBreachScriptOk

`func (o *IncidentField) GetBreachScriptOk() (*string, bool)`

GetBreachScriptOk returns a tuple with the BreachScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachScript

`func (o *IncidentField) SetBreachScript(v string)`

SetBreachScript sets BreachScript field to given value.

### HasBreachScript

`func (o *IncidentField) HasBreachScript() bool`

HasBreachScript returns a boolean if a field has been set.

### GetCaseInsensitive

`func (o *IncidentField) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *IncidentField) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *IncidentField) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *IncidentField) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetCliName

`func (o *IncidentField) GetCliName() string`

GetCliName returns the CliName field if non-nil, zero value otherwise.

### GetCliNameOk

`func (o *IncidentField) GetCliNameOk() (*string, bool)`

GetCliNameOk returns a tuple with the CliName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliName

`func (o *IncidentField) SetCliName(v string)`

SetCliName sets CliName field to given value.

### HasCliName

`func (o *IncidentField) HasCliName() bool`

HasCliName returns a boolean if a field has been set.

### GetCloseForm

`func (o *IncidentField) GetCloseForm() bool`

GetCloseForm returns the CloseForm field if non-nil, zero value otherwise.

### GetCloseFormOk

`func (o *IncidentField) GetCloseFormOk() (*bool, bool)`

GetCloseFormOk returns a tuple with the CloseForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseForm

`func (o *IncidentField) SetCloseForm(v bool)`

SetCloseForm sets CloseForm field to given value.

### HasCloseForm

`func (o *IncidentField) HasCloseForm() bool`

HasCloseForm returns a boolean if a field has been set.

### GetColumns

`func (o *IncidentField) GetColumns() []GridColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *IncidentField) GetColumnsOk() (*[]GridColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *IncidentField) SetColumns(v []GridColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *IncidentField) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetCommitMessage

`func (o *IncidentField) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *IncidentField) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *IncidentField) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *IncidentField) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetContent

`func (o *IncidentField) GetContent() bool`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentField) GetContentOk() (*bool, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentField) SetContent(v bool)`

SetContent sets Content field to given value.

### HasContent

`func (o *IncidentField) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDefaultRows

`func (o *IncidentField) GetDefaultRows() []map[string]map[string]interface{}`

GetDefaultRows returns the DefaultRows field if non-nil, zero value otherwise.

### GetDefaultRowsOk

`func (o *IncidentField) GetDefaultRowsOk() (*[]map[string]map[string]interface{}, bool)`

GetDefaultRowsOk returns a tuple with the DefaultRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRows

`func (o *IncidentField) SetDefaultRows(v []map[string]map[string]interface{})`

SetDefaultRows sets DefaultRows field to given value.

### HasDefaultRows

`func (o *IncidentField) HasDefaultRows() bool`

HasDefaultRows returns a boolean if a field has been set.

### GetDescription

`func (o *IncidentField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IncidentField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IncidentField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IncidentField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditForm

`func (o *IncidentField) GetEditForm() bool`

GetEditForm returns the EditForm field if non-nil, zero value otherwise.

### GetEditFormOk

`func (o *IncidentField) GetEditFormOk() (*bool, bool)`

GetEditFormOk returns a tuple with the EditForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditForm

`func (o *IncidentField) SetEditForm(v bool)`

SetEditForm sets EditForm field to given value.

### HasEditForm

`func (o *IncidentField) HasEditForm() bool`

HasEditForm returns a boolean if a field has been set.

### GetFieldCalcScript

`func (o *IncidentField) GetFieldCalcScript() string`

GetFieldCalcScript returns the FieldCalcScript field if non-nil, zero value otherwise.

### GetFieldCalcScriptOk

`func (o *IncidentField) GetFieldCalcScriptOk() (*string, bool)`

GetFieldCalcScriptOk returns a tuple with the FieldCalcScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCalcScript

`func (o *IncidentField) SetFieldCalcScript(v string)`

SetFieldCalcScript sets FieldCalcScript field to given value.

### HasFieldCalcScript

`func (o *IncidentField) HasFieldCalcScript() bool`

HasFieldCalcScript returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *IncidentField) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *IncidentField) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *IncidentField) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *IncidentField) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetGroup

`func (o *IncidentField) GetGroup() float64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IncidentField) GetGroupOk() (*float64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IncidentField) SetGroup(v float64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IncidentField) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHidden

`func (o *IncidentField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *IncidentField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *IncidentField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *IncidentField) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHighlight

`func (o *IncidentField) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *IncidentField) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *IncidentField) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *IncidentField) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *IncidentField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *IncidentField) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *IncidentField) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *IncidentField) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *IncidentField) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetItemVersion

`func (o *IncidentField) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *IncidentField) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *IncidentField) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *IncidentField) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLocked

`func (o *IncidentField) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IncidentField) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IncidentField) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IncidentField) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMergeStrategy

`func (o *IncidentField) GetMergeStrategy() string`

GetMergeStrategy returns the MergeStrategy field if non-nil, zero value otherwise.

### GetMergeStrategyOk

`func (o *IncidentField) GetMergeStrategyOk() (*string, bool)`

GetMergeStrategyOk returns a tuple with the MergeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStrategy

`func (o *IncidentField) SetMergeStrategy(v string)`

SetMergeStrategy sets MergeStrategy field to given value.

### HasMergeStrategy

`func (o *IncidentField) HasMergeStrategy() bool`

HasMergeStrategy returns a boolean if a field has been set.

### GetModified

`func (o *IncidentField) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentField) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentField) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentField) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeverSetAsRequired

`func (o *IncidentField) GetNeverSetAsRequired() bool`

GetNeverSetAsRequired returns the NeverSetAsRequired field if non-nil, zero value otherwise.

### GetNeverSetAsRequiredOk

`func (o *IncidentField) GetNeverSetAsRequiredOk() (*bool, bool)`

GetNeverSetAsRequiredOk returns a tuple with the NeverSetAsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverSetAsRequired

`func (o *IncidentField) SetNeverSetAsRequired(v bool)`

SetNeverSetAsRequired sets NeverSetAsRequired field to given value.

### HasNeverSetAsRequired

`func (o *IncidentField) HasNeverSetAsRequired() bool`

HasNeverSetAsRequired returns a boolean if a field has been set.

### GetNumericId

`func (o *IncidentField) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *IncidentField) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *IncidentField) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *IncidentField) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOwnerOnly

`func (o *IncidentField) GetOwnerOnly() bool`

GetOwnerOnly returns the OwnerOnly field if non-nil, zero value otherwise.

### GetOwnerOnlyOk

`func (o *IncidentField) GetOwnerOnlyOk() (*bool, bool)`

GetOwnerOnlyOk returns a tuple with the OwnerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOnly

`func (o *IncidentField) SetOwnerOnly(v bool)`

SetOwnerOnly sets OwnerOnly field to given value.

### HasOwnerOnly

`func (o *IncidentField) HasOwnerOnly() bool`

HasOwnerOnly returns a boolean if a field has been set.

### GetPackID

`func (o *IncidentField) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *IncidentField) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *IncidentField) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *IncidentField) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *IncidentField) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *IncidentField) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *IncidentField) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *IncidentField) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPlaceholder

`func (o *IncidentField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *IncidentField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *IncidentField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *IncidentField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetPrevName

`func (o *IncidentField) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *IncidentField) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *IncidentField) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *IncidentField) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *IncidentField) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *IncidentField) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *IncidentField) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *IncidentField) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *IncidentField) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *IncidentField) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *IncidentField) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *IncidentField) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetRequired

`func (o *IncidentField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IncidentField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IncidentField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IncidentField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRunScriptAfterUpdate

`func (o *IncidentField) GetRunScriptAfterUpdate() bool`

GetRunScriptAfterUpdate returns the RunScriptAfterUpdate field if non-nil, zero value otherwise.

### GetRunScriptAfterUpdateOk

`func (o *IncidentField) GetRunScriptAfterUpdateOk() (*bool, bool)`

GetRunScriptAfterUpdateOk returns a tuple with the RunScriptAfterUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScriptAfterUpdate

`func (o *IncidentField) SetRunScriptAfterUpdate(v bool)`

SetRunScriptAfterUpdate sets RunScriptAfterUpdate field to given value.

### HasRunScriptAfterUpdate

`func (o *IncidentField) HasRunScriptAfterUpdate() bool`

HasRunScriptAfterUpdate returns a boolean if a field has been set.

### GetScript

`func (o *IncidentField) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *IncidentField) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *IncidentField) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *IncidentField) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSelectValues

`func (o *IncidentField) GetSelectValues() []string`

GetSelectValues returns the SelectValues field if non-nil, zero value otherwise.

### GetSelectValuesOk

`func (o *IncidentField) GetSelectValuesOk() (*[]string, bool)`

GetSelectValuesOk returns a tuple with the SelectValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectValues

`func (o *IncidentField) SetSelectValues(v []string)`

SetSelectValues sets SelectValues field to given value.

### HasSelectValues

`func (o *IncidentField) HasSelectValues() bool`

HasSelectValues returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *IncidentField) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *IncidentField) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *IncidentField) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *IncidentField) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *IncidentField) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *IncidentField) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *IncidentField) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *IncidentField) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSla

`func (o *IncidentField) GetSla() int64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *IncidentField) GetSlaOk() (*int64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *IncidentField) SetSla(v int64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *IncidentField) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSortValues

`func (o *IncidentField) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *IncidentField) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *IncidentField) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *IncidentField) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSystem

`func (o *IncidentField) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IncidentField) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IncidentField) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IncidentField) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSystemAssociatedTypes

`func (o *IncidentField) GetSystemAssociatedTypes() []string`

GetSystemAssociatedTypes returns the SystemAssociatedTypes field if non-nil, zero value otherwise.

### GetSystemAssociatedTypesOk

`func (o *IncidentField) GetSystemAssociatedTypesOk() (*[]string, bool)`

GetSystemAssociatedTypesOk returns a tuple with the SystemAssociatedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAssociatedTypes

`func (o *IncidentField) SetSystemAssociatedTypes(v []string)`

SetSystemAssociatedTypes sets SystemAssociatedTypes field to given value.

### HasSystemAssociatedTypes

`func (o *IncidentField) HasSystemAssociatedTypes() bool`

HasSystemAssociatedTypes returns a boolean if a field has been set.

### GetThreshold

`func (o *IncidentField) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IncidentField) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IncidentField) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IncidentField) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetToServerVersion

`func (o *IncidentField) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *IncidentField) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *IncidentField) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *IncidentField) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetType

`func (o *IncidentField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnmapped

`func (o *IncidentField) GetUnmapped() bool`

GetUnmapped returns the Unmapped field if non-nil, zero value otherwise.

### GetUnmappedOk

`func (o *IncidentField) GetUnmappedOk() (*bool, bool)`

GetUnmappedOk returns a tuple with the Unmapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapped

`func (o *IncidentField) SetUnmapped(v bool)`

SetUnmapped sets Unmapped field to given value.

### HasUnmapped

`func (o *IncidentField) HasUnmapped() bool`

HasUnmapped returns a boolean if a field has been set.

### GetUnsearchable

`func (o *IncidentField) GetUnsearchable() bool`

GetUnsearchable returns the Unsearchable field if non-nil, zero value otherwise.

### GetUnsearchableOk

`func (o *IncidentField) GetUnsearchableOk() (*bool, bool)`

GetUnsearchableOk returns a tuple with the Unsearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsearchable

`func (o *IncidentField) SetUnsearchable(v bool)`

SetUnsearchable sets Unsearchable field to given value.

### HasUnsearchable

`func (o *IncidentField) HasUnsearchable() bool`

HasUnsearchable returns a boolean if a field has been set.

### GetUseAsKpi

`func (o *IncidentField) GetUseAsKpi() bool`

GetUseAsKpi returns the UseAsKpi field if non-nil, zero value otherwise.

### GetUseAsKpiOk

`func (o *IncidentField) GetUseAsKpiOk() (*bool, bool)`

GetUseAsKpiOk returns a tuple with the UseAsKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsKpi

`func (o *IncidentField) SetUseAsKpi(v bool)`

SetUseAsKpi sets UseAsKpi field to given value.

### HasUseAsKpi

`func (o *IncidentField) HasUseAsKpi() bool`

HasUseAsKpi returns a boolean if a field has been set.

### GetValidatedError

`func (o *IncidentField) GetValidatedError() string`

GetValidatedError returns the ValidatedError field if non-nil, zero value otherwise.

### GetValidatedErrorOk

`func (o *IncidentField) GetValidatedErrorOk() (*string, bool)`

GetValidatedErrorOk returns a tuple with the ValidatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedError

`func (o *IncidentField) SetValidatedError(v string)`

SetValidatedError sets ValidatedError field to given value.

### HasValidatedError

`func (o *IncidentField) HasValidatedError() bool`

HasValidatedError returns a boolean if a field has been set.

### GetValidationRegex

`func (o *IncidentField) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *IncidentField) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *IncidentField) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *IncidentField) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *IncidentField) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *IncidentField) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *IncidentField) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *IncidentField) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *IncidentField) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *IncidentField) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *IncidentField) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *IncidentField) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *IncidentField) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IncidentField) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IncidentField) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IncidentField) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


