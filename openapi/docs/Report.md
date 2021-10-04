# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**CronView** | Pointer to **bool** |  | [optional] 
**Dashboard** | Pointer to [**Dashboard**](Dashboard.md) |  | [optional] 
**Decoder** | Pointer to [**map[string]ReportFieldsDecoder**](ReportFieldsDecoder.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableHeader** | Pointer to **bool** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**EndingType** | Pointer to **string** | EndingType holds the type of schedule Ending | [optional] 
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**HumanCron** | Pointer to [**HumanCron**](HumanCron.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**LatestReportName** | Pointer to **string** |  | [optional] 
**LatestReportTime** | Pointer to **time.Time** |  | [optional] 
**LatestReportUsername** | Pointer to **string** |  | [optional] 
**LatestScheduledReportTime** | Pointer to **time.Time** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextScheduledTime** | Pointer to **time.Time** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**Orientation** | Pointer to **string** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PaperSize** | Pointer to **string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PrevType** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**Recurrent** | Pointer to **bool** |  | [optional] 
**ReportType** | Pointer to **string** |  | [optional] 
**RunOnce** | Pointer to **bool** |  | [optional] 
**RunningRoles** | Pointer to **[]string** |  | [optional] 
**RunningUser** | Pointer to **string** |  | [optional] 
**Scheduled** | Pointer to **bool** | is it scheduled | [optional] 
**Sections** | Pointer to [**[]Section**](Section.md) |  | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Times** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**TimezoneOffset** | Pointer to **int64** |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserAPIKey** | Pointer to **string** |  | [optional] 
**UserAPIKeyID** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *Report) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *Report) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *Report) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *Report) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Report) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Report) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Report) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Report) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCron

`func (o *Report) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *Report) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *Report) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *Report) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetCronView

`func (o *Report) GetCronView() bool`

GetCronView returns the CronView field if non-nil, zero value otherwise.

### GetCronViewOk

`func (o *Report) GetCronViewOk() (*bool, bool)`

GetCronViewOk returns a tuple with the CronView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronView

`func (o *Report) SetCronView(v bool)`

SetCronView sets CronView field to given value.

### HasCronView

`func (o *Report) HasCronView() bool`

HasCronView returns a boolean if a field has been set.

### GetDashboard

`func (o *Report) GetDashboard() Dashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *Report) GetDashboardOk() (*Dashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *Report) SetDashboard(v Dashboard)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *Report) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDecoder

`func (o *Report) GetDecoder() map[string]ReportFieldsDecoder`

GetDecoder returns the Decoder field if non-nil, zero value otherwise.

### GetDecoderOk

`func (o *Report) GetDecoderOk() (*map[string]ReportFieldsDecoder, bool)`

GetDecoderOk returns a tuple with the Decoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoder

`func (o *Report) SetDecoder(v map[string]ReportFieldsDecoder)`

SetDecoder sets Decoder field to given value.

### HasDecoder

`func (o *Report) HasDecoder() bool`

HasDecoder returns a boolean if a field has been set.

### GetDescription

`func (o *Report) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Report) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Report) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Report) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableHeader

`func (o *Report) GetDisableHeader() bool`

GetDisableHeader returns the DisableHeader field if non-nil, zero value otherwise.

### GetDisableHeaderOk

`func (o *Report) GetDisableHeaderOk() (*bool, bool)`

GetDisableHeaderOk returns a tuple with the DisableHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHeader

`func (o *Report) SetDisableHeader(v bool)`

SetDisableHeader sets DisableHeader field to given value.

### HasDisableHeader

`func (o *Report) HasDisableHeader() bool`

HasDisableHeader returns a boolean if a field has been set.

### GetEndingDate

`func (o *Report) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *Report) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *Report) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *Report) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetEndingType

`func (o *Report) GetEndingType() string`

GetEndingType returns the EndingType field if non-nil, zero value otherwise.

### GetEndingTypeOk

`func (o *Report) GetEndingTypeOk() (*string, bool)`

GetEndingTypeOk returns a tuple with the EndingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingType

`func (o *Report) SetEndingType(v string)`

SetEndingType sets EndingType field to given value.

### HasEndingType

`func (o *Report) HasEndingType() bool`

HasEndingType returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *Report) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *Report) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *Report) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *Report) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *Report) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Report) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Report) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Report) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetHumanCron

`func (o *Report) GetHumanCron() HumanCron`

GetHumanCron returns the HumanCron field if non-nil, zero value otherwise.

### GetHumanCronOk

`func (o *Report) GetHumanCronOk() (*HumanCron, bool)`

GetHumanCronOk returns a tuple with the HumanCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanCron

`func (o *Report) SetHumanCron(v HumanCron)`

SetHumanCron sets HumanCron field to given value.

### HasHumanCron

`func (o *Report) HasHumanCron() bool`

HasHumanCron returns a boolean if a field has been set.

### GetId

`func (o *Report) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemVersion

`func (o *Report) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *Report) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *Report) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *Report) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLatestReportName

`func (o *Report) GetLatestReportName() string`

GetLatestReportName returns the LatestReportName field if non-nil, zero value otherwise.

### GetLatestReportNameOk

`func (o *Report) GetLatestReportNameOk() (*string, bool)`

GetLatestReportNameOk returns a tuple with the LatestReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReportName

`func (o *Report) SetLatestReportName(v string)`

SetLatestReportName sets LatestReportName field to given value.

### HasLatestReportName

`func (o *Report) HasLatestReportName() bool`

HasLatestReportName returns a boolean if a field has been set.

### GetLatestReportTime

`func (o *Report) GetLatestReportTime() time.Time`

GetLatestReportTime returns the LatestReportTime field if non-nil, zero value otherwise.

### GetLatestReportTimeOk

`func (o *Report) GetLatestReportTimeOk() (*time.Time, bool)`

GetLatestReportTimeOk returns a tuple with the LatestReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReportTime

`func (o *Report) SetLatestReportTime(v time.Time)`

SetLatestReportTime sets LatestReportTime field to given value.

### HasLatestReportTime

`func (o *Report) HasLatestReportTime() bool`

HasLatestReportTime returns a boolean if a field has been set.

### GetLatestReportUsername

`func (o *Report) GetLatestReportUsername() string`

GetLatestReportUsername returns the LatestReportUsername field if non-nil, zero value otherwise.

### GetLatestReportUsernameOk

`func (o *Report) GetLatestReportUsernameOk() (*string, bool)`

GetLatestReportUsernameOk returns a tuple with the LatestReportUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReportUsername

`func (o *Report) SetLatestReportUsername(v string)`

SetLatestReportUsername sets LatestReportUsername field to given value.

### HasLatestReportUsername

`func (o *Report) HasLatestReportUsername() bool`

HasLatestReportUsername returns a boolean if a field has been set.

### GetLatestScheduledReportTime

`func (o *Report) GetLatestScheduledReportTime() time.Time`

GetLatestScheduledReportTime returns the LatestScheduledReportTime field if non-nil, zero value otherwise.

### GetLatestScheduledReportTimeOk

`func (o *Report) GetLatestScheduledReportTimeOk() (*time.Time, bool)`

GetLatestScheduledReportTimeOk returns a tuple with the LatestScheduledReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestScheduledReportTime

`func (o *Report) SetLatestScheduledReportTime(v time.Time)`

SetLatestScheduledReportTime sets LatestScheduledReportTime field to given value.

### HasLatestScheduledReportTime

`func (o *Report) HasLatestScheduledReportTime() bool`

HasLatestScheduledReportTime returns a boolean if a field has been set.

### GetLocked

`func (o *Report) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Report) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Report) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Report) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *Report) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Report) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Report) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Report) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Report) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Report) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Report) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Report) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextScheduledTime

`func (o *Report) GetNextScheduledTime() time.Time`

GetNextScheduledTime returns the NextScheduledTime field if non-nil, zero value otherwise.

### GetNextScheduledTimeOk

`func (o *Report) GetNextScheduledTimeOk() (*time.Time, bool)`

GetNextScheduledTimeOk returns a tuple with the NextScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledTime

`func (o *Report) SetNextScheduledTime(v time.Time)`

SetNextScheduledTime sets NextScheduledTime field to given value.

### HasNextScheduledTime

`func (o *Report) HasNextScheduledTime() bool`

HasNextScheduledTime returns a boolean if a field has been set.

### GetNumericId

`func (o *Report) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Report) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Report) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Report) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetOrientation

`func (o *Report) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Report) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Report) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Report) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPackID

`func (o *Report) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *Report) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *Report) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *Report) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *Report) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *Report) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *Report) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *Report) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPaperSize

`func (o *Report) GetPaperSize() string`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *Report) GetPaperSizeOk() (*string, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *Report) SetPaperSize(v string)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *Report) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.

### GetPrevName

`func (o *Report) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *Report) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *Report) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *Report) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrevType

`func (o *Report) GetPrevType() string`

GetPrevType returns the PrevType field if non-nil, zero value otherwise.

### GetPrevTypeOk

`func (o *Report) GetPrevTypeOk() (*string, bool)`

GetPrevTypeOk returns a tuple with the PrevType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevType

`func (o *Report) SetPrevType(v string)`

SetPrevType sets PrevType field to given value.

### HasPrevType

`func (o *Report) HasPrevType() bool`

HasPrevType returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Report) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Report) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Report) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Report) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *Report) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *Report) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *Report) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *Report) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetRecipients

`func (o *Report) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *Report) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *Report) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *Report) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecurrent

