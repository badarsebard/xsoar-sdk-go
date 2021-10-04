# Investigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardID** | Pointer to **int64** |  | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** | Category of the investigation | [optional] 
**ChildInvestigations** | Pointer to **[]string** | ChildInvestigations id&#39;s | [optional] 
**Closed** | Pointer to **time.Time** | When was this closed | [optional] 
**ClosingUserId** | Pointer to **string** | The user ID that closed this investigation | [optional] 
**Created** | Pointer to **time.Time** | When was this created | [optional] 
**CreatingUserId** | Pointer to **string** | The user ID that created this investigation | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**Details** | Pointer to **string** | User defined free text details | [optional] 
**Entitlements** | Pointer to **[]string** | One time entitlements | [optional] 
**EntryUsers** | Pointer to **[]string** | EntryUsers | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**HighPriority** | Pointer to **bool** | HighPriority - tasks of this investigation should run in high priority | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsChildInvestigation** | Pointer to **bool** | IsChildInvestigation | [optional] 
**IsDebug** | Pointer to **bool** | IsDebug ... | [optional] 
**LastOpen** | Pointer to **time.Time** |  | [optional] 
**MirrorAutoClose** | Pointer to **map[string]bool** | MirrorAutoClose will tell us to close the Chat Module channel if we close investigation | [optional] 
**MirrorTypes** | Pointer to **map[string]string** | MirrorTypes holds info about mirror direction and message type to be mirrored message type can be either &#39;all&#39; or &#39;chat&#39; direction can be either &#39;FromDemisto&#39;, &#39;ToDemisto&#39; or &#39;Both&#39; if this investigation is mirrored | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | The name of the investigation, which is unique to the project | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**OpenDuration** | Pointer to **int64** | Duration from open to close time | [optional] 
**ParentInvestigation** | Pointer to **string** | ParentInvestigation - parent id, in case this is a child investigation of another investigation | [optional] 
**PersistentEntitlements** | Pointer to **map[string]string** | Persistent entitlement per tag. Empty tag will also return an entitlement | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**RawCategory** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **map[string]string** | The reason for the status (resolve) | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**RunStatus** | Pointer to **string** | RunStatus of a job | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SlackMirrorAutoClose** | Pointer to **bool** | DEPRECATED - DeprecatedSlackMirrorAutoClose will tell us to close the Slack channel if we close investigation | [optional] 
**SlackMirrorType** | Pointer to **string** | DEPRECATED - DeprecatedSlackMirrorType holds info about mirror direction and message type to be mirror message type can be either &#39;all&#39; or &#39;chat&#39; direction can be either &#39;demisto2Slack&#39;, &#39;slack2Demisto&#39; or &#39;both&#39; if this investigation is mirrored to Slack | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **float64** | InvestigationStatus is the status type | [optional] 
**Systems** | Pointer to [**[]System**](System.md) | The systems involved | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**Type** | Pointer to **float64** |  | [optional] 
**Users** | Pointer to **[]string** | The users who share this investigation | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInvestigation

`func NewInvestigation() *Investigation`

NewInvestigation instantiates a new Investigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationWithDefaults

`func NewInvestigationWithDefaults() *Investigation`

NewInvestigationWithDefaults instantiates a new Investigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardID

`func (o *Investigation) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *Investigation) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *Investigation) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *Investigation) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAllRead

`func (o *Investigation) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *Investigation) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *Investigation) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *Investigation) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *Investigation) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *Investigation) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *Investigation) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *Investigation) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetCategory

`func (o *Investigation) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Investigation) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Investigation) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Investigation) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChildInvestigations

`func (o *Investigation) GetChildInvestigations() []string`

GetChildInvestigations returns the ChildInvestigations field if non-nil, zero value otherwise.

### GetChildInvestigationsOk

`func (o *Investigation) GetChildInvestigationsOk() (*[]string, bool)`

GetChildInvestigationsOk returns a tuple with the ChildInvestigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildInvestigations

`func (o *Investigation) SetChildInvestigations(v []string)`

SetChildInvestigations sets ChildInvestigations field to given value.

### HasChildInvestigations

`func (o *Investigation) HasChildInvestigations() bool`

HasChildInvestigations returns a boolean if a field has been set.

### GetClosed

`func (o *Investigation) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *Investigation) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *Investigation) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *Investigation) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosingUserId

`func (o *Investigation) GetClosingUserId() string`

GetClosingUserId returns the ClosingUserId field if non-nil, zero value otherwise.

### GetClosingUserIdOk

`func (o *Investigation) GetClosingUserIdOk() (*string, bool)`

GetClosingUserIdOk returns a tuple with the ClosingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUserId

`func (o *Investigation) SetClosingUserId(v string)`

SetClosingUserId sets ClosingUserId field to given value.

### HasClosingUserId

`func (o *Investigation) HasClosingUserId() bool`

HasClosingUserId returns a boolean if a field has been set.

### GetCreated

`func (o *Investigation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Investigation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Investigation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Investigation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatingUserId

`func (o *Investigation) GetCreatingUserId() string`

GetCreatingUserId returns the CreatingUserId field if non-nil, zero value otherwise.

### GetCreatingUserIdOk

`func (o *Investigation) GetCreatingUserIdOk() (*string, bool)`

GetCreatingUserIdOk returns a tuple with the CreatingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUserId

`func (o *Investigation) SetCreatingUserId(v string)`

SetCreatingUserId sets CreatingUserId field to given value.

### HasCreatingUserId

`func (o *Investigation) HasCreatingUserId() bool`

HasCreatingUserId returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *Investigation) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *Investigation) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *Investigation) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *Investigation) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDetails

`func (o *Investigation) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Investigation) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Investigation) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Investigation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEntitlements

`func (o *Investigation) GetEntitlements() []string`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *Investigation) GetEntitlementsOk() (*[]string, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *Investigation) SetEntitlements(v []string)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *Investigation) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntryUsers

`func (o *Investigation) GetEntryUsers() []string`

GetEntryUsers returns the EntryUsers field if non-nil, zero value otherwise.

### GetEntryUsersOk

`func (o *Investigation) GetEntryUsersOk() (*[]string, bool)`

GetEntryUsersOk returns a tuple with the EntryUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryUsers

`func (o *Investigation) SetEntryUsers(v []string)`

SetEntryUsers sets EntryUsers field to given value.

### HasEntryUsers

`func (o *Investigation) HasEntryUsers() bool`

HasEntryUsers returns a boolean if a field has been set.

### GetHasRole

`func (o *Investigation) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *Investigation) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *Investigation) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *Investigation) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighPriority

`func (o *Investigation) GetHighPriority() bool`

GetHighPriority returns the HighPriority field if non-nil, zero value otherwise.

### GetHighPriorityOk

`func (o *Investigation) GetHighPriorityOk() (*bool, bool)`

GetHighPriorityOk returns a tuple with the HighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriority

`func (o *Investigation) SetHighPriority(v bool)`

SetHighPriority sets HighPriority field to given value.

### HasHighPriority

`func (o *Investigation) HasHighPriority() bool`

HasHighPriority returns a boolean if a field has been set.

### GetHighlight

`func (o *Investigation) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Investigation) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Investigation) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Investigation) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Investigation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Investigation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Investigation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Investigation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsChildInvestigation

`func (o *Investigation) GetIsChildInvestigation() bool`

GetIsChildInvestigation returns the IsChildInvestigation field if non-nil, zero value otherwise.

### GetIsChildInvestigationOk

`func (o *Investigation) GetIsChildInvestigationOk() (*bool, bool)`

GetIsChildInvestigationOk returns a tuple with the IsChildInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChildInvestigation

`func (o *Investigation) SetIsChildInvestigation(v bool)`

SetIsChildInvestigation sets IsChildInvestigation field to given value.

### HasIsChildInvestigation

`func (o *Investigation) HasIsChildInvestigation() bool`

