# CreateIncidentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardID** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **string** | Account holds the tenant name so that slicing and dicing on the master can leverage bleve | [optional] 
**Activated** | Pointer to **time.Time** | When was this activated | [optional] 
**ActivatingingUserId** | Pointer to **string** | The user that activated this investigation | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Autime** | Pointer to **int64** | AlmostUniqueTime is an attempt to have a unique sortable ID for an incident | [optional] 
**Canvases** | Pointer to **[]string** | Canvases of the incident | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**CloseNotes** | Pointer to **string** | Notes for closing the incident | [optional] 
**CloseReason** | Pointer to **string** | The reason for closing the incident (select from existing predefined values) | [optional] 
**Closed** | Pointer to **time.Time** | When was this closed | [optional] 
**ClosingUserId** | Pointer to **string** | The user ID that closed this investigation | [optional] 
**CreateInvestigation** | Pointer to **bool** |  | [optional] 
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

### NewCreateIncidentRequest

`func NewCreateIncidentRequest() *CreateIncidentRequest`

NewCreateIncidentRequest instantiates a new CreateIncidentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIncidentRequestWithDefaults

`func NewCreateIncidentRequestWithDefaults() *CreateIncidentRequest`

NewCreateIncidentRequestWithDefaults instantiates a new CreateIncidentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardID

`func (o *CreateIncidentRequest) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *CreateIncidentRequest) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *CreateIncidentRequest) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *CreateIncidentRequest) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAccount

`func (o *CreateIncidentRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateIncidentRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateIncidentRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateIncidentRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActivated

`func (o *CreateIncidentRequest) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *CreateIncidentRequest) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *CreateIncidentRequest) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *CreateIncidentRequest) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivatingingUserId

`func (o *CreateIncidentRequest) GetActivatingingUserId() string`

GetActivatingingUserId returns the ActivatingingUserId field if non-nil, zero value otherwise.

### GetActivatingingUserIdOk

`func (o *CreateIncidentRequest) GetActivatingingUserIdOk() (*string, bool)`

GetActivatingingUserIdOk returns a tuple with the ActivatingingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatingingUserId

`func (o *CreateIncidentRequest) SetActivatingingUserId(v string)`

SetActivatingingUserId sets ActivatingingUserId field to given value.

### HasActivatingingUserId

`func (o *CreateIncidentRequest) HasActivatingingUserId() bool`

HasActivatingingUserId returns a boolean if a field has been set.

### GetAllRead

`func (o *CreateIncidentRequest) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *CreateIncidentRequest) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *CreateIncidentRequest) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *CreateIncidentRequest) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *CreateIncidentRequest) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *CreateIncidentRequest) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *CreateIncidentRequest) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *CreateIncidentRequest) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetAutime

`func (o *CreateIncidentRequest) GetAutime() int64`

GetAutime returns the Autime field if non-nil, zero value otherwise.

### GetAutimeOk

`func (o *CreateIncidentRequest) GetAutimeOk() (*int64, bool)`

GetAutimeOk returns a tuple with the Autime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutime

`func (o *CreateIncidentRequest) SetAutime(v int64)`

SetAutime sets Autime field to given value.

### HasAutime

`func (o *CreateIncidentRequest) HasAutime() bool`

HasAutime returns a boolean if a field has been set.

### GetCanvases

`func (o *CreateIncidentRequest) GetCanvases() []string`

GetCanvases returns the Canvases field if non-nil, zero value otherwise.

### GetCanvasesOk

`func (o *CreateIncidentRequest) GetCanvasesOk() (*[]string, bool)`

GetCanvasesOk returns a tuple with the Canvases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvases

`func (o *CreateIncidentRequest) SetCanvases(v []string)`

SetCanvases sets Canvases field to given value.

### HasCanvases

`func (o *CreateIncidentRequest) HasCanvases() bool`

HasCanvases returns a boolean if a field has been set.

### GetCategory

