# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**ClonedFrom** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsCommand** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsTitleTask** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PlaybookId** | Pointer to **string** |  | [optional] 
**PlaybookName** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**ScriptId** | Pointer to **string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | TaskType is the Task in the playbook context as a node | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *Task) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Task) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Task) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Task) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetClonedFrom

`func (o *Task) GetClonedFrom() string`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *Task) GetClonedFromOk() (*string, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *Task) SetClonedFrom(v string)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *Task) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetComment

`func (o *Task) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Task) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Task) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Task) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConditions

`func (o *Task) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Task) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Task) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Task) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHighlight

`func (o *Task) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Task) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Task) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Task) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCommand

`func (o *Task) GetIsCommand() bool`

GetIsCommand returns the IsCommand field if non-nil, zero value otherwise.

### GetIsCommandOk

`func (o *Task) GetIsCommandOk() (*bool, bool)`

GetIsCommandOk returns a tuple with the IsCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommand

`func (o *Task) SetIsCommand(v bool)`

SetIsCommand sets IsCommand field to given value.

### HasIsCommand

`func (o *Task) HasIsCommand() bool`

HasIsCommand returns a boolean if a field has been set.

### GetIsLocked

`func (o *Task) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Task) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Task) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Task) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsTitleTask

`func (o *Task) GetIsTitleTask() bool`

GetIsTitleTask returns the IsTitleTask field if non-nil, zero value otherwise.

### GetIsTitleTaskOk

`func (o *Task) GetIsTitleTaskOk() (*bool, bool)`

GetIsTitleTaskOk returns a tuple with the IsTitleTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTitleTask

`func (o *Task) SetIsTitleTask(v bool)`

SetIsTitleTask sets IsTitleTask field to given value.

### HasIsTitleTask

`func (o *Task) HasIsTitleTask() bool`

HasIsTitleTask returns a boolean if a field has been set.

### GetModified

`func (o *Task) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Task) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Task) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Task) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *Task) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Task) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Task) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Task) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPlaybookId

`func (o *Task) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *Task) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *Task) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *Task) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPlaybookName

`func (o *Task) GetPlaybookName() string`

GetPlaybookName returns the PlaybookName field if non-nil, zero value otherwise.

### GetPlaybookNameOk

`func (o *Task) GetPlaybookNameOk() (*string, bool)`

GetPlaybookNameOk returns a tuple with the PlaybookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookName

`func (o *Task) SetPlaybookName(v string)`

SetPlaybookName sets PlaybookName field to given value.

### HasPlaybookName

`func (o *Task) HasPlaybookName() bool`

HasPlaybookName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Task) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Task) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Task) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Task) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetScriptId

`func (o *Task) GetScriptId() string`

GetScriptId returns the ScriptId field if non-nil, zero value otherwise.

### GetScriptIdOk

`func (o *Task) GetScriptIdOk() (*string, bool)`

GetScriptIdOk returns a tuple with the ScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptId

`func (o *Task) SetScriptId(v string)`

SetScriptId sets ScriptId field to given value.

### HasScriptId

`func (o *Task) HasScriptId() bool`

HasScriptId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Task) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Task) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Task) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Task) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *Task) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Task) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Task) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Task) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetTags

`func (o *Task) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Task) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Task) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Task) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *Task) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Task) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Task) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Task) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


