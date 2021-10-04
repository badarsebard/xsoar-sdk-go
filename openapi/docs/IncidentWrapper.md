# IncidentWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardID** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **string** | Account holds the tenant name so that slicing and dicing on the master can leverage bleve | [optional] 
**Activated** | Pointer to **time.Time** | When was this activated | [optional] 
**ActivatingingUserId** | Pointer to **string** | The user that activated this investigation | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Attachment** | Pointer to [**[]Attachment**](Attachment.md) | Attachments | [optional] 
**Autime** | Pointer to **int64** | AlmostUniqueTime is an attempt to have a unique sortable ID for an incident | [optional] 
**Canvases** | Pointer to **[]string** | Canvases of the incident | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**ChangeStatus** | Pointer to **string** |  | [optional] 
**CloseNotes** | Pointer to **string** | Notes for closing the incident | [optional] 
**CloseReason** | Pointer to **string** | The reason for closing the incident (select from existing predefined values) | [optional] 
**Closed** | Pointer to **time.Time** | When was this closed | [optional] 
**ClosingUserId** | Pointer to **string** | The user ID that closed this investigation | [optional] 
**Created** | Pointer to **time.Time** | When was this created | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**DbotCurrentDirtyFields** | Pointer to **[]string** | For mirroring, manage a list of current dirty fields so that we can send delta to outgoing integration | [optional] 
**DbotDirtyFields** | Pointer to **[]string** | For mirroring, manage a list of dirty fields to not override them from the source of the incident | [optional] 
**DbotMirrorDirection** | Pointer to **string** | DBotMirrorDirection of how to mirror the incident (in/out/both) | [optional] 
**DbotMirrorId** | Pointer to **string** | DBotMirrorID of a remote system we are syncing with | [optional] 
**DbotMirrorInstance** | Pointer to **string** | DBotMirrorInstance name of a mirror integration instance | [optional] 
**DbotMirrorLastSync** | Pointer to **time.Time** | The last time we synced this incident even if we did not update anything | [optional] 
**DbotMirrorTags** | Pointer to **[]string** | The entry tags I want to sync to remote system | [optional] 
**Details** | Pointer to **string** | The details of the incident - reason, etc. | [optional] 
**DroppedCount** | Pointer to **int64** | DroppedCount ... | [optional] 
**DueDate** | Pointer to **time.Time** | SLA | [optional] 
**FeedBased** | Pointer to **bool** | If this incident was triggered by a feed job | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Insights** | Pointer to **int32** |  | [optional] 
**InvestigationId** | Pointer to **string** | Investigation that was opened as a result of the incoming event | [optional] 
**IsDebug** | Pointer to **bool** | IsDebug ... | [optional] 
**IsPlayground** | Pointer to **bool** | IsPlayGround | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | Labels related to incident - each label is composed of a type and value | [optional] 
**LastJobRunTime** | Pointer to **time.Time** | If this incident was triggered by a job, this would be the time the **previous** job started | [optional] 
**LastOpen** | Pointer to **time.Time** |  | [optional] 
**LinkedCount** | Pointer to **int64** | LinkedCount ... | [optional] 
**LinkedIncidents** | Pointer to **[]string** | LinkedIncidents incidents that were marked as linked by user | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | Incident Name - given by user | [optional] 
**NotifyTime** | Pointer to **time.Time** | Incdicates when last this field was changed with a value that supposed to send a notification | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Occurred** | Pointer to **time.Time** | When this incident has really occurred | [optional] 
**OpenDuration** | Pointer to **int64** | Duration incident was open | [optional] 
**Owner** | Pointer to **string** | The user who owns this incident | [optional] 
**Parent** | Pointer to **string** | Parent | [optional] 
**Phase** | Pointer to **string** | Phase | [optional] 
**PlaybookId** | Pointer to **string** | The associated playbook for this incident | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**RawCategory** | Pointer to **string** |  | [optional] 
**RawCloseReason** | Pointer to **string** | The reason for closing the incident (select from existing predefined values) | [optional] 
**RawJSON** | Pointer to **string** |  | [optional] 
**RawName** | Pointer to **string** | Incident RawName | [optional] 
**RawPhase** | Pointer to **string** | RawPhase | [optional] 
**RawType** | Pointer to **string** | Incident raw type | [optional] 
**Reason** | Pointer to **string** | The reason for the resolve | [optional] 
**Reminder** | Pointer to **time.Time** | When if at all to send a reminder | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**RunStatus** | Pointer to **string** | RunStatus of a job | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**Severity** | Pointer to **float64** | Severity is the incident severity | [optional] 
**Sla** | Pointer to **float64** | SLAState is the incident sla at closure time | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**SourceBrand** | Pointer to **string** | SourceBrand ... | [optional] 
**SourceInstance** | Pointer to **string** | SourceInstance ... | [optional] 
**Status** | Pointer to **float64** | IncidentStatus is the status of the incident | [optional] 
**TodoTaskIds** | Pointer to **[]string** | ToDoTaskIDs list of to do task ids | [optional] 
**Type** | Pointer to **string** | Incident type | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIncidentWrapper

