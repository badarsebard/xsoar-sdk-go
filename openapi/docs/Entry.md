# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndicatorTimeline** | Pointer to [**[]IndicatorTimelineFromEntry**](IndicatorTimelineFromEntry.md) |  | [optional] 
**InstanceID** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]RelationshipAPI**](RelationshipAPI.md) |  | [optional] 
**ShardID** | Pointer to **int64** |  | [optional] 
**AllRead** | Pointer to **bool** |  | [optional] 
**AllReadWrite** | Pointer to **bool** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to **map[string]interface{}** | The contents of the entry that is actually indexed - should not be used | [optional] 
**ContentsSize** | Pointer to **int64** | ContentsSize the total size of the contents | [optional] 
**Created** | Pointer to **time.Time** | When it was taken | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**CronView** | Pointer to **bool** |  | [optional] 
**DbotCreatedBy** | Pointer to **string** | Who has created this event - relevant only for manual incidents | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedBy** | Pointer to **string** |  | [optional] 
**DeletedFromFS** | Pointer to **bool** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**EndingType** | Pointer to **string** | EndingType holds the type of schedule Ending | [optional] 
**EntryTask** | Pointer to [**EntryTask**](EntryTask.md) |  | [optional] 
**ErrorSource** | Pointer to **string** | Source of the error | [optional] 
**File** | Pointer to **string** | Filename of associated content | [optional] 
**FileID** | Pointer to **string** | FileID is the file name when saved in the server | [optional] 
**FileMetadata** | Pointer to [**FileMetadata**](FileMetadata.md) |  | [optional] 
**Format** | Pointer to **string** | Holds information on how content is formatted | [optional] 
**HasRole** | Pointer to **bool** | Internal field to make queries on role faster | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**History** | Pointer to [**[]EntryHistory**](EntryHistory.md) | Edit history | [optional] 
**HumanCron** | Pointer to [**HumanCron**](HumanCron.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncidentCreationTime** | Pointer to **time.Time** | store the entry based on IncidentCreationTime | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**InvestigationId** | Pointer to **string** | The id of the investigation it belongs to | [optional] 
**IsTodo** | Pointer to **bool** | IsTodo | [optional] 
**Mirrored** | Pointer to **bool** | Only used for outbound mirroring to mark that it is already mirrored to remote system | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Note** | Pointer to **bool** | Note | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**ParentContent** | Pointer to **map[string]interface{}** | ParentEntry content - for reference | [optional] 
**ParentEntryTruncated** | Pointer to **bool** | ParentEntryTruncated - indicates weather entry content was truncated | [optional] 
**ParentId** | Pointer to **string** | ParentId is the ID of the parent entry | [optional] 
**Pinned** | Pointer to **bool** | Mark entry as pinned &#x3D; evidence | [optional] 
**PlaybookId** | Pointer to **string** | PlaybookID - if the entry is assigned as note to a playbook task, it will hold the playbook | [optional] 
**Polling** | Pointer to **bool** | Only used for polling entries | [optional] 
**PollingArgs** | Pointer to **map[string]map[string]interface{}** | ModuleArgs represents module args | [optional] 
**PollingCommand** | Pointer to **string** |  | [optional] 
**PreviousAllRead** | Pointer to **bool** |  | [optional] 
**PreviousAllReadWrite** | Pointer to **bool** |  | [optional] 
**PreviousRoles** | Pointer to **[]string** | Do not change this field manually | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**ReadOnly** | Pointer to **bool** | ReadOnly | [optional] 
**Recurrent** | Pointer to **bool** |  | [optional] 
**ReputationSize** | Pointer to **int64** | ReputationSize the total size of the reputation | [optional] 
**Reputations** | Pointer to [**[]EntryReputation**](EntryReputation.md) | EntryReputations the reputations calculated by regex match | [optional] 
**RetryTime** | Pointer to **time.Time** | When retry took place | [optional] 
**Roles** | Pointer to **[]string** | The role assigned to this investigation | [optional] 
**Scheduled** | Pointer to **bool** | is it scheduled | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**System** | Pointer to **string** | The name of the system associated with this entry | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**TagsRaw** | Pointer to **[]string** | TagsRaw | [optional] 
**TaskId** | Pointer to **string** | TaskID - used if the entry is assigned as note to a playbook task | [optional] 
**Times** | Pointer to **int64** |  | [optional] 
**TimesRan** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**TimezoneOffset** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **float64** | EntryType specifies the type of the entry | [optional] 
**User** | Pointer to **string** | The user who created  the entry | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**XsoarHasReadOnlyRole** | Pointer to **bool** |  | [optional] 
**XsoarPreviousReadOnlyRoles** | Pointer to **[]string** |  | [optional] 
**XsoarReadOnlyRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicatorTimeline

`func (o *Entry) GetIndicatorTimeline() []IndicatorTimelineFromEntry`

GetIndicatorTimeline returns the IndicatorTimeline field if non-nil, zero value otherwise.

### GetIndicatorTimelineOk

`func (o *Entry) GetIndicatorTimelineOk() (*[]IndicatorTimelineFromEntry, bool)`

GetIndicatorTimelineOk returns a tuple with the IndicatorTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorTimeline

`func (o *Entry) SetIndicatorTimeline(v []IndicatorTimelineFromEntry)`

SetIndicatorTimeline sets IndicatorTimeline field to given value.

### HasIndicatorTimeline

`func (o *Entry) HasIndicatorTimeline() bool`

HasIndicatorTimeline returns a boolean if a field has been set.

### GetInstanceID

`func (o *Entry) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *Entry) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *Entry) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *Entry) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetRelationships

`func (o *Entry) GetRelationships() []RelationshipAPI`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Entry) GetRelationshipsOk() (*[]RelationshipAPI, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Entry) SetRelationships(v []RelationshipAPI)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Entry) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetShardID

`func (o *Entry) GetShardID() int64`

GetShardID returns the ShardID field if non-nil, zero value otherwise.

### GetShardIDOk

`func (o *Entry) GetShardIDOk() (*int64, bool)`

GetShardIDOk returns a tuple with the ShardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardID

`func (o *Entry) SetShardID(v int64)`

SetShardID sets ShardID field to given value.

### HasShardID

`func (o *Entry) HasShardID() bool`

HasShardID returns a boolean if a field has been set.

### GetAllRead

`func (o *Entry) GetAllRead() bool`

GetAllRead returns the AllRead field if non-nil, zero value otherwise.

### GetAllReadOk

`func (o *Entry) GetAllReadOk() (*bool, bool)`

GetAllReadOk returns a tuple with the AllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRead

`func (o *Entry) SetAllRead(v bool)`

SetAllRead sets AllRead field to given value.

### HasAllRead

`func (o *Entry) HasAllRead() bool`

HasAllRead returns a boolean if a field has been set.

### GetAllReadWrite

`func (o *Entry) GetAllReadWrite() bool`

GetAllReadWrite returns the AllReadWrite field if non-nil, zero value otherwise.

### GetAllReadWriteOk

`func (o *Entry) GetAllReadWriteOk() (*bool, bool)`

GetAllReadWriteOk returns a tuple with the AllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReadWrite

`func (o *Entry) SetAllReadWrite(v bool)`

SetAllReadWrite sets AllReadWrite field to given value.

### HasAllReadWrite

`func (o *Entry) HasAllReadWrite() bool`

HasAllReadWrite returns a boolean if a field has been set.

### GetBrand

`func (o *Entry) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Entry) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Entry) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Entry) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCategory

`func (o *Entry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Entry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Entry) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Entry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContents

`func (o *Entry) GetContents() map[string]interface{}`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Entry) GetContentsOk() (*map[string]interface{}, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Entry) SetContents(v map[string]interface{})`

SetContents sets Contents field to given value.

### HasContents

`func (o *Entry) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetContentsSize

`func (o *Entry) GetContentsSize() int64`

GetContentsSize returns the ContentsSize field if non-nil, zero value otherwise.

### GetContentsSizeOk

`func (o *Entry) GetContentsSizeOk() (*int64, bool)`

GetContentsSizeOk returns a tuple with the ContentsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsSize

`func (o *Entry) SetContentsSize(v int64)`

SetContentsSize sets ContentsSize field to given value.

### HasContentsSize

`func (o *Entry) HasContentsSize() bool`

HasContentsSize returns a boolean if a field has been set.

### GetCreated

`func (o *Entry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Entry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Entry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Entry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCron

`func (o *Entry) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *Entry) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *Entry) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *Entry) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetCronView

`func (o *Entry) GetCronView() bool`

GetCronView returns the CronView field if non-nil, zero value otherwise.

### GetCronViewOk

`func (o *Entry) GetCronViewOk() (*bool, bool)`

GetCronViewOk returns a tuple with the CronView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronView

`func (o *Entry) SetCronView(v bool)`

SetCronView sets CronView field to given value.

### HasCronView

`func (o *Entry) HasCronView() bool`

HasCronView returns a boolean if a field has been set.

### GetDbotCreatedBy

`func (o *Entry) GetDbotCreatedBy() string`

GetDbotCreatedBy returns the DbotCreatedBy field if non-nil, zero value otherwise.

### GetDbotCreatedByOk

`func (o *Entry) GetDbotCreatedByOk() (*string, bool)`

GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbotCreatedBy

`func (o *Entry) SetDbotCreatedBy(v string)`

SetDbotCreatedBy sets DbotCreatedBy field to given value.

### HasDbotCreatedBy

`func (o *Entry) HasDbotCreatedBy() bool`

HasDbotCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *Entry) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Entry) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Entry) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Entry) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Entry) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Entry) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Entry) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Entry) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedFromFS

`func (o *Entry) GetDeletedFromFS() bool`

GetDeletedFromFS returns the DeletedFromFS field if non-nil, zero value otherwise.

### GetDeletedFromFSOk

`func (o *Entry) GetDeletedFromFSOk() (*bool, bool)`

GetDeletedFromFSOk returns a tuple with the DeletedFromFS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFromFS

`func (o *Entry) SetDeletedFromFS(v bool)`

SetDeletedFromFS sets DeletedFromFS field to given value.

### HasDeletedFromFS

`func (o *Entry) HasDeletedFromFS() bool`

HasDeletedFromFS returns a boolean if a field has been set.

### GetEndingDate

`func (o *Entry) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *Entry) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *Entry) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *Entry) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetEndingType

`func (o *Entry) GetEndingType() string`

GetEndingType returns the EndingType field if non-nil, zero value otherwise.

### GetEndingTypeOk

`func (o *Entry) GetEndingTypeOk() (*string, bool)`

GetEndingTypeOk returns a tuple with the EndingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingType

`func (o *Entry) SetEndingType(v string)`

SetEndingType sets EndingType field to given value.

### HasEndingType

`func (o *Entry) HasEndingType() bool`

HasEndingType returns a boolean if a field has been set.

### GetEntryTask

`func (o *Entry) GetEntryTask() EntryTask`

GetEntryTask returns the EntryTask field if non-nil, zero value otherwise.

### GetEntryTaskOk

`func (o *Entry) GetEntryTaskOk() (*EntryTask, bool)`

GetEntryTaskOk returns a tuple with the EntryTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTask

`func (o *Entry) SetEntryTask(v EntryTask)`

SetEntryTask sets EntryTask field to given value.

### HasEntryTask

`func (o *Entry) HasEntryTask() bool`

HasEntryTask returns a boolean if a field has been set.

### GetErrorSource

`func (o *Entry) GetErrorSource() string`

GetErrorSource returns the ErrorSource field if non-nil, zero value otherwise.

### GetErrorSourceOk

`func (o *Entry) GetErrorSourceOk() (*string, bool)`

GetErrorSourceOk returns a tuple with the ErrorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSource

`func (o *Entry) SetErrorSource(v string)`

SetErrorSource sets ErrorSource field to given value.

### HasErrorSource

`func (o *Entry) HasErrorSource() bool`