`func (o *Report) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *Report) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *Report) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *Report) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetReportType

`func (o *Report) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *Report) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *Report) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *Report) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetRunOnce

`func (o *Report) GetRunOnce() bool`

GetRunOnce returns the RunOnce field if non-nil, zero value otherwise.

### GetRunOnceOk

`func (o *Report) GetRunOnceOk() (*bool, bool)`

GetRunOnceOk returns a tuple with the RunOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnce

`func (o *Report) SetRunOnce(v bool)`

SetRunOnce sets RunOnce field to given value.

### HasRunOnce

`func (o *Report) HasRunOnce() bool`

HasRunOnce returns a boolean if a field has been set.

### GetRunningRoles

`func (o *Report) GetRunningRoles() []string`

GetRunningRoles returns the RunningRoles field if non-nil, zero value otherwise.

### GetRunningRolesOk

`func (o *Report) GetRunningRolesOk() (*[]string, bool)`

GetRunningRolesOk returns a tuple with the RunningRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningRoles

`func (o *Report) SetRunningRoles(v []string)`

SetRunningRoles sets RunningRoles field to given value.

### HasRunningRoles

`func (o *Report) HasRunningRoles() bool`

HasRunningRoles returns a boolean if a field has been set.

### GetRunningUser

`func (o *Report) GetRunningUser() string`

GetRunningUser returns the RunningUser field if non-nil, zero value otherwise.

### GetRunningUserOk

`func (o *Report) GetRunningUserOk() (*string, bool)`

GetRunningUserOk returns a tuple with the RunningUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningUser

`func (o *Report) SetRunningUser(v string)`

SetRunningUser sets RunningUser field to given value.

### HasRunningUser

`func (o *Report) HasRunningUser() bool`

HasRunningUser returns a boolean if a field has been set.

### GetScheduled

`func (o *Report) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *Report) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *Report) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *Report) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSections

`func (o *Report) GetSections() []Section`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Report) GetSectionsOk() (*[]Section, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Report) SetSections(v []Section)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Report) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSensitive

`func (o *Report) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *Report) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *Report) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *Report) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Report) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Report) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Report) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Report) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *Report) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *Report) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *Report) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *Report) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *Report) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Report) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Report) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Report) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetStartDate

`func (o *Report) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Report) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Report) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Report) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSystem

`func (o *Report) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Report) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Report) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Report) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *Report) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Report) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Report) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Report) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimes

`func (o *Report) GetTimes() int64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Report) GetTimesOk() (*int64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Report) SetTimes(v int64)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *Report) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimezone

`func (o *Report) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Report) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Report) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Report) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *Report) GetTimezoneOffset() int64`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *Report) GetTimezoneOffsetOk() (*int64, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *Report) SetTimezoneOffset(v int64)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *Report) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetToServerVersion

`func (o *Report) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *Report) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *Report) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *Report) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetType

`func (o *Report) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Report) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Report) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Report) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserAPIKey

`func (o *Report) GetUserAPIKey() string`

GetUserAPIKey returns the UserAPIKey field if non-nil, zero value otherwise.

### GetUserAPIKeyOk

`func (o *Report) GetUserAPIKeyOk() (*string, bool)`

GetUserAPIKeyOk returns a tuple with the UserAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAPIKey

`func (o *Report) SetUserAPIKey(v string)`

SetUserAPIKey sets UserAPIKey field to given value.

### HasUserAPIKey

`func (o *Report) HasUserAPIKey() bool`

HasUserAPIKey returns a boolean if a field has been set.

### GetUserAPIKeyID

`func (o *Report) GetUserAPIKeyID() string`

GetUserAPIKeyID returns the UserAPIKeyID field if non-nil, zero value otherwise.

### GetUserAPIKeyIDOk

`func (o *Report) GetUserAPIKeyIDOk() (*string, bool)`

GetUserAPIKeyIDOk returns a tuple with the UserAPIKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAPIKeyID

`func (o *Report) SetUserAPIKeyID(v string)`

SetUserAPIKeyID sets UserAPIKeyID field to given value.

### HasUserAPIKeyID

`func (o *Report) HasUserAPIKeyID() bool`

HasUserAPIKeyID returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *Report) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *Report) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *Report) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *Report) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *Report) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *Report) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *Report) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *Report) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *Report) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Report) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Report) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Report) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


