# Incident

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

### NewIncident

`func NewIncident() *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardID

`func (o *Incident) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *Incident) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *Incident) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *Incident) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAccount

`func (o *Incident) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Incident) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Incident) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Incident) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActivated

`func (o *Incident) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Incident) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Incident) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *Incident) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivatingingUserId

`func (o *Incident) GetActivatingingUserId() string`

GetActivatingingUserId returns the ActivatingingUserId field if non-nil, zero value otherwise.

### GetActivatingingUserIdOk

`func (o *Incident) GetActivatingingUserIdOk() (*string, bool)`

GetActivatingingUserIdOk returns a tuple with the ActivatingingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatingingUserId

`func (o *Incident) SetActivatingingUserId(v string)`

SetActivatingingUserId sets ActivatingingUserId field to given value.

### HasActivatingingUserId

`func (o *Incident) HasActivatingingUserId() bool`

HasActivatingingUserId returns a boolean if a field has been set.

### GetAllRead

`func (o *Incident) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *Incident) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *Incident) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *Incident) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *Incident) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *Incident) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *Incident) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *Incident) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetAttachment

`func (o *Incident) GetAttachment() []Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *Incident) GetAttachmentOk() (*[]Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *Incident) SetAttachment(v []Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *Incident) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetAutime

`func (o *Incident) GetAutime() int64`

GetAutime returns the Autime field if non-nil, zero value otherwise.

### GetAutimeOk

`func (o *Incident) GetAutimeOk() (*int64, bool)`

GetAutimeOk returns a tuple with the Autime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutime

`func (o *Incident) SetAutime(v int64)`

SetAutime sets Autime field to given value.

### HasAutime

`func (o *Incident) HasAutime() bool`

HasAutime returns a boolean if a field has been set.

### GetCanvases

`func (o *Incident) GetCanvases() []string`

GetCanvases returns the Canvases field if non-nil, zero value otherwise.

### GetCanvasesOk

`func (o *Incident) GetCanvasesOk() (*[]string, bool)`

GetCanvasesOk returns a tuple with the Canvases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvases

`func (o *Incident) SetCanvases(v []string)`

SetCanvases sets Canvases field to given value.

### HasCanvases

`func (o *Incident) HasCanvases() bool`

HasCanvases returns a boolean if a field has been set.

### GetCategory

`func (o *Incident) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Incident) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Incident) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Incident) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloseNotes

`func (o *Incident) GetCloseNotes() string`

GetCloseNotes returns the CloseNotes field if non-nil, zero value otherwise.

### GetCloseNotesOk

`func (o *Incident) GetCloseNotesOk() (*string, bool)`

GetCloseNotesOk returns a tuple with the CloseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseNotes

`func (o *Incident) SetCloseNotes(v string)`

SetCloseNotes sets CloseNotes field to given value.

### HasCloseNotes

`func (o *Incident) HasCloseNotes() bool`

HasCloseNotes returns a boolean if a field has been set.

### GetCloseReason

`func (o *Incident) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *Incident) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *Incident) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *Incident) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetClosed

`func (o *Incident) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *Incident) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *Incident) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *Incident) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosingUserId

`func (o *Incident) GetClosingUserId() string`

GetClosingUserId returns the ClosingUserId field if non-nil, zero value otherwise.

### GetClosingUserIdOk

`func (o *Incident) GetClosingUserIdOk() (*string, bool)`

GetClosingUserIdOk returns a tuple with the ClosingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUserId

`func (o *Incident) SetClosingUserId(v string)`

SetClosingUserId sets ClosingUserId field to given value.

### HasClosingUserId

`func (o *Incident) HasClosingUserId() bool`

HasClosingUserId returns a boolean if a field has been set.

### GetCreated

`func (o *Incident) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Incident) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Incident) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Incident) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *Incident) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *Incident) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *Incident) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *Incident) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDbotCurrentDirtyFields

`func (o *Incident) GetDbotCurrentDirtyFields() []string`

GetDbotCurrentDirtyFields returns the DbotCurrentDirtyFields field if non-nil, zero value otherwise.

### GetDbotCurrentDirtyFieldsOk

`func (o *Incident) GetDbotCurrentDirtyFieldsOk() (*[]string, bool)`

GetDbotCurrentDirtyFieldsOk returns a tuple with the DbotCurrentDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCurrentDirtyFields

`func (o *Incident) SetDbotCurrentDirtyFields(v []string)`

SetDbotCurrentDirtyFields sets DbotCurrentDirtyFields field to given value.