HasErrorSource returns a boolean if a field has been set.

### GetFile

`func (o *Entry) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Entry) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Entry) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *Entry) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileID

`func (o *Entry) GetFileID() string`

GetFileID returns the FileID field if non-nil, zero value otherwise.

### GetFileIDOk

`func (o *Entry) GetFileIDOk() (*string, bool)`

GetFileIDOk returns a tuple with the FileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileID

`func (o *Entry) SetFileID(v string)`

SetFileID sets FileID field to given value.

### HasFileID

`func (o *Entry) HasFileID() bool`

HasFileID returns a boolean if a field has been set.

### GetFileMetadata

`func (o *Entry) GetFileMetadata() FileMetadata`

GetFileMetadata returns the FileMetadata field if non-nil, zero value otherwise.

### GetFileMetadataOk

`func (o *Entry) GetFileMetadataOk() (*FileMetadata, bool)`

GetFileMetadataOk returns a tuple with the FileMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMetadata

`func (o *Entry) SetFileMetadata(v FileMetadata)`

SetFileMetadata sets FileMetadata field to given value.

### HasFileMetadata

`func (o *Entry) HasFileMetadata() bool`

HasFileMetadata returns a boolean if a field has been set.

### GetFormat

`func (o *Entry) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Entry) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Entry) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Entry) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHasRole

`func (o *Entry) GetHasRole() bool`

GetHasRole returns the HasRole field if non-nil, zero value otherwise.

### GetHasRoleOk

`func (o *Entry) GetHasRoleOk() (*bool, bool)`

GetHasRoleOk returns a tuple with the HasRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRole

`func (o *Entry) SetHasRole(v bool)`

SetHasRole sets HasRole field to given value.

### HasHasRole

`func (o *Entry) HasHasRole() bool`

HasHasRole returns a boolean if a field has been set.

### GetHighlight

`func (o *Entry) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Entry) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Entry) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Entry) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetHistory

`func (o *Entry) GetHistory() []EntryHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *Entry) GetHistoryOk() (*[]EntryHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *Entry) SetHistory(v []EntryHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *Entry) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetHumanCron

`func (o *Entry) GetHumanCron() HumanCron`

GetHumanCron returns the HumanCron field if non-nil, zero value otherwise.

### GetHumanCronOk

`func (o *Entry) GetHumanCronOk() (*HumanCron, bool)`

GetHumanCronOk returns a tuple with the HumanCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanCron

`func (o *Entry) SetHumanCron(v HumanCron)`

SetHumanCron sets HumanCron field to given value.

### HasHumanCron

`func (o *Entry) HasHumanCron() bool`

HasHumanCron returns a boolean if a field has been set.

### GetId

`func (o *Entry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentCreationTime

`func (o *Entry) GetIncidentCreationTime() time.Time`

GetIncidentCreationTime returns the IncidentCreationTime field if non-nil, zero value otherwise.

### GetIncidentCreationTimeOk

`func (o *Entry) GetIncidentCreationTimeOk() (*time.Time, bool)`

GetIncidentCreationTimeOk returns a tuple with the IncidentCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCreationTime

`func (o *Entry) SetIncidentCreationTime(v time.Time)`

SetIncidentCreationTime sets IncidentCreationTime field to given value.

### HasIncidentCreationTime

`func (o *Entry) HasIncidentCreationTime() bool`

HasIncidentCreationTime returns a boolean if a field has been set.

### GetInstance

`func (o *Entry) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Entry) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Entry) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Entry) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetInvestigationId

`func (o *Entry) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *Entry) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *Entry) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *Entry) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetIsTodo

`func (o *Entry) GetIsTodo() bool`

GetIsTodo returns the IsTodo field if non-nil, zero value otherwise.

### GetIsTodoOk

`func (o *Entry) GetIsTodoOk() (*bool, bool)`

GetIsTodoOk returns a tuple with the IsTodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTodo

`func (o *Entry) SetIsTodo(v bool)`

SetIsTodo sets IsTodo field to given value.

### HasIsTodo

`func (o *Entry) HasIsTodo() bool`

HasIsTodo returns a boolean if a field has been set.

### GetMirrored

`func (o *Entry) GetMirrored() bool`

GetMirrored returns the Mirrored field if non-nil, zero value otherwise.

### GetMirroredOk

`func (o *Entry) GetMirroredOk() (*bool, bool)`

GetMirroredOk returns a tuple with the Mirrored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrored

`func (o *Entry) SetMirrored(v bool)`

SetMirrored sets Mirrored field to given value.

### HasMirrored

`func (o *Entry) HasMirrored() bool`

HasMirrored returns a boolean if a field has been set.

### GetModified

`func (o *Entry) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Entry) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Entry) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Entry) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNote

`func (o *Entry) GetNote() bool`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Entry) GetNoteOk() (*bool, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Entry) SetNote(v bool)`

SetNote sets Note field to given value.

### HasNote

`func (o *Entry) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNumericId

`func (o *Entry) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Entry) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Entry) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Entry) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetParentContent

`func (o *Entry) GetParentContent() map[string]interface{}`

GetParentContent returns the ParentContent field if non-nil, zero value otherwise.

### GetParentContentOk

`func (o *Entry) GetParentContentOk() (*map[string]interface{}, bool)`

GetParentContentOk returns a tuple with the ParentContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContent

`func (o *Entry) SetParentContent(v map[string]interface{})`

SetParentContent sets ParentContent field to given value.

### HasParentContent

`func (o *Entry) HasParentContent() bool`

HasParentContent returns a boolean if a field has been set.

### GetParentEntryTruncated

`func (o *Entry) GetParentEntryTruncated() bool`

GetParentEntryTruncated returns the ParentEntryTruncated field if non-nil, zero value otherwise.

### GetParentEntryTruncatedOk

`func (o *Entry) GetParentEntryTruncatedOk() (*bool, bool)`

GetParentEntryTruncatedOk returns a tuple with the ParentEntryTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntryTruncated

`func (o *Entry) SetParentEntryTruncated(v bool)`

SetParentEntryTruncated sets ParentEntryTruncated field to given value.

### HasParentEntryTruncated

`func (o *Entry) HasParentEntryTruncated() bool`

HasParentEntryTruncated returns a boolean if a field has been set.

### GetParentId

`func (o *Entry) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Entry) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Entry) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Entry) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPinned

`func (o *Entry) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *Entry) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *Entry) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *Entry) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetPlaybookId

`func (o *Entry) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *Entry) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *Entry) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.

### HasPlaybookId

`func (o *Entry) HasPlaybookId() bool`

HasPlaybookId returns a boolean if a field has been set.

### GetPolling

`func (o *Entry) GetPolling() bool`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *Entry) GetPollingOk() (*bool, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *Entry) SetPolling(v bool)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *Entry) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetPollingArgs

`func (o *Entry) GetPollingArgs() map[string]map[string]interface{}`

GetPollingArgs returns the PollingArgs field if non-nil, zero value otherwise.

### GetPollingArgsOk

`func (o *Entry) GetPollingArgsOk() (*map[string]map[string]interface{}, bool)`

GetPollingArgsOk returns a tuple with the PollingArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingArgs

`func (o *Entry) SetPollingArgs(v map[string]map[string]interface{})`

SetPollingArgs sets PollingArgs field to given value.

### HasPollingArgs

`func (o *Entry) HasPollingArgs() bool`

HasPollingArgs returns a boolean if a field has been set.

### GetPollingCommand

`func (o *Entry) GetPollingCommand() string`

GetPollingCommand returns the PollingCommand field if non-nil, zero value otherwise.

### GetPollingCommandOk

`func (o *Entry) GetPollingCommandOk() (*string, bool)`

GetPollingCommandOk returns a tuple with the PollingCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingCommand

`func (o *Entry) SetPollingCommand(v string)`

SetPollingCommand sets PollingCommand field to given value.

### HasPollingCommand

`func (o *Entry) HasPollingCommand() bool`

HasPollingCommand returns a boolean if a field has been set.

