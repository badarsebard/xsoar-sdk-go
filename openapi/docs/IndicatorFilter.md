# IndicatorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**EarlyTimeInPage** | Pointer to **time.Time** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**FirstSeen** | Pointer to [**DateRangeFilter**](DateRangeFilter.md) |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**LastSeen** | Pointer to [**DateRangeFilter**](DateRangeFilter.md) |  | [optional] 
**LaterTimeInPage** | Pointer to **time.Time** |  | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**PrevPage** | Pointer to **bool** | MT support - these fields are for indicator search according to calculatedTime | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SearchAfter** | Pointer to **[]string** | Efficient next page, pass max sort value from previous page | [optional] 
**SearchAfterElastic** | Pointer to **[]string** | Efficient next page, pass max ES sort value from previous page | [optional] 
**SearchBefore** | Pointer to **[]string** | Efficient prev page, pass min sort value from next page | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** | Efficient prev page, pass min ES sort value from next page | [optional] 
**Size** | Pointer to **int64** | Size is limited to 1000, if not passed it defaults to 0, and no results will return | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | The sort order | [optional] 
**TimeFrame** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIndicatorFilter

`func NewIndicatorFilter() *IndicatorFilter`

NewIndicatorFilter instantiates a new IndicatorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorFilterWithDefaults

`func NewIndicatorFilterWithDefaults() *IndicatorFilter`

NewIndicatorFilterWithDefaults instantiates a new IndicatorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *IndicatorFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *IndicatorFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *IndicatorFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *IndicatorFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetEarlyTimeInPage

`func (o *IndicatorFilter) GetEarlyTimeInPage() time.Time`

GetEarlyTimeInPage returns the EarlyTimeInPage field if non-nil, zero value otherwise.

### GetEarlyTimeInPageOk

`func (o *IndicatorFilter) GetEarlyTimeInPageOk() (*time.Time, bool)`

GetEarlyTimeInPageOk returns a tuple with the EarlyTimeInPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyTimeInPage

`func (o *IndicatorFilter) SetEarlyTimeInPage(v time.Time)`

SetEarlyTimeInPage sets EarlyTimeInPage field to given value.

### HasEarlyTimeInPage

`func (o *IndicatorFilter) HasEarlyTimeInPage() bool`

HasEarlyTimeInPage returns a boolean if a field has been set.

### GetFields

`func (o *IndicatorFilter) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndicatorFilter) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndicatorFilter) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IndicatorFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFirstSeen

`func (o *IndicatorFilter) GetFirstSeen() DateRangeFilter`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *IndicatorFilter) GetFirstSeenOk() (*DateRangeFilter, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *IndicatorFilter) SetFirstSeen(v DateRangeFilter)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *IndicatorFilter) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetFromDate

`func (o *IndicatorFilter) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *IndicatorFilter) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *IndicatorFilter) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *IndicatorFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *IndicatorFilter) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *IndicatorFilter) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *IndicatorFilter) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *IndicatorFilter) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *IndicatorFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *IndicatorFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *IndicatorFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *IndicatorFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetLastSeen

`func (o *IndicatorFilter) GetLastSeen() DateRangeFilter`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *IndicatorFilter) GetLastSeenOk() (*DateRangeFilter, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *IndicatorFilter) SetLastSeen(v DateRangeFilter)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *IndicatorFilter) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLaterTimeInPage

`func (o *IndicatorFilter) GetLaterTimeInPage() time.Time`

GetLaterTimeInPage returns the LaterTimeInPage field if non-nil, zero value otherwise.

### GetLaterTimeInPageOk

`func (o *IndicatorFilter) GetLaterTimeInPageOk() (*time.Time, bool)`

GetLaterTimeInPageOk returns a tuple with the LaterTimeInPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaterTimeInPage

`func (o *IndicatorFilter) SetLaterTimeInPage(v time.Time)`

SetLaterTimeInPage sets LaterTimeInPage field to given value.

### HasLaterTimeInPage

`func (o *IndicatorFilter) HasLaterTimeInPage() bool`

HasLaterTimeInPage returns a boolean if a field has been set.

### GetPage

`func (o *IndicatorFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IndicatorFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IndicatorFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *IndicatorFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPeriod

`func (o *IndicatorFilter) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *IndicatorFilter) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *IndicatorFilter) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *IndicatorFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrevPage

`func (o *IndicatorFilter) GetPrevPage() bool`

GetPrevPage returns the PrevPage field if non-nil, zero value otherwise.

### GetPrevPageOk

`func (o *IndicatorFilter) GetPrevPageOk() (*bool, bool)`

GetPrevPageOk returns a tuple with the PrevPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPage

`func (o *IndicatorFilter) SetPrevPage(v bool)`

SetPrevPage sets PrevPage field to given value.

### HasPrevPage

`func (o *IndicatorFilter) HasPrevPage() bool`

HasPrevPage returns a boolean if a field has been set.

### GetQuery

`func (o *IndicatorFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IndicatorFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IndicatorFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IndicatorFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSearchAfter

`func (o *IndicatorFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *IndicatorFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *IndicatorFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *IndicatorFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *IndicatorFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *IndicatorFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *IndicatorFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *IndicatorFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *IndicatorFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *IndicatorFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *IndicatorFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *IndicatorFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *IndicatorFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *IndicatorFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *IndicatorFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *IndicatorFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *IndicatorFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IndicatorFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IndicatorFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *IndicatorFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *IndicatorFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IndicatorFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IndicatorFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IndicatorFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTimeFrame

`func (o *IndicatorFilter) GetTimeFrame() int64`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *IndicatorFilter) GetTimeFrameOk() (*int64, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *IndicatorFilter) SetTimeFrame(v int64)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *IndicatorFilter) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetToDate

`func (o *IndicatorFilter) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *IndicatorFilter) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *IndicatorFilter) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *IndicatorFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


