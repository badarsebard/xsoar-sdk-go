# IncidentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autorun** | Pointer to **bool** |  | [optional] 
**ClosureScript** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **int64** |  | [optional] 
**DaysR** | Pointer to **int64** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Detached** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**ExtractSettings** | Pointer to [**IncidentTypeExtractSettings**](IncidentTypeExtractSettings.md) |  | [optional] 
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Hours** | Pointer to **int64** |  | [optional] 
**HoursR** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Layout** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**OnChangeRepAlg** | Pointer to **float64** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PlaybookId** | Pointer to **string** |  | [optional] 
**PreProcessingScript** | Pointer to **string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**ReputationCalc** | Pointer to **float64** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**Sla** | Pointer to **int64** |  | [optional] 
**SlaReminder** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Weeks** | Pointer to **int64** |  | [optional] 
**WeeksR** | Pointer to **int64** |  | [optional] 

## Methods

### NewIncidentType

`func NewIncidentType() *IncidentType`

NewIncidentType instantiates a new IncidentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTypeWithDefaults

`func NewIncidentTypeWithDefaults() *IncidentType`

NewIncidentTypeWithDefaults instantiates a new IncidentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutorun

`func (o *IncidentType) GetAutorun() bool`

GetAutorun returns the Autorun field if non-nil, zero value otherwise.

### GetAutorunOk

`func (o *IncidentType) GetAutorunOk() (*bool, bool)`

GetAutorunOk returns a tuple with the Autorun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorun

`func (o *IncidentType) SetAutorun(v bool)`

SetAutorun sets Autorun field to given value.

### HasAutorun

`func (o *IncidentType) HasAutorun() bool`

HasAutorun returns a boolean if a field has been set.

### GetClosureScript

`func (o *IncidentType) GetClosureScript() string`

GetClosureScript returns the ClosureScript field if non-nil, zero value otherwise.

### GetClosureScriptOk

`func (o *IncidentType) GetClosureScriptOk() (*string, bool)`

GetClosureScriptOk returns a tuple with the ClosureScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosureScript

`func (o *IncidentType) SetClosureScript(v string)`

SetClosureScript sets ClosureScript field to given value.

### HasClosureScript

`func (o *IncidentType) HasClosureScript() bool`

HasClosureScript returns a boolean if a field has been set.

### GetColor

`func (o *IncidentType) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *IncidentType) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *IncidentType) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *IncidentType) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCommitMessage

`func (o *IncidentType) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *IncidentType) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *IncidentType) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *IncidentType) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetDays

`func (o *IncidentType) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *IncidentType) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *IncidentType) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *IncidentType) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDaysR

`func (o *IncidentType) GetDaysR() int64`

GetDaysR returns the DaysR field if non-nil, zero value otherwise.

### GetDaysROk

`func (o *IncidentType) GetDaysROk() (*int64, bool)`

GetDaysROk returns a tuple with the DaysR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysR

`func (o *IncidentType) SetDaysR(v int64)`

SetDaysR sets DaysR field to given value.

### HasDaysR

`func (o *IncidentType) HasDaysR() bool`

HasDaysR returns a boolean if a field has been set.

### GetDefault

`func (o *IncidentType) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IncidentType) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IncidentType) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IncidentType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDetached

`func (o *IncidentType) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *IncidentType) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *IncidentType) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *IncidentType) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetDisabled

`func (o *IncidentType) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *IncidentType) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *IncidentType) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *IncidentType) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExtractSettings

`func (o *IncidentType) GetExtractSettings() IncidentTypeExtractSettings`

GetExtractSettings returns the ExtractSettings field if non-nil, zero value otherwise.

### GetExtractSettingsOk

`func (o *IncidentType) GetExtractSettingsOk() (*IncidentTypeExtractSettings, bool)`

GetExtractSettingsOk returns a tuple with the ExtractSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractSettings

`func (o *IncidentType) SetExtractSettings(v IncidentTypeExtractSettings)`

SetExtractSettings sets ExtractSettings field to given value.

### HasExtractSettings

`func (o *IncidentType) HasExtractSettings() bool`

HasExtractSettings returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *IncidentType) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *IncidentType) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *IncidentType) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *IncidentType) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *IncidentType) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *IncidentType) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *IncidentType) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *IncidentType) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetHours

`func (o *IncidentType) GetHours() int64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *IncidentType) GetHoursOk() (*int64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *IncidentType) SetHours(v int64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *IncidentType) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetHoursR

`func (o *IncidentType) GetHoursR() int64`

GetHoursR returns the HoursR field if non-nil, zero value otherwise.

### GetHoursROk

`func (o *IncidentType) GetHoursROk() (*int64, bool)`

GetHoursROk returns a tuple with the HoursR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursR

`func (o *IncidentType) SetHoursR(v int64)`

SetHoursR sets HoursR field to given value.

### HasHoursR

`func (o *IncidentType) HasHoursR() bool`

HasHoursR returns a boolean if a field has been set.

### GetId

`func (o *IncidentType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemVersion

`func (o *IncidentType) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *IncidentType) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *IncidentType) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *IncidentType) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLayout

`func (o *IncidentType) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *IncidentType) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *IncidentType) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *IncidentType) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetLocked

`func (o *IncidentType) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IncidentType) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IncidentType) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IncidentType) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *IncidentType) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentType) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentType) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentType) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *IncidentType) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *IncidentType) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *IncidentType) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *IncidentType) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOnChangeRepAlg

`func (o *IncidentType) GetOnChangeRepAlg() float64`

GetOnChangeRepAlg returns the OnChangeRepAlg field if non-nil, zero value otherwise.

### GetOnChangeRepAlgOk

`func (o *IncidentType) GetOnChangeRepAlgOk() (*float64, bool)`

GetOnChangeRepAlgOk returns a tuple with the OnChangeRepAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnChangeRepAlg

`func (o *IncidentType) SetOnChangeRepAlg(v float64)`

SetOnChangeRepAlg sets OnChangeRepAlg field to given value.

### HasOnChangeRepAlg

`func (o *IncidentType) HasOnChangeRepAlg() bool`

HasOnChangeRepAlg returns a boolean if a field has been set.

### GetPackID

`func (o *IncidentType) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *IncidentType) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *IncidentType) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *IncidentType) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *IncidentType) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *IncidentType) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *IncidentType) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *IncidentType) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPlaybookId

`func (o *IncidentType) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *IncidentType) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *IncidentType) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *IncidentType) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPreProcessingScript

`func (o *IncidentType) GetPreProcessingScript() string`

GetPreProcessingScript returns the PreProcessingScript field if non-nil, zero value otherwise.

### GetPreProcessingScriptOk

`func (o *IncidentType) GetPreProcessingScriptOk() (*string, bool)`

GetPreProcessingScriptOk returns a tuple with the PreProcessingScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProcessingScript

`func (o *IncidentType) SetPreProcessingScript(v string)`

SetPreProcessingScript sets PreProcessingScript field to given value.

### HasPreProcessingScript

`func (o *IncidentType) HasPreProcessingScript() bool`

HasPreProcessingScript returns a boolean if a field has been set.

### GetPrevName

`func (o *IncidentType) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *IncidentType) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *IncidentType) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *IncidentType) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *IncidentType) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *IncidentType) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *IncidentType) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *IncidentType) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *IncidentType) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *IncidentType) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *IncidentType) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *IncidentType) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetReadonly

`func (o *IncidentType) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *IncidentType) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *IncidentType) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *IncidentType) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetReputationCalc

`func (o *IncidentType) GetReputationCalc() float64`

GetReputationCalc returns the ReputationCalc field if non-nil, zero value otherwise.

### GetReputationCalcOk

`func (o *IncidentType) GetReputationCalcOk() (*float64, bool)`

GetReputationCalcOk returns a tuple with the ReputationCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationCalc

`func (o *IncidentType) SetReputationCalc(v float64)`

SetReputationCalc sets ReputationCalc field to given value.

### HasReputationCalc

`func (o *IncidentType) HasReputationCalc() bool`

HasReputationCalc returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *IncidentType) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *IncidentType) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *IncidentType) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *IncidentType) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *IncidentType) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *IncidentType) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *IncidentType) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *IncidentType) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSla

`func (o *IncidentType) GetSla() int64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *IncidentType) GetSlaOk() (*int64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *IncidentType) SetSla(v int64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *IncidentType) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaReminder

`func (o *IncidentType) GetSlaReminder() int64`

GetSlaReminder returns the SlaReminder field if non-nil, zero value otherwise.

### GetSlaReminderOk

`func (o *IncidentType) GetSlaReminderOk() (*int64, bool)`

GetSlaReminderOk returns a tuple with the SlaReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaReminder

`func (o *IncidentType) SetSlaReminder(v int64)`

SetSlaReminder sets SlaReminder field to given value.

### HasSlaReminder

`func (o *IncidentType) HasSlaReminder() bool`

HasSlaReminder returns a boolean if a field has been set.

### GetSortValues

`func (o *IncidentType) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *IncidentType) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *IncidentType) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *IncidentType) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSystem

`func (o *IncidentType) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IncidentType) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IncidentType) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IncidentType) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetToServerVersion

`func (o *IncidentType) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *IncidentType) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *IncidentType) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *IncidentType) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *IncidentType) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *IncidentType) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *IncidentType) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *IncidentType) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *IncidentType) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *IncidentType) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *IncidentType) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *IncidentType) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *IncidentType) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IncidentType) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IncidentType) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IncidentType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWeeks

`func (o *IncidentType) GetWeeks() int64`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *IncidentType) GetWeeksOk() (*int64, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *IncidentType) SetWeeks(v int64)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *IncidentType) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetWeeksR

`func (o *IncidentType) GetWeeksR() int64`

GetWeeksR returns the WeeksR field if non-nil, zero value otherwise.

### GetWeeksROk

`func (o *IncidentType) GetWeeksROk() (*int64, bool)`

GetWeeksROk returns a tuple with the WeeksR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeksR

`func (o *IncidentType) SetWeeksR(v int64)`

SetWeeksR sets WeeksR field to given value.

### HasWeeksR

`func (o *IncidentType) HasWeeksR() bool`

HasWeeksR returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


