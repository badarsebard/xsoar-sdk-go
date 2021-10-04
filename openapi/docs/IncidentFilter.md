# IncidentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**AndOp** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **[]string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Files** | Pointer to **[]string** |  | [optional] 
**FromActivatedDate** | Pointer to **time.Time** |  | [optional] 
**FromClosedDate** | Pointer to **time.Time** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**FromDueDate** | Pointer to **time.Time** |  | [optional] 
**FromReminder** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **[]string** |  | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**IncludeTmp** | Pointer to **bool** |  | [optional] 
**Investigation** | Pointer to **[]string** |  | [optional] 
**Level** | Pointer to **[]float64** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**NotCategory** | Pointer to **[]string** |  | [optional] 
**NotInvestigation** | Pointer to **[]string** |  | [optional] 
**NotStatus** | Pointer to **[]float64** |  | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Parent** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **[]string** |  | [optional] 
**SearchAfter** | Pointer to **[]string** | Efficient next page, pass max sort value from previous page | [optional] 
**SearchAfterElastic** | Pointer to **[]string** | Efficient next page, pass max ES sort value from previous page | [optional] 
**SearchBefore** | Pointer to **[]string** | Efficient prev page, pass min sort value from next page | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** | Efficient prev page, pass min ES sort value from next page | [optional] 
**Size** | Pointer to **int64** | Size is limited to 1000, if not passed it defaults to 0, and no results will return | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | The sort order | [optional] 
**Status** | Pointer to **[]float64** |  | [optional] 
**Systems** | Pointer to **[]string** |  | [optional] 
**TimeFrame** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**ToActivatedDate** | Pointer to **time.Time** |  | [optional] 
**ToClosedDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**ToDueDate** | Pointer to **time.Time** |  | [optional] 
**ToReminder** | Pointer to **time.Time** |  | [optional] 
**TotalOnly** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 
**Urls** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIncidentFilter

`func NewIncidentFilter() *IncidentFilter`

NewIncidentFilter instantiates a new IncidentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFilterWithDefaults

`func NewIncidentFilterWithDefaults() *IncidentFilter`

NewIncidentFilterWithDefaults instantiates a new IncidentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *IncidentFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *IncidentFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *IncidentFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *IncidentFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetAndOp

`func (o *IncidentFilter) GetAndOp() bool`

GetAndOp returns the AndOp field if non-nil, zero value otherwise.

### GetAndOpOk

`func (o *IncidentFilter) GetAndOpOk() (*bool, bool)`

GetAndOpOk returns a tuple with the AndOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndOp

`func (o *IncidentFilter) SetAndOp(v bool)`

SetAndOp sets AndOp field to given value.

### HasAndOp

`func (o *IncidentFilter) HasAndOp() bool`

HasAndOp returns a boolean if a field has been set.

### GetCategory

`func (o *IncidentFilter) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IncidentFilter) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IncidentFilter) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IncidentFilter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDetails

`func (o *IncidentFilter) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IncidentFilter) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IncidentFilter) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IncidentFilter) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFields

`func (o *IncidentFilter) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IncidentFilter) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IncidentFilter) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IncidentFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFiles

`func (o *IncidentFilter) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *IncidentFilter) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *IncidentFilter) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *IncidentFilter) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFromActivatedDate

`func (o *IncidentFilter) GetFromActivatedDate() time.Time`

GetFromActivatedDate returns the FromActivatedDate field if non-nil, zero value otherwise.

### GetFromActivatedDateOk

`func (o *IncidentFilter) GetFromActivatedDateOk() (*time.Time, bool)`

GetFromActivatedDateOk returns a tuple with the FromActivatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromActivatedDate

`func (o *IncidentFilter) SetFromActivatedDate(v time.Time)`

SetFromActivatedDate sets FromActivatedDate field to given value.

### HasFromActivatedDate

`func (o *IncidentFilter) HasFromActivatedDate() bool`

HasFromActivatedDate returns a boolean if a field has been set.

### GetFromClosedDate

`func (o *IncidentFilter) GetFromClosedDate() time.Time`

GetFromClosedDate returns the FromClosedDate field if non-nil, zero value otherwise.

### GetFromClosedDateOk

`func (o *IncidentFilter) GetFromClosedDateOk() (*time.Time, bool)`

GetFromClosedDateOk returns a tuple with the FromClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromClosedDate

`func (o *IncidentFilter) SetFromClosedDate(v time.Time)`

SetFromClosedDate sets FromClosedDate field to given value.

### HasFromClosedDate

`func (o *IncidentFilter) HasFromClosedDate() bool`

HasFromClosedDate returns a boolean if a field has been set.

### GetFromDate