### HasDbotCurrentDirtyFields

`func (o *Incident) HasDbotCurrentDirtyFields() bool`

HasDbotCurrentDirtyFields returns a boolean if a field has been set.

### GetDbotDirtyFields

`func (o *Incident) GetDbotDirtyFields() []string`

GetDbotDirtyFields returns the DbotDirtyFields field if non-nil, zero value otherwise.

### GetDbotDirtyFieldsOk

`func (o *Incident) GetDbotDirtyFieldsOk() (*[]string, bool)`

GetDbotDirtyFieldsOk returns a tuple with the DbotDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotDirtyFields

`func (o *Incident) SetDbotDirtyFields(v []string)`

SetDbotDirtyFields sets DbotDirtyFields field to given value.

### HasDbotDirtyFields

`func (o *Incident) HasDbotDirtyFields() bool`

HasDbotDirtyFields returns a boolean if a field has been set.

### GetDbotMirrorDirection

`func (o *Incident) GetDbotMirrorDirection() string`

GetDbotMirrorDirection returns the DbotMirrorDirection field if non-nil, zero value otherwise.

### GetDbotMirrorDirectionOk

`func (o *Incident) GetDbotMirrorDirectionOk() (*string, bool)`

GetDbotMirrorDirectionOk returns a tuple with the DbotMirrorDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorDirection

`func (o *Incident) SetDbotMirrorDirection(v string)`

SetDbotMirrorDirection sets DbotMirrorDirection field to given value.

### HasDbotMirrorDirection

`func (o *Incident) HasDbotMirrorDirection() bool`

HasDbotMirrorDirection returns a boolean if a field has been set.

### GetDbotMirrorId

`func (o *Incident) GetDbotMirrorId() string`

GetDbotMirrorId returns the DbotMirrorId field if non-nil, zero value otherwise.

### GetDbotMirrorIdOk

`func (o *Incident) GetDbotMirrorIdOk() (*string, bool)`

GetDbotMirrorIdOk returns a tuple with the DbotMirrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorId

`func (o *Incident) SetDbotMirrorId(v string)`

SetDbotMirrorId sets DbotMirrorId field to given value.

### HasDbotMirrorId

`func (o *Incident) HasDbotMirrorId() bool`

HasDbotMirrorId returns a boolean if a field has been set.

### GetDbotMirrorInstance

`func (o *Incident) GetDbotMirrorInstance() string`

GetDbotMirrorInstance returns the DbotMirrorInstance field if non-nil, zero value otherwise.

### GetDbotMirrorInstanceOk

`func (o *Incident) GetDbotMirrorInstanceOk() (*string, bool)`

GetDbotMirrorInstanceOk returns a tuple with the DbotMirrorInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorInstance

`func (o *Incident) SetDbotMirrorInstance(v string)`

SetDbotMirrorInstance sets DbotMirrorInstance field to given value.

### HasDbotMirrorInstance

`func (o *Incident) HasDbotMirrorInstance() bool`

HasDbotMirrorInstance returns a boolean if a field has been set.

### GetDbotMirrorLastSync

`func (o *Incident) GetDbotMirrorLastSync() time.Time`

GetDbotMirrorLastSync returns the DbotMirrorLastSync field if non-nil, zero value otherwise.

### GetDbotMirrorLastSyncOk

`func (o *Incident) GetDbotMirrorLastSyncOk() (*time.Time, bool)`

GetDbotMirrorLastSyncOk returns a tuple with the DbotMirrorLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorLastSync

`func (o *Incident) SetDbotMirrorLastSync(v time.Time)`

SetDbotMirrorLastSync sets DbotMirrorLastSync field to given value.

### HasDbotMirrorLastSync

`func (o *Incident) HasDbotMirrorLastSync() bool`

HasDbotMirrorLastSync returns a boolean if a field has been set.

### GetDbotMirrorTags

`func (o *Incident) GetDbotMirrorTags() []string`

GetDbotMirrorTags returns the DbotMirrorTags field if non-nil, zero value otherwise.

### GetDbotMirrorTagsOk

`func (o *Incident) GetDbotMirrorTagsOk() (*[]string, bool)`

GetDbotMirrorTagsOk returns a tuple with the DbotMirrorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorTags

`func (o *Incident) SetDbotMirrorTags(v []string)`

SetDbotMirrorTags sets DbotMirrorTags field to given value.

### HasDbotMirrorTags

`func (o *Incident) HasDbotMirrorTags() bool`

HasDbotMirrorTags returns a boolean if a field has been set.