`func NewIncidentWrapper() *IncidentWrapper`

NewIncidentWrapper instantiates a new IncidentWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWrapperWithDefaults

`func NewIncidentWrapperWithDefaults() *IncidentWrapper`

NewIncidentWrapperWithDefaults instantiates a new IncidentWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardID

`func (o *IncidentWrapper) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *IncidentWrapper) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *IncidentWrapper) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *IncidentWrapper) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAccount

`func (o *IncidentWrapper) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IncidentWrapper) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IncidentWrapper) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IncidentWrapper) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActivated

`func (o *IncidentWrapper) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *IncidentWrapper) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *IncidentWrapper) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *IncidentWrapper) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivatingingUserId

`func (o *IncidentWrapper) GetActivatingingUserId() string`

GetActivatingingUserId returns the ActivatingingUserId field if non-nil, zero value otherwise.

### GetActivatingingUserIdOk

`func (o *IncidentWrapper) GetActivatingingUserIdOk() (*string, bool)`

GetActivatingingUserIdOk returns a tuple with the ActivatingingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatingingUserId

`func (o *IncidentWrapper) SetActivatingingUserId(v string)`

SetActivatingingUserId sets ActivatingingUserId field to given value.

### HasActivatingingUserId

`func (o *IncidentWrapper) HasActivatingingUserId() bool`

HasActivatingingUserId returns a boolean if a field has been set.

### GetAllRead

`func (o *IncidentWrapper) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *IncidentWrapper) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *IncidentWrapper) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *IncidentWrapper) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *IncidentWrapper) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *IncidentWrapper) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *IncidentWrapper) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *IncidentWrapper) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetAttachment

`func (o *IncidentWrapper) GetAttachment() []Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *IncidentWrapper) GetAttachmentOk() (*[]Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *IncidentWrapper) SetAttachment(v []Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *IncidentWrapper) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetAutime

`func (o *IncidentWrapper) GetAutime() int64`

GetAutime returns the Autime field if non-nil, zero value otherwise.

### GetAutimeOk

`func (o *IncidentWrapper) GetAutimeOk() (*int64, bool)`

GetAutimeOk returns a tuple with the Autime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutime

`func (o *IncidentWrapper) SetAutime(v int64)`

SetAutime sets Autime field to given value.

### HasAutime

`func (o *IncidentWrapper) HasAutime() bool`

HasAutime returns a boolean if a field has been set.

### GetCanvases

`func (o *IncidentWrapper) GetCanvases() []string`

GetCanvases returns the Canvases field if non-nil, zero value otherwise.

### GetCanvasesOk

`func (o *IncidentWrapper) GetCanvasesOk() (*[]string, bool)`

GetCanvasesOk returns a tuple with the Canvases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvases

`func (o *IncidentWrapper) SetCanvases(v []string)`

SetCanvases sets Canvases field to given value.

### HasCanvases

`func (o *IncidentWrapper) HasCanvases() bool`

HasCanvases returns a boolean if a field has been set.

### GetCategory

`func (o *IncidentWrapper) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IncidentWrapper) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IncidentWrapper) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IncidentWrapper) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChangeStatus

