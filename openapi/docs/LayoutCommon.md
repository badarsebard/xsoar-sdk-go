# LayoutCommon

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

### NewLayoutCommon

`func NewLayoutCommon() *LayoutCommon`

NewLayoutCommon instantiates a new LayoutCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutCommonWithDefaults

`func NewLayoutCommonWithDefaults() *LayoutCommon`

NewLayoutCommonWithDefaults instantiates a new LayoutCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *LayoutCommon) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *LayoutCommon) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *LayoutCommon) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *LayoutCommon) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *LayoutCommon) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *LayoutCommon) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *LayoutCommon) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *LayoutCommon) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *LayoutCommon) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *LayoutCommon) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *LayoutCommon) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *LayoutCommon) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *LayoutCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LayoutCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LayoutCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LayoutCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemVersion

`func (o *LayoutCommon) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *LayoutCommon) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *LayoutCommon) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *LayoutCommon) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetKind

`func (o *LayoutCommon) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LayoutCommon) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LayoutCommon) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LayoutCommon) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLocked

`func (o *LayoutCommon) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *LayoutCommon) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *LayoutCommon) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *LayoutCommon) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *LayoutCommon) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LayoutCommon) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *LayoutCommon) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *LayoutCommon) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *LayoutCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayoutCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayoutCommon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LayoutCommon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *LayoutCommon) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *LayoutCommon) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *LayoutCommon) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *LayoutCommon) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPackID

`func (o *LayoutCommon) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *LayoutCommon) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *LayoutCommon) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *LayoutCommon) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *LayoutCommon) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *LayoutCommon) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *LayoutCommon) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *LayoutCommon) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPrevKind

`func (o *LayoutCommon) GetPrevKind() string`

GetPrevKind returns the PrevKind field if non-nil, zero value otherwise.

### GetPrevKindOk

`func (o *LayoutCommon) GetPrevKindOk() (*string, bool)`

GetPrevKindOk returns a tuple with the PrevKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevKind

`func (o *LayoutCommon) SetPrevKind(v string)`

SetPrevKind sets PrevKind field to given value.

### HasPrevKind

`func (o *LayoutCommon) HasPrevKind() bool`

HasPrevKind returns a boolean if a field has been set.

### GetPrevTypeId

`func (o *LayoutCommon) GetPrevTypeId() string`

GetPrevTypeId returns the PrevTypeId field if non-nil, zero value otherwise.

### GetPrevTypeIdOk

`func (o *LayoutCommon) GetPrevTypeIdOk() (*string, bool)`

GetPrevTypeIdOk returns a tuple with the PrevTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTypeId

`func (o *LayoutCommon) SetPrevTypeId(v string)`

SetPrevTypeId sets PrevTypeId field to given value.

### HasPrevTypeId

`func (o *LayoutCommon) HasPrevTypeId() bool`

HasPrevTypeId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *LayoutCommon) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *LayoutCommon) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *LayoutCommon) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *LayoutCommon) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *LayoutCommon) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *LayoutCommon) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *LayoutCommon) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *LayoutCommon) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *LayoutCommon) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *LayoutCommon) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *LayoutCommon) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *LayoutCommon) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *LayoutCommon) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *LayoutCommon) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *LayoutCommon) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *LayoutCommon) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *LayoutCommon) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *LayoutCommon) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *LayoutCommon) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *LayoutCommon) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSystem

`func (o *LayoutCommon) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *LayoutCommon) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *LayoutCommon) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *LayoutCommon) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetToServerVersion

`func (o *LayoutCommon) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *LayoutCommon) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *LayoutCommon) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *LayoutCommon) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetTypeId

`func (o *LayoutCommon) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *LayoutCommon) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *LayoutCommon) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *LayoutCommon) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *LayoutCommon) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *LayoutCommon) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *LayoutCommon) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *LayoutCommon) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *LayoutCommon) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *LayoutCommon) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *LayoutCommon) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *LayoutCommon) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *LayoutCommon) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LayoutCommon) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LayoutCommon) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LayoutCommon) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


