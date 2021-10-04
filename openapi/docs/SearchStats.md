# SearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Committed** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SearchSource** | Pointer to **string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchStats

`func NewSearchStats() *SearchStats`

NewSearchStats instantiates a new SearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsWithDefaults

`func NewSearchStatsWithDefaults() *SearchStats`

NewSearchStatsWithDefaults instantiates a new SearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitted

`func (o *SearchStats) GetCommitted() time.Time`

GetCommitted returns the Committed field if non-nil, zero value otherwise.

### GetCommittedOk

`func (o *SearchStats) GetCommittedOk() (*time.Time, bool)`

GetCommittedOk returns a tuple with the Committed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitted

`func (o *SearchStats) SetCommitted(v time.Time)`

SetCommitted sets Committed field to given value.

### HasCommitted

`func (o *SearchStats) HasCommitted() bool`

HasCommitted returns a boolean if a field has been set.

### GetDuration

`func (o *SearchStats) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchStats) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchStats) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SearchStats) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFromDate

`func (o *SearchStats) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *SearchStats) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *SearchStats) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *SearchStats) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetHighlight

`func (o *SearchStats) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *SearchStats) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *SearchStats) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *SearchStats) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *SearchStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *SearchStats) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SearchStats) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SearchStats) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SearchStats) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNumericId

`func (o *SearchStats) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *SearchStats) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *SearchStats) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *SearchStats) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *SearchStats) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *SearchStats) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *SearchStats) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *SearchStats) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetQuery

`func (o *SearchStats) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchStats) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchStats) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchStats) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSearchSource

`func (o *SearchStats) GetSearchSource() string`

GetSearchSource returns the SearchSource field if non-nil, zero value otherwise.

### GetSearchSourceOk

`func (o *SearchStats) GetSearchSourceOk() (*string, bool)`

GetSearchSourceOk returns a tuple with the SearchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSource

`func (o *SearchStats) SetSearchSource(v string)`

SetSearchSource sets SearchSource field to given value.

### HasSearchSource

`func (o *SearchStats) HasSearchSource() bool`

HasSearchSource returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *SearchStats) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *SearchStats) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *SearchStats) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *SearchStats) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *SearchStats) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *SearchStats) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *SearchStats) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *SearchStats) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetToDate

`func (o *SearchStats) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *SearchStats) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *SearchStats) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *SearchStats) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetUsername

`func (o *SearchStats) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SearchStats) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SearchStats) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SearchStats) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVersion

`func (o *SearchStats) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchStats) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchStats) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SearchStats) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