`func (o *IncidentWrapper) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *IncidentWrapper) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *IncidentWrapper) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.

### HasChangeStatus

`func (o *IncidentWrapper) HasChangeStatus() bool`

HasChangeStatus returns a boolean if a field has been set.

### GetCloseNotes

`func (o *IncidentWrapper) GetCloseNotes() string`

GetCloseNotes returns the CloseNotes field if non-nil, zero value otherwise.

### GetCloseNotesOk

`func (o *IncidentWrapper) GetCloseNotesOk() (*string, bool)`

GetCloseNotesOk returns a tuple with the CloseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseNotes

`func (o *IncidentWrapper) SetCloseNotes(v string)`

SetCloseNotes sets CloseNotes field to given value.

### HasCloseNotes

`func (o *IncidentWrapper) HasCloseNotes() bool`

HasCloseNotes returns a boolean if a field has been set.

### GetCloseReason

`func (o *IncidentWrapper) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *IncidentWrapper) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *IncidentWrapper) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *IncidentWrapper) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetClosed

`func (o *IncidentWrapper) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *IncidentWrapper) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *IncidentWrapper) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *IncidentWrapper) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosingUserId

`func (o *IncidentWrapper) GetClosingUserId() string`

GetClosingUserId returns the ClosingUserId field if non-nil, zero value otherwise.

### GetClosingUserIdOk

`func (o *IncidentWrapper) GetClosingUserIdOk() (*string, bool)`

GetClosingUserIdOk returns a tuple with the ClosingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUserId

`func (o *IncidentWrapper) SetClosingUserId(v string)`

SetClosingUserId sets ClosingUserId field to given value.

### HasClosingUserId

`func (o *IncidentWrapper) HasClosingUserId() bool`

HasClosingUserId returns a boolean if a field has been set.

### GetCreated

`func (o *IncidentWrapper) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentWrapper) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentWrapper) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentWrapper) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *IncidentWrapper) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *IncidentWrapper) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *IncidentWrapper) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *IncidentWrapper) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDbotCurrentDirtyFields

`func (o *IncidentWrapper) GetDbotCurrentDirtyFields() []string`

GetDbotCurrentDirtyFields returns the DbotCurrentDirtyFields field if non-nil, zero value otherwise.

### GetDbotCurrentDirtyFieldsOk

`func (o *IncidentWrapper) GetDbotCurrentDirtyFieldsOk() (*[]string, bool)`

GetDbotCurrentDirtyFieldsOk returns a tuple with the DbotCurrentDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCurrentDirtyFields

`func (o *IncidentWrapper) SetDbotCurrentDirtyFields(v []string)`

SetDbotCurrentDirtyFields sets DbotCurrentDirtyFields field to given value.

### HasDbotCurrentDirtyFields

`func (o *IncidentWrapper) HasDbotCurrentDirtyFields() bool`

HasDbotCurrentDirtyFields returns a boolean if a field has been set.

### GetDbotDirtyFields

`func (o *IncidentWrapper) GetDbotDirtyFields() []string`

GetDbotDirtyFields returns the DbotDirtyFields field if non-nil, zero value otherwise.

### GetDbotDirtyFieldsOk

`func (o *IncidentWrapper) GetDbotDirtyFieldsOk() (*[]string, bool)`

GetDbotDirtyFieldsOk returns a tuple with the DbotDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotDirtyFields

`func (o *IncidentWrapper) SetDbotDirtyFields(v []string)`

SetDbotDirtyFields sets DbotDirtyFields field to given value.

### HasDbotDirtyFields

`func (o *IncidentWrapper) HasDbotDirtyFields() bool`

HasDbotDirtyFields returns a boolean if a field has been set.

### GetDbotMirrorDirection

`func (o *IncidentWrapper) GetDbotMirrorDirection() string`

