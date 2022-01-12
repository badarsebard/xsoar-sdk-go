# Layout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | Pointer to **string** |  | [optional] 
**FromServerVersion** | Pointer to **string** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ItemVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PrevKind** | Pointer to **string** |  | [optional] 
**PrevTypeId** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Sections** | Pointer to [**[]LayoutSection**](LayoutSection.md) |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**ToServerVersion** | Pointer to **string** |  | [optional] 
**TypeId** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewLayout

`func NewLayout() *Layout`

NewLayout instantiates a new Layout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutWithDefaults

`func NewLayoutWithDefaults() *Layout`

NewLayoutWithDefaults instantiates a new Layout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *Layout) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *Layout) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *Layout) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *Layout) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *Layout) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *Layout) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *Layout) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *Layout) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *Layout) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Layout) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Layout) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Layout) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Layout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Layout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Layout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Layout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemVersion

`func (o *Layout) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *Layout) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *Layout) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *Layout) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetKind

`func (o *Layout) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Layout) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Layout) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Layout) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLocked

`func (o *Layout) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Layout) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Layout) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Layout) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *Layout) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Layout) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Layout) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Layout) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Layout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Layout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Layout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Layout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *Layout) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Layout) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Layout) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Layout) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPackID

`func (o *Layout) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *Layout) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *Layout) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *Layout) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *Layout) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *Layout) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *Layout) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *Layout) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPrevKind

`func (o *Layout) GetPrevKind() string`

GetPrevKind returns the PrevKind field if non-nil, zero value otherwise.

### GetPrevKindOk

`func (o *Layout) GetPrevKindOk() (*string, bool)`

GetPrevKindOk returns a tuple with the PrevKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevKind

`func (o *Layout) SetPrevKind(v string)`

SetPrevKind sets PrevKind field to given value.

### HasPrevKind

`func (o *Layout) HasPrevKind() bool`

HasPrevKind returns a boolean if a field has been set.

### GetPrevTypeId

`func (o *Layout) GetPrevTypeId() string`

GetPrevTypeId returns the PrevTypeId field if non-nil, zero value otherwise.

### GetPrevTypeIdOk

`func (o *Layout) GetPrevTypeIdOk() (*string, bool)`

GetPrevTypeIdOk returns a tuple with the PrevTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTypeId

`func (o *Layout) SetPrevTypeId(v string)`

SetPrevTypeId sets PrevTypeId field to given value.

### HasPrevTypeId

`func (o *Layout) HasPrevTypeId() bool`

HasPrevTypeId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Layout) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Layout) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Layout) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Layout) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *Layout) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *Layout) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *Layout) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *Layout) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetSections

`func (o *Layout) GetSections() []LayoutSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Layout) GetSectionsOk() (*[]LayoutSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Layout) SetSections(v []LayoutSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Layout) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Layout) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Layout) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Layout) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Layout) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *Layout) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *Layout) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *Layout) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *Layout) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *Layout) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Layout) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Layout) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Layout) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSystem

`func (o *Layout) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Layout) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Layout) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Layout) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetToServerVersion

`func (o *Layout) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *Layout) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *Layout) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *Layout) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetTypeId

`func (o *Layout) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Layout) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Layout) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Layout) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *Layout) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *Layout) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *Layout) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *Layout) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *Layout) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *Layout) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *Layout) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *Layout) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *Layout) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Layout) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Layout) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Layout) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