### GetPreviousAllRead

`func (o *Entry) GetPreviousAllRead() bool`

GetPreviousAllRead returns the PreviousAllRead field if non-nil, zero value otherwise.

### GetPreviousAllReadOk

`func (o *Entry) GetPreviousAllReadOk() (*bool, bool)`

GetPreviousAllReadOk returns a tuple with the PreviousAllRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllRead

`func (o *Entry) SetPreviousAllRead(v bool)`

SetPreviousAllRead sets PreviousAllRead field to given value.

### HasPreviousAllRead

`func (o *Entry) HasPreviousAllRead() bool`

HasPreviousAllRead returns a boolean if a field has been set.

### GetPreviousAllReadWrite

`func (o *Entry) GetPreviousAllReadWrite() bool`

GetPreviousAllReadWrite returns the PreviousAllReadWrite field if non-nil, zero value otherwise.

### GetPreviousAllReadWriteOk

`func (o *Entry) GetPreviousAllReadWriteOk() (*bool, bool)`

GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllReadWrite

`func (o *Entry) SetPreviousAllReadWrite(v bool)`

SetPreviousAllReadWrite sets PreviousAllReadWrite field to given value.

### HasPreviousAllReadWrite

`func (o *Entry) HasPreviousAllReadWrite() bool`

HasPreviousAllReadWrite returns a boolean if a field has been set.

### GetPreviousRoles

`func (o *Entry) GetPreviousRoles() []string`

GetPreviousRoles returns the PreviousRoles field if non-nil, zero value otherwise.

### GetPreviousRolesOk

`func (o *Entry) GetPreviousRolesOk() (*[]string, bool)`

GetPreviousRolesOk returns a tuple with the PreviousRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRoles

`func (o *Entry) SetPreviousRoles(v []string)`

SetPreviousRoles sets PreviousRoles field to given value.

### HasPreviousRoles

`func (o *Entry) HasPreviousRoles() bool`

HasPreviousRoles returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Entry) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Entry) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Entry) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Entry) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetReadOnly

`func (o *Entry) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Entry) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Entry) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Entry) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRecurrent

`func (o *Entry) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *Entry) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *Entry) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *Entry) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetReputationSize

`func (o *Entry) GetReputationSize() int64`

GetReputationSize returns the ReputationSize field if non-nil, zero value otherwise.

### GetReputationSizeOk

`func (o *Entry) GetReputationSizeOk() (*int64, bool)`

GetReputationSizeOk returns a tuple with the ReputationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationSize

`func (o *Entry) SetReputationSize(v int64)`

SetReputationSize sets ReputationSize field to given value.

### HasReputationSize

`func (o *Entry) HasReputationSize() bool`

HasReputationSize returns a boolean if a field has been set.

### GetReputations

`func (o *Entry) GetReputations() []EntryReputation`

GetReputations returns the Reputations field if non-nil, zero value otherwise.

### GetReputationsOk

`func (o *Entry) GetReputationsOk() (*[]EntryReputation, bool)`

GetReputationsOk returns a tuple with the Reputations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputations

`func (o *Entry) SetReputations(v []EntryReputation)`

SetReputations sets Reputations field to given value.

### HasReputations

`func (o *Entry) HasReputations() bool`

HasReputations returns a boolean if a field has been set.

### GetRetryTime

`func (o *Entry) GetRetryTime() time.Time`

GetRetryTime returns the RetryTime field if non-nil, zero value otherwise.

### GetRetryTimeOk

`func (o *Entry) GetRetryTimeOk() (*time.Time, bool)`

GetRetryTimeOk returns a tuple with the RetryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTime

`func (o *Entry) SetRetryTime(v time.Time)`

SetRetryTime sets RetryTime field to given value.

### HasRetryTime

`func (o *Entry) HasRetryTime() bool`

HasRetryTime returns a boolean if a field has been set.

### GetRoles

`func (o *Entry) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Entry) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Entry) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Entry) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScheduled

`func (o *Entry) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *Entry) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *Entry) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *Entry) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Entry) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Entry) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Entry) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Entry) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *Entry) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Entry) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Entry) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Entry) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetStartDate

`func (o *Entry) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Entry) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Entry) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Entry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSystem

`func (o *Entry) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Entry) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Entry) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Entry) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *Entry) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Entry) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Entry) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Entry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagsRaw

`func (o *Entry) GetTagsRaw() []string`

GetTagsRaw returns the TagsRaw field if non-nil, zero value otherwise.

### GetTagsRawOk

`func (o *Entry) GetTagsRawOk() (*[]string, bool)`

GetTagsRawOk returns a tuple with the TagsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsRaw

`func (o *Entry) SetTagsRaw(v []string)`

SetTagsRaw sets TagsRaw field to given value.

### HasTagsRaw

`func (o *Entry) HasTagsRaw() bool`

HasTagsRaw returns a boolean if a field has been set.

### GetTaskId

`func (o *Entry) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Entry) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Entry) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Entry) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTimes

`func (o *Entry) GetTimes() int64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Entry) GetTimesOk() (*int64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Entry) SetTimes(v int64)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *Entry) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesRan

`func (o *Entry) GetTimesRan() int64`

GetTimesRan returns the TimesRan field if non-nil, zero value otherwise.

### GetTimesRanOk

`func (o *Entry) GetTimesRanOk() (*int64, bool)`

GetTimesRanOk returns a tuple with the TimesRan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesRan

`func (o *Entry) SetTimesRan(v int64)`

SetTimesRan sets TimesRan field to given value.

### HasTimesRan

`func (o *Entry) HasTimesRan() bool`

HasTimesRan returns a boolean if a field has been set.

### GetTimezone

`func (o *Entry) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Entry) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Entry) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Entry) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *Entry) GetTimezoneOffset() int64`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *Entry) GetTimezoneOffsetOk() (*int64, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *Entry) SetTimezoneOffset(v int64)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *Entry) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetType

`func (o *Entry) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entry) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entry) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *Entry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *Entry) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Entry) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Entry) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Entry) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *Entry) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Entry) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Entry) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Entry) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXsoarHasReadOnlyRole

`func (o *Entry) GetXsoarHasReadOnlyRole() bool`

GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field if non-nil, zero value otherwise.

### GetXsoarHasReadOnlyRoleOk

`func (o *Entry) GetXsoarHasReadOnlyRoleOk() (*bool, bool)`

GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarHasReadOnlyRole

`func (o *Entry) SetXsoarHasReadOnlyRole(v bool)`

SetXsoarHasReadOnlyRole sets XsoarHasReadOnlyRole field to given value.

### HasXsoarHasReadOnlyRole

`func (o *Entry) HasXsoarHasReadOnlyRole() bool`

HasXsoarHasReadOnlyRole returns a boolean if a field has been set.

### GetXsoarPreviousReadOnlyRoles

`func (o *Entry) GetXsoarPreviousReadOnlyRoles() []string`

GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarPreviousReadOnlyRolesOk

`func (o *Entry) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool)`

GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarPreviousReadOnlyRoles

`func (o *Entry) SetXsoarPreviousReadOnlyRoles(v []string)`

SetXsoarPreviousReadOnlyRoles sets XsoarPreviousReadOnlyRoles field to given value.

### HasXsoarPreviousReadOnlyRoles

`func (o *Entry) HasXsoarPreviousReadOnlyRoles() bool`

HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.

### GetXsoarReadOnlyRoles

`func (o *Entry) GetXsoarReadOnlyRoles() []string`

GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field if non-nil, zero value otherwise.

### GetXsoarReadOnlyRolesOk

`func (o *Entry) GetXsoarReadOnlyRolesOk() (*[]string, bool)`

GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsoarReadOnlyRoles

`func (o *Entry) SetXsoarReadOnlyRoles(v []string)`

SetXsoarReadOnlyRoles sets XsoarReadOnlyRoles field to given value.

### HasXsoarReadOnlyRoles

`func (o *Entry) HasXsoarReadOnlyRoles() bool`

HasXsoarReadOnlyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