GetDbotMirrorDirection returns the DbotMirrorDirection field if non-nil, zero value otherwise.

### GetDbotMirrorDirectionOk

`func (o *IncidentWrapper) GetDbotMirrorDirectionOk() (*string, bool)`

GetDbotMirrorDirectionOk returns a tuple with the DbotMirrorDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorDirection

`func (o *IncidentWrapper) SetDbotMirrorDirection(v string)`

SetDbotMirrorDirection sets DbotMirrorDirection field to given value.

### HasDbotMirrorDirection

`func (o *IncidentWrapper) HasDbotMirrorDirection() bool`

HasDbotMirrorDirection returns a boolean if a field has been set.

### GetDbotMirrorId

`func (o *IncidentWrapper) GetDbotMirrorId() string`

GetDbotMirrorId returns the DbotMirrorId field if non-nil, zero value otherwise.

### GetDbotMirrorIdOk

`func (o *IncidentWrapper) GetDbotMirrorIdOk() (*string, bool)`

GetDbotMirrorIdOk returns a tuple with the DbotMirrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorId

`func (o *IncidentWrapper) SetDbotMirrorId(v string)`

SetDbotMirrorId sets DbotMirrorId field to given value.

### HasDbotMirrorId

`func (o *IncidentWrapper) HasDbotMirrorId() bool`

HasDbotMirrorId returns a boolean if a field has been set.

### GetDbotMirrorInstance

`func (o *IncidentWrapper) GetDbotMirrorInstance() string`

GetDbotMirrorInstance returns the DbotMirrorInstance field if non-nil, zero value otherwise.

### GetDbotMirrorInstanceOk

`func (o *IncidentWrapper) GetDbotMirrorInstanceOk() (*string, bool)`

GetDbotMirrorInstanceOk returns a tuple with the DbotMirrorInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorInstance

`func (o *IncidentWrapper) SetDbotMirrorInstance(v string)`

SetDbotMirrorInstance sets DbotMirrorInstance field to given value.

### HasDbotMirrorInstance

`func (o *IncidentWrapper) HasDbotMirrorInstance() bool`

HasDbotMirrorInstance returns a boolean if a field has been set.

### GetDbotMirrorLastSync

`func (o *IncidentWrapper) GetDbotMirrorLastSync() time.Time`

GetDbotMirrorLastSync returns the DbotMirrorLastSync field if non-nil, zero value otherwise.

### GetDbotMirrorLastSyncOk

`func (o *IncidentWrapper) GetDbotMirrorLastSyncOk() (*time.Time, bool)`

GetDbotMirrorLastSyncOk returns a tuple with the DbotMirrorLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorLastSync

`func (o *IncidentWrapper) SetDbotMirrorLastSync(v time.Time)`

SetDbotMirrorLastSync sets DbotMirrorLastSync field to given value.

### HasDbotMirrorLastSync

`func (o *IncidentWrapper) HasDbotMirrorLastSync() bool`

HasDbotMirrorLastSync returns a boolean if a field has been set.

### GetDbotMirrorTags

`func (o *IncidentWrapper) GetDbotMirrorTags() []string`

GetDbotMirrorTags returns the DbotMirrorTags field if non-nil, zero value otherwise.

### GetDbotMirrorTagsOk

`func (o *IncidentWrapper) GetDbotMirrorTagsOk() (*[]string, bool)`

GetDbotMirrorTagsOk returns a tuple with the DbotMirrorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorTags

`func (o *IncidentWrapper) SetDbotMirrorTags(v []string)`

SetDbotMirrorTags sets DbotMirrorTags field to given value.

### HasDbotMirrorTags

`func (o *IncidentWrapper) HasDbotMirrorTags() bool`

HasDbotMirrorTags returns a boolean if a field has been set.

### GetDetails