HasIsChildInvestigation returns a boolean if a field has been set.

### GetIsDebug

`func (o *Investigation) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *Investigation) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *Investigation) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *Investigation) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetLastOpen

`func (o *Investigation) GetLastOpen() time.Time`

GetLastOpen returns the LastOpen field if non-nil, zero value otherwise.

### GetLastOpenOk

`func (o *Investigation) GetLastOpenOk() (*time.Time, bool)`

GetLastOpenOk returns a tuple with the LastOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpen

`func (o *Investigation) SetLastOpen(v time.Time)`

SetLastOpen sets LastOpen field to given value.

### HasLastOpen

`func (o *Investigation) HasLastOpen() bool`

HasLastOpen returns a boolean if a field has been set.

### GetMirrorAutoClose

`func (o *Investigation) GetMirrorAutoClose() map[string]bool`

GetMirrorAutoClose returns the MirrorAutoClose field if non-nil, zero value otherwise.

### GetMirrorAutoCloseOk

`func (o *Investigation) GetMirrorAutoCloseOk() (*map[string]bool, bool)`

GetMirrorAutoCloseOk returns a tuple with the MirrorAutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorAutoClose

`func (o *Investigation) SetMirrorAutoClose(v map[string]bool)`

SetMirrorAutoClose sets MirrorAutoClose field to given value.

### HasMirrorAutoClose

`func (o *Investigation) HasMirrorAutoClose() bool`

HasMirrorAutoClose returns a boolean if a field has been set.

### GetMirrorTypes

`func (o *Investigation) GetMirrorTypes() map[string]string`

GetMirrorTypes returns the MirrorTypes field if non-nil, zero value otherwise.

### GetMirrorTypesOk

`func (o *Investigation) GetMirrorTypesOk() (*map[string]string, bool)`

GetMirrorTypesOk returns a tuple with the MirrorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTypes

`func (o *Investigation) SetMirrorTypes(v map[string]string)`

SetMirrorTypes sets MirrorTypes field to given value.

### HasMirrorTypes

`func (o *Investigation) HasMirrorTypes() bool`

HasMirrorTypes returns a boolean if a field has been set.

### GetModified

`func (o *Investigation) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Investigation) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Investigation) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Investigation) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Investigation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Investigation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Investigation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Investigation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *Investigation) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Investigation) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Investigation) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Investigation) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOpenDuration

`func (o *Investigation) GetOpenDuration() int64`

GetOpenDuration returns the OpenDuration field if non-nil, zero value otherwise.

### GetOpenDurationOk

`func (o *Investigation) GetOpenDurationOk() (*int64, bool)`

GetOpenDurationOk returns a tuple with the OpenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDuration

`func (o *Investigation) SetOpenDuration(v int64)`

SetOpenDuration sets OpenDuration field to given value.

### HasOpenDuration

`func (o *Investigation) HasOpenDuration() bool`

HasOpenDuration returns a boolean if a field has been set.

### GetParentInvestigation

`func (o *Investigation) GetParentInvestigation() string`

GetParentInvestigation returns the ParentInvestigation field if non-nil, zero value otherwise.

### GetParentInvestigationOk

`func (o *Investigation) GetParentInvestigationOk() (*string, bool)`

GetParentInvestigationOk returns a tuple with the ParentInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInvestigation

`func (o *Investigation) SetParentInvestigation(v string)`

SetParentInvestigation sets ParentInvestigation field to given value.

### HasParentInvestigation

`func (o *Investigation) HasParentInvestigation() bool`

HasParentInvestigation returns a boolean if a field has been set.

### GetPersistentEntitlements

`func (o *Investigation) GetPersistentEntitlements() map[string]string`

GetPersistentEntitlements returns the PersistentEntitlements field if non-nil, zero value otherwise.

### GetPersistentEntitlementsOk

`func (o *Investigation) GetPersistentEntitlementsOk() (*map[string]string, bool)`

GetPersistentEntitlementsOk returns a tuple with the PersistentEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentEntitlements

`func (o *Investigation) SetPersistentEntitlements(v map[string]string)`

SetPersistentEntitlements sets PersistentEntitlements field to given value.

### HasPersistentEntitlements

`func (o *Investigation) HasPersistentEntitlements() bool`

HasPersistentEntitlements returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *Investigation) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *Investigation) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *Investigation) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *Investigation) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *Investigation) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *Investigation) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *Investigation) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *Investigation) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *Investigation) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *Investigation) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *Investigation) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *Investigation) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Investigation) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Investigation) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Investigation) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Investigation) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRawCategory

`func (o *Investigation) GetRawCategory() string`

GetRawCategory returns the RawCategory field if non-nil, zero value otherwise.

### GetRawCategoryOk

`func (o *Investigation) GetRawCategoryOk() (*string, bool)`

GetRawCategoryOk returns a tuple with the RawCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCategory

`func (o *Investigation) SetRawCategory(v string)`

SetRawCategory sets RawCategory field to given value.

### HasRawCategory

`func (o *Investigation) HasRawCategory() bool`

HasRawCategory returns a boolean if a field has been set.

### GetReason

`func (o *Investigation) GetReason() map[string]string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Investigation) GetReasonOk() (*map[string]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Investigation) SetReason(v map[string]string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Investigation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRoles

`func (o *Investigation) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Investigation) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Investigation) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Investigation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunStatus

`func (o *Investigation) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *Investigation) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *Investigation) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *Investigation) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Investigation) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Investigation) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Investigation) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Investigation) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSlackMirrorAutoClose

`func (o *Investigation) GetSlackMirrorAutoClose() bool`

GetSlackMirrorAutoClose returns the SlackMirrorAutoClose field if non-nil, zero value otherwise.

### GetSlackMirrorAutoCloseOk

`func (o *Investigation) GetSlackMirrorAutoCloseOk() (*bool, bool)`

GetSlackMirrorAutoCloseOk returns a tuple with the SlackMirrorAutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackMirrorAutoClose

`func (o *Investigation) SetSlackMirrorAutoClose(v bool)`

SetSlackMirrorAutoClose sets SlackMirrorAutoClose field to given value.

### HasSlackMirrorAutoClose

`func (o *Investigation) HasSlackMirrorAutoClose() bool`

HasSlackMirrorAutoClose returns a boolean if a field has been set.

### GetSlackMirrorType

`func (o *Investigation) GetSlackMirrorType() string`

GetSlackMirrorType returns the SlackMirrorType field if non-nil, zero value otherwise.

### GetSlackMirrorTypeOk

`func (o *Investigation) GetSlackMirrorTypeOk() (*string, bool)`

GetSlackMirrorTypeOk returns a tuple with the SlackMirrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackMirrorType

`func (o *Investigation) SetSlackMirrorType(v string)`

SetSlackMirrorType sets SlackMirrorType field to given value.

### HasSlackMirrorType

`func (o *Investigation) HasSlackMirrorType() bool`

HasSlackMirrorType returns a boolean if a field has been set.

### GetSortValues

`func (o *Investigation) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Investigation) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Investigation) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Investigation) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetStatus

`func (o *Investigation) GetStatus() float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Investigation) GetStatusOk() (*float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Investigation) SetStatus(v float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Investigation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystems

`func (o *Investigation) GetSystems() []System`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *Investigation) GetSystemsOk() (*[]System, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *Investigation) SetSystems(v []System)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *Investigation) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetTags

`func (o *Investigation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Investigation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Investigation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Investigation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Investigation) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Investigation) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Investigation) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *Investigation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsers

`func (o *Investigation) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Investigation) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Investigation) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Investigation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetVersion

`func (o *Investigation) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Investigation) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Investigation) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Investigation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *Investigation) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *Investigation) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *Investigation) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *Investigation) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *Investigation) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *Investigation) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *Investigation) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *Investigation) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *Investigation) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *Investigation) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *Investigation) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *Investigation) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


