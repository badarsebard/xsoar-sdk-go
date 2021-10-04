# RelationshipFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**AttachIndicators** | Pointer to **bool** |  | [optional] 
**Entities** | Pointer to **[]string** |  | [optional] 
**EntityFamilies** | Pointer to **[]string** |  | [optional] 
**EntityTypes** | Pointer to **[]string** |  | [optional] 
**ExcludedEntities** | Pointer to **[]string** |  | [optional] 
**ExcludedEntityTypes** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**RelationshipNames** | Pointer to **[]string** |  | [optional] 
**SearchAfter** | Pointer to **[]string** | Efficient next page, pass max sort value from previous page | [optional] 
**SearchAfterElastic** | Pointer to **[]string** | Efficient next page, pass max ES sort value from previous page | [optional] 
**SearchBefore** | Pointer to **[]string** | Efficient prev page, pass min sort value from next page | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** | Efficient prev page, pass min ES sort value from next page | [optional] 
**Size** | Pointer to **int64** | Size is limited to 1000, if not passed it defaults to 0, and no results will return | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | The sort order | [optional] 
**TimeFrame** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRelationshipFilter

`func NewRelationshipFilter() *RelationshipFilter`

NewRelationshipFilter instantiates a new RelationshipFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipFilterWithDefaults

`func NewRelationshipFilterWithDefaults() *RelationshipFilter`

NewRelationshipFilterWithDefaults instantiates a new RelationshipFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *RelationshipFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *RelationshipFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *RelationshipFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *RelationshipFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetAttachIndicators

`func (o *RelationshipFilter) GetAttachIndicators() bool`

GetAttachIndicators returns the AttachIndicators field if non-nil, zero value otherwise.

### GetAttachIndicatorsOk

`func (o *RelationshipFilter) GetAttachIndicatorsOk() (*bool, bool)`

GetAttachIndicatorsOk returns a tuple with the AttachIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachIndicators

`func (o *RelationshipFilter) SetAttachIndicators(v bool)`

SetAttachIndicators sets AttachIndicators field to given value.

### HasAttachIndicators

`func (o *RelationshipFilter) HasAttachIndicators() bool`

HasAttachIndicators returns a boolean if a field has been set.

### GetEntities

`func (o *RelationshipFilter) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RelationshipFilter) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RelationshipFilter) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RelationshipFilter) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetEntityFamilies

`func (o *RelationshipFilter) GetEntityFamilies() []string`

GetEntityFamilies returns the EntityFamilies field if non-nil, zero value otherwise.

### GetEntityFamiliesOk

`func (o *RelationshipFilter) GetEntityFamiliesOk() (*[]string, bool)`

GetEntityFamiliesOk returns a tuple with the EntityFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityFamilies

`func (o *RelationshipFilter) SetEntityFamilies(v []string)`

SetEntityFamilies sets EntityFamilies field to given value.

### HasEntityFamilies

`func (o *RelationshipFilter) HasEntityFamilies() bool`

HasEntityFamilies returns a boolean if a field has been set.

### GetEntityTypes

`func (o *RelationshipFilter) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *RelationshipFilter) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *RelationshipFilter) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *RelationshipFilter) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetExcludedEntities

`func (o *RelationshipFilter) GetExcludedEntities() []string`

GetExcludedEntities returns the ExcludedEntities field if non-nil, zero value otherwise.

### GetExcludedEntitiesOk

`func (o *RelationshipFilter) GetExcludedEntitiesOk() (*[]string, bool)`

GetExcludedEntitiesOk returns a tuple with the ExcludedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntities

`func (o *RelationshipFilter) SetExcludedEntities(v []string)`

SetExcludedEntities sets ExcludedEntities field to given value.

### HasExcludedEntities

`func (o *RelationshipFilter) HasExcludedEntities() bool`

HasExcludedEntities returns a boolean if a field has been set.

### GetExcludedEntityTypes

`func (o *RelationshipFilter) GetExcludedEntityTypes() []string`

GetExcludedEntityTypes returns the ExcludedEntityTypes field if non-nil, zero value otherwise.

### GetExcludedEntityTypesOk

`func (o *RelationshipFilter) GetExcludedEntityTypesOk() (*[]string, bool)`

GetExcludedEntityTypesOk returns a tuple with the ExcludedEntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntityTypes

`func (o *RelationshipFilter) SetExcludedEntityTypes(v []string)`

SetExcludedEntityTypes sets ExcludedEntityTypes field to given value.

### HasExcludedEntityTypes

`func (o *RelationshipFilter) HasExcludedEntityTypes() bool`

HasExcludedEntityTypes returns a boolean if a field has been set.

### GetFields

`func (o *RelationshipFilter) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RelationshipFilter) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RelationshipFilter) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RelationshipFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFromDate

`func (o *RelationshipFilter) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *RelationshipFilter) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *RelationshipFilter) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *RelationshipFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *RelationshipFilter) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *RelationshipFilter) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *RelationshipFilter) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *RelationshipFilter) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *RelationshipFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *RelationshipFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *RelationshipFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *RelationshipFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetPage

`func (o *RelationshipFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RelationshipFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RelationshipFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *RelationshipFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPeriod

`func (o *RelationshipFilter) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RelationshipFilter) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RelationshipFilter) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RelationshipFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuery

`func (o *RelationshipFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RelationshipFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RelationshipFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RelationshipFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRelationshipNames

`func (o *RelationshipFilter) GetRelationshipNames() []string`

GetRelationshipNames returns the RelationshipNames field if non-nil, zero value otherwise.

### GetRelationshipNamesOk

`func (o *RelationshipFilter) GetRelationshipNamesOk() (*[]string, bool)`

GetRelationshipNamesOk returns a tuple with the RelationshipNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipNames

`func (o *RelationshipFilter) SetRelationshipNames(v []string)`

SetRelationshipNames sets RelationshipNames field to given value.

### HasRelationshipNames

`func (o *RelationshipFilter) HasRelationshipNames() bool`

HasRelationshipNames returns a boolean if a field has been set.

### GetSearchAfter

`func (o *RelationshipFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *RelationshipFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *RelationshipFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *RelationshipFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *RelationshipFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *RelationshipFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *RelationshipFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *RelationshipFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *RelationshipFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *RelationshipFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *RelationshipFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *RelationshipFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *RelationshipFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *RelationshipFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *RelationshipFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *RelationshipFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *RelationshipFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RelationshipFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RelationshipFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RelationshipFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *RelationshipFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RelationshipFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RelationshipFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RelationshipFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTimeFrame

`func (o *RelationshipFilter) GetTimeFrame() int64`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *RelationshipFilter) GetTimeFrameOk() (*int64, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *RelationshipFilter) SetTimeFrame(v int64)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *RelationshipFilter) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetToDate

`func (o *RelationshipFilter) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *RelationshipFilter) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *RelationshipFilter) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *RelationshipFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


