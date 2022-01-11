# StatsTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**Groups**](Groups.md) |  | [optional] 
**Text** | Pointer to **string** | Describe the complete text for the text widget, after placeholders injection. | [optional] 

## Methods

### NewStatsTextResponse

`func NewStatsTextResponse() *StatsTextResponse`

NewStatsTextResponse instantiates a new StatsTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsTextResponseWithDefaults

`func NewStatsTextResponseWithDefaults() *StatsTextResponse`

NewStatsTextResponseWithDefaults instantiates a new StatsTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *StatsTextResponse) GetGroups() Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *StatsTextResponse) GetGroupsOk() (*Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *StatsTextResponse) SetGroups(v Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *StatsTextResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetText

`func (o *StatsTextResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StatsTextResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StatsTextResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StatsTextResponse) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


