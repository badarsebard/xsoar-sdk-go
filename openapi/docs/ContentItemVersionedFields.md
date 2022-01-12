# ContentItemVersionedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | Pointer to **string** |  | [optional] 
**FromServerVersion** | Pointer to **string** |  | [optional] 
**ItemVersion** | Pointer to **string** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**ToServerVersion** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 

## Methods

### NewContentItemVersionedFields

`func NewContentItemVersionedFields() *ContentItemVersionedFields`

NewContentItemVersionedFields instantiates a new ContentItemVersionedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentItemVersionedFieldsWithDefaults

`func NewContentItemVersionedFieldsWithDefaults() *ContentItemVersionedFields`

NewContentItemVersionedFieldsWithDefaults instantiates a new ContentItemVersionedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *ContentItemVersionedFields) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ContentItemVersionedFields) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ContentItemVersionedFields) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ContentItemVersionedFields) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *ContentItemVersionedFields) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *ContentItemVersionedFields) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *ContentItemVersionedFields) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *ContentItemVersionedFields) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetItemVersion

`func (o *ContentItemVersionedFields) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *ContentItemVersionedFields) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *ContentItemVersionedFields) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *ContentItemVersionedFields) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetPackID

`func (o *ContentItemVersionedFields) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *ContentItemVersionedFields) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *ContentItemVersionedFields) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *ContentItemVersionedFields) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *ContentItemVersionedFields) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *ContentItemVersionedFields) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *ContentItemVersionedFields) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *ContentItemVersionedFields) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *ContentItemVersionedFields) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *ContentItemVersionedFields) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *ContentItemVersionedFields) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *ContentItemVersionedFields) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetShouldCommit

`func (o *ContentItemVersionedFields) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *ContentItemVersionedFields) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *ContentItemVersionedFields) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *ContentItemVersionedFields) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetToServerVersion

`func (o *ContentItemVersionedFields) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *ContentItemVersionedFields) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *ContentItemVersionedFields) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *ContentItemVersionedFields) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *ContentItemVersionedFields) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *ContentItemVersionedFields) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *ContentItemVersionedFields) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *ContentItemVersionedFields) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *ContentItemVersionedFields) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *ContentItemVersionedFields) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *ContentItemVersionedFields) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *ContentItemVersionedFields) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


