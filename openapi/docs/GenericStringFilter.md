# GenericStringFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SearchAfter** | Pointer to **[]string** | Efficient next page, pass max sort value from previous page | [optional] 
**SearchAfterElastic** | Pointer to **[]string** | Efficient next page, pass max ES sort value from previous page | [optional] 
**SearchBefore** | Pointer to **[]string** | Efficient prev page, pass min sort value from next page | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** | Efficient prev page, pass min ES sort value from next page | [optional] 
**Size** | Pointer to **int64** | Size is limited to 1000, if not passed it defaults to 0, and no results will return | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | The sort order | [optional] 

## Methods

### NewGenericStringFilter

`func NewGenericStringFilter() *GenericStringFilter`

NewGenericStringFilter instantiates a new GenericStringFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericStringFilterWithDefaults

`func NewGenericStringFilterWithDefaults() *GenericStringFilter`

NewGenericStringFilterWithDefaults instantiates a new GenericStringFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *GenericStringFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *GenericStringFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *GenericStringFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *GenericStringFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *GenericStringFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *GenericStringFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *GenericStringFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *GenericStringFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetPage

`func (o *GenericStringFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GenericStringFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GenericStringFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *GenericStringFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *GenericStringFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GenericStringFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GenericStringFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GenericStringFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSearchAfter

`func (o *GenericStringFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *GenericStringFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *GenericStringFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *GenericStringFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *GenericStringFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *GenericStringFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *GenericStringFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *GenericStringFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *GenericStringFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *GenericStringFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *GenericStringFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *GenericStringFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *GenericStringFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *GenericStringFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *GenericStringFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *GenericStringFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *GenericStringFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GenericStringFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GenericStringFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GenericStringFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *GenericStringFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GenericStringFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GenericStringFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GenericStringFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


