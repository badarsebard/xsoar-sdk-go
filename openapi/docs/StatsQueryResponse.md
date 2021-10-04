# StatsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | data array or object describing the statistics data based on type. | [optional] 
**WidgetCellId** | Pointer to **string** |  | [optional] 

## Methods

### NewStatsQueryResponse

`func NewStatsQueryResponse() *StatsQueryResponse`

NewStatsQueryResponse instantiates a new StatsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsQueryResponseWithDefaults

`func NewStatsQueryResponseWithDefaults() *StatsQueryResponse`

NewStatsQueryResponseWithDefaults instantiates a new StatsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StatsQueryResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatsQueryResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatsQueryResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *StatsQueryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWidgetCellId

`func (o *StatsQueryResponse) GetWidgetCellId() string`

GetWidgetCellId returns the WidgetCellId field if non-nil, zero value otherwise.

### GetWidgetCellIdOk

`func (o *StatsQueryResponse) GetWidgetCellIdOk() (*string, bool)`

GetWidgetCellIdOk returns a tuple with the WidgetCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetCellId

`func (o *StatsQueryResponse) SetWidgetCellId(v string)`

SetWidgetCellId sets WidgetCellId field to given value.

### HasWidgetCellId

`func (o *StatsQueryResponse) HasWidgetCellId() bool`

HasWidgetCellId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


