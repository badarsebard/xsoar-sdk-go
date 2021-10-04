# PlaybookInputQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**QueryEntity** | Pointer to **string** |  | [optional] 
**QueryState** | Pointer to [**QueryState**](QueryState.md) |  | [optional] 
**Results** | Pointer to **map[string]interface{}** |  | [optional] 
**RunFromLastJobTime** | Pointer to **bool** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPlaybookInputQuery

`func NewPlaybookInputQuery() *PlaybookInputQuery`

NewPlaybookInputQuery instantiates a new PlaybookInputQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookInputQueryWithDefaults

`func NewPlaybookInputQueryWithDefaults() *PlaybookInputQuery`

NewPlaybookInputQueryWithDefaults instantiates a new PlaybookInputQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *PlaybookInputQuery) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *PlaybookInputQuery) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *PlaybookInputQuery) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *PlaybookInputQuery) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *PlaybookInputQuery) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *PlaybookInputQuery) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *PlaybookInputQuery) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *PlaybookInputQuery) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetPeriod

`func (o *PlaybookInputQuery) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *PlaybookInputQuery) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *PlaybookInputQuery) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *PlaybookInputQuery) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuery

`func (o *PlaybookInputQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PlaybookInputQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PlaybookInputQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PlaybookInputQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryEntity

`func (o *PlaybookInputQuery) GetQueryEntity() string`

GetQueryEntity returns the QueryEntity field if non-nil, zero value otherwise.

### GetQueryEntityOk

`func (o *PlaybookInputQuery) GetQueryEntityOk() (*string, bool)`

GetQueryEntityOk returns a tuple with the QueryEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEntity

`func (o *PlaybookInputQuery) SetQueryEntity(v string)`

SetQueryEntity sets QueryEntity field to given value.

### HasQueryEntity

`func (o *PlaybookInputQuery) HasQueryEntity() bool`

HasQueryEntity returns a boolean if a field has been set.

### GetQueryState

`func (o *PlaybookInputQuery) GetQueryState() QueryState`

GetQueryState returns the QueryState field if non-nil, zero value otherwise.

### GetQueryStateOk

`func (o *PlaybookInputQuery) GetQueryStateOk() (*QueryState, bool)`

GetQueryStateOk returns a tuple with the QueryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryState

`func (o *PlaybookInputQuery) SetQueryState(v QueryState)`

SetQueryState sets QueryState field to given value.

### HasQueryState

`func (o *PlaybookInputQuery) HasQueryState() bool`

HasQueryState returns a boolean if a field has been set.

### GetResults

`func (o *PlaybookInputQuery) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlaybookInputQuery) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlaybookInputQuery) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *PlaybookInputQuery) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRunFromLastJobTime

`func (o *PlaybookInputQuery) GetRunFromLastJobTime() bool`

GetRunFromLastJobTime returns the RunFromLastJobTime field if non-nil, zero value otherwise.

### GetRunFromLastJobTimeOk

`func (o *PlaybookInputQuery) GetRunFromLastJobTimeOk() (*bool, bool)`

GetRunFromLastJobTimeOk returns a tuple with the RunFromLastJobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunFromLastJobTime

`func (o *PlaybookInputQuery) SetRunFromLastJobTime(v bool)`

SetRunFromLastJobTime sets RunFromLastJobTime field to given value.

### HasRunFromLastJobTime

`func (o *PlaybookInputQuery) HasRunFromLastJobTime() bool`

HasRunFromLastJobTime returns a boolean if a field has been set.

### GetToDate

`func (o *PlaybookInputQuery) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *PlaybookInputQuery) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *PlaybookInputQuery) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *PlaybookInputQuery) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