`func (o *CreateIncidentRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateIncidentRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateIncidentRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateIncidentRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloseNotes

`func (o *CreateIncidentRequest) GetCloseNotes() string`

GetCloseNotes returns the CloseNotes field if non-nil, zero value otherwise.

### GetCloseNotesOk

`func (o *CreateIncidentRequest) GetCloseNotesOk() (*string, bool)`

GetCloseNotesOk returns a tuple with the CloseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseNotes

`func (o *CreateIncidentRequest) SetCloseNotes(v string)`

SetCloseNotes sets CloseNotes field to given value.

### HasCloseNotes

`func (o *CreateIncidentRequest) HasCloseNotes() bool`

HasCloseNotes returns a boolean if a field has been set.

### GetCloseReason

`func (o *CreateIncidentRequest) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *CreateIncidentRequest) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *CreateIncidentRequest) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *CreateIncidentRequest) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetClosed

`func (o *CreateIncidentRequest) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *CreateIncidentRequest) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *CreateIncidentRequest) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *CreateIncidentRequest) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosingUserId

`func (o *CreateIncidentRequest) GetClosingUserId() string`

GetClosingUserId returns the ClosingUserId field if non-nil, zero value otherwise.

### GetClosingUserIdOk

`func (o *CreateIncidentRequest) GetClosingUserIdOk() (*string, bool)`

GetClosingUserIdOk returns a tuple with the ClosingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUserId

`func (o *CreateIncidentRequest) SetClosingUserId(v string)`

SetClosingUserId sets ClosingUserId field to given value.

### HasClosingUserId

`func (o *CreateIncidentRequest) HasClosingUserId() bool`

HasClosingUserId returns a boolean if a field has been set.

### GetCreateInvestigation

`func (o *CreateIncidentRequest) GetCreateInvestigation() bool`

GetCreateInvestigation returns the CreateInvestigation field if non-nil, zero value otherwise.

### GetCreateInvestigationOk

`func (o *CreateIncidentRequest) GetCreateInvestigationOk() (*bool, bool)`

GetCreateInvestigationOk returns a tuple with the CreateInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInvestigation

`func (o *CreateIncidentRequest) SetCreateInvestigation(v bool)`

SetCreateInvestigation sets CreateInvestigation field to given value.

### HasCreateInvestigation

`func (o *CreateIncidentRequest) HasCreateInvestigation() bool`

HasCreateInvestigation returns a boolean if a field has been set.

### GetCreated

`func (o *CreateIncidentRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateIncidentRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateIncidentRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateIncidentRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *CreateIncidentRequest) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *CreateIncidentRequest) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *CreateIncidentRequest) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *CreateIncidentRequest) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDbotCurrentDirtyFields

`func (o *CreateIncidentRequest) GetDbotCurrentDirtyFields() []string`

GetDbotCurrentDirtyFields returns the DbotCurrentDirtyFields field if non-nil, zero value otherwise.

### GetDbotCurrentDirtyFieldsOk

`func (o *CreateIncidentRequest) GetDbotCurrentDirtyFieldsOk() (*[]string, bool)`

GetDbotCurrentDirtyFieldsOk returns a tuple with the DbotCurrentDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCurrentDirtyFields

`func (o *CreateIncidentRequest) SetDbotCurrentDirtyFields(v []string)`

SetDbotCurrentDirtyFields sets DbotCurrentDirtyFields field to given value.

### HasDbotCurrentDirtyFields

`func (o *CreateIncidentRequest) HasDbotCurrentDirtyFields() bool`

HasDbotCurrentDirtyFields returns a boolean if a field has been set.

### GetDbotDirtyFields

`func (o *CreateIncidentRequest) GetDbotDirtyFields() []string`

GetDbotDirtyFields returns the DbotDirtyFields field if non-nil, zero value otherwise.

### GetDbotDirtyFieldsOk

`func (o *CreateIncidentRequest) GetDbotDirtyFieldsOk() (*[]string, bool)`

GetDbotDirtyFieldsOk returns a tuple with the DbotDirtyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotDirtyFields

`func (o *CreateIncidentRequest) SetDbotDirtyFields(v []string)`

SetDbotDirtyFields sets DbotDirtyFields field to given value.

### HasDbotDirtyFields

`func (o *CreateIncidentRequest) HasDbotDirtyFields() bool`

HasDbotDirtyFields returns a boolean if a field has been set.

