# ModuleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beta** | Pointer to **bool** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**CanGetSamples** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Cmdline** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**[]ConfigField**](ConfigField.md) |  | [optional] 
**DefaultClassifier** | Pointer to **string** |  | [optional] 
**DefaultMapperIn** | Pointer to **string** |  | [optional] 
**DefaultMapperOut** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DetailedDescription** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Executable** | Pointer to **string** |  | [optional] 
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HideEngines** | Pointer to **bool** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**IntegrationScript** | Pointer to [**IntegrationScript**](IntegrationScript.md) |  | [optional] 
**IsPasswordProtected** | Pointer to **bool** |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**ScriptNotVisible** | Pointer to **bool** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**SourceModuleID** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewModuleConfiguration

`func NewModuleConfiguration() *ModuleConfiguration`

NewModuleConfiguration instantiates a new ModuleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleConfigurationWithDefaults

`func NewModuleConfigurationWithDefaults() *ModuleConfiguration`

NewModuleConfigurationWithDefaults instantiates a new ModuleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *ModuleConfiguration) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ModuleConfiguration) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ModuleConfiguration) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ModuleConfiguration) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetBrand

`func (o *ModuleConfiguration) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ModuleConfiguration) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ModuleConfiguration) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ModuleConfiguration) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCanGetSamples

`func (o *ModuleConfiguration) GetCanGetSamples() bool`

GetCanGetSamples returns the CanGetSamples field if non-nil, zero value otherwise.

### GetCanGetSamplesOk

`func (o *ModuleConfiguration) GetCanGetSamplesOk() (*bool, bool)`

GetCanGetSamplesOk returns a tuple with the CanGetSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGetSamples

`func (o *ModuleConfiguration) SetCanGetSamples(v bool)`

SetCanGetSamples sets CanGetSamples field to given value.

### HasCanGetSamples

`func (o *ModuleConfiguration) HasCanGetSamples() bool`

HasCanGetSamples returns a boolean if a field has been set.

### GetCategory

`func (o *ModuleConfiguration) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModuleConfiguration) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModuleConfiguration) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModuleConfiguration) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCmdline

`func (o *ModuleConfiguration) GetCmdline() string`

GetCmdline returns the Cmdline field if non-nil, zero value otherwise.

### GetCmdlineOk

`func (o *ModuleConfiguration) GetCmdlineOk() (*string, bool)`

GetCmdlineOk returns a tuple with the Cmdline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdline

`func (o *ModuleConfiguration) SetCmdline(v string)`

SetCmdline sets Cmdline field to given value.

### HasCmdline

`func (o *ModuleConfiguration) HasCmdline() bool`

HasCmdline returns a boolean if a field has been set.

### GetCommitMessage

`func (o *ModuleConfiguration) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ModuleConfiguration) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ModuleConfiguration) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ModuleConfiguration) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetConfiguration

`func (o *ModuleConfiguration) GetConfiguration() []ConfigField`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ModuleConfiguration) GetConfigurationOk() (*[]ConfigField, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ModuleConfiguration) SetConfiguration(v []ConfigField)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ModuleConfiguration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDefaultClassifier

`func (o *ModuleConfiguration) GetDefaultClassifier() string`

GetDefaultClassifier returns the DefaultClassifier field if non-nil, zero value otherwise.

### GetDefaultClassifierOk

`func (o *ModuleConfiguration) GetDefaultClassifierOk() (*string, bool)`

GetDefaultClassifierOk returns a tuple with the DefaultClassifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClassifier

`func (o *ModuleConfiguration) SetDefaultClassifier(v string)`

SetDefaultClassifier sets DefaultClassifier field to given value.

### HasDefaultClassifier

`func (o *ModuleConfiguration) HasDefaultClassifier() bool`

HasDefaultClassifier returns a boolean if a field has been set.

### GetDefaultMapperIn

`func (o *ModuleConfiguration) GetDefaultMapperIn() string`

GetDefaultMapperIn returns the DefaultMapperIn field if non-nil, zero value otherwise.

### GetDefaultMapperInOk

`func (o *ModuleConfiguration) GetDefaultMapperInOk() (*string, bool)`

GetDefaultMapperInOk returns a tuple with the DefaultMapperIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapperIn

`func (o *ModuleConfiguration) SetDefaultMapperIn(v string)`

SetDefaultMapperIn sets DefaultMapperIn field to given value.

### HasDefaultMapperIn

`func (o *ModuleConfiguration) HasDefaultMapperIn() bool`

HasDefaultMapperIn returns a boolean if a field has been set.

### GetDefaultMapperOut

`func (o *ModuleConfiguration) GetDefaultMapperOut() string`

GetDefaultMapperOut returns the DefaultMapperOut field if non-nil, zero value otherwise.

### GetDefaultMapperOutOk

`func (o *ModuleConfiguration) GetDefaultMapperOutOk() (*string, bool)`

GetDefaultMapperOutOk returns a tuple with the DefaultMapperOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapperOut

`func (o *ModuleConfiguration) SetDefaultMapperOut(v string)`

SetDefaultMapperOut sets DefaultMapperOut field to given value.

### HasDefaultMapperOut

`func (o *ModuleConfiguration) HasDefaultMapperOut() bool`

HasDefaultMapperOut returns a boolean if a field has been set.

### GetDeprecated

`func (o *ModuleConfiguration) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ModuleConfiguration) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ModuleConfiguration) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ModuleConfiguration) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *ModuleConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetailedDescription

`func (o *ModuleConfiguration) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *ModuleConfiguration) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *ModuleConfiguration) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *ModuleConfiguration) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetDisplay

`func (o *ModuleConfiguration) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ModuleConfiguration) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ModuleConfiguration) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ModuleConfiguration) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetExecutable

`func (o *ModuleConfiguration) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *ModuleConfiguration) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *ModuleConfiguration) SetExecutable(v string)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *ModuleConfiguration) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *ModuleConfiguration) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *ModuleConfiguration) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *ModuleConfiguration) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *ModuleConfiguration) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHidden

`func (o *ModuleConfiguration) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ModuleConfiguration) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ModuleConfiguration) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ModuleConfiguration) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHideEngines

`func (o *ModuleConfiguration) GetHideEngines() bool`

GetHideEngines returns the HideEngines field if non-nil, zero value otherwise.

### GetHideEnginesOk

`func (o *ModuleConfiguration) GetHideEnginesOk() (*bool, bool)`

GetHideEnginesOk returns a tuple with the HideEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideEngines

`func (o *ModuleConfiguration) SetHideEngines(v bool)`

SetHideEngines sets HideEngines field to given value.

### HasHideEngines

`func (o *ModuleConfiguration) HasHideEngines() bool`

HasHideEngines returns a boolean if a field has been set.

### GetHighlight

`func (o *ModuleConfiguration) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *ModuleConfiguration) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *ModuleConfiguration) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *ModuleConfiguration) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetIcon

`func (o *ModuleConfiguration) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ModuleConfiguration) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ModuleConfiguration) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ModuleConfiguration) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *ModuleConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModuleConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ModuleConfiguration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModuleConfiguration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModuleConfiguration) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ModuleConfiguration) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIntegrationScript

`func (o *ModuleConfiguration) GetIntegrationScript() IntegrationScript`

GetIntegrationScript returns the IntegrationScript field if non-nil, zero value otherwise.

### GetIntegrationScriptOk

`func (o *ModuleConfiguration) GetIntegrationScriptOk() (*IntegrationScript, bool)`

GetIntegrationScriptOk returns a tuple with the IntegrationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationScript

`func (o *ModuleConfiguration) SetIntegrationScript(v IntegrationScript)`

SetIntegrationScript sets IntegrationScript field to given value.

### HasIntegrationScript

`func (o *ModuleConfiguration) HasIntegrationScript() bool`

HasIntegrationScript returns a boolean if a field has been set.

### GetIsPasswordProtected

`func (o *ModuleConfiguration) GetIsPasswordProtected() bool`

GetIsPasswordProtected returns the IsPasswordProtected field if non-nil, zero value otherwise.

### GetIsPasswordProtectedOk

`func (o *ModuleConfiguration) GetIsPasswordProtectedOk() (*bool, bool)`

GetIsPasswordProtectedOk returns a tuple with the IsPasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordProtected

`func (o *ModuleConfiguration) SetIsPasswordProtected(v bool)`

