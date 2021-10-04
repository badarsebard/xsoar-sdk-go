# GenericStringDateFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
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

### NewGenericStringDateFilter

`func NewGenericStringDateFilter() *GenericStringDateFilter`

NewGenericStringDateFilter instantiates a new GenericStringDateFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericStringDateFilterWithDefaults

`func NewGenericStringDateFilterWithDefaults() *GenericStringDateFilter`

NewGenericStringDateFilterWithDefaults instantiates a new GenericStringDateFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *GenericStringDateFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *GenericStringDateFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *GenericStringDateFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *GenericStringDateFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetFields

`func (o *GenericStringDateFilter) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GenericStringDateFilter) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GenericStringDateFilter) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GenericStringDateFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFromDate

`func (o *GenericStringDateFilter) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *GenericStringDateFilter) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *GenericStringDateFilter) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *GenericStringDateFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *GenericStringDateFilter) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *GenericStringDateFilter) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *GenericStringDateFilter) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *GenericStringDateFilter) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *GenericStringDateFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *GenericStringDateFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *GenericStringDateFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *GenericStringDateFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetPage

`func (o *GenericStringDateFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GenericStringDateFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GenericStringDateFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *GenericStringDateFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPeriod

`func (o *GenericStringDateFilter) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GenericStringDateFilter) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GenericStringDateFilter) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GenericStringDateFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuery

`func (o *GenericStringDateFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GenericStringDateFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GenericStringDateFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GenericStringDateFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSearchAfter

`func (o *GenericStringDateFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *GenericStringDateFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *GenericStringDateFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *GenericStringDateFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *GenericStringDateFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *GenericStringDateFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *GenericStringDateFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *GenericStringDateFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *GenericStringDateFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *GenericStringDateFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *GenericStringDateFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *GenericStringDateFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *GenericStringDateFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *GenericStringDateFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *GenericStringDateFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *GenericStringDateFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *GenericStringDateFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GenericStringDateFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GenericStringDateFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GenericStringDateFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *GenericStringDateFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GenericStringDateFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GenericStringDateFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GenericStringDateFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GenericStringDateFilter) GetTimeFrame() int64`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GenericStringDateFilter) GetTimeFrameOk() (*int64, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GenericStringDateFilter) SetTimeFrame(v int64)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GenericStringDateFilter) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetToDate

`func (o *GenericStringDateFilter) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *GenericStringDateFilter) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *GenericStringDateFilter) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *GenericStringDateFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