### GetDbotMirrorDirection

`func (o *CreateIncidentRequest) GetDbotMirrorDirection() string`

GetDbotMirrorDirection returns the DbotMirrorDirection field if non-nil, zero value otherwise.

### GetDbotMirrorDirectionOk

`func (o *CreateIncidentRequest) GetDbotMirrorDirectionOk() (*string, bool)`

GetDbotMirrorDirectionOk returns a tuple with the DbotMirrorDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorDirection

`func (o *CreateIncidentRequest) SetDbotMirrorDirection(v string)`

SetDbotMirrorDirection sets DbotMirrorDirection field to given value.

### HasDbotMirrorDirection

`func (o *CreateIncidentRequest) HasDbotMirrorDirection() bool`

HasDbotMirrorDirection returns a boolean if a field has been set.

### GetDbotMirrorId

`func (o *CreateIncidentRequest) GetDbotMirrorId() string`

GetDbotMirrorId returns the DbotMirrorId field if non-nil, zero value otherwise.

### GetDbotMirrorIdOk

`func (o *CreateIncidentRequest) GetDbotMirrorIdOk() (*string, bool)`

GetDbotMirrorIdOk returns a tuple with the DbotMirrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorId

`func (o *CreateIncidentRequest) SetDbotMirrorId(v string)`

SetDbotMirrorId sets DbotMirrorId field to given value.

### HasDbotMirrorId

`func (o *CreateIncidentRequest) HasDbotMirrorId() bool`

HasDbotMirrorId returns a boolean if a field has been set.

### GetDbotMirrorInstance

`func (o *CreateIncidentRequest) GetDbotMirrorInstance() string`

GetDbotMirrorInstance returns the DbotMirrorInstance field if non-nil, zero value otherwise.

### GetDbotMirrorInstanceOk

`func (o *CreateIncidentRequest) GetDbotMirrorInstanceOk() (*string, bool)`

GetDbotMirrorInstanceOk returns a tuple with the DbotMirrorInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorInstance

`func (o *CreateIncidentRequest) SetDbotMirrorInstance(v string)`

SetDbotMirrorInstance sets DbotMirrorInstance field to given value.

### HasDbotMirrorInstance

`func (o *CreateIncidentRequest) HasDbotMirrorInstance() bool`

HasDbotMirrorInstance returns a boolean if a field has been set.

### GetDbotMirrorLastSync

`func (o *CreateIncidentRequest) GetDbotMirrorLastSync() time.Time`

GetDbotMirrorLastSync returns the DbotMirrorLastSync field if non-nil, zero value otherwise.

### GetDbotMirrorLastSyncOk

`func (o *CreateIncidentRequest) GetDbotMirrorLastSyncOk() (*time.Time, bool)`

GetDbotMirrorLastSyncOk returns a tuple with the DbotMirrorLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorLastSync

`func (o *CreateIncidentRequest) SetDbotMirrorLastSync(v time.Time)`

SetDbotMirrorLastSync sets DbotMirrorLastSync field to given value.

### HasDbotMirrorLastSync

`func (o *CreateIncidentRequest) HasDbotMirrorLastSync() bool`

HasDbotMirrorLastSync returns a boolean if a field has been set.

### GetDbotMirrorTags

`func (o *CreateIncidentRequest) GetDbotMirrorTags() []string`

GetDbotMirrorTags returns the DbotMirrorTags field if non-nil, zero value otherwise.

### GetDbotMirrorTagsOk

`func (o *CreateIncidentRequest) GetDbotMirrorTagsOk() (*[]string, bool)`

GetDbotMirrorTagsOk returns a tuple with the DbotMirrorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotMirrorTags

`func (o *CreateIncidentRequest) SetDbotMirrorTags(v []string)`

SetDbotMirrorTags sets DbotMirrorTags field to given value.

### HasDbotMirrorTags

`func (o *CreateIncidentRequest) HasDbotMirrorTags() bool`

HasDbotMirrorTags returns a boolean if a field has been set.

### GetDetails