`func (o *IncidentWrapper) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IncidentWrapper) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IncidentWrapper) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IncidentWrapper) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDroppedCount

`func (o *IncidentWrapper) GetDroppedCount() int64`

GetDroppedCount returns the DroppedCount field if non-nil, zero value otherwise.

### GetDroppedCountOk

`func (o *IncidentWrapper) GetDroppedCountOk() (*int64, bool)`

GetDroppedCountOk returns a tuple with the DroppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedCount

`func (o *IncidentWrapper) SetDroppedCount(v int64)`

SetDroppedCount sets DroppedCount field to given value.

### HasDroppedCount

`func (o *IncidentWrapper) HasDroppedCount() bool`

HasDroppedCount returns a boolean if a field has been set.

### GetDueDate

`func (o *IncidentWrapper) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *IncidentWrapper) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *IncidentWrapper) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *IncidentWrapper) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetFeedBased

`func (o *IncidentWrapper) GetFeedBased() bool`

GetFeedBased returns the FeedBased field if non-nil, zero value otherwise.

### GetFeedBasedOk

`func (o *IncidentWrapper) GetFeedBasedOk() (*bool, bool)`

GetFeedBasedOk returns a tuple with the FeedBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedBased

`func (o *IncidentWrapper) SetFeedBased(v bool)`

SetFeedBased sets FeedBased field to given value.

### HasFeedBased

`func (o *IncidentWrapper) HasFeedBased() bool`

HasFeedBased returns a boolean if a field has been set.

### GetHasRole

`func (o *IncidentWrapper) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *IncidentWrapper) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *IncidentWrapper) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *IncidentWrapper) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *IncidentWrapper) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *IncidentWrapper) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *IncidentWrapper) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *IncidentWrapper) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *IncidentWrapper) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentWrapper) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentWrapper) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentWrapper) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsights

`func (o *IncidentWrapper) GetInsights() int32`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *IncidentWrapper) GetInsightsOk() (*int32, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *IncidentWrapper) SetInsights(v int32)`

SetInsights sets Insights field to given value.

### HasInsights

`func (o *IncidentWrapper) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### GetInvestigationId

`func (o *IncidentWrapper) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *IncidentWrapper) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *IncidentWrapper) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *IncidentWrapper) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsDebug

`func (o *IncidentWrapper) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *IncidentWrapper) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *IncidentWrapper) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *IncidentWrapper) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetIsPlayground

`func (o *IncidentWrapper) GetIsPlayground() bool`

GetIsPlayground returns the IsPlayground field if non-nil, zero value otherwise.

### GetIsPlaygroundOk

`func (o *IncidentWrapper) GetIsPlaygroundOk() (*bool, bool)`

GetIsPlaygroundOk returns a tuple with the IsPlayground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayground

`func (o *IncidentWrapper) SetIsPlayground(v bool)`

SetIsPlayground sets IsPlayground field to given value.

### HasIsPlayground

`func (o *IncidentWrapper) HasIsPlayground() bool`

HasIsPlayground returns a boolean if a field has been set.

### GetLabels

`func (o *IncidentWrapper) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IncidentWrapper) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IncidentWrapper) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IncidentWrapper) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastJobRunTime

`func (o *IncidentWrapper) GetLastJobRunTime() time.Time`

GetLastJobRunTime returns the LastJobRunTime field if non-nil, zero value otherwise.

### GetLastJobRunTimeOk

`func (o *IncidentWrapper) GetLastJobRunTimeOk() (*time.Time, bool)`

GetLastJobRunTimeOk returns a tuple with the LastJobRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastJobRunTime

`func (o *IncidentWrapper) SetLastJobRunTime(v time.Time)`

SetLastJobRunTime sets LastJobRunTime field to given value.

### HasLastJobRunTime

`func (o *IncidentWrapper) HasLastJobRunTime() bool`

HasLastJobRunTime returns a boolean if a field has been set.

### GetLastOpen

`func (o *IncidentWrapper) GetLastOpen() time.Time`

GetLastOpen returns the LastOpen field if non-nil, zero value otherwise.

### GetLastOpenOk

`func (o *IncidentWrapper) GetLastOpenOk() (*time.Time, bool)`

GetLastOpenOk returns a tuple with the LastOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpen

`func (o *IncidentWrapper) SetLastOpen(v time.Time)`

SetLastOpen sets LastOpen field to given value.

### HasLastOpen

`func (o *IncidentWrapper) HasLastOpen() bool`

HasLastOpen returns a boolean if a field has been set.

### GetLinkedCount

`func (o *IncidentWrapper) GetLinkedCount() int64`

GetLinkedCount returns the LinkedCount field if non-nil, zero value otherwise.

### GetLinkedCountOk

`func (o *IncidentWrapper) GetLinkedCountOk() (*int64, bool)`

GetLinkedCountOk returns a tuple with the LinkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCount

`func (o *IncidentWrapper) SetLinkedCount(v int64)`

SetLinkedCount sets LinkedCount field to given value.

### HasLinkedCount

`func (o *IncidentWrapper) HasLinkedCount() bool`

HasLinkedCount returns a boolean if a field has been set.

### GetLinkedIncidents

`func (o *IncidentWrapper) GetLinkedIncidents() []string`

GetLinkedIncidents returns the LinkedIncidents field if non-nil, zero value otherwise.

### GetLinkedIncidentsOk

`func (o *IncidentWrapper) GetLinkedIncidentsOk() (*[]string, bool)`

GetLinkedIncidentsOk returns a tuple with the LinkedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIncidents

`func (o *IncidentWrapper) SetLinkedIncidents(v []string)`

SetLinkedIncidents sets LinkedIncidents field to given value.

### HasLinkedIncidents

`func (o *IncidentWrapper) HasLinkedIncidents() bool`

HasLinkedIncidents returns a boolean if a field has been set.

### GetModified

`func (o *IncidentWrapper) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentWrapper) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentWrapper) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentWrapper) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentWrapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentWrapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentWrapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentWrapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyTime

`func (o *IncidentWrapper) GetNotifyTime() time.Time`

GetNotifyTime returns the NotifyTime field if non-nil, zero value otherwise.

### GetNotifyTimeOk

`func (o *IncidentWrapper) GetNotifyTimeOk() (*time.Time, bool)`

GetNotifyTimeOk returns a tuple with the NotifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTime

`func (o *IncidentWrapper) SetNotifyTime(v time.Time)`

SetNotifyTime sets NotifyTime field to given value.

### HasNotifyTime

`func (o *IncidentWrapper) HasNotifyTime() bool`

HasNotifyTime returns a boolean if a field has been set.

### GetNumericId

`func (o *IncidentWrapper) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *IncidentWrapper) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *IncidentWrapper) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *IncidentWrapper) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOccurred

`func (o *IncidentWrapper) GetOccurred() time.Time`

GetOccurred returns the Occurred field if non-nil, zero value otherwise.

### GetOccurredOk

`func (o *IncidentWrapper) GetOccurredOk() (*time.Time, bool)`

GetOccurredOk returns a tuple with the Occurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurred

`func (o *IncidentWrapper) SetOccurred(v time.Time)`

SetOccurred sets Occurred field to given value.

### HasOccurred

`func (o *IncidentWrapper) HasOccurred() bool`

HasOccurred returns a boolean if a field has been set.

### GetOpenDuration

`func (o *IncidentWrapper) GetOpenDuration() int64`

GetOpenDuration returns the OpenDuration field if non-nil, zero value otherwise.

### GetOpenDurationOk

`func (o *IncidentWrapper) GetOpenDurationOk() (*int64, bool)`

GetOpenDurationOk returns a tuple with the OpenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDuration

`func (o *IncidentWrapper) SetOpenDuration(v int64)`

SetOpenDuration sets OpenDuration field to given value.

### HasOpenDuration

`func (o *IncidentWrapper) HasOpenDuration() bool`

HasOpenDuration returns a boolean if a field has been set.

### GetOwner

`func (o *IncidentWrapper) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IncidentWrapper) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IncidentWrapper) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IncidentWrapper) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *IncidentWrapper) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IncidentWrapper) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IncidentWrapper) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IncidentWrapper) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPhase

`func (o *IncidentWrapper) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IncidentWrapper) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IncidentWrapper) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IncidentWrapper) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlaybookId

