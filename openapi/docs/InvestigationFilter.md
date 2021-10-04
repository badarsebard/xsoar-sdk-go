# InvestigationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **map[string][]string** | Cache of join functions | [optional] 
**AndOp** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **[]string** |  | [optional] 
**FromCloseDate** | Pointer to **time.Time** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **[]string** |  | [optional] 
**IdsOnly** | Pointer to **bool** |  | [optional] 
**IgnoreWorkers** | Pointer to **bool** | Do not use workers mechanism while searching bleve | [optional] 
**IncludeChildInv** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**NotCategory** | Pointer to **[]string** |  | [optional] 
**NotIDs** | Pointer to **[]string** |  | [optional] 
**Page** | Pointer to **int64** | 0-based page | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**Reason** | Pointer to **[]string** |  | [optional] 
**SearchAfter** | Pointer to **[]string** | Efficient next page, pass max sort value from previous page | [optional] 
**SearchAfterElastic** | Pointer to **[]string** | Efficient next page, pass max ES sort value from previous page | [optional] 
**SearchBefore** | Pointer to **[]string** | Efficient prev page, pass min sort value from next page | [optional] 
**SearchBeforeElastic** | Pointer to **[]string** | Efficient prev page, pass min ES sort value from next page | [optional] 
**Size** | Pointer to **int64** | Size is limited to 1000, if not passed it defaults to 0, and no results will return | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | The sort order | [optional] 
**Status** | Pointer to **[]float64** |  | [optional] 
**TimeFrame** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**ToCloseDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **[]float64** |  | [optional] 
**User** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInvestigationFilter

`func NewInvestigationFilter() *InvestigationFilter`

NewInvestigationFilter instantiates a new InvestigationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationFilterWithDefaults

`func NewInvestigationFilterWithDefaults() *InvestigationFilter`

NewInvestigationFilterWithDefaults instantiates a new InvestigationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *InvestigationFilter) GetCache() map[string][]string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *InvestigationFilter) GetCacheOk() (*map[string][]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *InvestigationFilter) SetCache(v map[string][]string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *InvestigationFilter) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetAndOp

`func (o *InvestigationFilter) GetAndOp() bool`

GetAndOp returns the AndOp field if non-nil, zero value otherwise.

### GetAndOpOk

`func (o *InvestigationFilter) GetAndOpOk() (*bool, bool)`

GetAndOpOk returns a tuple with the AndOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndOp

`func (o *InvestigationFilter) SetAndOp(v bool)`

SetAndOp sets AndOp field to given value.

### HasAndOp

`func (o *InvestigationFilter) HasAndOp() bool`

HasAndOp returns a boolean if a field has been set.

### GetCategory

`func (o *InvestigationFilter) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InvestigationFilter) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InvestigationFilter) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InvestigationFilter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFromCloseDate

`func (o *InvestigationFilter) GetFromCloseDate() time.Time`

GetFromCloseDate returns the FromCloseDate field if non-nil, zero value otherwise.

### GetFromCloseDateOk

`func (o *InvestigationFilter) GetFromCloseDateOk() (*time.Time, bool)`

GetFromCloseDateOk returns a tuple with the FromCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCloseDate

`func (o *InvestigationFilter) SetFromCloseDate(v time.Time)`

SetFromCloseDate sets FromCloseDate field to given value.

### HasFromCloseDate

`func (o *InvestigationFilter) HasFromCloseDate() bool`

HasFromCloseDate returns a boolean if a field has been set.

### GetFromDate

`func (o *InvestigationFilter) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *InvestigationFilter) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *InvestigationFilter) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *InvestigationFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *InvestigationFilter) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *InvestigationFilter) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *InvestigationFilter) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *InvestigationFilter) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetId

`func (o *InvestigationFilter) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvestigationFilter) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvestigationFilter) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *InvestigationFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdsOnly

`func (o *InvestigationFilter) GetIdsOnly() bool`

GetIdsOnly returns the IdsOnly field if non-nil, zero value otherwise.

### GetIdsOnlyOk

`func (o *InvestigationFilter) GetIdsOnlyOk() (*bool, bool)`

GetIdsOnlyOk returns a tuple with the IdsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsOnly

`func (o *InvestigationFilter) SetIdsOnly(v bool)`

SetIdsOnly sets IdsOnly field to given value.

### HasIdsOnly

`func (o *InvestigationFilter) HasIdsOnly() bool`

HasIdsOnly returns a boolean if a field has been set.

### GetIgnoreWorkers

`func (o *InvestigationFilter) GetIgnoreWorkers() bool`

GetIgnoreWorkers returns the IgnoreWorkers field if non-nil, zero value otherwise.

### GetIgnoreWorkersOk

`func (o *InvestigationFilter) GetIgnoreWorkersOk() (*bool, bool)`

GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWorkers

`func (o *InvestigationFilter) SetIgnoreWorkers(v bool)`

SetIgnoreWorkers sets IgnoreWorkers field to given value.

### HasIgnoreWorkers

`func (o *InvestigationFilter) HasIgnoreWorkers() bool`

HasIgnoreWorkers returns a boolean if a field has been set.

### GetIncludeChildInv

`func (o *InvestigationFilter) GetIncludeChildInv() bool`

GetIncludeChildInv returns the IncludeChildInv field if non-nil, zero value otherwise.

### GetIncludeChildInvOk

`func (o *InvestigationFilter) GetIncludeChildInvOk() (*bool, bool)`

GetIncludeChildInvOk returns a tuple with the IncludeChildInv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChildInv

`func (o *InvestigationFilter) SetIncludeChildInv(v bool)`

SetIncludeChildInv sets IncludeChildInv field to given value.

### HasIncludeChildInv

`func (o *InvestigationFilter) HasIncludeChildInv() bool`

HasIncludeChildInv returns a boolean if a field has been set.

### GetName

`func (o *InvestigationFilter) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationFilter) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationFilter) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *InvestigationFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotCategory

`func (o *InvestigationFilter) GetNotCategory() []string`

GetNotCategory returns the NotCategory field if non-nil, zero value otherwise.

### GetNotCategoryOk

`func (o *InvestigationFilter) GetNotCategoryOk() (*[]string, bool)`

GetNotCategoryOk returns a tuple with the NotCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCategory

`func (o *InvestigationFilter) SetNotCategory(v []string)`

SetNotCategory sets NotCategory field to given value.

### HasNotCategory

`func (o *InvestigationFilter) HasNotCategory() bool`

HasNotCategory returns a boolean if a field has been set.

### GetNotIDs

`func (o *InvestigationFilter) GetNotIDs() []string`

GetNotIDs returns the NotIDs field if non-nil, zero value otherwise.

### GetNotIDsOk

`func (o *InvestigationFilter) GetNotIDsOk() (*[]string, bool)`

GetNotIDsOk returns a tuple with the NotIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIDs

`func (o *InvestigationFilter) SetNotIDs(v []string)`

SetNotIDs sets NotIDs field to given value.

### HasNotIDs

`func (o *InvestigationFilter) HasNotIDs() bool`

HasNotIDs returns a boolean if a field has been set.

### GetPage

`func (o *InvestigationFilter) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InvestigationFilter) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InvestigationFilter) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *InvestigationFilter) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPeriod

`func (o *InvestigationFilter) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvestigationFilter) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvestigationFilter) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *InvestigationFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetReason

`func (o *InvestigationFilter) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvestigationFilter) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvestigationFilter) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvestigationFilter) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSearchAfter

`func (o *InvestigationFilter) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *InvestigationFilter) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *InvestigationFilter) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *InvestigationFilter) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchAfterElastic

`func (o *InvestigationFilter) GetSearchAfterElastic() []string`

GetSearchAfterElastic returns the SearchAfterElastic field if non-nil, zero value otherwise.

### GetSearchAfterElasticOk

`func (o *InvestigationFilter) GetSearchAfterElasticOk() (*[]string, bool)`

GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfterElastic

`func (o *InvestigationFilter) SetSearchAfterElastic(v []string)`

SetSearchAfterElastic sets SearchAfterElastic field to given value.

### HasSearchAfterElastic

`func (o *InvestigationFilter) HasSearchAfterElastic() bool`

HasSearchAfterElastic returns a boolean if a field has been set.

### GetSearchBefore

`func (o *InvestigationFilter) GetSearchBefore() []string`

GetSearchBefore returns the SearchBefore field if non-nil, zero value otherwise.

### GetSearchBeforeOk

`func (o *InvestigationFilter) GetSearchBeforeOk() (*[]string, bool)`

GetSearchBeforeOk returns a tuple with the SearchBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBefore

`func (o *InvestigationFilter) SetSearchBefore(v []string)`

SetSearchBefore sets SearchBefore field to given value.

### HasSearchBefore

`func (o *InvestigationFilter) HasSearchBefore() bool`

HasSearchBefore returns a boolean if a field has been set.

### GetSearchBeforeElastic

`func (o *InvestigationFilter) GetSearchBeforeElastic() []string`

GetSearchBeforeElastic returns the SearchBeforeElastic field if non-nil, zero value otherwise.

### GetSearchBeforeElasticOk

`func (o *InvestigationFilter) GetSearchBeforeElasticOk() (*[]string, bool)`

GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBeforeElastic

`func (o *InvestigationFilter) SetSearchBeforeElastic(v []string)`

SetSearchBeforeElastic sets SearchBeforeElastic field to given value.

### HasSearchBeforeElastic

`func (o *InvestigationFilter) HasSearchBeforeElastic() bool`

HasSearchBeforeElastic returns a boolean if a field has been set.

### GetSize

`func (o *InvestigationFilter) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InvestigationFilter) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InvestigationFilter) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *InvestigationFilter) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *InvestigationFilter) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *InvestigationFilter) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *InvestigationFilter) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *InvestigationFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStatus

`func (o *InvestigationFilter) GetStatus() []float64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvestigationFilter) GetStatusOk() (*[]float64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvestigationFilter) SetStatus(v []float64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvestigationFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeFrame

`func (o *InvestigationFilter) GetTimeFrame() int64`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *InvestigationFilter) GetTimeFrameOk() (*int64, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *InvestigationFilter) SetTimeFrame(v int64)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *InvestigationFilter) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetToCloseDate

`func (o *InvestigationFilter) GetToCloseDate() time.Time`

GetToCloseDate returns the ToCloseDate field if non-nil, zero value otherwise.

### GetToCloseDateOk

`func (o *InvestigationFilter) GetToCloseDateOk() (*time.Time, bool)`

GetToCloseDateOk returns a tuple with the ToCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCloseDate

`func (o *InvestigationFilter) SetToCloseDate(v time.Time)`

SetToCloseDate sets ToCloseDate field to given value.

### HasToCloseDate

`func (o *InvestigationFilter) HasToCloseDate() bool`

HasToCloseDate returns a boolean if a field has been set.

### GetToDate

`func (o *InvestigationFilter) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *InvestigationFilter) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *InvestigationFilter) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *InvestigationFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetType

`func (o *InvestigationFilter) GetType() []float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvestigationFilter) GetTypeOk() (*[]float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvestigationFilter) SetType(v []float64)`

SetType sets Type field to given value.

### HasType

`func (o *InvestigationFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *InvestigationFilter) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InvestigationFilter) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InvestigationFilter) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *InvestigationFilter) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