SetIsPasswordProtected sets IsPasswordProtected field to given value.

### HasIsPasswordProtected

`func (o *ModuleConfiguration) HasIsPasswordProtected() bool`

HasIsPasswordProtected returns a boolean if a field has been set.

### GetItemVersion

`func (o *ModuleConfiguration) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *ModuleConfiguration) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *ModuleConfiguration) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *ModuleConfiguration) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLocked

`func (o *ModuleConfiguration) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ModuleConfiguration) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ModuleConfiguration) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ModuleConfiguration) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *ModuleConfiguration) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ModuleConfiguration) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ModuleConfiguration) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ModuleConfiguration) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *ModuleConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModuleConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *ModuleConfiguration) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *ModuleConfiguration) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *ModuleConfiguration) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *ModuleConfiguration) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPackID

`func (o *ModuleConfiguration) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *ModuleConfiguration) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *ModuleConfiguration) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *ModuleConfiguration) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *ModuleConfiguration) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *ModuleConfiguration) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *ModuleConfiguration) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *ModuleConfiguration) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPath

`func (o *ModuleConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModuleConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModuleConfiguration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModuleConfiguration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrevName

`func (o *ModuleConfiguration) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *ModuleConfiguration) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *ModuleConfiguration) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *ModuleConfiguration) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *ModuleConfiguration) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *ModuleConfiguration) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *ModuleConfiguration) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *ModuleConfiguration) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPrivate

`func (o *ModuleConfiguration) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ModuleConfiguration) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ModuleConfiguration) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ModuleConfiguration) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *ModuleConfiguration) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *ModuleConfiguration) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *ModuleConfiguration) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *ModuleConfiguration) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetReadonly

`func (o *ModuleConfiguration) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ModuleConfiguration) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ModuleConfiguration) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ModuleConfiguration) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetScriptNotVisible

`func (o *ModuleConfiguration) GetScriptNotVisible() bool`

GetScriptNotVisible returns the ScriptNotVisible field if non-nil, zero value otherwise.

### GetScriptNotVisibleOk

`func (o *ModuleConfiguration) GetScriptNotVisibleOk() (*bool, bool)`

GetScriptNotVisibleOk returns a tuple with the ScriptNotVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptNotVisible

`func (o *ModuleConfiguration) SetScriptNotVisible(v bool)`

SetScriptNotVisible sets ScriptNotVisible field to given value.

### HasScriptNotVisible

`func (o *ModuleConfiguration) HasScriptNotVisible() bool`

HasScriptNotVisible returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ModuleConfiguration) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ModuleConfiguration) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ModuleConfiguration) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ModuleConfiguration) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *ModuleConfiguration) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *ModuleConfiguration) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *ModuleConfiguration) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *ModuleConfiguration) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *ModuleConfiguration) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *ModuleConfiguration) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *ModuleConfiguration) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *ModuleConfiguration) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceModuleID

`func (o *ModuleConfiguration) GetSourceModuleID() string`

GetSourceModuleID returns the SourceModuleID field if non-nil, zero value otherwise.

### GetSourceModuleIDOk

`func (o *ModuleConfiguration) GetSourceModuleIDOk() (*string, bool)`

GetSourceModuleIDOk returns a tuple with the SourceModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceModuleID

`func (o *ModuleConfiguration) SetSourceModuleID(v string)`

SetSourceModuleID sets SourceModuleID field to given value.

### HasSourceModuleID

`func (o *ModuleConfiguration) HasSourceModuleID() bool`

HasSourceModuleID returns a boolean if a field has been set.

### GetSystem

`func (o *ModuleConfiguration) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ModuleConfiguration) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ModuleConfiguration) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ModuleConfiguration) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetToServerVersion

`func (o *ModuleConfiguration) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *ModuleConfiguration) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *ModuleConfiguration) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *ModuleConfiguration) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *ModuleConfiguration) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *ModuleConfiguration) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *ModuleConfiguration) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *ModuleConfiguration) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *ModuleConfiguration) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *ModuleConfiguration) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *ModuleConfiguration) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *ModuleConfiguration) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *ModuleConfiguration) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModuleConfiguration) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModuleConfiguration) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModuleConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