`func (o *IncidentWrapper) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *IncidentWrapper) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *IncidentWrapper) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *IncidentWrapper) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *IncidentWrapper) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *IncidentWrapper) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *IncidentWrapper) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *IncidentWrapper) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *IncidentWrapper) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *IncidentWrapper) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *IncidentWrapper) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *IncidentWrapper) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *IncidentWrapper) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *IncidentWrapper) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *IncidentWrapper) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *IncidentWrapper) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *IncidentWrapper) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *IncidentWrapper) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *IncidentWrapper) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *IncidentWrapper) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRawCategory

`func (o *IncidentWrapper) GetRawCategory() string`

GetRawCategory returns the RawCategory field if non-nil, zero value otherwise.

### GetRawCategoryOk

`func (o *IncidentWrapper) GetRawCategoryOk() (*string, bool)`

GetRawCategoryOk returns a tuple with the RawCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCategory

`func (o *IncidentWrapper) SetRawCategory(v string)`

SetRawCategory sets RawCategory field to given value.

### HasRawCategory

`func (o *IncidentWrapper) HasRawCategory() bool`

HasRawCategory returns a boolean if a field has been set.

### GetRawCloseReason

`func (o *IncidentWrapper) GetRawCloseReason() string`

GetRawCloseReason returns the RawCloseReason field if non-nil, zero value otherwise.

### GetRawCloseReasonOk

`func (o *IncidentWrapper) GetRawCloseReasonOk() (*string, bool)`

GetRawCloseReasonOk returns a tuple with the RawCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCloseReason

`func (o *IncidentWrapper) SetRawCloseReason(v string)`

SetRawCloseReason sets RawCloseReason field to given value.

### HasRawCloseReason

`func (o *IncidentWrapper) HasRawCloseReason() bool`

HasRawCloseReason returns a boolean if a field has been set.

### GetRawJSON

`func (o *IncidentWrapper) GetRawJSON() string`

GetRawJSON returns the RawJSON field if non-nil, zero value otherwise.

### GetRawJSONOk

`func (o *IncidentWrapper) GetRawJSONOk() (*string, bool)`

GetRawJSONOk returns a tuple with the RawJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawJSON

`func (o *IncidentWrapper) SetRawJSON(v string)`

SetRawJSON sets RawJSON field to given value.

### HasRawJSON

`func (o *IncidentWrapper) HasRawJSON() bool`

HasRawJSON returns a boolean if a field has been set.

### GetRawName

`func (o *IncidentWrapper) GetRawName() string`

GetRawName returns the RawName field if non-nil, zero value otherwise.

### GetRawNameOk

`func (o *IncidentWrapper) GetRawNameOk() (*string, bool)`

GetRawNameOk returns a tuple with the RawName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawName

`func (o *IncidentWrapper) SetRawName(v string)`

SetRawName sets RawName field to given value.

### HasRawName

`func (o *IncidentWrapper) HasRawName() bool`

HasRawName returns a boolean if a field has been set.

### GetRawPhase

`func (o *IncidentWrapper) GetRawPhase() string`

GetRawPhase returns the RawPhase field if non-nil, zero value otherwise.

### GetRawPhaseOk

`func (o *IncidentWrapper) GetRawPhaseOk() (*string, bool)`

GetRawPhaseOk returns a tuple with the RawPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPhase

`func (o *IncidentWrapper) SetRawPhase(v string)`

SetRawPhase sets RawPhase field to given value.

### HasRawPhase

`func (o *IncidentWrapper) HasRawPhase() bool`

HasRawPhase returns a boolean if a field has been set.

### GetRawType

`func (o *IncidentWrapper) GetRawType() string`

GetRawType returns the RawType field if non-nil, zero value otherwise.

### GetRawTypeOk

`func (o *IncidentWrapper) GetRawTypeOk() (*string, bool)`

GetRawTypeOk returns a tuple with the RawType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawType

`func (o *IncidentWrapper) SetRawType(v string)`

SetRawType sets RawType field to given value.

### HasRawType

`func (o *IncidentWrapper) HasRawType() bool`

HasRawType returns a boolean if a field has been set.

### GetReason

`func (o *IncidentWrapper) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IncidentWrapper) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IncidentWrapper) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IncidentWrapper) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReminder

