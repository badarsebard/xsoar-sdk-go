# ReportQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomGroupBy** | Pointer to [**[]map[string]CustomGroup**](map[string]CustomGroup.md) | CustomGroups - a custom group for each group-by element | [optional] 
**Filter** | Pointer to **map[string]interface{}** |  | [optional] 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewReportQuery

`func NewReportQuery() *ReportQuery`

NewReportQuery instantiates a new ReportQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportQueryWithDefaults

`func NewReportQueryWithDefaults() *ReportQuery`

NewReportQueryWithDefaults instantiates a new ReportQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomGroupBy

`func (o *ReportQuery) GetCustomGroupBy() []map[string]CustomGroup`

GetCustomGroupBy returns the CustomGroupBy field if non-nil, zero value otherwise.

### GetCustomGroupByOk

`func (o *ReportQuery) GetCustomGroupByOk() (*[]map[string]CustomGroup, bool)`

GetCustomGroupByOk returns a tuple with the CustomGroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGroupBy

`func (o *ReportQuery) SetCustomGroupBy(v []map[string]CustomGroup)`

SetCustomGroupBy sets CustomGroupBy field to given value.

### HasCustomGroupBy

`func (o *ReportQuery) HasCustomGroupBy() bool`

HasCustomGroupBy returns a boolean if a field has been set.

### GetFilter

`func (o *ReportQuery) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReportQuery) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReportQuery) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReportQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *ReportQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ReportQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ReportQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ReportQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetKeys

`func (o *ReportQuery) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReportQuery) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReportQuery) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReportQuery) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetType

`func (o *ReportQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportQuery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportQuery) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