`func (o *IncidentFilter) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *IncidentFilter) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *IncidentFilter) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *IncidentFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *IncidentFilter) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *IncidentFilter) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *IncidentFilter) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *IncidentFilter) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetFromDueDate

`func (o *IncidentFilter) GetFromDueDate() time.Time`

GetFromDueDate returns the FromDueDate field if non-nil, zero value otherwise.

### GetFromDueDateOk

`func (o *IncidentFilter) GetFromDueDateOk() (*time.Time, bool)`

GetFromDueDateOk returns a tuple with the FromDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDueDate

`func (o *IncidentFilter) SetFromDueDate(v time.Time)`

SetFromDueDate sets FromDueDate field to given value.

### HasFromDueDate

`func (o *IncidentFilter) HasFromDueDate() bool`

HasFromDueDate returns a boolean if a field has been set.

### GetFromReminder

`func (o *IncidentFilter) GetFromReminder() time.Time`

GetFromReminder returns the FromReminder field if non-nil, zero value otherwise.

### GetFromReminderOk

`func (o *IncidentFilter) GetFromReminderOk() (*time.Time, bool)`

GetFromReminderOk returns a tuple with the FromReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromReminder

`func (o *IncidentFilter) SetFromReminder(v time.Time)`

SetFromReminder sets FromReminder field to given value.

### HasFromReminder

`func (o *IncidentFilter) HasFromReminder() bool`

HasFromReminder returns a boolean if a field has been set.

### GetId

`func (o *IncidentFilter) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentFilter) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentFilter) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *IncidentFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *IncidentFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *IncidentFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *IncidentFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetIncludeTmp

`func (o *IncidentFilter) GetIncludeTmp() bool`

GetIncludeTmp returns the IncludeTmp field if non-nil, zero value otherwise.

### GetIncludeTmpOk

`func (o *IncidentFilter) GetIncludeTmpOk() (*bool, bool)`

GetIncludeTmpOk returns a tuple with the IncludeTmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTmp

`func (o *IncidentFilter) SetIncludeTmp(v bool)`

SetIncludeTmp sets IncludeTmp field to given value.

### HasIncludeTmp

`func (o *IncidentFilter) HasIncludeTmp() bool`

HasIncludeTmp returns a boolean if a field has been set.

### GetInvestigation

`func (o *IncidentFilter) GetInvestigation() []string`

GetInvestigation returns the Investigation field if non-nil, zero value otherwise.

### GetInvestigationOk

`func (o *IncidentFilter) GetInvestigationOk() (*[]string, bool)`

GetInvestigationOk returns a tuple with the Investigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigation

`func (o *IncidentFilter) SetInvestigation(v []string)`

SetInvestigation sets Investigation field to given value.

### HasInvestigation

`func (o *IncidentFilter) HasInvestigation() bool`

HasInvestigation returns a boolean if a field has been set.

### GetLevel

`func (o *IncidentFilter) GetLevel() []float64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IncidentFilter) GetLevelOk() (*[]float64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IncidentFilter) SetLevel(v []float64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IncidentFilter) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *IncidentFilter) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentFilter) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentFilter) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotCategory

`func (o *IncidentFilter) GetNotCategory() []string`

GetNotCategory returns the NotCategory field if non-nil, zero value otherwise.

### GetNotCategoryOk

`func (o *IncidentFilter) GetNotCategoryOk() (*[]string, bool)`

GetNotCategoryOk returns a tuple with the NotCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCategory

`func (o *IncidentFilter) SetNotCategory(v []string)`

SetNotCategory sets NotCategory field to given value.

### HasNotCategory

`func (o *IncidentFilter) HasNotCategory() bool`

HasNotCategory returns a boolean if a field has been set.

### GetNotInvestigation

`func (o *IncidentFilter) GetNotInvestigation() []string`

GetNotInvestigation returns the NotInvestigation field if non-nil, zero value otherwise.

### GetNotInvestigationOk

`func (o *IncidentFilter) GetNotInvestigationOk() (*[]string, bool)`

GetNotInvestigationOk returns a tuple with the NotInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInvestigation

`func (o *IncidentFilter) SetNotInvestigation(v []string)`

SetNotInvestigation sets NotInvestigation field to given value.

### HasNotInvestigation

`func (o *IncidentFilter) HasNotInvestigation() bool`

HasNotInvestigation returns a boolean if a field has been set.

### GetNotStatus

`func (o *IncidentFilter) GetNotStatus() []float64`

GetNotStatus returns the NotStatus field if non-nil, zero value otherwise.

### GetNotStatusOk

`func (o *IncidentFilter) GetNotStatusOk() (*[]float64, bool)`

GetNotStatusOk returns a tuple with the NotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStatus

`func (o *IncidentFilter) SetNotStatus(v []float64)`

SetNotStatus sets NotStatus field to given value.

### HasNotStatus

`func (o *IncidentFilter) HasNotStatus() bool`

HasNotStatus returns a boolean if a field has been set.

### GetPage

`func (o *IncidentFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IncidentFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IncidentFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *IncidentFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetParent

`func (o *IncidentFilter) GetParent() []string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IncidentFilter) GetParentOk() (*[]string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IncidentFilter) SetParent(v []string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IncidentFilter) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPeriod

`func (o *IncidentFilter) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *IncidentFilter) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *IncidentFilter) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *IncidentFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuery

`func (o *IncidentFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IncidentFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IncidentFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IncidentFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReason

`func (o *IncidentFilter) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IncidentFilter) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IncidentFilter) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IncidentFilter) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSearchAfter

`func (o *IncidentFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *IncidentFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *IncidentFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *IncidentFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *IncidentFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *IncidentFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *IncidentFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *IncidentFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *IncidentFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *IncidentFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *IncidentFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *IncidentFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *IncidentFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *IncidentFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *IncidentFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *IncidentFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *IncidentFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IncidentFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IncidentFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *IncidentFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *IncidentFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IncidentFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IncidentFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IncidentFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStatus

`func (o *IncidentFilter) GetStatus() []float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentFilter) GetStatusOk() (*[]float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentFilter) SetStatus(v []float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystems

`func (o *IncidentFilter) GetSystems() []string`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *IncidentFilter) GetSystemsOk() (*[]string, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *IncidentFilter) SetSystems(v []string)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *IncidentFilter) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetTimeFrame

`func (o *IncidentFilter) GetTimeFrame() int64`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *IncidentFilter) GetTimeFrameOk() (*int64, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *IncidentFilter) SetTimeFrame(v int64)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *IncidentFilter) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetToActivatedDate

`func (o *IncidentFilter) GetToActivatedDate() time.Time`

GetToActivatedDate returns the ToActivatedDate field if non-nil, zero value otherwise.

### GetToActivatedDateOk

`func (o *IncidentFilter) GetToActivatedDateOk() (*time.Time, bool)`

GetToActivatedDateOk returns a tuple with the ToActivatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToActivatedDate

`func (o *IncidentFilter) SetToActivatedDate(v time.Time)`

SetToActivatedDate sets ToActivatedDate field to given value.

### HasToActivatedDate

`func (o *IncidentFilter) HasToActivatedDate() bool`

HasToActivatedDate returns a boolean if a field has been set.

### GetToClosedDate

`func (o *IncidentFilter) GetToClosedDate() time.Time`

GetToClosedDate returns the ToClosedDate field if non-nil, zero value otherwise.

### GetToClosedDateOk

`func (o *IncidentFilter) GetToClosedDateOk() (*time.Time, bool)`

GetToClosedDateOk returns a tuple with the ToClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToClosedDate

`func (o *IncidentFilter) SetToClosedDate(v time.Time)`

SetToClosedDate sets ToClosedDate field to given value.

### HasToClosedDate

`func (o *IncidentFilter) HasToClosedDate() bool`

HasToClosedDate returns a boolean if a field has been set.

### GetToDate

`func (o *IncidentFilter) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *IncidentFilter) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *IncidentFilter) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *IncidentFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetToDueDate

`func (o *IncidentFilter) GetToDueDate() time.Time`

GetToDueDate returns the ToDueDate field if non-nil, zero value otherwise.

### GetToDueDateOk

`func (o *IncidentFilter) GetToDueDateOk() (*time.Time, bool)`

GetToDueDateOk returns a tuple with the ToDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDueDate

`func (o *IncidentFilter) SetToDueDate(v time.Time)`

SetToDueDate sets ToDueDate field to given value.

### HasToDueDate

`func (o *IncidentFilter) HasToDueDate() bool`

HasToDueDate returns a boolean if a field has been set.

### GetToReminder

`func (o *IncidentFilter) GetToReminder() time.Time`

GetToReminder returns the ToReminder field if non-nil, zero value otherwise.

### GetToReminderOk

`func (o *IncidentFilter) GetToReminderOk() (*time.Time, bool)`

GetToReminderOk returns a tuple with the ToReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToReminder

`func (o *IncidentFilter) SetToReminder(v time.Time)`

SetToReminder sets ToReminder field to given value.

### HasToReminder

`func (o *IncidentFilter) HasToReminder() bool`

HasToReminder returns a boolean if a field has been set.

### GetTotalOnly

`func (o *IncidentFilter) GetTotalOnly() bool`

GetTotalOnly returns the TotalOnly field if non-nil, zero value otherwise.

### GetTotalOnlyOk

`func (o *IncidentFilter) GetTotalOnlyOk() (*bool, bool)`

GetTotalOnlyOk returns a tuple with the TotalOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnly

`func (o *IncidentFilter) SetTotalOnly(v bool)`

SetTotalOnly sets TotalOnly field to given value.

### HasTotalOnly

`func (o *IncidentFilter) HasTotalOnly() bool`

HasTotalOnly returns a boolean if a field has been set.

### GetType

`func (o *IncidentFilter) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentFilter) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentFilter) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrls

`func (o *IncidentFilter) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *IncidentFilter) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *IncidentFilter) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *IncidentFilter) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetUsers

`func (o *IncidentFilter) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IncidentFilter) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IncidentFilter) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IncidentFilter) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