`func (o *IncidentWrapper) GetReminder() time.Time`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *IncidentWrapper) GetReminderOk() (*time.Time, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *IncidentWrapper) SetReminder(v time.Time)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *IncidentWrapper) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetRoles

`func (o *IncidentWrapper) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IncidentWrapper) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IncidentWrapper) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IncidentWrapper) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunStatus

`func (o *IncidentWrapper) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *IncidentWrapper) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *IncidentWrapper) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *IncidentWrapper) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *IncidentWrapper) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *IncidentWrapper) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *IncidentWrapper) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *IncidentWrapper) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSeverity

`func (o *IncidentWrapper) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentWrapper) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentWrapper) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IncidentWrapper) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSla

`func (o *IncidentWrapper) GetSla() float64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *IncidentWrapper) GetSlaOk() (*float64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *IncidentWrapper) SetSla(v float64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *IncidentWrapper) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSortValues

`func (o *IncidentWrapper) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *IncidentWrapper) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *IncidentWrapper) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *IncidentWrapper) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceBrand

`func (o *IncidentWrapper) GetSourceBrand() string`

GetSourceBrand returns the SourceBrand field if non-nil, zero value otherwise.

### GetSourceBrandOk

`func (o *IncidentWrapper) GetSourceBrandOk() (*string, bool)`

GetSourceBrandOk returns a tuple with the SourceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrand

`func (o *IncidentWrapper) SetSourceBrand(v string)`

SetSourceBrand sets SourceBrand field to given value.

### HasSourceBrand

`func (o *IncidentWrapper) HasSourceBrand() bool`

HasSourceBrand returns a boolean if a field has been set.

### GetSourceInstance

`func (o *IncidentWrapper) GetSourceInstance() string`

GetSourceInstance returns the SourceInstance field if non-nil, zero value otherwise.

### GetSourceInstanceOk

`func (o *IncidentWrapper) GetSourceInstanceOk() (*string, bool)`

GetSourceInstanceOk returns a tuple with the SourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstance

`func (o *IncidentWrapper) SetSourceInstance(v string)`

SetSourceInstance sets SourceInstance field to given value.

### HasSourceInstance

`func (o *IncidentWrapper) HasSourceInstance() bool`

HasSourceInstance returns a boolean if a field has been set.

### GetStatus

`func (o *IncidentWrapper) GetStatus() float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentWrapper) GetStatusOk() (*float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentWrapper) SetStatus(v float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTodoTaskIds

`func (o *IncidentWrapper) GetTodoTaskIds() []string`

GetTodoTaskIds returns the TodoTaskIds field if non-nil, zero value otherwise.

### GetTodoTaskIdsOk

`func (o *IncidentWrapper) GetTodoTaskIdsOk() (*[]string, bool)`

GetTodoTaskIdsOk returns a tuple with the TodoTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoTaskIds

`func (o *IncidentWrapper) SetTodoTaskIds(v []string)`

SetTodoTaskIds sets TodoTaskIds field to given value.

### HasTodoTaskIds

`func (o *IncidentWrapper) HasTodoTaskIds() bool`

HasTodoTaskIds returns a boolean if a field has been set.

### GetType

`func (o *IncidentWrapper) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentWrapper) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentWrapper) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentWrapper) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *IncidentWrapper) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IncidentWrapper) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IncidentWrapper) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IncidentWrapper) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *IncidentWrapper) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *IncidentWrapper) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *IncidentWrapper) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *IncidentWrapper) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *IncidentWrapper) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *IncidentWrapper) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *IncidentWrapper) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *IncidentWrapper) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *IncidentWrapper) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *IncidentWrapper) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *IncidentWrapper) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *IncidentWrapper) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