`func (o *CreateIncidentRequest) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreateIncidentRequest) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreateIncidentRequest) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CreateIncidentRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDroppedCount

`func (o *CreateIncidentRequest) GetDroppedCount() int64`

GetDroppedCount returns the DroppedCount field if non-nil, zero value otherwise.

### GetDroppedCountOk

`func (o *CreateIncidentRequest) GetDroppedCountOk() (*int64, bool)`

GetDroppedCountOk returns a tuple with the DroppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedCount

`func (o *CreateIncidentRequest) SetDroppedCount(v int64)`

SetDroppedCount sets DroppedCount field to given value.

### HasDroppedCount

`func (o *CreateIncidentRequest) HasDroppedCount() bool`

HasDroppedCount returns a boolean if a field has been set.

### GetDueDate

`func (o *CreateIncidentRequest) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreateIncidentRequest) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreateIncidentRequest) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CreateIncidentRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetFeedBased

`func (o *CreateIncidentRequest) GetFeedBased() bool`

GetFeedBased returns the FeedBased field if non-nil, zero value otherwise.

### GetFeedBasedOk

`func (o *CreateIncidentRequest) GetFeedBasedOk() (*bool, bool)`

GetFeedBasedOk returns a tuple with the FeedBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedBased

`func (o *CreateIncidentRequest) SetFeedBased(v bool)`

SetFeedBased sets FeedBased field to given value.

### HasFeedBased

`func (o *CreateIncidentRequest) HasFeedBased() bool`

HasFeedBased returns a boolean if a field has been set.

### GetHasRole

`func (o *CreateIncidentRequest) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *CreateIncidentRequest) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *CreateIncidentRequest) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *CreateIncidentRequest) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *CreateIncidentRequest) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *CreateIncidentRequest) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *CreateIncidentRequest) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *CreateIncidentRequest) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *CreateIncidentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIncidentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIncidentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIncidentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvestigationId

`func (o *CreateIncidentRequest) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *CreateIncidentRequest) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *CreateIncidentRequest) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *CreateIncidentRequest) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsDebug

`func (o *CreateIncidentRequest) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *CreateIncidentRequest) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *CreateIncidentRequest) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *CreateIncidentRequest) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetIsPlayground

`func (o *CreateIncidentRequest) GetIsPlayground() bool`

GetIsPlayground returns the IsPlayground field if non-nil, zero value otherwise.

### GetIsPlaygroundOk

`func (o *CreateIncidentRequest) GetIsPlaygroundOk() (*bool, bool)`

GetIsPlaygroundOk returns a tuple with the IsPlayground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayground

`func (o *CreateIncidentRequest) SetIsPlayground(v bool)`

SetIsPlayground sets IsPlayground field to given value.

### HasIsPlayground

`func (o *CreateIncidentRequest) HasIsPlayground() bool`

HasIsPlayground returns a boolean if a field has been set.

### GetLabels

`func (o *CreateIncidentRequest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateIncidentRequest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateIncidentRequest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateIncidentRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastJobRunTime

`func (o *CreateIncidentRequest) GetLastJobRunTime() time.Time`

GetLastJobRunTime returns the LastJobRunTime field if non-nil, zero value otherwise.

### GetLastJobRunTimeOk

`func (o *CreateIncidentRequest) GetLastJobRunTimeOk() (*time.Time, bool)`

GetLastJobRunTimeOk returns a tuple with the LastJobRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastJobRunTime

`func (o *CreateIncidentRequest) SetLastJobRunTime(v time.Time)`

SetLastJobRunTime sets LastJobRunTime field to given value.

### HasLastJobRunTime

`func (o *CreateIncidentRequest) HasLastJobRunTime() bool`

HasLastJobRunTime returns a boolean if a field has been set.

### GetLastOpen

`func (o *CreateIncidentRequest) GetLastOpen() time.Time`

GetLastOpen returns the LastOpen field if non-nil, zero value otherwise.

### GetLastOpenOk

`func (o *CreateIncidentRequest) GetLastOpenOk() (*time.Time, bool)`

GetLastOpenOk returns a tuple with the LastOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpen

`func (o *CreateIncidentRequest) SetLastOpen(v time.Time)`

SetLastOpen sets LastOpen field to given value.

### HasLastOpen

`func (o *CreateIncidentRequest) HasLastOpen() bool`

HasLastOpen returns a boolean if a field has been set.

### GetLinkedCount

`func (o *CreateIncidentRequest) GetLinkedCount() int64`

GetLinkedCount returns the LinkedCount field if non-nil, zero value otherwise.

### GetLinkedCountOk

`func (o *CreateIncidentRequest) GetLinkedCountOk() (*int64, bool)`

GetLinkedCountOk returns a tuple with the LinkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCount

`func (o *CreateIncidentRequest) SetLinkedCount(v int64)`

SetLinkedCount sets LinkedCount field to given value.

### HasLinkedCount

`func (o *CreateIncidentRequest) HasLinkedCount() bool`

HasLinkedCount returns a boolean if a field has been set.

### GetLinkedIncidents

`func (o *CreateIncidentRequest) GetLinkedIncidents() []string`

GetLinkedIncidents returns the LinkedIncidents field if non-nil, zero value otherwise.

### GetLinkedIncidentsOk

`func (o *CreateIncidentRequest) GetLinkedIncidentsOk() (*[]string, bool)`

GetLinkedIncidentsOk returns a tuple with the LinkedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIncidents

`func (o *CreateIncidentRequest) SetLinkedIncidents(v []string)`

SetLinkedIncidents sets LinkedIncidents field to given value.

### HasLinkedIncidents

`func (o *CreateIncidentRequest) HasLinkedIncidents() bool`

HasLinkedIncidents returns a boolean if a field has been set.

### GetModified

`func (o *CreateIncidentRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CreateIncidentRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CreateIncidentRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CreateIncidentRequest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *CreateIncidentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIncidentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIncidentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIncidentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyTime

`func (o *CreateIncidentRequest) GetNotifyTime() time.Time`

GetNotifyTime returns the NotifyTime field if non-nil, zero value otherwise.

### GetNotifyTimeOk

`func (o *CreateIncidentRequest) GetNotifyTimeOk() (*time.Time, bool)`

GetNotifyTimeOk returns a tuple with the NotifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTime

`func (o *CreateIncidentRequest) SetNotifyTime(v time.Time)`

SetNotifyTime sets NotifyTime field to given value.

### HasNotifyTime

`func (o *CreateIncidentRequest) HasNotifyTime() bool`

HasNotifyTime returns a boolean if a field has been set.

### GetNumericId

`func (o *CreateIncidentRequest) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *CreateIncidentRequest) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *CreateIncidentRequest) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *CreateIncidentRequest) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOccurred

`func (o *CreateIncidentRequest) GetOccurred() time.Time`

GetOccurred returns the Occurred field if non-nil, zero value otherwise.

### GetOccurredOk

`func (o *CreateIncidentRequest) GetOccurredOk() (*time.Time, bool)`

GetOccurredOk returns a tuple with the Occurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurred

`func (o *CreateIncidentRequest) SetOccurred(v time.Time)`

SetOccurred sets Occurred field to given value.

### HasOccurred

`func (o *CreateIncidentRequest) HasOccurred() bool`

HasOccurred returns a boolean if a field has been set.

### GetOpenDuration

`func (o *CreateIncidentRequest) GetOpenDuration() int64`

GetOpenDuration returns the OpenDuration field if non-nil, zero value otherwise.

### GetOpenDurationOk

`func (o *CreateIncidentRequest) GetOpenDurationOk() (*int64, bool)`

GetOpenDurationOk returns a tuple with the OpenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDuration

`func (o *CreateIncidentRequest) SetOpenDuration(v int64)`

SetOpenDuration sets OpenDuration field to given value.

### HasOpenDuration

`func (o *CreateIncidentRequest) HasOpenDuration() bool`

HasOpenDuration returns a boolean if a field has been set.

### GetOwner

`func (o *CreateIncidentRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateIncidentRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateIncidentRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateIncidentRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *CreateIncidentRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CreateIncidentRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CreateIncidentRequest) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CreateIncidentRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPhase

`func (o *CreateIncidentRequest) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CreateIncidentRequest) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CreateIncidentRequest) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CreateIncidentRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlaybookId

`func (o *CreateIncidentRequest) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *CreateIncidentRequest) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *CreateIncidentRequest) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *CreateIncidentRequest) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *CreateIncidentRequest) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *CreateIncidentRequest) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *CreateIncidentRequest) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *CreateIncidentRequest) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *CreateIncidentRequest) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *CreateIncidentRequest) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *CreateIncidentRequest) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *CreateIncidentRequest) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *CreateIncidentRequest) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *CreateIncidentRequest) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *CreateIncidentRequest) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *CreateIncidentRequest) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *CreateIncidentRequest) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *CreateIncidentRequest) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *CreateIncidentRequest) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *CreateIncidentRequest) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRawCategory

`func (o *CreateIncidentRequest) GetRawCategory() string`

GetRawCategory returns the RawCategory field if non-nil, zero value otherwise.

### GetRawCategoryOk

`func (o *CreateIncidentRequest) GetRawCategoryOk() (*string, bool)`

GetRawCategoryOk returns a tuple with the RawCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCategory

`func (o *CreateIncidentRequest) SetRawCategory(v string)`

SetRawCategory sets RawCategory field to given value.

### HasRawCategory

`func (o *CreateIncidentRequest) HasRawCategory() bool`

HasRawCategory returns a boolean if a field has been set.

### GetRawCloseReason

`func (o *CreateIncidentRequest) GetRawCloseReason() string`

GetRawCloseReason returns the RawCloseReason field if non-nil, zero value otherwise.

### GetRawCloseReasonOk

`func (o *CreateIncidentRequest) GetRawCloseReasonOk() (*string, bool)`

GetRawCloseReasonOk returns a tuple with the RawCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCloseReason

`func (o *CreateIncidentRequest) SetRawCloseReason(v string)`

SetRawCloseReason sets RawCloseReason field to given value.

### HasRawCloseReason

`func (o *CreateIncidentRequest) HasRawCloseReason() bool`

HasRawCloseReason returns a boolean if a field has been set.

### GetRawJSON

`func (o *CreateIncidentRequest) GetRawJSON() string`

GetRawJSON returns the RawJSON field if non-nil, zero value otherwise.

### GetRawJSONOk

`func (o *CreateIncidentRequest) GetRawJSONOk() (*string, bool)`

GetRawJSONOk returns a tuple with the RawJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawJSON

`func (o *CreateIncidentRequest) SetRawJSON(v string)`

SetRawJSON sets RawJSON field to given value.

### HasRawJSON

`func (o *CreateIncidentRequest) HasRawJSON() bool`

HasRawJSON returns a boolean if a field has been set.

### GetRawName

`func (o *CreateIncidentRequest) GetRawName() string`

GetRawName returns the RawName field if non-nil, zero value otherwise.

### GetRawNameOk

`func (o *CreateIncidentRequest) GetRawNameOk() (*string, bool)`

GetRawNameOk returns a tuple with the RawName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawName

`func (o *CreateIncidentRequest) SetRawName(v string)`

SetRawName sets RawName field to given value.

### HasRawName

`func (o *CreateIncidentRequest) HasRawName() bool`

HasRawName returns a boolean if a field has been set.

### GetRawPhase

`func (o *CreateIncidentRequest) GetRawPhase() string`

GetRawPhase returns the RawPhase field if non-nil, zero value otherwise.

### GetRawPhaseOk

`func (o *CreateIncidentRequest) GetRawPhaseOk() (*string, bool)`

GetRawPhaseOk returns a tuple with the RawPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPhase

`func (o *CreateIncidentRequest) SetRawPhase(v string)`

SetRawPhase sets RawPhase field to given value.

### HasRawPhase

`func (o *CreateIncidentRequest) HasRawPhase() bool`

HasRawPhase returns a boolean if a field has been set.

### GetRawType

`func (o *CreateIncidentRequest) GetRawType() string`

GetRawType returns the RawType field if non-nil, zero value otherwise.

### GetRawTypeOk

`func (o *CreateIncidentRequest) GetRawTypeOk() (*string, bool)`

GetRawTypeOk returns a tuple with the RawType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawType

`func (o *CreateIncidentRequest) SetRawType(v string)`

SetRawType sets RawType field to given value.

### HasRawType

`func (o *CreateIncidentRequest) HasRawType() bool`

HasRawType returns a boolean if a field has been set.

### GetReason

`func (o *CreateIncidentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateIncidentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateIncidentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateIncidentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReminder

`func (o *CreateIncidentRequest) GetReminder() time.Time`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *CreateIncidentRequest) GetReminderOk() (*time.Time, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *CreateIncidentRequest) SetReminder(v time.Time)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *CreateIncidentRequest) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetRoles

`func (o *CreateIncidentRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateIncidentRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateIncidentRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateIncidentRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunStatus

`func (o *CreateIncidentRequest) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *CreateIncidentRequest) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *CreateIncidentRequest) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *CreateIncidentRequest) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *CreateIncidentRequest) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *CreateIncidentRequest) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *CreateIncidentRequest) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *CreateIncidentRequest) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSeverity

`func (o *CreateIncidentRequest) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateIncidentRequest) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateIncidentRequest) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CreateIncidentRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSla

`func (o *CreateIncidentRequest) GetSla() float64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CreateIncidentRequest) GetSlaOk() (*float64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CreateIncidentRequest) SetSla(v float64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CreateIncidentRequest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSortValues

`func (o *CreateIncidentRequest) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *CreateIncidentRequest) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *CreateIncidentRequest) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *CreateIncidentRequest) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceBrand

`func (o *CreateIncidentRequest) GetSourceBrand() string`

GetSourceBrand returns the SourceBrand field if non-nil, zero value otherwise.

### GetSourceBrandOk

`func (o *CreateIncidentRequest) GetSourceBrandOk() (*string, bool)`

GetSourceBrandOk returns a tuple with the SourceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrand

`func (o *CreateIncidentRequest) SetSourceBrand(v string)`

SetSourceBrand sets SourceBrand field to given value.

### HasSourceBrand

`func (o *CreateIncidentRequest) HasSourceBrand() bool`

HasSourceBrand returns a boolean if a field has been set.

### GetSourceInstance

`func (o *CreateIncidentRequest) GetSourceInstance() string`

GetSourceInstance returns the SourceInstance field if non-nil, zero value otherwise.

### GetSourceInstanceOk

`func (o *CreateIncidentRequest) GetSourceInstanceOk() (*string, bool)`

GetSourceInstanceOk returns a tuple with the SourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstance

`func (o *CreateIncidentRequest) SetSourceInstance(v string)`

SetSourceInstance sets SourceInstance field to given value.

### HasSourceInstance

`func (o *CreateIncidentRequest) HasSourceInstance() bool`

HasSourceInstance returns a boolean if a field has been set.

### GetStatus

`func (o *CreateIncidentRequest) GetStatus() float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateIncidentRequest) GetStatusOk() (*float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateIncidentRequest) SetStatus(v float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateIncidentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTodoTaskIds

`func (o *CreateIncidentRequest) GetTodoTaskIds() []string`

GetTodoTaskIds returns the TodoTaskIds field if non-nil, zero value otherwise.

### GetTodoTaskIdsOk

`func (o *CreateIncidentRequest) GetTodoTaskIdsOk() (*[]string, bool)`

GetTodoTaskIdsOk returns a tuple with the TodoTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoTaskIds

`func (o *CreateIncidentRequest) SetTodoTaskIds(v []string)`

SetTodoTaskIds sets TodoTaskIds field to given value.

### HasTodoTaskIds

`func (o *CreateIncidentRequest) HasTodoTaskIds() bool`

HasTodoTaskIds returns a boolean if a field has been set.

### GetType

`func (o *CreateIncidentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIncidentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIncidentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateIncidentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CreateIncidentRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateIncidentRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateIncidentRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateIncidentRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *CreateIncidentRequest) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *CreateIncidentRequest) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *CreateIncidentRequest) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *CreateIncidentRequest) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *CreateIncidentRequest) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *CreateIncidentRequest) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *CreateIncidentRequest) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *CreateIncidentRequest) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *CreateIncidentRequest) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *CreateIncidentRequest) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *CreateIncidentRequest) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *CreateIncidentRequest) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


