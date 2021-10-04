# InsightCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Scores** | Pointer to [**map[string]DBotScore**](DBotScore.md) |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewInsightCache

`func NewInsightCache() *InsightCache`

NewInsightCache instantiates a new InsightCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightCacheWithDefaults

`func NewInsightCacheWithDefaults() *InsightCache`

NewInsightCacheWithDefaults instantiates a new InsightCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlight

`func (o *InsightCache) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *InsightCache) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *InsightCache) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *InsightCache) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *InsightCache) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightCache) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightCache) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InsightCache) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *InsightCache) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InsightCache) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InsightCache) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *InsightCache) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNumericId

`func (o *InsightCache) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *InsightCache) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *InsightCache) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *InsightCache) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *InsightCache) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *InsightCache) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *InsightCache) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *InsightCache) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetScores

`func (o *InsightCache) GetScores() map[string]DBotScore`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *InsightCache) GetScoresOk() (*map[string]DBotScore, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *InsightCache) SetScores(v map[string]DBotScore)`

SetScores sets Scores field to given value.

### HasScores

`func (o *InsightCache) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *InsightCache) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *InsightCache) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *InsightCache) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *InsightCache) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *InsightCache) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *InsightCache) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *InsightCache) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *InsightCache) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetVersion

`func (o *InsightCache) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InsightCache) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InsightCache) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InsightCache) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


