# Evidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardID** | Pointer to **int64** |  | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**Description** | Pointer to **string** | The description for the resolve | [optional] 
**EntryId** | Pointer to **string** | The entry ID | [optional] 
**Fetched** | Pointer to **time.Time** | when the evidence entry was fetched | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncidentId** | Pointer to **string** | The incident ID | [optional] 
**MarkedBy** | Pointer to **string** | the user that marked this evidence | [optional] 
**MarkedDate** | Pointer to **time.Time** | when this evidence was marked | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Occurred** | Pointer to **time.Time** | When this evidence has occurred | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**TagsRaw** | Pointer to **[]string** | TagsRaw | [optional] 
**TaskId** | Pointer to **string** | when the evidence entry was fetched | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEvidence

`func NewEvidence() *Evidence`

NewEvidence instantiates a new Evidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceWithDefaults

`func NewEvidenceWithDefaults() *Evidence`

NewEvidenceWithDefaults instantiates a new Evidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardID

`func (o *Evidence) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *Evidence) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *Evidence) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *Evidence) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAllRead

`func (o *Evidence) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *Evidence) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *Evidence) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *Evidence) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *Evidence) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *Evidence) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *Evidence) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *Evidence) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *Evidence) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *Evidence) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *Evidence) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *Evidence) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *Evidence) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Evidence) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Evidence) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Evidence) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntryId

`func (o *Evidence) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *Evidence) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *Evidence) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *Evidence) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetFetched

`func (o *Evidence) GetFetched() time.Time`

GetFetched returns the Fetched field if non-nil, zero value otherwise.

### GetFetchedOk

`func (o *Evidence) GetFetchedOk() (*time.Time, bool)`

GetFetchedOk returns a tuple with the Fetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetched

`func (o *Evidence) SetFetched(v time.Time)`

SetFetched sets Fetched field to given value.

### HasFetched

`func (o *Evidence) HasFetched() bool`

HasFetched returns a boolean if a field has been set.

### GetHasRole

`func (o *Evidence) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *Evidence) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *Evidence) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *Evidence) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *Evidence) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Evidence) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Evidence) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Evidence) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Evidence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Evidence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Evidence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Evidence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentId

`func (o *Evidence) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *Evidence) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *Evidence) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *Evidence) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetMarkedBy

`func (o *Evidence) GetMarkedBy() string`

GetMarkedBy returns the MarkedBy field if non-nil, zero value otherwise.

### GetMarkedByOk

`func (o *Evidence) GetMarkedByOk() (*string, bool)`

GetMarkedByOk returns a tuple with the MarkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedBy

`func (o *Evidence) SetMarkedBy(v string)`

SetMarkedBy sets MarkedBy field to given value.

### HasMarkedBy

`func (o *Evidence) HasMarkedBy() bool`

HasMarkedBy returns a boolean if a field has been set.

### GetMarkedDate

`func (o *Evidence) GetMarkedDate() time.Time`

GetMarkedDate returns the MarkedDate field if non-nil, zero value otherwise.

### GetMarkedDateOk

`func (o *Evidence) GetMarkedDateOk() (*time.Time, bool)`

GetMarkedDateOk returns a tuple with the MarkedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedDate

`func (o *Evidence) SetMarkedDate(v time.Time)`

SetMarkedDate sets MarkedDate field to given value.

### HasMarkedDate

`func (o *Evidence) HasMarkedDate() bool`

HasMarkedDate returns a boolean if a field has been set.

### GetModified

`func (o *Evidence) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Evidence) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Evidence) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Evidence) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNumericId

`func (o *Evidence) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Evidence) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Evidence) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Evidence) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOccurred

`func (o *Evidence) GetOccurred() time.Time`

GetOccurred returns the Occurred field if non-nil, zero value otherwise.

### GetOccurredOk

`func (o *Evidence) GetOccurredOk() (*time.Time, bool)`

GetOccurredOk returns a tuple with the Occurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurred

`func (o *Evidence) SetOccurred(v time.Time)`

SetOccurred sets Occurred field to given value.

### HasOccurred

`func (o *Evidence) HasOccurred() bool`

HasOccurred returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *Evidence) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *Evidence) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *Evidence) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *Evidence) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *Evidence) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *Evidence) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *Evidence) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *Evidence) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *Evidence) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *Evidence) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *Evidence) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *Evidence) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Evidence) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Evidence) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Evidence) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Evidence) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRoles

`func (o *Evidence) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Evidence) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Evidence) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Evidence) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Evidence) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Evidence) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Evidence) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Evidence) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *Evidence) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Evidence) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Evidence) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Evidence) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetTags

`func (o *Evidence) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Evidence) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Evidence) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Evidence) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagsRaw

`func (o *Evidence) GetTagsRaw() []string`

GetTagsRaw returns the TagsRaw field if non-nil, zero value otherwise.

### GetTagsRawOk

`func (o *Evidence) GetTagsRawOk() (*[]string, bool)`

GetTagsRawOk returns a tuple with the TagsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsRaw

`func (o *Evidence) SetTagsRaw(v []string)`

SetTagsRaw sets TagsRaw field to given value.

### HasTagsRaw

`func (o *Evidence) HasTagsRaw() bool`

HasTagsRaw returns a boolean if a field has been set.

### GetTaskId

`func (o *Evidence) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Evidence) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Evidence) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Evidence) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetVersion

`func (o *Evidence) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Evidence) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Evidence) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Evidence) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *Evidence) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *Evidence) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *Evidence) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *Evidence) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *Evidence) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *Evidence) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *Evidence) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *Evidence) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *Evidence) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *Evidence) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *Evidence) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *Evidence) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