### GetDetails

`func (o *Incident) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Incident) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Incident) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Incident) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDroppedCount

`func (o *Incident) GetDroppedCount() int64`

GetDroppedCount returns the DroppedCount field if non-nil, zero value otherwise.

### GetDroppedCountOk

`func (o *Incident) GetDroppedCountOk() (*int64, bool)`

GetDroppedCountOk returns a tuple with the DroppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedCount

`func (o *Incident) SetDroppedCount(v int64)`

SetDroppedCount sets DroppedCount field to given value.

### HasDroppedCount

`func (o *Incident) HasDroppedCount() bool`

HasDroppedCount returns a boolean if a field has been set.

### GetDueDate

`func (o *Incident) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Incident) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Incident) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Incident) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetFeedBased

`func (o *Incident) GetFeedBased() bool`

GetFeedBased returns the FeedBased field if non-nil, zero value otherwise.

### GetFeedBasedOk

`func (o *Incident) GetFeedBasedOk() (*bool, bool)`

GetFeedBasedOk returns a tuple with the FeedBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedBased

`func (o *Incident) SetFeedBased(v bool)`

SetFeedBased sets FeedBased field to given value.

### HasFeedBased

`func (o *Incident) HasFeedBased() bool`

HasFeedBased returns a boolean if a field has been set.

### GetHasRole

`func (o *Incident) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *Incident) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *Incident) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *Incident) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *Incident) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Incident) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Incident) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Incident) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Incident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvestigationId

`func (o *Incident) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *Incident) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *Incident) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *Incident) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsDebug

`func (o *Incident) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *Incident) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *Incident) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *Incident) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetIsPlayground

`func (o *Incident) GetIsPlayground() bool`

GetIsPlayground returns the IsPlayground field if non-nil, zero value otherwise.

### GetIsPlaygroundOk

`func (o *Incident) GetIsPlaygroundOk() (*bool, bool)`

GetIsPlaygroundOk returns a tuple with the IsPlayground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayground

`func (o *Incident) SetIsPlayground(v bool)`

SetIsPlayground sets IsPlayground field to given value.

### HasIsPlayground

`func (o *Incident) HasIsPlayground() bool`

HasIsPlayground returns a boolean if a field has been set.

### GetLabels

`func (o *Incident) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Incident) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Incident) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Incident) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastJobRunTime

`func (o *Incident) GetLastJobRunTime() time.Time`

GetLastJobRunTime returns the LastJobRunTime field if non-nil, zero value otherwise.

### GetLastJobRunTimeOk

`func (o *Incident) GetLastJobRunTimeOk() (*time.Time, bool)`

GetLastJobRunTimeOk returns a tuple with the LastJobRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastJobRunTime

`func (o *Incident) SetLastJobRunTime(v time.Time)`

SetLastJobRunTime sets LastJobRunTime field to given value.

### HasLastJobRunTime

`func (o *Incident) HasLastJobRunTime() bool`

HasLastJobRunTime returns a boolean if a field has been set.

### GetLastOpen

`func (o *Incident) GetLastOpen() time.Time`

GetLastOpen returns the LastOpen field if non-nil, zero value otherwise.

### GetLastOpenOk

`func (o *Incident) GetLastOpenOk() (*time.Time, bool)`

GetLastOpenOk returns a tuple with the LastOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpen

`func (o *Incident) SetLastOpen(v time.Time)`

SetLastOpen sets LastOpen field to given value.

### HasLastOpen

`func (o *Incident) HasLastOpen() bool`

HasLastOpen returns a boolean if a field has been set.

### GetLinkedCount

`func (o *Incident) GetLinkedCount() int64`

GetLinkedCount returns the LinkedCount field if non-nil, zero value otherwise.

### GetLinkedCountOk

`func (o *Incident) GetLinkedCountOk() (*int64, bool)`

GetLinkedCountOk returns a tuple with the LinkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCount

`func (o *Incident) SetLinkedCount(v int64)`

SetLinkedCount sets LinkedCount field to given value.

### HasLinkedCount

`func (o *Incident) HasLinkedCount() bool`

HasLinkedCount returns a boolean if a field has been set.

### GetLinkedIncidents

`func (o *Incident) GetLinkedIncidents() []string`

GetLinkedIncidents returns the LinkedIncidents field if non-nil, zero value otherwise.

### GetLinkedIncidentsOk

`func (o *Incident) GetLinkedIncidentsOk() (*[]string, bool)`

GetLinkedIncidentsOk returns a tuple with the LinkedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIncidents

`func (o *Incident) SetLinkedIncidents(v []string)`

SetLinkedIncidents sets LinkedIncidents field to given value.

### HasLinkedIncidents

`func (o *Incident) HasLinkedIncidents() bool`

HasLinkedIncidents returns a boolean if a field has been set.

### GetModified

`func (o *Incident) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Incident) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Incident) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Incident) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Incident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Incident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Incident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Incident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyTime

`func (o *Incident) GetNotifyTime() time.Time`

GetNotifyTime returns the NotifyTime field if non-nil, zero value otherwise.

### GetNotifyTimeOk

`func (o *Incident) GetNotifyTimeOk() (*time.Time, bool)`

GetNotifyTimeOk returns a tuple with the NotifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTime

`func (o *Incident) SetNotifyTime(v time.Time)`

SetNotifyTime sets NotifyTime field to given value.

### HasNotifyTime

`func (o *Incident) HasNotifyTime() bool`

HasNotifyTime returns a boolean if a field has been set.

### GetNumericId

`func (o *Incident) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Incident) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Incident) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Incident) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOccurred

`func (o *Incident) GetOccurred() time.Time`

GetOccurred returns the Occurred field if non-nil, zero value otherwise.

### GetOccurredOk

`func (o *Incident) GetOccurredOk() (*time.Time, bool)`

GetOccurredOk returns a tuple with the Occurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurred

`func (o *Incident) SetOccurred(v time.Time)`

SetOccurred sets Occurred field to given value.

### HasOccurred

`func (o *Incident) HasOccurred() bool`

HasOccurred returns a boolean if a field has been set.

### GetOpenDuration

`func (o *Incident) GetOpenDuration() int64`

GetOpenDuration returns the OpenDuration field if non-nil, zero value otherwise.

### GetOpenDurationOk

`func (o *Incident) GetOpenDurationOk() (*int64, bool)`

GetOpenDurationOk returns a tuple with the OpenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDuration

`func (o *Incident) SetOpenDuration(v int64)`

SetOpenDuration sets OpenDuration field to given value.

### HasOpenDuration

`func (o *Incident) HasOpenDuration() bool`

HasOpenDuration returns a boolean if a field has been set.

### GetOwner

`func (o *Incident) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Incident) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Incident) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Incident) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *Incident) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Incident) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Incident) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Incident) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPhase

`func (o *Incident) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Incident) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Incident) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Incident) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlaybookId

`func (o *Incident) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *Incident) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *Incident) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *Incident) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *Incident) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *Incident) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *Incident) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *Incident) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *Incident) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *Incident) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *Incident) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *Incident) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *Incident) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *Incident) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *Incident) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *Incident) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Incident) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Incident) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Incident) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Incident) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRawCategory

`func (o *Incident) GetRawCategory() string`

GetRawCategory returns the RawCategory field if non-nil, zero value otherwise.

### GetRawCategoryOk

`func (o *Incident) GetRawCategoryOk() (*string, bool)`

GetRawCategoryOk returns a tuple with the RawCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCategory

`func (o *Incident) SetRawCategory(v string)`

SetRawCategory sets RawCategory field to given value.

### HasRawCategory

`func (o *Incident) HasRawCategory() bool`

HasRawCategory returns a boolean if a field has been set.

### GetRawCloseReason

`func (o *Incident) GetRawCloseReason() string`

GetRawCloseReason returns the RawCloseReason field if non-nil, zero value otherwise.

### GetRawCloseReasonOk

`func (o *Incident) GetRawCloseReasonOk() (*string, bool)`

GetRawCloseReasonOk returns a tuple with the RawCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCloseReason

`func (o *Incident) SetRawCloseReason(v string)`

SetRawCloseReason sets RawCloseReason field to given value.

### HasRawCloseReason

`func (o *Incident) HasRawCloseReason() bool`

HasRawCloseReason returns a boolean if a field has been set.

### GetRawJSON

`func (o *Incident) GetRawJSON() string`

GetRawJSON returns the RawJSON field if non-nil, zero value otherwise.

### GetRawJSONOk

`func (o *Incident) GetRawJSONOk() (*string, bool)`

GetRawJSONOk returns a tuple with the RawJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawJSON

`func (o *Incident) SetRawJSON(v string)`

SetRawJSON sets RawJSON field to given value.

### HasRawJSON

`func (o *Incident) HasRawJSON() bool`

HasRawJSON returns a boolean if a field has been set.

### GetRawName

`func (o *Incident) GetRawName() string`

GetRawName returns the RawName field if non-nil, zero value otherwise.

### GetRawNameOk

`func (o *Incident) GetRawNameOk() (*string, bool)`

GetRawNameOk returns a tuple with the RawName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawName

`func (o *Incident) SetRawName(v string)`

SetRawName sets RawName field to given value.

### HasRawName

`func (o *Incident) HasRawName() bool`

HasRawName returns a boolean if a field has been set.

### GetRawPhase

`func (o *Incident) GetRawPhase() string`

GetRawPhase returns the RawPhase field if non-nil, zero value otherwise.

### GetRawPhaseOk

`func (o *Incident) GetRawPhaseOk() (*string, bool)`

GetRawPhaseOk returns a tuple with the RawPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPhase

`func (o *Incident) SetRawPhase(v string)`

SetRawPhase sets RawPhase field to given value.

### HasRawPhase

`func (o *Incident) HasRawPhase() bool`

HasRawPhase returns a boolean if a field has been set.

### GetRawType

`func (o *Incident) GetRawType() string`

GetRawType returns the RawType field if non-nil, zero value otherwise.

### GetRawTypeOk

`func (o *Incident) GetRawTypeOk() (*string, bool)`

GetRawTypeOk returns a tuple with the RawType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawType

`func (o *Incident) SetRawType(v string)`

SetRawType sets RawType field to given value.

### HasRawType

`func (o *Incident) HasRawType() bool`

HasRawType returns a boolean if a field has been set.

### GetReason

`func (o *Incident) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Incident) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Incident) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Incident) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReminder

`func (o *Incident) GetReminder() time.Time`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *Incident) GetReminderOk() (*time.Time, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *Incident) SetReminder(v time.Time)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *Incident) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetRoles

`func (o *Incident) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Incident) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Incident) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Incident) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunStatus

`func (o *Incident) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *Incident) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *Incident) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *Incident) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Incident) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Incident) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Incident) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Incident) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSeverity

`func (o *Incident) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Incident) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Incident) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Incident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSla

`func (o *Incident) GetSla() float64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *Incident) GetSlaOk() (*float64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *Incident) SetSla(v float64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *Incident) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSortValues

`func (o *Incident) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Incident) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Incident) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Incident) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceBrand

`func (o *Incident) GetSourceBrand() string`

GetSourceBrand returns the SourceBrand field if non-nil, zero value otherwise.

### GetSourceBrandOk

`func (o *Incident) GetSourceBrandOk() (*string, bool)`

GetSourceBrandOk returns a tuple with the SourceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrand

`func (o *Incident) SetSourceBrand(v string)`

SetSourceBrand sets SourceBrand field to given value.

### HasSourceBrand

`func (o *Incident) HasSourceBrand() bool`

HasSourceBrand returns a boolean if a field has been set.

### GetSourceInstance

`func (o *Incident) GetSourceInstance() string`

GetSourceInstance returns the SourceInstance field if non-nil, zero value otherwise.

### GetSourceInstanceOk

`func (o *Incident) GetSourceInstanceOk() (*string, bool)`

GetSourceInstanceOk returns a tuple with the SourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstance

`func (o *Incident) SetSourceInstance(v string)`

SetSourceInstance sets SourceInstance field to given value.

### HasSourceInstance

`func (o *Incident) HasSourceInstance() bool`

HasSourceInstance returns a boolean if a field has been set.

### GetStatus

`func (o *Incident) GetStatus() float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Incident) GetStatusOk() (*float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Incident) SetStatus(v float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Incident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTodoTaskIds

`func (o *Incident) GetTodoTaskIds() []string`

GetTodoTaskIds returns the TodoTaskIds field if non-nil, zero value otherwise.

### GetTodoTaskIdsOk

`func (o *Incident) GetTodoTaskIdsOk() (*[]string, bool)`

GetTodoTaskIdsOk returns a tuple with the TodoTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoTaskIds

`func (o *Incident) SetTodoTaskIds(v []string)`

SetTodoTaskIds sets TodoTaskIds field to given value.

### HasTodoTaskIds

`func (o *Incident) HasTodoTaskIds() bool`

HasTodoTaskIds returns a boolean if a field has been set.

### GetType

`func (o *Incident) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Incident) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Incident) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Incident) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *Incident) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Incident) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Incident) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Incident) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *Incident) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *Incident) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *Incident) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *Incident) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *Incident) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *Incident) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *Incident) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *Incident) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *Incident) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *Incident) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *Incident) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *Incident) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


